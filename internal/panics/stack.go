package panics

import (
	"strings"
)

func ExtractFirst(lines []string, from int) ([]string, int) {
	var stack []string
	var found bool
	for i := from; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "panic:") {
			if found {
				return stack, i
			}
			found = true
		}
		if found {
			stack = append(stack, line)
		}
	}
	return stack, len(lines)
}

func Extract(s string) [][]string {
	lines := strings.Split(s, "\n")
	var panics [][]string

	cursor := 0
	for cursor < len(lines) {
		stack, next := ExtractFirst(lines, cursor)
		if len(stack) > 0 {
			panics = append(panics, stack)
		}
		cursor = next
	}
	return panics
}
