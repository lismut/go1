package main

import (
//	"fmt"
//	"io"
	"os"
//	"path/filepath"
	"io/ioutil"
//	"strings"
	"errors"
)

type eleminfo struct {
	data interface {}
	isFirst bool
	isLast bool
	level int
	fi os.FileInfo
	path string
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
	currPath := path
	stk := new(stack)
	levelParent := 0
	var info eleminfo
	for all := 1; all > 1; all = stk.Size {
		if (info.fi.IsDir() || currPath == path) {
			startList, _ := ioutil.ReadDir(currPath)
			for i := len(startList); i > 0 ; i-- {
				infoN := new(eleminfo)
				infoN.isFirst = i == len(startList)
				infoN.isLast = i == 1
				infoN.fi = startList[i-1]
				infoN.path = currPath
				infoN.level = levelParent + 1 
				stk.Push(*infoN)
			}
		}
		info,_ := stk.Pop()
		if info.fi.IsDir() {
			currPath = currPath + info.fi.Name()
		}
		levelParent = info.level

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
