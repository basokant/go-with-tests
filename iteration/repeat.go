package iteration

import "strings"

// Repeat returns a strng that repeats char repeatCount times.
func Repeat(char string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
