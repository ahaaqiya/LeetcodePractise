package main

import "LeetCodeHOT100/hotMethod"

func main()  {
	arr := []int{1,2,3,4,5}
	head := hotMethod.CreatLinklistHead(arr)
	hotMethod.PrintLinklist(head)
}
