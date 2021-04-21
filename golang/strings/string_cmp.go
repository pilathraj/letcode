package main

import (
"fmt"
 "strings"
)

func main(){
  s := "Pilathraj"
  if s == "pilathraj"{
    fmt.Println("Equal")
  }
  if strings.EqualFold(s, "Pilathraj") {
    fmt.Println("Equal fold")
  }
}
