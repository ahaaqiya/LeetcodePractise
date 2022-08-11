package main

import "fmt"

func main()  {
	fmt.Println('b'-'a')
}

type Trie struct {
	end bool
	children [26]*Trie
}


func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	node := this
	for _,ch := range word{
		order := ch - 'a'
		if node.children[order] == nil{
			node.children[order] = &Trie{}
		}
		node = node.children[order]
	}
	node.end = true
}


func (this *Trie) Search(word string) bool {
	node := this
	for _,ch := range word{
		order := ch - 'a'
		if node.children[order]==nil{
			return false
		}
		node = node.children[order]
	}
	if node !=nil && node.end{
		return true
	}
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _,ch := range prefix{
		order := ch - 'a'
		if node.children[order]==nil{
			return false
		}
		node = node.children[order]
	}
	if node!=nil{
		return true
	}
	return false
}
