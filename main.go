package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	"errors"
)

type eleminfo struct {
	data interface {}
	isFirst bool
	isLast bool
	level int
	name string
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


func dirTree(out *os.File, path string, printFiles bool) error {
	startList, err := ioutil.ReadDir(path)
	stk := new(stack)
	levelParent := 0
	for all := 1; all > 0; all = stk.Size {
		for i := len(startList); i > 0 ; i-- {
			info := new(eleminfo)
			info.isFirst = i == len(startList)
			info.isLast = i == 1
			info.name = startList[i-1].Name()
			info.level = levelParent 
			stk.Push(info)
		}
		info := stk.Pop()
		startList, err := ioutil.ReadDir(info.name)
		levelParent = info.level

	}

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
