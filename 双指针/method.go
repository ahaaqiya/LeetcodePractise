package main

import "sort"

func main()  {

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int,0)
	for i:=0;i<len(nums)-3;i++{
		if nums[i]>target && target>0{
			break
		}
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		n1 := nums[i]
		for j:=i+1;j<len(nums)-2;j++{
			if nums[j]>target-nums[i] && target-nums[i]>0{
				break
			}
			if j>i+1 && nums[j]==nums[j-1]{
				continue
			}
			n2 := nums[j]
			left,right := j+1,len(nums)-1
			for left<right {
				n3,n4 := nums[left],nums[right]
				if n1+n2+n3+n4==target{
					res = append(res,[]int{n1,n2,n3,n4})
					for left<right && nums[left]==n3{
						left++
					}
					for left<right && nums[right]==n4{
						right--
					}
				}else if n1+n2+n3+n4<target{
					left++
				}else {
					right--
				}
			}
		}
	}
	return res
}

func findMin(a,b int) int {
	if a>b{
		return b
	}else {
		return a
	}
}

func findMax(a,b int) int {
	if a>b{
		return a
	}else {
		return b
	}
}

func removeElement(nums []int, val int) int {
	slow,fast := 0,0
	for fast < len(nums){
		if nums[fast]!=val{
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func replaceSpace(s string) string {
	count := 0
	fast := len(s)-1
	arr := []byte(s)
	for i:=0;i<len(arr);i++{
		if arr[i]==' '{
			count++
		}
	}
	tmp := make([]byte,count*2)
	arr = append(arr,tmp...)
	slow := len(arr)-1
	for fast>=0{
		if arr[fast]==' '{
			arr[slow] = '0'
			arr[slow-1] = '2'
			arr[slow-2] = '%'
			slow = slow - 3
			fast--
		}else {
			arr[slow] = arr[fast]
			fast--
			slow--
		}
	}
	return string(arr)
}

func moveZeroes(nums []int)  {
	fast,slow := 0,0
	for fast<len(nums){
		if nums[fast] != 0{
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow<len(nums){
		nums[slow] = 0
	}
}
