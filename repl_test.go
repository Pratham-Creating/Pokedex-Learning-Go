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
			input:    "Hello World",
			expected: []string{"Hello", "World"},
		},
		{
			input:    "This is my job",
			expected: []string{"This", "is", "my", "job"},
		},
	}

	for _, c := range cases {
		original := CleanInput(c.input)

		for i := range original {

			if c.expected[i] != original[i] {
				t.Errorf("Test Failed:- String is not same as the Slice")
				t.Fail()

			}
		}
	}
}
