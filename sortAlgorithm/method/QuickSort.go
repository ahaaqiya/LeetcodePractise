package method

func QuickSort(nums []int) {
	QuickSortHelp(nums,0,len(nums)-1)
}

func QuickSortHelp(nums []int,left,right int)  {
	if left<right{
		index := QuickSortWaKeng(nums,left,right)
		QuickSortHelp(nums,left,index-1)
		QuickSortHelp(nums,index+1,right)
	}
}
//挖坑思想
func QuickSortWaKeng(nums []int,left,right int) int {
	fin := nums[left]
	for left<right{
		for left<right && nums[right]>=fin{
			right--
		}
		if left<right {
			nums[left] = nums[right]
		}
		for left<right && nums[left]<=fin{
			left++
		}
		if left<right{
			nums[right]=nums[left]
		}
	}
	nums[left] = fin
	return left
}