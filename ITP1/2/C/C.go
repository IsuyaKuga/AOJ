package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
)

func swap(x, y int) (int, int) {
  return y, x
}

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  nums := strings.Fields(stdin.Text())
  a, _ := strconv.Atoi(nums[0])
  b, _ := strconv.Atoi(nums[1])
  c, _ := strconv.Atoi(nums[2])

  if c < b {
    b, c = swap(b, c)
  }
  if b < a {
    a, b = swap(a, b)
  }
  if c < b {
    b, c = swap(b, c)
  }

  fmt.Println(a,b,c)
}
