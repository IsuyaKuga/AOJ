package main

import (
	. "fmt"
)

func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	if c < b {
		b, c = c, b
	}
	if b < a {
		a, b = b, a
	}
	if c < b {
		b, c = c, b
	}
	Println(a, b, c)
}
