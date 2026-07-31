package forticlient

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// MultValue describes the nested structure in the results
type SortInputModel struct {
	SortBy        string            `json:"sortby"` // mkey may be string or int
	SortDirection string            `json:"sortdirection"`
	URL           string            `json:"url"`
	ManualOrder   []interface{}     `json:"manual_order"`
	WSParams      map[string]string `json:"ws_params"` // workspace parameters
	URLParams     map[string]string `json:"url_params"`
}

func (input_model *SortInputModel) update() {
	// todo: need update
	if !strings.ContainsAny(input_model.URL, "[ | {") {
		return
	}

	// Find all placeholder patterns {.*?}
	re := regexp.MustCompile(`{.*?}`)
	placeholders := re.FindAllString(input_model.URL, -1)
	placeholderCount := len(placeholders)

	if placeholderCount > 0 {
		// Multiple placeholders found, handle each one individually
		updatedURL := input_model.URL
		for _, placeholder := range placeholders {
			// Extract content between {} - remove the braces
			paramKey := placeholder[1 : len(placeholder)-1]

			// Try to find the parameter in URLParams
			if input_model.URLParams != nil {
				if value, exists := input_model.URLParams[paramKey]; exists {
					// Replace this specific placeholder with the found value
					updatedURL = strings.ReplaceAll(updatedURL, placeholder, fmt.Sprintf("%v", value))
				} else {
					// Parameter not found, log and keep original placeholder
					log.Printf("[WARNING] URL parameter '%s' (from placeholder '%s') not found in URLParams", paramKey, placeholder)
				}
			} else {
				log.Printf("[WARNING] URLParams is nil, cannot replace placeholder '%s'", placeholder)
			}
		}

		input_model.URL = updatedURL
	}
}
