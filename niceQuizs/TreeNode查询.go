package main

import "fmt"

// 一开始没想到多叉树是怎么遍历的，然后参考了一段不知道什么语言的代码，感觉写得乱七八糟的，
// 唯一看懂的就这句 for (var i = 0, len = treeNodes.length; i < len; i++)，然后突然就知道要怎么写了

//var parseTreeJson = function(treeNodes){
//if (!treeNodes || !treeNodes.length) return;
//
//for (var i = 0, len = treeNodes.length; i < len; i++) {
//
//var childs = treeNodes[i].children;
//
//console.log(treeNodes[i].id);
//
//if(childs && childs.length > 0){
//parseTreeJson(childs);
//}
//}
//}

// 链接： https://segmentfault.com/a/1190000003004435

type TreeNode struct {
	id       string
	lable    string
	children []*TreeNode
}

func findNodeById(root *TreeNode, id string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.id == id {
		return root
	}
	for i := 0; i < len(root.children); i++ {
		temp := findNodeById(root.children[i], id)
		if temp != nil && temp.id == id {
			return temp
		}
	}
	return nil
}

func main() {
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
