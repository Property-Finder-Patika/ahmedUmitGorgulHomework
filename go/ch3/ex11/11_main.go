package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var stdout io.Writer = os.Stdout // modified during testing

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
	if strings.ContainsAny(s, "+-") {
		buf.WriteByte(s[0])
		s = s[1:]
		n--
	}
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	if d := strings.Index(s, "."); d != -1 {
		n = d
	}

	for e := i + 3; e < n; {
		buf.WriteString("," + s[i:e])
		i, e = e, e+3
	}

	if len(s[i:]) > 0 {
		buf.WriteString("," + s[i:])
	}
	return buf.String()
}
