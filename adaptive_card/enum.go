package adaptive_card

import "strings"

var (
	spacingValues             = []string{"default", "none", "small", "medium", "large", "extraLarge", "padding"}
	horizontalAlignmentValues = []string{"left", "center", "right"}
	columnWidthValues         = []string{"auto", "stretch"}
	containerStyleValues      = []string{"default", "emphasis"}
	imageStyleValues          = []string{"default", "person"}
	imageSizeValues           = []string{"auto", "stretch", "small", "medium", "large"}
	textSizeValues            = []string{"default", "small", "medium", "large", "extraLarge"}
	textWeightValues          = []string{"default", "lighter", "bolder"}
	textColorValues           = []string{"default", "dark", "light", "accent", "good", "warning", "attention"}
)

// normalize returns the value in allowed that matches v case-insensitively, or "" if none matches.
func normalize(v string, allowed []string) string {
	for _, a := range allowed {
		if strings.EqualFold(v, a) {
			return a
		}
	}
	return ""
}
