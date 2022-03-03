package method

func BubbleSort(nums []int)  {
	for i:=0;i<len(nums);i++{
		flag := true
		for j:=0;j<len(nums)-1-i;j++{
			if nums[j]>nums[j+1]{
				flag=false
				swap(j,j+1,nums)
			}
		}
		if flag==true{
			return
		}
	}
}

func swap(a,b int,nums []int)  {
	tmp := nums[a];
	nums[a] = nums[b]
	nums[b] = tmp
}