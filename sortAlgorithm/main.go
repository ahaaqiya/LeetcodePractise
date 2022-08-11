package main

import (
	"fmt"
	"sortAlgorithm/method"
)

func main()  {
	nums := []int{5,1,3,6,4,7,5,8}
	method.MergeSort(nums)
	fmt.Println(nums)
}

func sortArrayByParity(nums []int) []int {
	i,j := 0,len(nums)-1
	for i<=j{
		if nums[i]%2==0 && nums[j]%2==1{
			i++
			j--
		}else if nums[i]%2==0 && nums[j]%2==0{
			i++
		}else if nums[i]%2==1 && nums[j]%2==0{
			nums[i],nums[j] = nums[j],nums[i]
		}else {
			j--
		}
	}
	return nums
}