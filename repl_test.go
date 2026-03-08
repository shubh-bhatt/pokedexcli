package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length not equal")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if !reflect.DeepEqual(expectedWord, word) {
				t.Fatalf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
