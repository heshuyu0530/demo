package main

import (
	"fmt"

	"example.com/demo/leetcode"
)

func main(){
	nums := []int{1,1,1,1,1}
	fmt.Print(leetcode.FindTargetSumWays(nums,3))
}