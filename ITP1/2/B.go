package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  nums := strings.Fields(stdin.Text())
  a, _ := strconv.Atoi(nums[0])
  b, _ := strconv.Atoi(nums[1])
  c, _ := strconv.Atoi(nums[2])
  fmt.Println(a,b,c)

  if a < b && b < c {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
