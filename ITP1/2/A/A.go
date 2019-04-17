package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  nums := strings.Fields(stdin.Text())
  a, _ := strconv.Atoi(nums[0])
  b, _ := strconv.Atoi(nums[1])

  if a < b {
    fmt.Println("a < b")
  } else if a == b {
    fmt.Println("a == b")
  } else {
    fmt.Println("a > b")
  }
}
