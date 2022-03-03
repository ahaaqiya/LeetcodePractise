package method

func ShellSort(nums []int)  {
	interval := len(nums)
	for interval>1{
		interval = interval/2
		for index:=0;index<interval;index++{
			for i:=index+interval;i<len(nums);i = i+interval{
				tmp := nums[i]
				var j int
				for j=i-interval; j>=index; j = j-interval {
					if tmp < nums[j]{
						nums[j+interval] = nums[j]
						continue
					}
					break
				}
				nums[j+interval] = tmp
			}
		}
	}
}