package main

import (
    "fmt"
)

func sum(nums ...int){
    fmt.Printf("%v\n", nums)
}

func main(){
    sum(1, 2)
    nums := []int{1,2,3,4,5,6}
    sum(nums...)
    nums2 := [5]int{1,2,3,4,5}
    nums3 := nums2[:3]
    nums3[0] = 7
    nums3 = append(nums3, 1000)
    fmt.Printf("%v\n", nums2)
}
