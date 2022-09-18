package fns

import (
	"regexp"
	"strings"
	"unicode"
)

var reSnakeCase = regexp.MustCompile(`[^a-z0-9]+`)

func SnakeCase(s string) string {
	lowering := false
	combining := false
	var b strings.Builder
	for i, c := range s {
		if unicode.IsNumber(c) || unicode.IsLower(c) {
			b.WriteRune(c)
			lowering = false
			combining = false
			continue
		} else if unicode.IsUpper(c) {
			if !lowering && i > 0 {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(c))
			lowering = true
			combining = false
		} else {
			if !combining {
				b.WriteRune('_')
				combining = true
			}
		}
	}
	return b.String()
}
