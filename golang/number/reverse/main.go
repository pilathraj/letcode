package main

import (
	"fmt"
)

func main() {
	rev := 0
	n := 132134
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	fmt.Println(rev)
}
