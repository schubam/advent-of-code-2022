package tree

import (
	"fmt"
	"strings"
)

const DIR = "dir"
const FILE = "file"

type Node struct {
	Type     string
	Name     string
	Size     int32
	children []*Node
	parent   *Node
}

type Tree struct {
	Root    *Node
	Current *Node
}

func NewTree(name string) *Tree {
	n := &Node{Type: DIR, Name: name, parent: nil}
	return &Tree{Root: n, Current: n}
}

func (t *Tree) Print() {
	t.Root.Print(0)
}

func (root *Node) Print(indent int) {
	space := strings.Repeat(" ", indent)
	if root.Type == DIR && root.Size > 100000 {
		fmt.Printf("%s - %s (%s, size=%d)\n", space, root.Name, root.Type, root.Size)
	}
	for _, n := range root.children {
		n.Print(indent + 2)
	}
}

func (t *Tree) CreateDir(name string) {
	//fmt.Println("create dir:", name)
	n := &Node{Type: DIR, Name: name, parent: t.Current}
	t.Current.children = append(t.Current.children, n)
}

func (t *Tree) CreateFile(name string, size int32) {
	//fmt.Println("create file:", name)
	n := &Node{Type: FILE, Name: name, parent: t.Current, Size: size}
	t.Current.children = append(t.Current.children, n)

	checkNode := n
	for checkNode != nil {
		if checkNode.parent != nil {
			checkNode.parent.Size += n.Size
			checkNode = checkNode.parent
		} else {
			checkNode = nil
		}
	}
}

func (t *Tree) Traverse(name string) {
	//fmt.Println("traverse to", name)
	switch name {
	case "..":
		t.Current = t.Current.parent
	case "/":
		t.Current = t.Root
	default:
		{
			for _, n := range t.Current.children {
				if n.Name == name && n.Type == DIR {
					t.Current = n
					return
				}
			}
		}
	}
}

const all = int32(70000000)
const needed = int32(30000000)

func (t *Tree) FindSmallestDir() int32 {
	free := all - t.Root.Size
	need := needed - free

	return SmallestDirectorySize(t.Root, need, all)
}

func SmallestDirectorySize(n *Node, need int32, min int32) int32 {
	if n.Type == DIR && n.Size >= need && n.Size <= min {
		min = n.Size
	}

	for _, c := range n.children {
		min = SmallestDirectorySize(c, need, min)
	}

	return min
}

func DirectorySizes(n *Node) int32 {
	var res int32

	if n.Type == DIR && n.Size <= 100000 {
		res += n.Size
	}

	for _, c := range n.children {
		res += DirectorySizes(c)
	}

	return res
}
