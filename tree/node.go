package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode
}

func createNode(value int) *treeNode{  //返回局部变量依然可以使用，直到返回的指针不再使用，垃圾回收机制将其释放
	return &treeNode{value:value}
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		return
	}
	node.value = value
}

func (node treeNode) print()  {
	fmt.Printf("%d ",node.value)
}

func (node *treeNode) traverseNode()  {
	if node ==nil {
		return
	}
	node.left.traverseNode()
	node.print()
	node.right.traverseNode()
}

func main() {
	var root treeNode

	root = treeNode{value:3}//root = treeNode{value:3,nil}
	//root.value=5
	root.left = &treeNode{}//root.left.value=3
	root.right = &treeNode{6,nil,nil}//root.right.value=6
	root.right.left=new(treeNode)

	root.left.right = createNode(2)

	fmt.Println(root)
	root.print()
	root.setValue(666)
	fmt.Println(root)
	root.print()

	/*nodes := []treeNode{
		{value:8},
		{value:18},
		{9,nil,&root},
	}

	fmt.Println(nodes)*/

	root.traverseNode()


}
