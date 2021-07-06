package main

import (
	"fmt"
	"math"
)

func main() {	
	n :=1011 
	v :=0
	i :=0
	for n >0 {
	  l := n%10;
	  v += int(math.Pow(2.0, float64(i)))*l
	  n = n/10
	  i++
	}
	fmt.Println(v)
}
