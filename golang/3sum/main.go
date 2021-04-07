package main

import (
	"fmt"
	"sort"
	"reflect"
)

func threeSum(nums []int) [][]int {
    var temp, rs [][]int
    l := len(nums)
    if l < 0 ||l >= 3000{
        return rs
    }
    for i := 0; i <l-2; i++{     
        for j := i+1; j <l-1; j++{
            for k := j+1; k<l; k++ {
                if nums[i] + nums[j] + nums[k]== 0 {
                    a := []int{nums[i], nums[j], nums[k]}
                    sort.Ints(a)                    
                    temp = append(temp, a)                    
                }
            }
        }
    }
    for _, item := range(temp){
        if !containArray(rs, item){
        rs = append(rs, item)
        }
    }
    return rs
}
func containArray(s [][]int, e []int) bool {
    for _, a := range s {        
        if reflect.DeepEqual(a,e) {
            return true
        }
    }
    return false
}

func main() {
	fmt.Println(threeSum([]int{}))
}
