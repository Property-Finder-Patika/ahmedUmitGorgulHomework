package main

import (
	"bytes"
	"os"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		args     []string
		expected string
	}{
		{[]string{"isanagram", "foo", "bar"}, "foo and bar anagram degil\n"},
		{[]string{"isanagram", "foo", "foo"}, "foo and foo anagram degil\n"},
		{[]string{"isanagram", "asdf", "fdsa"}, "asdf and fdsa anagram\n"},
		{[]string{"isanagram", "asdf", "fdsa", "fdsa"}, "uygunsuz be\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		stdout = new(bytes.Buffer)
		main()
		got := stdout.(*bytes.Buffer).String()
		if got != test.expected {
			t.Errorf("Result = %q, Expected %q", got, test.expected)
		}
	}
}
