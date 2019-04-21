package main

import (
	. "fmt"
	"sort"
)

func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	s := []int{a, b, c}
	sort.Ints(s)
	Println(s[0], s[1], s[2])
}
