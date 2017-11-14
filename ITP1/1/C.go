package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := strings.Fields(stdin.Text())
  a, _ := strconv.Atoi(text[0])
  b, _ := strconv.Atoi(text[1])
  fmt.Println(a*b, 2*a+2*b)
}
