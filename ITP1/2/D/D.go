package main

import (
	. "fmt"
)

func main() {
	var W, H, x, y, r int
	Scan(&W, &H, &x, &y, &r)
	if (0 <= x-r) && (x+r <= W) && (0 <= y-r) && (y+r <= H) {
		Println("Yes")
	} else {
		Println("No")
	}
}
