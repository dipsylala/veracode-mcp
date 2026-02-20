// TypeScript types matching the Go MCPFindingsResponse structure
export interface MCPFindingsResponse {
  application: MCPApplication;
  sandbox?: MCPSandbox;
  summary: MCPFindingsSummary;
  findings: MCPFinding[];
  pagination?: MCPPagination;
  policy_filter?: boolean;
}

export interface MCPApplication {
  name: string;
  id: string;
  business_criticality?: string;
}

export interface MCPSandbox {
  name: string;
  id: string;
}

export interface MCPFinding {
  flaw_id: string;
  scan_type: string;
  status: string;
  mitigation_status: string;
  violates_policy: boolean;
  severity: string;
  severity_score: number;
  cwe_id: number;
  description: string;
  references?: string[];
  file_path?: string;
  line_number?: number;
  url?: string;
  component?: MCPComponent;
  vulnerability?: MCPVulnerability;
  mitigations?: MCPMitigation[];
  first_found?: string;
  last_seen?: string;
  module?: string;
  procedure?: string;
  attack_vector?: string;
  context_type?: string;
  count?: number;
}

export interface MCPComponent {
  name: string;
  version: string;
  library?: string;
}

export interface MCPVulnerability {
  cve_id: string;
  cvss_score: number;
  exploitable: boolean;
}

export interface MCPMitigation {
  action: string;
  comment: string;
  submitter: string;
  date: string;
}

export interface MCPFindingsSummary {
  total_findings: number;
  open_findings: number;
  policy_violations: number;
  by_severity: Record<string, number>;
  by_status: Record<string, number>;
  by_mitigation_status: Record<string, number>;
}

export interface MCPPagination {
  current_page: number;
  total_pages: number;
  page_size: number;
  total_elements: number;
  has_next: boolean;
  has_previous: boolean;
}
