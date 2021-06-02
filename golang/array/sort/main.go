package main

import (
	"fmt"
)

/*func main() {
	fmt.Println("Hello, playground")
	a := [5]int{2,4,1,3,5}
	s :=0
	e := len(a)-1
	for s<e {   // O(N)
		  if a[s] > a[e]{
			  a[s], a[e] = a[e], a[s]
			  s +=1	  
		  } else {
		    e -= 1
		  }
	}
	fmt.Println(a)  //[1 2 4 3 5]	
}*/
func main() {
	a := []int{3, 4, 2, 1, 5}
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}

