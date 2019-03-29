package main

import "fmt"

// SIZE 哈希表大小
const SIZE = 15

// Node 链表中元素的结构
type Node struct {
	Value int
	Next  *Node
}

// HashTable 哈希表的结构
type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element

	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]

			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
		}
		fmt.Println()
	}
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("size of hash:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
