package main

type Node struct {
	Val bool
	IsLeaf bool
	TopLeft *Node
	TopRight *Node
	BottomLeft *Node
	BottomRight *Node
}
var g [][]int

func construct(grid [][]int) *Node {
	g = grid
	return constructHelp(0,0,len(grid)-1,len(grid)-1)
}

func constructHelp(a,b,c,d int) *Node{
	standard := g[a][b]
	ok := true
	for i:=a;i<=c && ok;i++{
		for j:=b;j<=d && ok;j++{
			if g[i][j] != standard{
				ok = false
			}
		}
	}
	if ok{
		return &Node{standard==1,true,nil,nil,nil,nil}
	}
	//root := new(Node)
	var root Node
	root.Val = standard==1
	root.IsLeaf = false
	midX,midY:= c-a+1,d-b+1
	root.TopLeft= constructHelp(a,b,a+midX/2-1,b+midY/2-1)
	root.TopRight = constructHelp(a,b+midY/2,a+midX/2-1,d)
	root.BottomLeft = constructHelp(a+midX/2,b,c,b+midY/2-1)
	root.BottomRight = constructHelp(a+midX/2,b+midY/2,c,d)
	return &root
}