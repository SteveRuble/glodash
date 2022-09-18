package fns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{
			input:    "PascalCase",
			expected: "pascal_case",
		},
		{
			input:    "camelCase",
			expected: "camel_case",
		},
		{
			input:    "kebab-case",
			expected: "kebab_case",
		},
		{
			input:    "spaced out",
			expected: "spaced_out",
		},
		{
			input:    "a mess... of non&ense",
			expected: "a_mess_of_non_ense",
		},
	} {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, SnakeCase(tt.input))
		})
	}
}
