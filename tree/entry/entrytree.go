package main

import (

"GitHub/learngo/tree"
"fmt"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3} //root = Node{value:3,nil}
	//root.value=5
	root.Left = &tree.Node{}             //root.left.value=3
	root.Right = &tree.Node{6, nil, nil} //root.right.value=6
	root.Right.Left = new(tree.Node)

	root.Left.Right = tree.Create(2)

	fmt.Println(root)
	root.Print()
	root.SetValue(666)
	fmt.Println(root)
	//root.print()

	/*nodes := []Node{
		{value:8},
		{value:18},
		{9,nil,&root},
	}

	fmt.Println(nodes)*/

	root.TraverseNode()

}