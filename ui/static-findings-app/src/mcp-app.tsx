import { useApp } from "@modelcontextprotocol/ext-apps/react";
import { StrictMode, useState } from "react";
import { createRoot } from "react-dom/client";
import type { MCPFindingsResponse, MCPFinding, MCPMitigation } from "./types";
import styles from "./mcp-app.module.css";

function StaticFindingsApp() {
  const [resultsData, setResultsData] = useState<MCPFindingsResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const { app, error: appError } = useApp({
    appInfo: { name: "Static Findings", version: "1.0.0" },
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

      app.onhostcontextchanged = (params) => {
        console.info("Host context changed:", params);
      };

      app.onerror = (err) => {
        console.error("App error:", err);
        setError(err.message);
      };
    },
  });

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

  return <StaticFindingsView data={resultsData} />;
}

interface StaticFindingsViewProps {
  data: MCPFindingsResponse;
}

function StaticFindingsView({ data }: StaticFindingsViewProps) {
  const { application, summary, findings } = data;

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h1>Static Findings: {application.name}</h1>
        
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
          {Object.entries(summary.by_severity).map(([severity, count]) => {
            const numCount = count as number;
            return numCount > 0 && (
              <div key={severity} className={`${styles.severityItem} ${styles[severity]}`}>
                <strong>{numCount}</strong> {severity}
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
                <th className={styles.severityHeader}>Severity</th>
                <th>CWE</th>
                <th>File</th>
                <th>Status</th>
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
        <td className={styles.severityCell}>
          <span className={styles.expander}>{isExpanded ? '▼' : '▶'}</span>
          <span className={`${styles.severityBadge} ${styles[finding.severity.toLowerCase()]}`}>
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
          <div className={styles.filePath}>
            {finding.file_path || '-'}{finding.file_path && finding.line_number ? `:${finding.line_number}` : ''}
          </div>
        </td>
        <td>{finding.status}</td>
        <td>{finding.mitigation_status}</td>
      </tr>
      {isExpanded && (
        <tr className={styles.expandedRow}>
          <td colSpan={4}>
            <div className={styles.expandedContent}>
              <div className={styles.expandedSection}>
                <h4 className={styles.expandedHeader}>Description</h4>
                <div className={styles.description}>{finding.description}</div>
              </div>              {finding.mitigations && finding.mitigations.length > 0 && (
                <div className={styles.expandedSection}>
                  <h4 className={styles.expandedHeader}>Mitigations ({finding.mitigations.length})</h4>
                  {finding.mitigations.map((mitigation: MCPMitigation, idx: number) => (
                    <div key={idx} className={styles.mitigation}>
                      <div className={styles.mitigationHeader}>
                        <strong>{mitigation.action}</strong>
                        <span className={styles.mitigationMeta}>
                          by {mitigation.submitter} on {new Date(mitigation.date).toLocaleDateString()}
                        </span>
                      </div>
                      {mitigation.comment && (
                        <div className={styles.mitigationComment}>{mitigation.comment}</div>
                      )}
                    </div>
                  ))}
                </div>
              )}            </div>
          </td>
        </tr>
      )}
    </>
  );
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <StaticFindingsApp />
  </StrictMode>,
);
