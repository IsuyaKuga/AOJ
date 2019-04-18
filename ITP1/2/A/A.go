package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b := nextInt(), nextInt()
	if a == b {
		Printf("a == b\n")
	} else if a < b {
		Printf("a < b\n")
	} else {
		Printf("a > b\n")
	}
}
