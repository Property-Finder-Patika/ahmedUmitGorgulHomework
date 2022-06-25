package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer)
	main()
	got := stdout.(*bytes.Buffer).String()
	if strings.Contains(got, "NaN") {
		t.Errorf("SVG included NaN value.\n%s", got)
	}
}
