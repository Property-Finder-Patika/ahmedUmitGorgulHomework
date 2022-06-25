package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCorner(t *testing.T) {
	var tests = []struct {
		i        int
		j        int
		expected bool
	}{
		{50, 50, false},
		{50, 49, true},
		{49, 50, true},
	}

	for _, test := range tests {
		_, _, result := corner(test.i, test.j)
		if result != test.expected {
			t.Errorf("i = %v, j = %v, result = %v", test.i, test.j, test.expected)
		}
	}
}

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer)
	main()
	got := stdout.(*bytes.Buffer).String()
	if strings.Contains(got, "nan") {
		t.Errorf("svg included nan.\n%s", got)
	}
}
