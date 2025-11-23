package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    when YOU wIsH  uPON a STAR   ",
			expected: []string{"when", "you", "wish", "upon", "a", "star"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: Length of actual slice does not match expected.  Length of actual: %v, Length of expected :%v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("FAIL: Word %v is not the same as %v", word, expectedWord)
			}
		}
	}

}
