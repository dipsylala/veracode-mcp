package server

import (
	"encoding/json"
	"log"
)

// UI capability detection and client analysis functionality.
// This determines whether the client supports MCP Apps UI features.

// detectUICapability analyzes client capabilities to determine UI support.
// It looks for the MCP Apps UI extension and the specific MIME type support.
func (s *MCPServer) detectUICapability(capabilities ClientCapabilities) bool {
	// Check for force override first
	if s.forceMCPApp {
		log.Printf("\n=== FORCE MCP APP MODE ENABLED ===")
		log.Printf("✓✓✓ Forcing MCP Apps UI mode via --force-mcp-app flag ✓✓✓")
		log.Printf("   structuredContent will always be sent regardless of client capabilities")
		log.Printf("=== END UI SUPPORT CHECK ===\n")
		return true
	}

	log.Printf("\n=== CHECKING FOR UI SUPPORT ===")
	log.Printf("Extensions field present: %v", capabilities.Extensions != nil)

	if extensions := capabilities.Extensions; extensions != nil {
		extensionsJSON, _ := json.MarshalIndent(extensions, "", "  ")
		log.Printf("Extensions object:\n%s", string(extensionsJSON))

		// Check all keys in extensions for debugging
		log.Printf("Extension keys found:")
		for key := range extensions {
			log.Printf("  - %s", key)
		}

		// Look for the MCP Apps UI extension
		if uiExt, ok := extensions[UIExtensionKey].(map[string]interface{}); ok {
			log.Printf("✓ Found '%s' extension", UIExtensionKey)
			uiExtJSON, _ := json.MarshalIndent(uiExt, "", "  ")
			log.Printf("UI Extension content:\n%s", string(uiExtJSON))

			// Check for supported MIME types
			if mimeTypes, ok := uiExt["mimeTypes"].([]interface{}); ok {
				log.Printf("MimeTypes found: %v", mimeTypes)
				for i, mt := range mimeTypes {
					log.Printf("  [%d] %v (type: %T)", i, mt, mt)
					if mtStr, ok := mt.(string); ok && mtStr == UICapabilityMimeType {
						log.Printf("✓✓✓ Client supports MCP Apps UI! ✓✓✓")
						log.Printf("✓ UI mode enabled - will send brief summaries to LLM")
						log.Printf("=== END UI SUPPORT CHECK ===\n")
						return true
					}
				}
			} else {
				log.Printf("⚠ 'mimeTypes' field not found or not an array in UI extension")
			}
		} else {
			log.Printf("⚠ '%s' key not found in extensions", UIExtensionKey)
		}
	} else {
		log.Printf("⚠ No 'extensions' field in client capabilities")
	}

	log.Printf("\n❌ Client does NOT support MCP Apps UI (will use text-only mode)")
	log.Printf("   This means full JSON will be sent to the LLM instead of brief summaries")
	log.Printf("=== END UI SUPPORT CHECK ===\n")
	return false
}
