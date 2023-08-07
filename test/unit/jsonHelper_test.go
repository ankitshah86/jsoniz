package unit_test

import (
	"testing"

	"github.com/ankitshah86/jsoniz/internal/helpers"
)

func TestValidateJson(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid json",
			input:    `{"foo": "bar"}`,
			expected: true,
		},
		{
			name:     "invalid json",
			input:    `{"foo": "bar"`,
			expected: false,
		},
		{
			name:     "empty json",
			input:    `{}`,
			expected: true,
		},
		{
			name:     "json with whitespace",
			input:    `   {   "foo"   :   "bar"   }   `,
			expected: true,
		},
		{
			name:     "json with comments",
			input:    `{"foo": "bar" /* comment */}`,
			expected: false,
		},
		{
			name: "valid nested json",
			input: `{
					"foo": "bar",
					"baz": {
						"qux": "quux",
						"corge": {
							"grault": "garply",
							"waldo": {
								"fred": "plugh",
								"xyzzy": "thud"
							}
						}
					}
				}`,
			expected: true,
		},
		{
			name: "invalid nested json",
			input: `{
					"foo": "bar",
					"baz": {
						"qux": "quux",
						"corge": {
							"grault": "garply",
							"waldo": {
								"fred": "plugh",
								"xyzzy": "thud"
							}
						}
					}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.ValidateJson(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateJson(%q) returned %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
