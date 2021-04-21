package main

import "fmt"

func changeSlice(s []int){
  s[2] = 22
}
func changeArray(a [3]int){
  a[2] =22
}

func main(){
  s := []int{1,2,3}
  a := [3]int{1,2,3}
  fmt.Printf("Slice: %v, Array: %v\n", s, a)
  changeSlice(s)
  changeArray(a)
  fmt.Printf("Slice: %v, Array: %v\n", s, a)  
}
