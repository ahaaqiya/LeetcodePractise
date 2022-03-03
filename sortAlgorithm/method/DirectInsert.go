package method

func DirectInsert(nums []int)  {
	for i:=1;i<len(nums);i++{
		tmp := nums[i]
		var j int
		for j=i-1;j>=0;j--{
			if tmp < nums[j]{
				nums[j+1] = nums[j]
				continue
			}
			break
		}
		nums[j+1] = tmp
	}
}