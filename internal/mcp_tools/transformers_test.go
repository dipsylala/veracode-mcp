package mcp_tools

import (
	"testing"
)

func TestExtractReferences(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantRefs    []Reference
	}{
		{
			name:        "Extract escaped HTML anchor tags",
			description: `Some text <span>References: <a href=\"https://cwe.mitre.org/data/definitions/117.html\">CWE</a> <a href=\"https://owasp.org/www-community/attacks/Log_Injection\">OWASP</a></span>`,
			wantRefs: []Reference{
				{Name: "CWE", URL: "https://cwe.mitre.org/data/definitions/117.html"},
				{Name: "OWASP", URL: "https://owasp.org/www-community/attacks/Log_Injection"},
			},
		},
		{
			name:        "Extract plain URLs as fallback",
			description: "Some text with URLs: https://cwe.mitre.org/data/definitions/117.html and https://owasp.org/",
			wantRefs: []Reference{
				{Name: "https://cwe.mitre.org/data/definitions/117.html", URL: "https://cwe.mitre.org/data/definitions/117.html"},
				{Name: "https://owasp.org/", URL: "https://owasp.org/"},
			},
		},
		{
			name:        "No references",
			description: "Just plain text with no URLs",
			wantRefs:    []Reference{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			references := ExtractReferences(tt.description)

			if len(references) != len(tt.wantRefs) {
				t.Errorf("ExtractReferences() got %d references, want %d", len(references), len(tt.wantRefs))
				return
			}

			for i, ref := range references {
				if ref.Name != tt.wantRefs[i].Name {
					t.Errorf("ExtractReferences() ref[%d].Name = %v, want %v", i, ref.Name, tt.wantRefs[i].Name)
				}
				if ref.URL != tt.wantRefs[i].URL {
					t.Errorf("ExtractReferences() ref[%d].URL = %v, want %v", i, ref.URL, tt.wantRefs[i].URL)
				}
			}
		})
	}
}

func TestCleanDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		want        string
	}{
		{
			name:        "Remove references section with anchor tags",
			description: `Some vulnerability description here. References: <a href=\"https://cwe.mitre.org/data/definitions/117.html\">CWE</a> <a href=\"https://owasp.org/\">OWASP</a>`,
			want:        "Some vulnerability description here.",
		},
		{
			name:        "Remove references section wrapped in span",
			description: `<span>This is a test. </span><span>References: <a href=\"https://cwe.mitre.org\">CWE</a></span>`,
			want:        "This is a test.",
		},
		{
			name:        "Remove HTML tags and clean whitespace",
			description: `<span>Multiple   spaces   here</span>`,
			want:        "Multiple spaces here",
		},
		{
			name:        "Handle mixed case 'references'",
			description: `Description text. REFERENCES: <a href=\"http://example.com\">Link</a>`,
			want:        "Description text.",
		},
		{
			name:        "Preserve description without references",
			description: `Just a normal description with no links`,
			want:        "Just a normal description with no links",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanDescription(tt.description)
			if result != tt.want {
				t.Errorf("CleanDescription() = %q, want %q", result, tt.want)
			}
		})
	}
}
