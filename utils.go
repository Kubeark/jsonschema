package jsonschema

import (
	"regexp"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts the provided string into snake case using dashes.
// This is useful for Schema IDs and definitions to be coherent with
// common JSON Schema examples.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

func splitIgnoringEscapedCommas(input string) []string {
	var result []string
	var current strings.Builder
	escaped := false

	for _, char := range input {
		if char == '\\' {
			escaped = !escaped
			current.WriteRune(char) // Keep the backslash for now
			continue
		}

		if char == ',' && !escaped {
			// Split at non-escaped comma
			result = append(result, strings.ReplaceAll(current.String(), `\,`, `,`))
			current.Reset()
		} else {
			current.WriteRune(char)
		}

		// Reset escaped flag
		escaped = false
	}

	// Add the final segment
	result = append(result, strings.ReplaceAll(current.String(), `\,`, `,`))
	return result
}

// NewProperties is a helper method to instantiate a new properties ordered
// map.
func NewProperties() *orderedmap.OrderedMap[string, *Schema] {
	return orderedmap.New[string, *Schema]()
}
