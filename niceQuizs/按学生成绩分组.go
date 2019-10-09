package main

import (
	"fmt"
)

type student struct {
	name  string
	score int
}

func groupBy(students []student) map[string][]student {
	m := make(map[string][]student)
	for i := range students {
		score := students[i].score
		if score < 60 {
			m["C"] = append(m["C"], students[i])
		} else if score >= 60 && score < 80 {
			m["B"] = append(m["B"], students[i])
		} else {
			m["A"] = append(m["A"], students[i])
		}

	}
	return m
}
func main() {
	intput := []student{
		{name: "张三", score: 84},
		{name: "李四", score: 58},
		{name: "王五", score: 99},
		{name: "赵六", score: 69},
	}

	output := groupBy(intput)

	fmt.Println(output)

}
