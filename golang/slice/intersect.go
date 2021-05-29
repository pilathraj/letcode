package main

import (
	"fmt"
)
func  contains(a []int, find int) bool{
   for _, v := range(a){
    if v== find {
     return true
    }
   }
  return false;
}

func intersect(a []int, b []int) int{
  var c []int
  for _, v1 := range(a){
   for _, v2 := range(b){
     if v1==v2 && !contains(c, v2){
      c = append(c, v2)
     }
   }
  }
   
  return len(c)
}

func main() {
       a := []int{89, 24, 75, 11, 23}
       b := []int{89, 2, 4}
      fmt.Println(intersect(a, b))
	
}
