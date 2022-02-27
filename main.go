package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	flatList := []*Flat{
		{Path: "~/Foo/Videos/Clasic/museum.mp4"},
		{Path: "~/Foo/Documents/doc/virtual.doc"},
		{Path: "~/Foo/Documents/doc/real.doc"},
		{Path: "~/Foo/Documents/doc/intuitive.doc"},
		{Path: "~/Foo/Documents/doc/example.doc"},
		{Path: "~/Foo/Documents/Pdf/analysis.pdf"},
		{Path: "~/Bar/Exist/Pdf/analysis.pdf"},
		{Path: "~/Mico/Sql/Pdf/analysis.pdf"},
		{Path: "~/Data/Variable/Pdf/analysis.pdf"},
	}
	treeList := GetTree(flatList)

	b, err := json.Marshal(treeList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func GetTree(flatList []*Flat) []*Tree {

	list := []*Tree{}
	for _, item := range flatList {
		splitted := getSplittedPath(item.Path)
		var parent *Tree
		for _, text := range splitted {
			node, isNil := getNode(list, parent, text)
			if !isNil {
				parent = node
			} else {
				if parent == nil {
					parent = createData(text, item.Path)
					list = append(list, parent)
				} else {
					data := createData(text, item.Path)
					parent.Childs = append(parent.Childs, data)
					parent = data
				}
			}
		}
	}

	return list
}

func createData(key string, path string) *Tree {
	tree := &Tree{
		Key:    key,
		Path:   path,
		Type:   getType(key),
		Childs: make([]*Tree, 0),
	}

	return tree
}

func getType(key string) string {
	val := "File"
	if !strings.Contains(key, ".") {
		val = "Folder"
	}

	return val
}

func getSplittedPath(path string) []string {
	cleanPath := strings.Replace(path, "~/", "", -1)
	splitted := strings.Split(cleanPath, "/")

	return splitted
}

func getNode(list []*Tree, parent *Tree, name string) (*Tree, bool) {
	tree := &Tree{}
	isNil := true
	for _, value := range list {
		if name == value.Key {
			tree = value
			isNil = false
		}
	}

	if parent != nil {
		for _, value := range parent.Childs {
			if name == value.Key {
				tree = value
				isNil = false
			}
		}
	}

	return tree, isNil
}

type Flat struct {
	Path string
}

type Tree struct {
	Key    string
	Path   string
	Type   string
	Childs []*Tree
}
