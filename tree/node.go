package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func Create(value int) *Node { //返回局部变量依然可以使用，直到返回的指针不再使用，垃圾回收机制将其释放
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}

func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

func (node *Node) TraverseNode() {
	if node == nil {
		return
	}
	node.Left.TraverseNode()
	node.Print()
	node.Right.TraverseNode()
}


