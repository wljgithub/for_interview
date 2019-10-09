package main

import "fmt"

// 这段代码是我在网上找的，我没学过，一下子不知道要怎么遍历树

// 现在才知道这种方式叫前序遍历，这是Java版的前序遍历

// 链接： https://blog.csdn.net/My_Jobs/article/details/43451187

//public void preOrderTraverse1(TreeNode root) {
//if (root != null) {
//System.out.print(root.val+"  ");
//preOrderTraverse1(root.left);
//preOrderTraverse1(root.right);
//}

type TreeNode struct {
	id    string
	left  *TreeNode
	right *TreeNode
}

// 下面这段代码是参考了Java版的前序遍历之后，写出来的
func findNodeById(root *TreeNode, id string) *TreeNode {
	if root != nil {
		if root.id == id {
			return root
		}
		fmt.Println(root.id)
		left := findNodeById(root.left, id)
		right := findNodeById(root.right, id)
		if left != nil {
			return left
		} else {
			return right
		}
	}
	return nil
}

func main() {
	// 当我写完再看一下题目的时候，发现题目是一颗多叉树
	// 我画图的时候只有两个节点，光看我笔记上的图，以为是课二叉树
	root := &TreeNode{id: "1"}
	root.left = &TreeNode{id: "2"}
	root.right = &TreeNode{id: "3"}
	root.right.left = &TreeNode{id: "4"}
	root.right.right = &TreeNode{id: "5"}

	node := findNodeById(root, "3")
	fmt.Println("find:", node.id)

	node = findNodeById(root, "5")
	fmt.Println("find:", node.id)

}
