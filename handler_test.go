package urlshort

import (
	"testing"
	"urlshort"
)

func TestParseYaml(t *testing.T) {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	pathUrls, err := urlshort.ParseYAML([]byte(yml))
	if err != nil {
		t.Error("Failed to parse YAML")
	}

	if len(pathUrls) != 2 {
		t.Error("Not 2 elements found")
	}
}
