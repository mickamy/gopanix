package panics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFirst(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		input  []string
		from   int
		output []string
		next   int
	}{
		{
			name: "no panic",
			input: []string{
				"line 1",
				"line 2",
				"line 3",
			},
			from:   0,
			output: nil,
			next:   3,
		},
		{
			name: "single panic",
			input: []string{
				"line 1",
				"line 2",
				"panic: something went wrong",
				"line 4",
				"line 5",
			},
			from: 0,
			output: []string{
				"panic: something went wrong",
				"line 4",
				"line 5",
			},
			next: 5,
		},
		{
			name: "multiple panics",
			input: []string{
				"line 1",
				"line 2",
				"panic: first panic",
				"line 4",
				"line 5",
				"panic: second panic",
				"line 7",
				"line 8",
			},
			from: 0,
			output: []string{
				"panic: first panic",
				"line 4",
				"line 5",
			},
			next: 5,
		},
		{
			name: "panic at the end",
			input: []string{
				"line 1",
				"line 2",
				"panic: something went wrong",
				"line 4",
				"line 5",
				"line 6",
			},
			from: 0,
			output: []string{
				"panic: something went wrong",
				"line 4",
				"line 5",
				"line 6",
			},
			next: 6,
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			output, next := ExtractFirst(tc.input, tc.from)
			assert.Len(t, output, len(tc.output), "expected output to be the same length")
			for i := range output {
				assert.Equal(t, tc.output[i], output[i], "expected output to be the same")
			}
			assert.Equal(t, tc.next, next, "expected next to be the same")
		})
	}
}

func TestExtract(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		input  string
		output [][]string
	}{
		{
			name:   "empty input",
			input:  "",
			output: nil,
		},
		{
			name: "no panic",
			input: `line 1
line 2
line 3`,
			output: nil,
		},
		{
			name: "single panic",
			input: `line 1
line 2
panic: something went wrong
line 4
line 5`,
			output: [][]string{
				{
					"panic: something went wrong",
					"line 4",
					"line 5",
				},
			},
		},
		{
			name: "multiple panics",
			input: `line 1
line 2
panic: first panic
line 4
line 5
panic: second panic
line 7
line 8`,
			output: [][]string{
				{
					"panic: first panic",
					"line 4",
					"line 5",
				},
				{
					"panic: second panic",
					"line 7",
					"line 8",
				},
			},
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			output := Extract(tc.input)
			assert.Len(t, output, len(tc.output), "expected output to be the same length")
			for i := range output {
				assert.Len(t, output[i], len(tc.output[i]), "expected output to be the same length")
				for j := range output[i] {
					assert.Equal(t, tc.output[i][j], output[i][j], "expected output to be the same")
				}
			}
		})
	}
}
