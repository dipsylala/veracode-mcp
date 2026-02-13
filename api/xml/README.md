# Veracode XML API Client

This package provides a Go client for the Veracode XML API endpoints.

## Features

- HMAC authentication
- XML request/response handling
- Type-safe API models

## Endpoints

### Get Mitigation Info

Retrieves mitigation information for specific flaws in a build.

**Endpoint:** `getmitigationinfo.do`

**Parameters:**

- `build_id` - The build ID to query
- `flaw_id_list` - Comma-delimited list of flaw IDs

**Example:**

```go
import (
    "context"
    "github.com/dipsylala/veracodemcp-go/api/xml"
)

client, err := xml.NewClient()
if err != nil {
    // handle error
}

ctx := context.Background()
buildID := int64(53691725)
flawIDs := []int64{50}

mitigationInfo, err := client.GetMitigationInfo(ctx, buildID, flawIDs)
if err != nil {
    // handle error
}

for _, issue := range mitigationInfo.Issues {
    fmt.Printf("Flaw %d: %s\n", issue.FlawID, issue.Category)
    for _, action := range issue.MitigationActions {
        fmt.Printf("  Action: %s by %s on %s\n", action.Action, action.Reviewer, action.Date)
        if action.Comment != "" {
            fmt.Printf("  Comment: %s\n", action.Comment)
        }
    }
}
```

## Response Structure

The XML response is unmarshaled into the following types:

- `MitigationInfo` - Root element containing build info, issues, and errors
- `Issue` - A flaw with its mitigation actions
- `MitigationAction` - A single mitigation action (e.g., accepted, rejected, false positive)
- `Error` - Errors for flaws that could not be processed

## Mitigation Action Types

The following action types are supported:

- `comment` - Comment on the flaw
- `fp` - False Positive
- `library` - Library issue
- `acceptrisk` - Accept Risk
- `appdesign` - Application Design
- `osenv` - Operating System Environment
- `netenv` - Network Environment
- `rejected` - Mitigation Rejected
- `accepted` - Mitigation Accepted
- `remediated` - Flaw Remediated
- `noactiontaken` - No Action Taken
- `conforms` - Conforms to Policy
- `deviates` - Deviates from Policy
- `defer` - Deferred

## Schema

The XML response follows the schema defined at:
https://analysiscenter.veracode.com/resource/mitigationinfo.xsd
