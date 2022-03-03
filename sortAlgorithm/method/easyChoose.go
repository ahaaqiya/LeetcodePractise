package method

func EasyChooseSort(nums []int)  {
	var min int
	for i:=0;i<len(nums);i++{
		min = i
		for j:=i+1;j<len(nums);j++{
			if nums[min]>nums[j]{
				min = j
			}
		}
		if min != i{
			swap(i,min,nums)
		}
	}
}