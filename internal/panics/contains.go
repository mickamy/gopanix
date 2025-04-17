package panics

import (
	"strings"
)

func Contains(s string) bool {
	return strings.Contains(s, "panic:")
}
