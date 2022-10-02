package action

import (
	"strings"
	"testing"
)

func TestExtensionFields(t *testing.T) {
	tests := map[string]string{
		"json+one-two+":      "json|+one|-two|+",
		"":                   "",
		"json":               "json",
		"json-":              "json|-",
		"json--+":            "json|-|-|+",
		"json-one-two+three": "json|-one|-two|+three",
	}
	for test, expected := range tests {
		t.Run(test, func(t *testing.T) {
			if actual := strings.Join(extensionFields(test), "|"); actual != expected {
				t.Errorf("expected: '%v' actual: '%v'", expected, actual)
			}
		})
	}
}
