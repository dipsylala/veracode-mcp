// TypeScript types matching the Go IaC results structure

export interface IACResultsResponse {
  application: IACApplication;
  summary: IACSummary;
  pagination: IACPagination;
  filters?: IACFilters;
  findings: IACFinding[];
}

export interface IACApplication {
  name: string;
  path: string;
}

export interface IACSummary {
  total_checks: number;
  fail: number;
  pass: number;
  unique_checks: number;
  unique_targets: number;
  by_severity: {
    critical: number;
    high: number;
    medium: number;
    low: number;
  };
}

export interface IACPagination {
  current_page: number;
  page_size: number;
  total_elements: number;
  total_pages: number;
  has_next: boolean;
  has_previous: boolean;
}

export interface IACFilters {
  target?: string;
  check_id?: string;
  severity_gte?: string;
  status?: string;
}

export interface IACFinding {
  check_id: string;
  title: string;
  description: string;
  message: string;
  severity: string;
  status: string;
  target: string;
  type: string;
  resolution: string;
  primary_url: string;
  references: string[];
  provider: string;
  policy_status?: string;
}
