# Shared UI Styles

This directory contains shared CSS modules used across all Veracode MCP UI applications.

## Files

### `base.module.css`

Core styles shared by all UI apps. Includes:

- **Layout**: Container, header, h1 styling
- **Pagination & Filters**: Pagination info box, filters display with badges
- **Summary Section**: Summary items, labels, values
- **Severity Breakdown**: Severity count display for both SAST/DAST (6-level) and SCA (4-level) scales
- **Table Structure**: Table container, headers, cells, hover effects
- **Expandable Rows**: Expander icons, expanded row styling, content sections
- **Severity Badges**: Badge styling for both SAST/DAST and SCA severity scales
- **Policy & Status Icons**: Policy violations, status indicators, mitigation icons
- **Mitigation Display**: Mitigation comment boxes
- **File Paths**: Monospace text for file paths
- **State Displays**: Error, loading, and empty state styles

## Usage

Each app imports the shared base and adds only app-specific styles:

```css
/* Import shared base styles */
@import '../../shared/base.module.css';

/* App-specific styles below */
.myAppSpecificClass {
  /* ... */
}
```

## Severity Scales

### SAST/DAST (6-level scale)

- Very High (.veryhigh) - #fee background, #c00 text
- High (.high) - #fed background, #c50 text
- Medium (.medium) - #ffd background, #960 text
- Low (.low) - #ffe background, #660 text
- Very Low (.verylow) - #ffe background, #880 text
- Informational (.info) - #eef background, #66f text

### SCA (4-level scale)

- Critical (.critical) - #fdd background, #900 text
- High (.high) - #fed background, #c50 text
- Medium (.medium) - #ffd background, #960 text
- Low (.low) - #ffe background, #660 text

## Standardization Decisions

The following were standardized across all apps:

1. **Table cell padding**: Standardized to 8px (was 4px in some apps)
2. **Pagination info color**: Standardized to #666 (was #000 in some apps)
3. **Severity badge min-width**: Standardized to 80px (was 75px in some apps)
4. **Expanded row td padding**: Standardized to 16px (was 4px in some apps)

## App-Specific Styles

Each app maintains only its unique styling needs:

### Static Findings App

- moduleCell, fileCell styling

### Dynamic Findings App

- url styling for URL display

### Pipeline Results App

- No app-specific styles (uses shared base entirely)

### Local SCA Results App

- CVE links (cveLink)
- Component display (componentName, version)
- EPSS column (epssHeader, epssCell, epssScore)
- Fix status (fixAvailable, noFix)
- Two-column expanded layout (expandedTwoColumn)
- Details list (detailsList, detailItem, detailLabel, detailValue)
- CVSS display (cvssContainer, cvssItem, cvssVersion, cvssMetrics, etc.)
- EPSS detailed display (epssDetails, epssMetricLarge, etc.)
- CWE links (cweList, cweLink)
- Related CVE links (relatedCveList, relatedCveLink)
- Location display (locationList, locationItem, locationPath, locationAccess)

## Benefits

- **Consistency**: All apps share the same base styling
- **Maintainability**: Update shared styles in one place
- **Smaller files**: App-specific CSS files are much smaller
- **Easier onboarding**: New developers can understand app-specific customizations quickly
- **Standards enforcement**: Shared base ensures consistent UX patterns
