import { useApp } from "@modelcontextprotocol/ext-apps/react";
import {
  applyHostStyleVariables,
  applyDocumentTheme,
  type McpUiHostContext,
} from "@modelcontextprotocol/ext-apps";
import { StrictMode, useState, useEffect } from "react";
import { createRoot } from "react-dom/client";
import type { IACResultsResponse, IACFinding } from "./types";
import styles from "./mcp-app.module.css";

function normalizeSeverity(severity: string): string {
  return severity.toLowerCase().replace(/\s+/g, '');
}

function LocalIACResultsApp() {
  const [resultsData, setResultsData] = useState<IACResultsResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const { app, error: appError } = useApp({
    appInfo: { name: "Local IaC Results", version: "1.0.0" },
    capabilities: {},
    onAppCreated: (app) => {
      app.ontoolinput = async (input) => {
        console.info("Received tool call input:", input);
      };

      app.ontoolresult = async (result) => {
        console.info("Received tool call result:", result);
        try {
          if (result.structuredContent) {
            const data = result.structuredContent as unknown as IACResultsResponse;
            setResultsData(data);
            setError(null);
          } else {
            const jsonContent = result.content?.find((c) => c.type === "text" && c.text.includes("{"));
            if (jsonContent && jsonContent.type === "text") {
              const data = JSON.parse(jsonContent.text) as IACResultsResponse;
              setResultsData(data);
              setError(null);
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
        if (ctx.theme) applyDocumentTheme(ctx.theme);
        if (ctx.styles?.variables) applyHostStyleVariables(ctx.styles.variables);
      };

      app.onerror = (err) => {
        console.error("App error:", err);
        setError(err.message);
      };
    },
  });

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

  return <LocalIACResultsView data={resultsData} />;
}

interface LocalIACResultsViewProps {
  data: IACResultsResponse;
}

function LocalIACResultsView({ data }: LocalIACResultsViewProps) {
  const { application, summary, findings, pagination, filters } = data;

  if (findings.length === 0) {
    return (
      <div className={styles.container}>
        <div className={styles.header}>
          <h1>Local IaC Results: {application.name}</h1>
          {filters && <ActiveFilters filters={filters} />}
        </div>
        <div className={styles.empty}>No findings{filters ? ' matching filters' : ''}</div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1>Local IaC Results: {application.name}</h1>

        {filters && <ActiveFilters filters={filters} />}

        {pagination && (
          <div className={styles.paginationInfo}>
            Showing {findings.length} findings on page {pagination.current_page + 1} of {pagination.total_pages} (Total: {pagination.total_elements} findings)
          </div>
        )}

        <div className={styles.summary}>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Unique Checks</span>
            <span className={styles.summaryValue}>{summary.unique_checks}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Unique Targets</span>
            <span className={styles.summaryValue}>{summary.unique_targets}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Fail</span>
            <span className={styles.summaryValue}>{summary.fail}</span>
          </div>
          <div className={styles.summaryItem}>
            <span className={styles.summaryLabel}>Pass</span>
            <span className={styles.summaryValue}>{summary.pass}</span>
          </div>
        </div>

        <div className={styles.severityBreakdown}>
          {(['critical', 'high', 'medium', 'low'] as const).map(severity => {
            const count = summary.by_severity[severity];
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
              <th>Check ID</th>
              <th className={styles.severityHeader}>Severity</th>
              <th className={styles.statusHeader}>Status</th>
              <th>Target</th>
              <th>Title</th>
            </tr>
          </thead>
          <tbody>
            {findings.map((finding: IACFinding, index: number) => (
              <FindingRow key={index} finding={finding} />
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

interface ActiveFiltersProps {
  filters: NonNullable<IACResultsResponse['filters']>;
}

function ActiveFilters({ filters }: ActiveFiltersProps) {
  return (
    <div className={styles.filtersInfo}>
      <strong>Active Filters:</strong>
      {filters.target && <span className={styles.filterBadge}>Target: {filters.target}</span>}
      {filters.check_id && <span className={styles.filterBadge}>Check: {filters.check_id}</span>}
      {filters.severity_gte && <span className={styles.filterBadge}>Min Severity: {filters.severity_gte}</span>}
      {filters.status && <span className={styles.filterBadge}>Status: {filters.status}</span>}
    </div>
  );
}

interface FindingRowProps {
  finding: IACFinding;
}

function FindingRow({ finding }: FindingRowProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const isFail = finding.status === 'fail';

  return (
    <>
      <tr className={styles.findingRow} onClick={() => setIsExpanded(!isExpanded)}>
        <td className={styles.expanderCell}>
          <span className={styles.expander}>{isExpanded ? '▼' : '▶'}</span>
        </td>
        <td>
          <a
            href={finding.primary_url}
            target="_blank"
            rel="noopener noreferrer"
            onClick={(e) => e.stopPropagation()}
            className={styles.checkId}
          >
            {finding.check_id}
          </a>
        </td>
        <td className={styles.severityCell}>
          <span className={`${styles.severityBadge} ${styles[normalizeSeverity(finding.severity)]}`}>
            {finding.severity}
          </span>
        </td>
        <td className={styles.statusCell}>
          <span className={isFail ? styles.statusFail : styles.statusPass}>
            {finding.status.toUpperCase()}
          </span>
        </td>
        <td>
          <div className={styles.targetPath}>{finding.target}</div>
        </td>
        <td>
          <div className={styles.findingTitle}>{finding.title}</div>
        </td>
      </tr>
      {isExpanded && (
        <tr className={styles.expandedRow}>
          <td colSpan={6}>
            <div className={styles.expandedContent}>
              <div className={styles.expandedTwoColumn}>
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Description</h4>
                  <div className={styles.description}>{finding.description}</div>
                </div>
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Resolution</h4>
                  <div className={styles.resolution}>{finding.resolution}</div>
                </div>
              </div>

              {finding.message && finding.message !== finding.title && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Message</h4>
                  <div className={styles.description}>{finding.message}</div>
                </div>
              )}

              {finding.references && finding.references.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>References</h4>
                  <div className={styles.referenceList}>
                    {finding.references.map((ref: string, idx: number) => (
                      <a
                        key={idx}
                        href={ref}
                        target="_blank"
                        rel="noopener noreferrer"
                        className={styles.referenceLink}
                      >
                        {ref}
                      </a>
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
    <LocalIACResultsApp />
  </StrictMode>,
);
