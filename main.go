package main

import (
	//	"fmt"
	//	"io"
	"fmt"
	"os"
	//	"path/filepath"
	"io/ioutil"
	//	"strings"
	"errors"
)

type eleminfo struct {
	data   interface{}
	isLast bool
	level  int
	fi     os.FileInfo
	path   string
}

type element struct {
	data eleminfo
	next *element
}

type stack struct {
	head *element
	Size int
}

func (stk *stack) Push(data eleminfo) {
	element := new(element)
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++
}

func (stk *stack) Pop() (eleminfo, error) {
	if stk.head == nil {
		var s eleminfo
		return s, errors.New("The method failed")
	}
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	return r, nil
}

var closed [100]bool

func printTree(out *os.File, info eleminfo) {
	for i := 0; i < info.level; i++ {
		if closed[i] {
			out.WriteString("\t")
		} else {
			out.WriteString("│\t")
		}
	}
	if info.isLast {
		out.WriteString("└───" + info.fi.Name() + "\n")
		closed[info.level] = true
	} else {
		out.WriteString("├───" + info.fi.Name() + "\n")
		closed[info.level] = false
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	currPath := path
	stk := new(stack)
	levelParent := -1
	var info eleminfo
	info.path = path
	isFirst := true
	for all := 1; all > 0; all = stk.Size {
		if isFirst || info.fi.IsDir() {
			isFirst = false
			startList, _ := ioutil.ReadDir(currPath)
			fmt.Println(out, len(startList))
			for i := len(startList); i > 0; i-- {
				var infoN eleminfo
				//infoN.isFirst = i == len(startList)
				infoN.isLast = i == 1
				infoN.fi = startList[i-1]
				infoN.path = currPath
				infoN.level = levelParent + 1
				if printFiles || infoN.fi.IsDir() {
					stk.Push(infoN)
				}
			}
		}
		info, _ := stk.Pop()
		if info.fi.IsDir() {
			currPath = currPath + info.fi.Name()
			levelParent = info.level
		}
		printTree(out, info)
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
