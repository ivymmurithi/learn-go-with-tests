package iteration

import (
	"strings"
)

func Repeat(letter string, count int) string {
	var repeated strings.Builder

	for i := 0; i < count; i++ {
		repeated.WriteString(letter)
		
	}
	return repeated.String()
}