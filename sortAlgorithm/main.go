package main

import (
	"fmt"
	"sortAlgorithm/method"
)

func main()  {
	nums := []int{5,1,3,6,4,7,5,8}
	method.HeapSort(nums)
	fmt.Println(nums)
}
