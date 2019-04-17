package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
)

func main () {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  inputs := strings.Fields(stdin.Text())
  W, _ := strconv.Atoi(inputs[0])
  H, _ := strconv.Atoi(inputs[1])
  x, _ := strconv.Atoi(inputs[2])
  y, _ := strconv.Atoi(inputs[3])
  r, _ := strconv.Atoi(inputs[4])

  if (0 <= x-r) && (x+r <= W) && (0 <= y-r) && (y+r <= H) {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
