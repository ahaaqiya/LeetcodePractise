package main

import "fmt"

func main()  {
	nums := []int{10,5,2,6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums,k))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	left,right := 0,0
	sum := 1
	for right<len(nums){
		sum = sum * nums[right]
		right++
		for left<right && sum>=k{
			sum = sum/nums[left]
			left++
		}
		count += right-left
	}

	return count
}