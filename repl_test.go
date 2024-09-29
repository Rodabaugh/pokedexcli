package main

import (
	"testing"
)

func TestCleanUserInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input: " Hello  World ",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HeLlo world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		result := cleanUserInput(cs.input)
		if len(result) != len(cs.expected) {
			t.Errorf("Lengths are not equal: %v vs %v", len(result), len(cs.expected))
			continue
		}
		for i := range result {
			resultWord := result[i]
			expectedWord := cs.expected[i]
			if resultWord != expectedWord {
				t.Errorf("Words don't match: %v != %v", expectedWord, resultWord)
			}
		}
	}
}
