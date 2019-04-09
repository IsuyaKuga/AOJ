package main
import (
"bufio"
"fmt"
"os"
"strconv"
)
var sc = bufio.NewScanner(os.Stdin)
func nextInt() int {
sc.Scan()
i, e := strconv.Atoi(sc.Text())
if e != nil {
panic(e)
}
return i
}
func main() {
sc.Split(bufio.ScanWords)
x := nextInt()
h := x/3600
m := x%3600/60
s := x%3600%60
fmt.Printf("%d:%d:%d\n",h,m,s)
}
