import { useApp } from "@modelcontextprotocol/ext-apps/react";
import {
  applyHostStyleVariables,
  applyDocumentTheme,
  type McpUiHostContext,
} from "@modelcontextprotocol/ext-apps";
import { StrictMode, useState, useEffect } from "react";
import { createRoot } from "react-dom/client";
import type { MCPFindingsResponse, MCPFinding, MCPMitigation } from "./types";
import styles from "./mcp-app.module.css";

function DynamicFindingsApp() {
  const [resultsData, setResultsData] = useState<MCPFindingsResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const { app, error: appError } = useApp({
    appInfo: { name: "Dynamic Findings", version: "1.0.0" },
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
            // structuredContent is Record<string, unknown>, convert to our type
            const data = result.structuredContent as unknown as MCPFindingsResponse;
            setResultsData(data);
            setError(null);
            console.info("Loaded results from structuredContent:", data);
          } else {
            // Fallback: try parsing JSON from text content
            const jsonContent = result.content?.find((c) => c.type === "text" && c.text.includes("{"));
            if (jsonContent && jsonContent.type === "text") {
              const data = JSON.parse(jsonContent.text) as MCPFindingsResponse;
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

  return <DynamicFindingsView data={resultsData} />;
}

interface DynamicFindingsViewProps {
  data: MCPFindingsResponse;
}

function DynamicFindingsView({ data }: DynamicFindingsViewProps) {
  const { application, summary, findings, pagination } = data;

  // Define severity order (5 to 0)
  const severityOrder = ['Very High', 'High', 'Medium', 'Low', 'Very Low', 'Info'];

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1>Dynamic Findings: {application.name}</h1>
        
        {pagination && (
          <div className={styles.paginationInfo}>
            Showing {findings.length} findings on page {pagination.current_page + 1} of {pagination.total_pages} (Total: {pagination.total_elements} findings)
          </div>
        )}
        
        <div className={styles.summary}>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Total Findings</span>
            <span className={styles.summaryValue}>{summary.total_findings}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Open Findings</span>
            <span className={styles.summaryValue}>{summary.open_findings}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Policy Violations</span>
            <span className={styles.summaryValue}>{summary.policy_violations}</span>
          </div>
        </div>

        <div className={styles.severityBreakdown}>
          {severityOrder.map((severity) => {
            const count = summary.by_severity[severity] as number || 0;
            return count > 0 && (
              <div key={severity} className={`${styles.severityItem} ${styles[severity.toLowerCase().replace(' ', '')]}`}>
                <strong>{count}</strong> {severity}
              </div>
            );
          })}
        </div>
      </div>

      {findings.length === 0 ? (
        <div className={styles.empty}>No findings to display</div>
      ) : (
        <div className={styles.tableContainer}>
          <table className={styles.table}>
            <thead>
              <tr>
                <th className={styles.expanderHeader}></th>
                <th>Flaw</th>
                <th className={styles.severityHeader}>Severity</th>
                <th>CWE</th>
                <th>URL</th>
                <th>Status</th>
                <th>Policy</th>
                <th>Mitigation</th>
              </tr>
            </thead>
            <tbody>
              {findings.map((finding: MCPFinding, index: number) => (
                <FindingRow key={index} finding={finding} />
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

interface FindingRowProps {
  finding: MCPFinding;
}

function FindingRow({ finding }: FindingRowProps) {
  const [isExpanded, setIsExpanded] = useState(false);

  return (
    <>
      <tr className={styles.findingRow} onClick={() => setIsExpanded(!isExpanded)}>
        <td className={styles.expanderCell}>
          <span className={styles.expander}>{isExpanded ? '▼' : '▶'}</span>
        </td>
        <td>{finding.flaw_id}</td>
        <td className={styles.severityCell}>
          <span className={`${styles.severityBadge} ${styles[finding.severity.toLowerCase().replace(' ', '')]}`}>
            {finding.severity}
          </span>
        </td>
        <td>
          <a
            href={`https://cwe.mitre.org/data/definitions/${finding.cwe_id}.html`}
            target="_blank"
            rel="noopener noreferrer"
            onClick={(e) => e.stopPropagation()}
          >
            {finding.cwe_id}
          </a>
        </td>
        <td>
          <div className={styles.url} title={finding.url || undefined}>
            {finding.url || '-'}
          </div>
        </td>
        <td className={styles.statusCell}>
          {finding.status === 'CLOSED' && (
            <span className={styles.statusIconClosed} title="Closed">✓</span>
          )}
          {finding.status === 'NEW' && (
            <span className={styles.statusIcon} title="New Finding">🆕</span>
          )}
        </td>
        <td className={styles.policyCell}>
          {finding.violates_policy && (
            <span className={styles.policyViolation} title="Violates Policy">⚠️</span>
          )}
        </td>
        <td className={styles.mitigationCell}>
          {finding.mitigation_status === 'PROPOSED' && (
            <span className={styles.mitigationIcon} title="Mitigation Proposed">⏳</span>
          )}
          {finding.mitigation_status === 'APPROVED' && (
            <span className={styles.mitigationIcon} title="Mitigation Approved">👍</span>
          )}
          {finding.mitigation_status === 'REJECTED' && (
            <span className={styles.mitigationIcon} title="Mitigation Rejected">👎</span>
          )}
        </td>
      </tr>
      {isExpanded && (
        <tr className={styles.expandedRow}>
          <td colSpan={8}>
            <div className={styles.expandedContent}>
              {finding.attack_vector && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Attack Vector</h4>
                  <div className={styles.description}>{finding.attack_vector}</div>
                </div>
              )}
              <div className={styles.expandedSection}>
                <h4 className={styles.expandedHeader}>Description</h4>
                <div className={styles.description}>{finding.description}</div>
              </div>
              {finding.mitigations && finding.mitigations.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Mitigations ({finding.mitigations.length})</h4>
                  {[...finding.mitigations]
                    .sort((a, b) => {
                      const dateA = new Date(a.date).getTime();
                      const dateB = new Date(b.date).getTime();
                      return dateB - dateA; // Most recent first
                    })
                    .map((mitigation: MCPMitigation, idx: number) => (
                    <div key={idx} className={styles.mitigation}>
                      <div className={styles.mitigationHeader}>
                        <strong>{mitigation.action}</strong>
                        <span className={styles.mitigationMeta}>
                          by {mitigation.submitter} on {new Date(mitigation.date).toLocaleString('en-US', { hour12: false })}
                        </span>
                      </div>
                      {mitigation.comment && (
                        <div className={styles.mitigationComment}>{mitigation.comment}</div>
                      )}
                    </div>
                  ))}
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
    <DynamicFindingsApp />
  </StrictMode>,
);
