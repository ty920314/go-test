package main

import "tygo/tree"

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//nodes := []treeNode{
	//	{value:3},
	//	{},
	//	{6,nil,&root},
	//}
	root.Traverse()
}
