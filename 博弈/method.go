package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index,value := range nums{
		if key,have := m[target-value];have{
			return []int{index,key}
		}else{
			m[value] = index
		}
	}
	return nil
}
