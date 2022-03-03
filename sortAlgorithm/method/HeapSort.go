package method

func HeapSort(nums []int) {
	length := len(nums)
	arr := make([]int,length+1)
	for i:=0;i<length;i++{
		arr[i+1] = nums[i]
	}
	//下沉建堆
	for i:=length/2;i>=1;i--{
		sink(arr,i,length)
	}
	index := length
	for index>1{
		swap(index,1,arr)
		index--
		sink(arr,1,index)
	}
	for i:=1;i<=length;i++{
		nums[i-1] = arr[i]
	}
}

//大根堆
func sink(nums []int,index int,len int)  {
	for index*2 <= len{
		j := 2*index
		if j+1<=len && nums[j+1]>nums[j]{
			j++
		}
		if nums[j]>nums[index]{
			swap(j,index,nums)
		}else{
			break
		}
		index = j
	}
}
