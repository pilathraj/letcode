package main

import "fmt"

func reverse(str string) string {
  var result string
  for _, v:= range(str) {
    result = v + result    
  }
  return result
}

func main() {
  s := []int{1,2,3}
  s1 := "Pilathraj"
  for i, j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
  s2 :=[]rune(s1)
  for i, j := 0, len(s2)-1; i<j; i,j = i+1, j-1 {
   s2[i], s2[j] = s2[j], s2[i]
  }
  fmt.Println(s)
  fmt.Println(string(s2))
  fmt.Println(reverse(s1))
}
