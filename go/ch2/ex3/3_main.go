package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var pc [256]byte
var stdout io.Writer = os.Stdout

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	n := 0
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func main() {
	n, _ := strconv.ParseUint(os.Args[1], 0, 64)
	fmt.Fprintf(stdout, "result %v\n", PopCountByLoop(n))
}
