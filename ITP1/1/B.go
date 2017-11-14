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
  fmt.Println(num*num*num)
}
