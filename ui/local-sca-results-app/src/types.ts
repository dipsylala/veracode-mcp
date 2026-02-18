// TypeScript types matching the Go SCA results structure

export interface SCAResultsResponse {
  application: SCAApplication;
  summary: SCASummary;
  pagination: SCAPagination;
  filters?: SCAFilters;
  findings: SCAFinding[];
}

export interface SCAApplication {
  name: string;
  path: string;
}

export interface SCASummary {
  total_matches: number;
  unique_vulnerabilities: number;
  vulnerable_components: number;
  by_severity: {
    critical: number;
    high: number;
    medium: number;
    low: number;
  };
  epss_data_available: number;
}

export interface SCAPagination {
  current_page: number;
  page_size: number;
  total_elements: number;
  total_pages: number;
  has_next: boolean;
  has_previous: boolean;
}

export interface SCAFilters {
  cve?: string;
  component_name?: string;
  severity_gte?: string;
}

export interface SCAFinding {
  vulnerability_id: string;
  severity: string;
  description: string;
  risk_score: number;
  data_source: string;
  component: SCAComponent;
  cvss?: SCACVSS[];
  cwes?: string[];
  epss?: SCAEPSS;
  fix: SCAFix;
  related_cves?: string[];
}

export interface SCAComponent {
  name: string;
  version: string;
  type: string;
  language: string;
  purl: string;
  licenses: string[];
  locations: SCALocation[];
}

export interface SCALocation {
  path: string;
  accessPath: string;
}

export interface SCACVSS {
  version: string;
  vector: string;
  base_score: number;
  exploitability_score: number;
  impact_score: number;
}

export interface SCAEPSS {
  cve: string;
  score: number;
  percentile: number;
  date: string;
}

export interface SCAFix {
  state: string;
  versions: string[];
  recommended_version: string;
}
