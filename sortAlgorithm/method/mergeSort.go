package method

func MergeSort(nums []int)  {
	divid(nums,0,len(nums)-1)
}

func divid(nums []int,left,right int)  {
	if left<right{
		mid := left + (right-left)>>1
		divid(nums,left,mid)
		divid(nums,mid+1,right)
		merge(nums,left,right,mid)
	}
}
func merge(arr []int,left,right int,mid int)  {
	tmpArr := make([]int,right-left+1)
	i,j := left,mid+1
	index := 0
	for i <= mid && j <= right{
		if arr[i]<=arr[j]{
			tmpArr[index] = arr[i]
			i++
		}else{
			tmpArr[index] = arr[j]
			j++
		}
		index++
	}
	for i<=mid{
			tmpArr[index] = arr[i]
			i++
			index++
	}
	for j<=right{
		tmpArr[index] = arr[j]
		j++
		index++
	}
	index = 0
	for left<=right{
		arr[left] = tmpArr[index]
		left++
		index++
	}
}
