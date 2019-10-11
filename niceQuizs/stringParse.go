package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type node struct {
	Type     string
	Children []*node
}

var keyWords map[string]int

func buildTree(str string) *node {
	if str == "" {
		return nil
	}
	node := createNode(str)
	childrens := getChildrens(str)
	for i := 0; i < len(childrens); i++ {
		node.Children = append(node.Children, buildTree(childrens[i]))
	}
	return node
}

func createNode(str string) *node {
	var nodeType string
	var index = 0
	for i := range str {
		if (str[i] <= 90 && str[i] >= 65) || (str[i] <= 122 && str[i] >= 97) {
			index++
		} else {
			break
		}
	}
	nodeType = str[:index]
	// 判断解析的类型属不属于关键字
	if _, ok := keyWords[nodeType]; !ok {
		return nil
	}
	return &node{Type: nodeType}
}

// 根据字符串解析child
func getChildrens(str string) []string {
	if str == "" {
		return nil
	}
	str = strings.Replace(str, " ", "", -1)
	var index int
	for i := range str {
		if (str[i] <= 90 && str[i] >= 65) || (str[i] <= 122 && str[i] >= 97) {
			index++
		} else {
			break
		}
	}
	firstStr := str[:index]
	switch firstStr {
	case "int":
		return nil
	case "string":
		return nil
	case "bool":
		return nil
	}

	left, right := 0, len(str)-1
	for {
		if str[left] != '<' {
			left++
		}
		if str[right] != '>' {
			right--
		}
		if str[left] == '<' && str[right] == '>' {
			break
		}
	}
	str = str[left+1 : right]
	if firstStr == "Array" {
		return []string{str}
	}
	index = 0
	var separateSign int
	if firstStr == "Map" {
		for i := range str {
			if (str[i] <= 90 && str[i] >= 65) || (str[i] <= 122 && str[i] >= 97) {
				index++
			} else {
				break
			}
		}
		first := str[:index]
		if first != "Map" {
			separateSign = strings.Index(str, ",")
		} else {
			separateSign = strings.LastIndex(str, ",")
		}
	}
	return []string{str[:separateSign], str[separateSign+1:]}
}

func main() {
	keyWords = make(map[string]int)
	keyWords["int"] = 0
	keyWords["bool"] = 0
	keyWords["string"] = 0
	keyWords["Array"] = 0
	keyWords["Map"] = 0

	testCases := []string{
		"int",
		"bool",
		"string",
		"Array<int>",
		"Array<Array<int>>",
		"Array<Array<Array<int>>>",
		"Map<string, Map<string, bool>>",
		"Map<Map<string, bool>, bool>",
	}

	for _, v := range testCases {
		node := buildTree(v)
		data, _ := json.Marshal(node)
		fmt.Println(string(data))
	}
	// output:
	//{"Type":"int","Children":null}
	//{"Type":"bool","Children":null}
	//{"Type":"string","Children":null}
	//{"Type":"Array","Children":[{"Type":"int","Children":null}]}
	//{"Type":"Array","Children":[{"Type":"Array","Children":[{"Type":"int","Children":null}]}]}
	//{"Type":"Array","Children":[{"Type":"Array","Children":[{"Type":"Array","Children":[{"Type":"int","Children":null}]}]}]}
	//{"Type":"Map","Children":[{"Type":"string","Children":null},{"Type":"Map","Children":[{"Type":"string","Children":null},{"Type":"bool","Children":null}]}]}
	//{"Type":"Map","Children":[{"Type":"Map","Children":[{"Type":"string","Children":null},{"Type":"bool","Children":null}]},{"Type":"bool","Children":null}]}

}

