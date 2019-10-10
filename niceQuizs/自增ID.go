package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	id       string
	_type    string
	name     string
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

func getIncName(srcName string, root *TreeNode) string {
	var sequences = make(map[int]int)
	var s = new(stack)
	s.push(root)
	for !s.isEmpty() {
		top := s.popup()
		if top._type == srcName {
			// 传入一个widget，返回末尾的数字
			seq := getSeq(top)

			// 将数字存到map中
			sequences[seq] = 1
		}

		children := top.children
		if children != nil && len(children) > 0 {
			for i := 0; i < len(children); i++ {
				s.push(children[i])
			}
		}
	}
	// 从0到map的长度遍历，判断数字是否连续
	for i := 0; i < len(sequences); i++ {
		if _, ok := sequences[i]; !ok {
			return strconv.Itoa(i)
		}
	}
	return strconv.Itoa(len(sequences))

}

// 取出一个widget末尾的数字
func getSeq(node *TreeNode) int {
	if node.name == "button" || node.name == "view" {
		return 0
	}
	var sum int
	var power = 1
	for i := len(node.name) - 1; i >= 0; i-- {
		if node.name[i] >= 48 && node.name[i] <= 57 {
			sum += (int(node.name[i]) - 48) * power
			power *= 10
		} else {
			break
		}
	}
	return sum
}

func main() {
	// 验证我写的函数，取出一个widget背后的数字
	node := &TreeNode{name: "view_123"}
	fmt.Println(getSeq(node)) //output: 123

	node = &TreeNode{name: "button"}
	fmt.Println(getSeq(node)) //output:0

	root := &TreeNode{id: "1", name: "view", _type: "View"}
	root.children = []*TreeNode{
		&TreeNode{id: "2", name: "button", _type: "Button"},
		&TreeNode{id: "3", name: "view_1", _type: "View",
			children: []*TreeNode{
				&TreeNode{id: "4", name: "button_1", _type: "Button"},
				&TreeNode{id: "5", name: "view_2", _type: "View"},
				&TreeNode{id: "6", name: "view_4", _type: "View"}}}}


	// 测试了view_1,view_2,view_4的情况
	id:=getIncName("View",root)
	fmt.Println(id) //output: 3
}
