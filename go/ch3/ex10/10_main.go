package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Fprintf(stdout, "  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	for e := i + 3; e < n; {
		buf.WriteString("," + s[i:e])
		i, e = e, e+3
	}
	buf.WriteString("," + s[i:])
	return buf.String()
}
