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
	fi os.FileInfo
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

var closed [10]bool
func dirTree(out *os.File, info eleminfo) {
	for i := 0; i < info.level; i++ {
		if closed[i] {
			out.Write("\t")
		} else {
			out.Write("│\t")
		}
	}
	if info.isLast {
		out.Write("└───" + info.fi.Name() + "\n")
		closed[info.level] = true
	} else {
		out.Write("├───" + info.fi.Name() + "\n")
		closed[info.level] = false
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	current := path
	stk := new(stack)
	levelParent := 0
	var startList []os.FileInfo 
	for i := len(startList); i > 0 ; i-- {
		infoN := new(eleminfo)
		infoN.isFirst = i == len(startList)
		infoN.isLast = i == 1
		infoN.fi = startList[i-1]
		infoN.level = info. 
		if info.fi.IsDir || printFiles {
			stk.Push(info)
		}
	}
	for all := 1; all > 0; all = stk.Size {
		
		startList, err := ioutil.ReadDir(current)
		
		for i := len(startList); i > 0 ; i-- {
			infoN := new(eleminfo)
			infoN.isFirst = i == len(startList)
			infoN.isLast = i == 1
			infoN.fi = startList[i-1]
			infoN.level = info. 
			if info.fi.IsDir || printFiles {
				stk.Push(info)
			}
		}
		info := stk.Pop()
		    
		startList, err := ioutil.ReadDir(info.fi.
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
