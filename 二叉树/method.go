package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findMin(a,b int) int{
	if a>b{
		return b
	}else {
		return a
	}
}

func findMax(a,b int) int{
	if a>b {
		return a
	}else {
		return b
	}
}

func abs(a int) int{
	if a>=0{
		return a
	}else {
		return -a
	}
}

func main()  {
	inorder  := []int{9,3,15,20,7}
	posorder := []int{9,15,7,20,3}
	root := buildTree(inorder,posorder)
	fmt.Println(isBalanced(root))

}

func maxLevelSum(root *TreeNode) int {
	if root==nil{
		return 0
	}
	levelSum := make([][]int,0)
	treeList := make([]*TreeNode,0)
	treeList = append(treeList,root)
	level := 1
	for len(treeList)>0{
		var sum int
		tmp := make([]*TreeNode,0)
		for len(treeList)>0{
			node := treeList[len(treeList)-1]
			sum += node.Val
			if node.Left!=nil{
				tmp = append(tmp,node.Left)
			}
			if node.Right!=nil{
				tmp = append(tmp,node.Right)
			}
			treeList = treeList[:len(treeList)-1]
		}
		levelSum = append(levelSum,[]int{level,sum})
		level++
		treeList = tmp
	}
	sort.Slice(levelSum, func(i, j int) bool {
		if levelSum[i][1]==levelSum[j][1]{
			return levelSum[i][0]<levelSum[j][0]
		}
		return levelSum[i][1]>levelSum[j][1]
	})
	fmt.Println(levelSum)
	return 0
}

func minCameraCover(root *TreeNode) int {
	result := 0
	var traceback func(root *TreeNode) int
	//0 无覆盖， 1 有摄像头， 2  有覆盖
	traceback = func(root *TreeNode) int{
		if root==nil{
			return 2
		}
		left := traceback(root.Left)
		right := traceback(root.Right)
		if left==2 && right==2{
			return 0
		}
		if right==0 || left==0{
			result++
			return 1
		}
		if right==1 || left==1{
			return 2
		}
		return -1
	}
	if traceback(root)==0{
		result++
	}
	return result
}

func convertBST(root *TreeNode) *TreeNode {
	var prev *TreeNode
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root==nil{
			return nil
		}
		travel(root.Right)
		if prev != nil {
			root.Val = root.Val + prev.Val
		}
		prev = root
		travel(root.Left)
		return root
	}
	return travel(root)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	index := len(nums)/2
	val := nums[index]
	root := &TreeNode{val,nil,nil}
	root.Left = sortedArrayToBST(nums[:index])
	root.Right = sortedArrayToBST(nums[index+1:])
	return root
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val<low{
		right := trimBST(root.Right,low,high)
		return right
	}
	if root.Val>high{
		left := trimBST(root.Left,low,high)
		return left
	}
	root.Left = trimBST(root.Left,low,high)
	root.Right = trimBST(root.Right,low,high)
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}
	if root.Val <key{
		root.Right = deleteNode(root.Right,key)
	}
	if root.Val == key{
		if root.Left==nil && root.Right==nil{
			root = nil
		}else if root.Left != nil && root.Right == nil{
			root = root.Left
		}else if root.Right != nil && root.Left == nil{
			root = root.Right
		}else {
			tmp := root.Left
			root.Left=nil
			start := root.Right
			root = start
			for start.Left != nil{
				start = start.Left
			}
			start.Left = tmp
		}
	}
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return &TreeNode{val,nil,nil}
	}
	if root.Val>val{
		root.Left = insertIntoBST(root.Left,val)
	}
	if root.Val<val{
		root.Right = insertIntoBST(root.Right,val)
	}
	return root
}

func findNode(root *TreeNode,val int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val==val{
		return root
	}
	leftRes := findNode(root.Left,val)
	rightRes := findNode(root.Right,val)
	if leftRes!=nil && rightRes==nil{
		return leftRes
	}
	if rightRes!=nil && leftRes==nil{
		return rightRes
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==p || root==q || root==nil{
		return root
	}
	leftRes := lowestCommonAncestor(root.Left,p,q)
	rightRes := lowestCommonAncestor(root.Right,p,q)
	if leftRes!=nil && rightRes!=nil{
		return root
	}
	if leftRes==nil && rightRes!=nil{
		return rightRes
	}
	if rightRes==nil && leftRes!=nil{
		return leftRes
	}
	return nil
}

func findMode(root *TreeNode) []int {
	var prev *TreeNode
	res := make([]int,0)
	count := 0
	max := math.MinInt
	var travel func(root *TreeNode) []int
	travel = func(root *TreeNode) []int {
		if root==nil{
			return nil
		}
		travel(root.Left)
		if prev != nil && prev.Val==root.Val{
			count++
		}else{
			count=1
		}
		if count==max{
			res = append(res,root.Val)
		}else if count>max{
			res = make([]int,0)
			max = count
			res = append(res,root.Val)
		}
		prev = root
		travel(root.Right)
		return res
	}
	return travel(root)
}

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	var travel func(root *TreeNode) int
	res := math.MaxInt
	travel = func(root *TreeNode) int{
		if root==nil{
			return res
		}
		travel(root.Left)
		if prev != nil && abs(root.Val - prev.Val)<res{
			res = abs(root.Val - prev.Val)
		}
		prev = root
		travel(root.Right)
		return res
	}
	return travel(root)
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var judge func(root *TreeNode) bool
	judge = func(root *TreeNode) bool{
		if root==nil{
			return true
		}
		leftRes := judge(root.Left)
		if prev != nil && prev.Val>=root.Val{
			return false
		}
		prev = root
		rightRes := judge(root.Right)
		return rightRes && leftRes
	}
	return judge(root)
}


func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val==val{
		return root
	}
	if root.Val>val{
		return searchBST(root.Left,val)
	}else {
		return searchBST(root.Right,val)
	}
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1==nil && root2==nil{
		return nil
	}
	if root1==nil{
		return root2
	}
	if root2==nil{
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left,root2.Left)
	root1.Right = mergeTrees(root1.Right,root2.Right)
	return root1
}

func findMaxOfNums(nums []int) (res int,index int) {
	if len(nums)==0{
		res = 0
		index = 0
		return
	}
	res = math.MinInt
	for i:=0;i<len(nums);i++{
		if nums[i]>res{
			res = nums[i]
			index = i
		}
	}
	return res,index
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)==0 {
		return nil
	}
	arr,index := findMaxOfNums(nums)
	leftArr :=  nums[:index]
	rightArr := nums[index+1:]
	node := &TreeNode{arr,nil,nil}
	node.Left = constructMaximumBinaryTree(leftArr)
	node.Right = constructMaximumBinaryTree(rightArr)
	return node
}

/*func binaryTreePaths(root *TreeNode) []string {
	res := make([]string,0)
	var str string
	TreePathsHelp(root,str,&res)
	return res
}

func TreePathsHelp(root *TreeNode,str string,res *[]string)  {
	if root!=nil && root.Left==nil && root.Right==nil{
		var tmp string
		tmp = str
		if tmp==""{
			tmp = tmp + strconv.Itoa(root.Val)
		}else {
			tmp = tmp + "->" + strconv.Itoa(root.Val)
		}
		*res = append(*res,tmp)
		return
	}
	if root==nil{
		return
	}
	if str==""{
		str = str + strconv.Itoa(root.Val)
	}else {
		str = str + "->" + strconv.Itoa(root.Val)
	}
	TreePathsHelp(root.Left,str,res)
	TreePathsHelp(root.Right,str,res)
}
*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder)==0{
		return nil
	}
	head := postorder[len(postorder)-1]
	root := &TreeNode{head,nil,nil}
	if len(postorder)==1{
		return root
	}
	var index int
	for i:=0;i<len(inorder);i++{
		if head==inorder[i]{
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[:index],postorder[:index])
	root.Right = buildTree(inorder[index+1:],postorder[index:len(postorder)-1])
	return root
}




func pathSum(root *TreeNode, targetSum int) [][]int {
	if root==nil{
		return nil
	}
	nodeList := []int{root.Val}
	ans := make([][]int,0)
	pathSumHelp(root,targetSum-root.Val,nodeList,&ans)
	return ans
}

func pathSumHelp(root *TreeNode, targetSum int,nodeList []int,ans *[][]int) {
	if root.Left==nil && root.Right==nil && targetSum==0{
		tmp := make([]int,len(nodeList))
		copy(tmp,nodeList)
		*ans = append(*ans,tmp)
		return
	}

	if root.Left != nil{
		nodeList = append(nodeList,root.Left.Val)
		pathSumHelp(root.Left,targetSum-root.Left.Val,nodeList,ans)
		nodeList = nodeList[:len(nodeList)-1]
	}
	if root.Right != nil{
		nodeList = append(nodeList,root.Right.Val)
		pathSumHelp(root.Right,targetSum-root.Right.Val,nodeList,ans)
		nodeList = nodeList[:len(nodeList)-1]
	}
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root==nil{
		return false
	}
	return hasPathSumHelp1(root,targetSum - root.Val)
}

func hasPathSumHelp1(root *TreeNode, count int) bool {
	if root.Left==nil && root.Right==nil && count==0{
		return true
	}
	if root.Left!=nil{
		count = count - root.Left.Val
		if hasPathSumHelp1(root.Left,count){
			return true
		}
		count = count + root.Left.Val
	}
	if root.Right!=nil{
		count = count - root.Right.Val
		if hasPathSumHelp1(root.Right,count){
			return true
		}
		count = count + root.Right.Val
	}
	return false
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	var res int
	queue.PushBack(root)
	for queue.Len()>0{
		length := queue.Len()
		for i:=0;i<length;i++{
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i==0{
				res = node.Val
			}
			if node.Left != nil{
				queue.PushBack(node.Left)
			}
			if node.Right != nil{
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}

/*var maxDeep int
var value int

func findBottomLeftValue(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil && root.Right==nil{
		return root.Val
	}
	findLeftValue(root,maxDeep)
	return value
}

func findLeftValue(root *TreeNode,deep int)  {
	if root.Left==nil && root.Right==nil{
		if deep>maxDeep{
			value = root.Val
			maxDeep = deep
		}
	}
	if root.Left!=nil{
		deep++
		findLeftValue(root.Left,deep)
		deep--
	}
	if root.Right!=nil{
		deep++
		findLeftValue(root.Right,deep)
		deep--
	}
}*/

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string,0)
	var help func(root *TreeNode,str string)
	help = func(root *TreeNode, str string) {
		if root.Left==nil && root.Right==nil{
			v := str + strconv.Itoa(root.Val)
			res = append(res,v)
			return
		}
		str = str + strconv.Itoa(root.Val) + "->"
		if root.Left!=nil{
			help(root.Left,str)
		}
		if root.Right!=nil{
			help(root.Right,str)
		}
	}
	help(root,"")
	return res
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root==nil{
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)
	rightValue := sumOfLeftLeaves(root.Right)

	value := 0
	if root.Left != nil && root.Left.Left==nil && root.Left.Right==nil {
		value = root.Left.Val
	}
	sum := leftValue + rightValue + value
	return sum
}

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	if isBalancedHelp(root)!=-1{
		return true
	}
	return false
}

func isBalancedHelp(root *TreeNode) int {
	if root==nil{
		return 0
	}
	left := isBalancedHelp(root.Left)
	if left==-1{
		return -1
	}
	right := isBalancedHelp(root.Right)
	if right==-1{
		return -1
	}
	if abs(right-left)>1{
		return -1
	}
	return findMax(right,left)+1
}

func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	leftH,rightH := 0,0
	leftNode,rightNode := root.Left,root.Right
	if leftNode != nil{
		leftNode = leftNode.Left
		leftH++
	}
	if rightNode != nil{
		rightNode = rightNode.Right
		rightH++
	}
	if rightH==leftH {
		return (2 << leftH) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}



func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil && root.Right!=nil{
		return 1+minDepth(root.Right)
	}
	if root.Right==nil && root.Left!=nil{
		return 1+minDepth(root.Left)
	}
	return findMin(minDepth(root.Right),minDepth(root.Left))+1
}

/*func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}*/

func isSymmetric(root *TreeNode) bool {
	if root==nil {
		return true
	}
	return isSymmetricHelp(root.Left,root.Right)
}

func isSymmetricHelp(left *TreeNode,right *TreeNode) bool {
	var a,b bool
	if left==nil && right==nil{
		return true
	}
	if left==nil && right!=nil || left!=nil && right==nil{
		return false
	}
	if left.Val==right.Val{
		a = isSymmetricHelp(left.Left,right.Right)
		b = isSymmetricHelp(left.Right,right.Left)
	}else {
		return false
	}
	return a && b
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	left := maxDepth(root.Left)+1
	right := maxDepth(root.Right)+1
	if left<right && left!=1{
		return left
	}else {
		return right
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Left==nil && root.Right==nil{
		return root
	}
	root.Left,root.Right = root.Right,root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}



func inorderTraversal(root *TreeNode) []int {
	ans := make([]int,0)
	if root==nil{
		return ans
	}
	st := list.New()
	st.PushBack(root)
	var node *TreeNode
	for st.Len()>0{
		e := st.Back()
		st.Remove(e)
		if e.Value==nil{
			e=st.Back()
			st.Remove(e)
			node = e.Value.(*TreeNode)
			ans = append(ans,node.Val)
			continue
		}
		node = e.Value.(*TreeNode)
		if node.Right!=nil{
			st.PushBack(node.Right)
		}
		st.PushBack(root)
		st.PushBack(nil)
		if node.Left!=nil{
			st.PushBack(node.Left)
		}
	}
	return ans
}

func levelOrder(root *TreeNode) [][]int {
	resList := make([][]int,0)
	if root==nil{
		return resList
	}
	var node *TreeNode
	nodeList := make([]*TreeNode,0)
	nodeList = append(nodeList,root)
	for len(nodeList)>0{
		var tmpNodeList []*TreeNode
		var tmpResList []int
		for len(nodeList)>0{
			node = nodeList[0]
			tmpResList = append(tmpResList,node.Val)
			if node.Left != nil{
				tmpNodeList = append(tmpNodeList,node.Left)
			}
			if node.Right != nil{
				tmpNodeList = append(tmpNodeList,node.Right)
			}
			nodeList = nodeList[1:]
		}
		nodeList = tmpNodeList
		resList = append(resList,tmpResList)
	}
	return resList
}