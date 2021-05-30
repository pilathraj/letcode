package main

import (
	"fmt"
	"sort"
	"math"
)

func findClosest(a []int, x int) int{
  closestSum := int(^uint(0) >> 1)
  n := len(a)
  sort.Ints(a)
  for i:= 0; i< n-2; i++{
    next := i+1
    last := n -1
    for next < last {
      sum := a[i]+a[next]+a[last]
      if math.Abs(float64(x-sum)) < math.Abs(float64(x-closestSum)){
	closestSum = sum
      }
      if x >sum {
   	last -=1
      }else{
        next +=1
      }
    }
  }
   return closestSum 
}

func main() {	
	fmt.Println(findClosest([]int{-1 , 2, 1, -4},1)) //TC O(N^2)
	
}
