package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	a, b, c := nextInt(), nextInt(), nextInt()

	if a < b && b < c {
		Println("Yes")
	} else {
		Println("No")
	}
}
