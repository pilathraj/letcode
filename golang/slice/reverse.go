package main

import "fmt"

func main() {
  s := []int{1,2,3}
  for i, j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
  fmt.Println(s)
}
