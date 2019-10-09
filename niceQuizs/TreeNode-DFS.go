package main

import "fmt"

// https://segmentfault.com/a/1190000003004435
// 看了这篇文章，看他还用了深度优先和广度优先的方式实现多叉树遍历，我也想实现一个，但是感觉他写的乱七八糟的

// 所以我 又参考了这篇 https://zhuanlan.zhihu.com/p/29425290
// 然后实现一个golang版的

//主要参考了这段代码，然后自己实现一个栈，用栈来实现深度优先

//Stack<Node> stack = new Stack<Node>();
//List<Node> result = new ArrayList<Node>();
//stack.push(root);
//while (!stack.isEmpty()) {
//Node top = stack.pop();
//result.add(top);
//List<Node> children = top.getChildren();
//if (children != null && children.size() > 0) {
//for (int i = children.size() - 1; i >= 0; i--) {
//stack.push(children.get(i));
//}
//}
//}

type TreeNode struct {
	id       string
	lable    string
	children []*TreeNode
}

type stack []*TreeNode

func (s *stack) popup() *TreeNode {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last

}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func findNodeById(root *TreeNode, id string) *TreeNode {
	var s = new(stack)
	s.push(root)
	for !s.isEmpty() {
		top := s.popup()
		if top.id == id {
			return top
		}
		children := top.children
		if children != nil && len(children) > 0 {
			for i := 0; i < len(children); i++ {
				if children[i].id == id {
					return children[i]
				}
				s.push(children[i])
			}
		}
	}
	return nil
}
func main() {
	// 测试我实现的栈
	//var myStack = new(stack)
	//myStack.push(&TreeNode{id:"1",lable:"first"})
	//
	//fmt.Println(myStack.isEmpty())
	//
	//fmt.Println(*myStack.popup())

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
