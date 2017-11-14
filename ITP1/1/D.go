package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()
  num, _ := strconv.Atoi(text)
  h := num / 3600
  m := num / 60 % 60
  s := num - (h * 3600 + m * 60)
  fmt.Printf("%d:%d:%d\n", h, m, s)
}
