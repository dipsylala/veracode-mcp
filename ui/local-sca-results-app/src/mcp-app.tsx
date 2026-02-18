import { useApp } from "@modelcontextprotocol/ext-apps/react";
import {
  applyHostStyleVariables,
  applyDocumentTheme,
  type McpUiHostContext,
} from "@modelcontextprotocol/ext-apps";
import { StrictMode, useState, useEffect } from "react";
import { createRoot } from "react-dom/client";
import type { SCAResultsResponse, SCAFinding, SCACVSS, SCALocation } from "./types";
import styles from "./mcp-app.module.css";

// Normalize severity names for CSS class names
function normalizeSeverity(severity: string): string {
  return severity.toLowerCase().replace(/\s+/g, '');
}

function LocalSCAResultsApp() {
  const [resultsData, setResultsData] = useState<SCAResultsResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const { app, error: appError } = useApp({
    appInfo: { name: "Local SCA Results", version: "1.0.0" },
    capabilities: {},
    onAppCreated: (app) => {
      app.ontoolinput = async (input) => {
        console.info("Received tool call input:", input);
      };

      app.ontoolresult = async (result) => {
        console.info("Received tool call result:", result);
        try {
          // Extract data from structuredContent (MCP Apps pattern)
          if (result.structuredContent) {
            const data = result.structuredContent as unknown as SCAResultsResponse;
            setResultsData(data);
            setError(null);
            console.info("Loaded results from structuredContent:", data);
          } else {
            // Fallback: try parsing JSON from text content
            const jsonContent = result.content?.find((c) => c.type === "text" && c.text.includes("{"));
            if (jsonContent && jsonContent.type === "text") {
              const data = JSON.parse(jsonContent.text) as SCAResultsResponse;
              setResultsData(data);
              setError(null);
              console.info("Loaded results from text content:", data);
            } else {
              setError("No data found in tool result");
            }
          }
        } catch (e) {
          console.error("Failed to parse results:", e);
          setError(e instanceof Error ? e.message : "Failed to parse results");
        }
      };

      app.onhostcontextchanged = (ctx: McpUiHostContext) => {
        console.info("Host context changed:", ctx);
        
        // Apply host styling
        if (ctx.theme) applyDocumentTheme(ctx.theme);
        if (ctx.styles?.variables) applyHostStyleVariables(ctx.styles.variables);
      };

      app.onerror = (err) => {
        console.error("App error:", err);
        setError(err.message);
      };
    },
  });

  // Apply host styles after app is connected
  useEffect(() => {
    if (app) {
      const ctx = app.getHostContext();
      if (ctx) {
        if (ctx.theme) applyDocumentTheme(ctx.theme);
        if (ctx.styles?.variables) applyHostStyleVariables(ctx.styles.variables);
      }
    }
  }, [app]);

  if (appError) {
    return <div className={styles.error}><strong>ERROR:</strong> {appError.message}</div>;
  }

  if (!app) {
    return <div className={styles.loading}>Connecting...</div>;
  }

  if (error) {
    return <div className={styles.error}><strong>ERROR:</strong> {error}</div>;
  }

  if (!resultsData) {
    return <div className={styles.loading}>Loading results...</div>;
  }

  return <LocalSCAResultsView data={resultsData} />;
}

interface LocalSCAResultsViewProps {
  data: SCAResultsResponse;
}

function LocalSCAResultsView({ data }: LocalSCAResultsViewProps) {
  const { application, summary, findings, pagination, filters } = data;

  // If no findings, show simple message
  if (findings.length === 0) {
    return (
      <div className={styles.container}>
        <div className={styles.header}>
          <h1>Local SCA Results: {application.name}</h1>
          {filters && (
            <div className={styles.filtersInfo}>
              <strong>Active Filters:</strong>
              {filters.cve && <span className={styles.filterBadge}>CVE: {filters.cve}</span>}
              {filters.component_name && <span className={styles.filterBadge}>Component: {filters.component_name}</span>}
              {filters.severity_gte && <span className={styles.filterBadge}>Min Severity: {filters.severity_gte}</span>}
            </div>
          )}
        </div>
        <div className={styles.empty}>No vulnerabilities found{filters ? ' matching filters' : ''}</div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1>Local SCA Results: {application.name}</h1>
        
        {filters && (
          <div className={styles.filtersInfo}>
            <strong>Active Filters:</strong>
            {filters.cve && <span className={styles.filterBadge}>CVE: {filters.cve}</span>}
            {filters.component_name && <span className={styles.filterBadge}>Component: {filters.component_name}</span>}
            {filters.severity_gte && <span className={styles.filterBadge}>Min Severity: {filters.severity_gte}</span>}
          </div>
        )}
        
        {pagination && (
          <div className={styles.paginationInfo}>
            Showing {findings.length} findings on page {pagination.current_page + 1} of {pagination.total_pages} (Total: {pagination.total_elements} findings)
          </div>
        )}
        
        <div className={styles.summary}>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Unique Vulnerabilities</span>
            <span className={styles.summaryValue}>{summary.unique_vulnerabilities}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Vulnerable Components</span>
            <span className={styles.summaryValue}>{summary.vulnerable_components}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Total Matches</span>
            <span className={styles.summaryValue}>{summary.total_matches}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>With EPSS Data</span>
            <span className={styles.summaryValue}>{summary.epss_data_available}</span>
          </div>
        </div>

        <div className={styles.severityBreakdown}>
          {['critical', 'high', 'medium', 'low']
            .map(severity => {
              const count = summary.by_severity[severity as keyof typeof summary.by_severity] as number;
              return count > 0 && (
                <div key={severity} className={`${styles.severityItem} ${styles[normalizeSeverity(severity)]}`}>
                  <strong>{count}</strong> {severity}
                </div>
              );
            })}
        </div>
      </div>

      <div className={styles.tableContainer}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th className={styles.expanderHeader}></th>
              <th>CVE</th>
              <th className={styles.severityHeader}>Severity</th>
              <th>Component</th>
              <th>Version</th>
              <th className={styles.epssHeader}>EPSS</th>
              <th>Fix Available</th>
            </tr>
          </thead>
          <tbody>
            {findings.map((finding: SCAFinding, index: number) => (
              <FindingRow key={index} finding={finding} />
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

interface FindingRowProps {
  finding: SCAFinding;
}

function FindingRow({ finding }: FindingRowProps) {
  const [isExpanded, setIsExpanded] = useState(false);

  const hasEPSS = finding.epss && finding.epss.score > 0;
  const hasFix = finding.fix.recommended_version || (finding.fix.versions && finding.fix.versions.length > 0);
  
  // Use EPSS CVE if available, otherwise use vulnerability_id
  const scaID = (finding.epss && finding.epss.cve) ? finding.epss.cve : finding.vulnerability_id;
  
  // It's a CVE only if it came from EPSS
  const isCVE = !!(finding.epss && finding.epss.cve);

  return (
    <>
      <tr className={styles.findingRow} onClick={() => setIsExpanded(!isExpanded)}>
        <td className={styles.expanderCell}>
          <span className={styles.expander}>{isExpanded ? '▼' : '▶'}</span>
        </td>
        <td>
          {isCVE ? (
            <a
              href={`https://nvd.nist.gov/vuln/detail/${scaID}`}
              target="_blank"
              rel="noopener noreferrer"
              onClick={(e) => e.stopPropagation()}
              className={styles.cveLink}
            >
              {scaID}
            </a>
          ) : (
            <span className={styles.componentName}>{scaID}</span>
          )}
        </td>
        <td className={styles.severityCell}>
          <span className={`${styles.severityBadge} ${styles[normalizeSeverity(finding.severity)]}`}>
            {finding.severity}
          </span>
        </td>
        <td>
          <div className={styles.componentName}>{finding.component.name}</div>
        </td>
        <td>
          <div className={styles.version}>{finding.component.version}</div>
        </td>
        <td className={styles.epssCell}>
          {hasEPSS && (
            <div className={styles.epssScore} title={`EPSS: ${(finding.epss!.score * 100).toFixed(2)}% (${finding.epss!.percentile.toFixed(1)}th percentile)`}>
              {(finding.epss!.score * 100).toFixed(2)}%
            </div>
          )}
          {!hasEPSS && <span className={styles.noData}>-</span>}
        </td>
        <td>
          {hasFix ? (
            <span className={styles.fixAvailable}>✓</span>
          ) : (
            <span className={styles.noFix}>-</span>
          )}
        </td>
      </tr>
      {isExpanded && (
        <tr className={styles.expandedRow}>
          <td colSpan={7}>
            <div className={styles.expandedContent}>
              <div className={styles.expandedSection}>
                <h4 className={styles.expandedHeader}>Description</h4>
                <div className={styles.description}>{finding.description}</div>
              </div>

              <div className={styles.expandedTwoColumn}>
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Component Details</h4>
                  <div className={styles.detailsList}>
                    <div className={styles.detailItem}>
                      <span className={styles.detailLabel}>Type:</span>
                      <span className={styles.detailValue}>{finding.component.type}</span>
                    </div>
                    <div className={styles.detailItem}>
                      <span className={styles.detailLabel}>Language:</span>
                      <span className={styles.detailValue}>{finding.component.language}</span>
                    </div>
                    {finding.component.purl && (
                      <div className={styles.detailItem}>
                        <span className={styles.detailLabel}>PURL:</span>
                        <span className={styles.detailValue} style={{ fontFamily: 'var(--font-mono, monospace)', fontSize: '11px' }}>
                          {finding.component.purl}
                        </span>
                      </div>
                    )}
                    {finding.component.licenses && finding.component.licenses.length > 0 && (
                      <div className={styles.detailItem}>
                        <span className={styles.detailLabel}>Licenses:</span>
                        <span className={styles.detailValue}>{finding.component.licenses.join(', ')}</span>
                      </div>
                    )}
                  </div>
                </div>

                {hasFix && (
                  <div className={styles.expandedSection}>
                    <h4 className={styles.expandedHeader}>Fix Information</h4>
                    <div className={styles.detailsList}>
                      <div className={styles.detailItem}>
                        <span className={styles.detailLabel}>Status:</span>
                        <span className={styles.detailValue}>{finding.fix.state}</span>
                      </div>
                      {finding.fix.recommended_version && (
                        <div className={styles.detailItem}>
                          <span className={styles.detailLabel}>Recommended:</span>
                          <span className={`${styles.detailValue} ${styles.recommendedVersion}`}>
                            {finding.fix.recommended_version}
                          </span>
                        </div>
                      )}
                      {finding.fix.versions && finding.fix.versions.length > 0 && (
                        <div className={styles.detailItem}>
                          <span className={styles.detailLabel}>Fixed in:</span>
                          <span className={styles.detailValue}>
                            {finding.fix.versions.slice(0, 5).join(', ')}
                            {finding.fix.versions.length > 5 && ' ...'}
                          </span>
                        </div>
                      )}
                    </div>
                  </div>
                )}
              </div>

              {(finding.cvss && finding.cvss.length > 0) || hasEPSS ? (
                <div className={styles.expandedTwoColumn}>
                  {finding.cvss && finding.cvss.length > 0 && (
                    <div className={styles.expandedSection}>
                      <h4 className={styles.expandedHeader}>CVSS Scores</h4>
                      <div className={styles.cvssContainer}>
                        {finding.cvss.map((cvss: SCACVSS, idx: number) => (
                          <div key={idx} className={styles.cvssItem}>
                            <div className={styles.cvssVersion}>CVSS {cvss.version}</div>
                            <div className={styles.cvssMetrics}>
                              <div className={styles.cvssMetric}>
                                <span className={styles.metricLabel}>Base:</span>
                                <span className={styles.metricValue}>{cvss.base_score.toFixed(1)}</span>
                              </div>
                              <div className={styles.cvssMetric}>
                                <span className={styles.metricLabel}>Exploitability:</span>
                                <span className={styles.metricValue}>{cvss.exploitability_score.toFixed(1)}</span>
                              </div>
                              <div className={styles.cvssMetric}>
                                <span className={styles.metricLabel}>Impact:</span>
                                <span className={styles.metricValue}>{cvss.impact_score.toFixed(1)}</span>
                              </div>
                            </div>
                            {cvss.vector && (
                              <div className={styles.cvssVector}>Vector: {cvss.vector}</div>
                            )}
                          </div>
                        ))}
                      </div>
                    </div>
                  )}

                  {hasEPSS && (
                    <div className={styles.expandedSection}>
                      <h4 className={styles.expandedHeader}>EPSS (Exploit Prediction Scoring System)</h4>
                      <div className={styles.epssDetails}>
                        <div className={styles.epssMetricLarge}>
                          <span className={styles.epssLabel}>Exploitation Probability:</span>
                          <span className={styles.epssValueLarge}>{(finding.epss!.score * 100).toFixed(2)}%</span>
                        </div>
                        <div className={styles.epssMetric}>
                          <span className={styles.epssLabel}>Percentile:</span>
                          <span className={styles.epssValue}>{finding.epss!.percentile.toFixed(1)}th</span>
                        </div>
                        <div className={styles.epssNote}>
                          EPSS estimates the probability that this vulnerability will be exploited in the wild within 30 days.
                          Updated: {new Date(finding.epss!.date).toLocaleDateString()}
                        </div>
                      </div>
                    </div>
                  )}
                </div>
              ) : null}

              {finding.cwes && finding.cwes.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>CWEs</h4>
                  <div className={styles.cweList}>
                    {finding.cwes.map((cwe: string, idx: number) => {
                      const cweId = cwe.replace('CWE-', '');
                      return (
                        <a
                          key={idx}
                          href={`https://cwe.mitre.org/data/definitions/${cweId}.html`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className={styles.cweLink}
                        >
                          {cwe}
                        </a>
                      );
                    })}
                  </div>
                </div>
              )}

              {finding.related_cves && finding.related_cves.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Related CVEs</h4>
                  <div className={styles.relatedCveList}>
                    {finding.related_cves.map((cve: string, idx: number) => (
                      <a
                        key={idx}
                        href={`https://nvd.nist.gov/vuln/detail/${cve}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className={styles.relatedCveLink}
                      >
                        {cve}
                      </a>
                    ))}
                  </div>
                </div>
              )}

              {finding.component.locations && finding.component.locations.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Locations ({finding.component.locations.length})</h4>
                  <div className={styles.locationList}>
                    {finding.component.locations.map((loc: SCALocation, idx: number) => (
                      <div key={idx} className={styles.locationItem}>
                        <div className={styles.locationPath}>{loc.path}</div>
                        {loc.accessPath && loc.accessPath !== loc.path && (
                          <div className={styles.locationAccess}>→ {loc.accessPath}</div>
                        )}
                      </div>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </td>
        </tr>
      )}
    </>
  );
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <LocalSCAResultsApp />
  </StrictMode>,
);
