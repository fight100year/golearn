package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getFileContent(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

type jsonData []struct {
	Name       string   `json:"ImportPath"`
	Deps       []string `json:"Deps"`
	countOfDep int
}

type node struct {
	name       string
	subnode    map[string]*node
	parentnode map[string]*node
}

func createNode(name string) *node {
	n := node{name: name}
	n.subnode = make(map[string]*node, 0)
	n.parentnode = make(map[string]*node, 0)

	return &n
}

var nodes = map[string]*node{}

func main() {
	content, err := getFileContent("info")
	if err != nil {
		fmt.Println("read file failed:", err)
		return
	}

	var c jsonData
	err = json.Unmarshal(content, &c)
	if err != nil {
		fmt.Println("parse content failed:", err)
		return
	}

	level := []string{}

	for _, x := range c {
		x.countOfDep = len(x.Deps)
		// fmt.Printf("%s\n\t%d\t%v", x.Name, x.countOfDep, x.Deps)
		// fmt.Println()

		if len(x.Deps) == 0 {
			level = append(level, x.Name)
		}

		n, ok := nodes[x.Name]
		if !ok {
			n = createNode(x.Name)
			nodes[x.Name] = n
		}

		for _, subName := range x.Deps {
			subNode, ok := nodes[subName]
			if !ok {
				subNode = createNode(subName)
				nodes[subName] = subNode
			}

			n.subnode[subName] = subNode

			subNode.parentnode[x.Name] = n
		}
	}

	fmt.Printf("root:%d, %v\n", len(level), level)

	if len(c) != len(nodes) {
		fmt.Println("failed")
	} else {
		fmt.Printf("总包数:%d\n", len(c))
	}

	for {
		fmt.Print("\n输入要查询的包:")
		var name string
		_, _ = fmt.Scanln(&name)

		if name == "q" {
			return
		}
		if name == "r" {
			fmt.Printf("root:%d, %v\n", len(level), level)
			continue
		}
		if strings.HasPrefix(name, "search") {
			name = string([]rune(name)[len("search"):])
			if len(name) == 0 {
				continue
			}

			for key := range nodes {
				if strings.HasPrefix(key, name) {
					fmt.Println(key)
				}
			}

			continue
		}

		x, ok := nodes[name]
		if !ok {
			fmt.Print("请输入正确的包名")
			continue
		}

		fmt.Printf("parent:%d\n", len(x.parentnode))
		for key := range x.parentnode {
			fmt.Println("\t", key)
		}
		fmt.Printf("\nsubnode:%d\n", len(x.subnode))
		for key := range x.subnode {
			fmt.Println("\t", key)
		}
	}
}
