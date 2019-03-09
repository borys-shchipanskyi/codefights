package main

import (
	"fmt"
	//. "codefights/interviewPractice/solutions"
	//. "codefights/arcade/core"
	"strings"
)

type Dir struct {
	Name   string
	SubDir []*Dir
	IsFile bool
}

func getDirAsTree(input string) *Dir {
	dirs := strings.Split(input, "\f")
	if len(dirs) > 1 {
		for _, dir := range dirs {
			fi := strings.Index(dir, "\t")
			if fi != -1 {
				deep := strings.LastIndex(dir, "\t") + 1 - fi
				fmt.Println(dir, deep)
			} else {

				fmt.Println(dir)
			}

		}
	}

	return &Dir{input, nil, true}
}






func main() {
	fmt.Println("hello")
	getDirAsTree("user\f\tpictures\f\t\tphoto.png\f\t\tcamera\f\tdocuments\f\t\tlectures\f\t\t\tnotes.txt")
	//s :=ArrayPacking([]int{23,85,0})
	//s :=AdditionWithoutCarrying(456, 1734)

	//t := getTree()
	//PrintTree(t, "root")
	//fmt.Println("-----")
	//t := CreateTree(StrToMap([]byte(`{"value":1,"left":{"value":0,"left":{"value":3,"left":null,"right":null},"right":{"value":1,"left":null,"right":null}},"right":{"value":4,"left":null,"right":null}}`)))
	//
	//	s := DigitTreeSum(t)
	//PrintTree(t, "root")
	//fmt.Println(s)

}

