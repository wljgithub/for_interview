package main

import "fmt"

type TreeNode struct {
	id       string
	lable    string
	children []*TreeNode
}

type queue []*TreeNode

func (s *queue) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *queue) poll() *TreeNode {
	first := (*s)[0]
	(*s) = (*s)[1:]
	return first
}

func (s *queue) add(node *TreeNode) {
	(*s) = append(*s, node)
}

// 用队列可以实现广度优先，其实思路跟刚才用栈实现的差不多，区别只是把栈换成了队列，队列是先进先出的，然后就成了广度优先
func findNodeById(root *TreeNode, id string) *TreeNode {
	var q = new(queue)
	q.add(root)
	for !q.isEmpty() {
		first := q.poll()
		if first.id == id {
			return first
		}
		children := first.children
		if children != nil && len(children) > 0 {
			for i := 0; i < len(children); i++ {
				if children[i].id == id {
					return children[i]
				}
				q.add(children[i])
			}
		}
	}
	return nil
}
func main() {
	// 测试队列
	//q:=new(queue)
	//fmt.Println(q.isEmpty())
	//
	//q.add(&TreeNode{id:"1"})
	//q.poll()
	//fmt.Println(q.isEmpty())

	root := &TreeNode{id: "1", lable: "first"}
	root.children = []*TreeNode{
		&TreeNode{id: "2", lable: "second"},
		&TreeNode{id: "3", lable: "third",
			children: []*TreeNode{
				&TreeNode{id: "4", lable: "fourth"},
				&TreeNode{id: "5", lable: "fifth"}}},
	}
	node := findNodeById(root, "1")
	fmt.Println(node.id)

	node = findNodeById(root, "3")
	fmt.Println(node.id)

	node = findNodeById(root, "5")
	fmt.Println(node.id)

}
