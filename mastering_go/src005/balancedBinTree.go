// AVL二叉树,发明者: G. M. Adelson-Velsky和E. M. Landis
// 高度平衡树,任何节点的两子树的高度最大差别为1
// 递归版本

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// AVL ...
// 一个节点
type AVL struct {
	Parent *AVL
	Left   *AVL
	Right  *AVL
	data   int
	height int
}

func (avl *AVL) add(i int) *AVL {
	if i < avl.data {
		if avl.Left != nil {
			return avl.Left.add(i)
		}

		node := AVL{data: i, height: avl.height + 1, Parent: avl}
		avl.Left = &node

		return &node
	}

	if i > avl.data {
		if avl.Right != nil {
			return avl.Right.add(i)
		}

		node := AVL{data: i, height: avl.height + 1, Parent: avl}
		avl.Right = &node

		return &node
	}

	// i == avl.data
	fmt.Println("add a exists data:", i)
	return avl
}

// check 查找失衡点
// 从插入点的爷爷节点开始查找
func (avl *AVL) check() *AVL {
	if avl.Parent == nil || avl.Parent.Parent == nil {
		return nil
	}

	checkNode := avl.Parent.Parent
	leftDepth := checkNode.Left.depth()
	rightDepth := checkNode.Right.depth()
	if dif := leftDepth - rightDepth; dif == 1 || dif == 0 || dif == -1 {
		return avl.Parent.check()
	}

	return checkNode
}

// depth 计算节点的深度
// 返回值不是概念中的深度
//  返回值 = 概念上的深度 + 节点的高度
func (avl *AVL) depth() int {
	if avl.Left == nil && avl.Right == nil {
		return avl.height
	}

	leftMax, rightMax := 0, 0
	if avl.Left != nil {
		leftMax = avl.Left.depth()
	}
	if avl.Right != nil {
		rightMax = avl.Right.depth()
	}

	return max(leftMax, rightMax)
}

// max 计算最大值
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// balanceWay 计算重新平衡的方式
// 前置条件是已经失衡
// AVL重新平衡的方式是 LL RR LR RL
func (avl *AVL) balanceWay(checkNode *AVL) string {
	parent := avl.Parent
	grandpa := parent.Parent

	if grandpa == checkNode {
		if grandpa.Left == parent {
			if avl == parent.Left {
				return "LL"
			}
			return "LR"
		}
		if avl == parent.Left {
			return "RL"
		}
		return "RR"

	}

	return avl.Parent.balanceWay(checkNode)
}

// balance 将失衡的二叉树平衡掉
func (avl *AVL) balance(newNode, checkNode *AVL) *AVL {
	switch newNode.balanceWay(checkNode) {
	case "LL":
		return avl.LL(checkNode)
	case "RR":
		return avl.RR(checkNode)
	case "LR":
		return avl.LR(checkNode)
	case "RL":
		return avl.RL(checkNode)
	default:
		fmt.Println("unknown balance way")
	}

	return avl
}

// LL 右旋
// 以失衡点为视角
func (avl *AVL) LL(checkNode *AVL) *AVL {
	parent := checkNode.Parent
	leftSon := checkNode.Left
	leftSonRightSon := leftSon.Right

	leftSon.Parent = parent
	checkNode.Parent = leftSon
	leftSon.Right = checkNode
	if leftSonRightSon != nil {
		checkNode.Left = leftSonRightSon
		leftSonRightSon.Parent = checkNode
	}
	if parent != nil {
		parent.Left = leftSon
	}

	// udpate height
	checkNode.height++
	if checkNode.Right != nil {
		checkNode.Right.heightInc()
	}
	leftSon.height--
	leftSon.Left.heightDec()

	return leftSon.root()
}

// RR 左旋
func (avl *AVL) RR(checkNode *AVL) *AVL {
	parent := checkNode.Parent
	rightSon := checkNode.Right
	rightSonLeftSon := rightSon.Left

	rightSon.Parent = parent
	checkNode.Parent = rightSon
	rightSon.Left = checkNode
	if rightSonLeftSon != nil {
		checkNode.Right = rightSonLeftSon
		rightSonLeftSon.Parent = checkNode
	}
	if parent != nil {
		parent.Right = rightSon
	}

	// udpate height
	checkNode.height++
	if checkNode.Left != nil {
		checkNode.Left.heightInc()
	}
	rightSon.height--
	rightSon.Right.heightDec()

	return rightSon.root()
}

// LR 先左旋 再右旋
func (avl *AVL) LR(checkNode *AVL) *AVL {
	parent := checkNode.Parent
	leftSon := checkNode.Left
	leftSonRightSon := leftSon.Right
	leftSonRightSonLeftSon := leftSonRightSon.Left
	leftSonRightSonRightSon := leftSonRightSon.Right

	leftSonRightSon.Parent = parent
	leftSonRightSon.Left = leftSon
	leftSonRightSon.Right = checkNode
	leftSon.Parent = leftSonRightSon
	checkNode.Parent = leftSonRightSon
	if leftSonRightSonLeftSon != nil {
		leftSonRightSonLeftSon.Parent = leftSon
		leftSon.Right = leftSonRightSonLeftSon
	}
	if leftSonRightSonRightSon != nil {
		leftSonRightSonRightSon.Parent = checkNode
		checkNode.Left = leftSonRightSonRightSon
	}
	if parent != nil {
		parent.Left = leftSonRightSon
	}

	// update height
	checkNode.height++
	if checkNode.Right != nil {
		checkNode.Right.heightInc()
	}
	leftSonRightSon.height -= 2
	if leftSonRightSonLeftSon != nil {
		leftSonRightSonLeftSon.heightDec()
	}
	if leftSonRightSonRightSon != nil {
		leftSonRightSonRightSon.heightDec()
	}

	return leftSonRightSon.root()
}

// RL 先右旋 再左旋
func (avl *AVL) RL(checkNode *AVL) *AVL {
	parent := checkNode.Parent
	rightSon := checkNode.Right
	rightSonLeftSon := rightSon.Left
	rightSonLeftSonLeftSon := rightSonLeftSon.Left
	rightSonLeftSonRightSon := rightSonLeftSon.Right

	rightSonLeftSon.Parent = parent
	rightSonLeftSon.Left = checkNode
	rightSonLeftSon.Right = rightSon
	checkNode.Parent = rightSonLeftSon
	rightSon.Parent = rightSonLeftSon

	if rightSonLeftSonLeftSon != nil {
		rightSonLeftSonLeftSon.Parent = checkNode
		checkNode.Right = rightSonLeftSonLeftSon
	}
	if rightSonLeftSonRightSon != nil {
		rightSonLeftSonRightSon.Parent = rightSon
		rightSon.Left = rightSonLeftSonRightSon
	}
	if parent != nil {
		parent.Right = rightSonLeftSon
	}

	// update height
	checkNode.height++
	if checkNode.Left != nil {
		checkNode.Left.heightInc()
	}
	rightSonLeftSon.height -= 2
	if rightSonLeftSonLeftSon != nil {
		rightSonLeftSonLeftSon.heightDec()
	}
	if rightSonLeftSonRightSon != nil {
		rightSonLeftSonRightSon.heightDec()
	}

	return rightSonLeftSon.root()
}

func (avl *AVL) heightInc() {
	avl.height++
	if avl.Left != nil {
		avl.Left.heightInc()
	}
	if avl.Right != nil {
		avl.Right.heightInc()
	}
}

func (avl *AVL) heightDec() {
	avl.height--
	if avl.Left != nil {
		avl.Left.heightDec()
	}
	if avl.Right != nil {
		avl.Right.heightDec()
	}
}

func (avl *AVL) root() *AVL {
	if avl.Parent == nil {
		return avl
	}

	return avl.Parent.root()
}

func (avl *AVL) print() {
	fmt.Println("root:", avl.data)
	avl.printSub()
	fmt.Println()
}

func (avl *AVL) printSub() {
	if avl.Left != nil {
		avl.Left.printSub()
	}
	fmt.Print(avl.data, " ")
	if avl.Right != nil {
		avl.Right.printSub()
	}
}

func main() {
	var newNode, checkNode *AVL
	node := AVL{data: 1}
	root := &node

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		newNode = node.add(rand.Intn(50))
		checkNode = newNode.check()
		if checkNode != nil {
			root = root.balance(newNode, checkNode)
		}
		root.printSub()
		time.Sleep(time.Second)
	}
}
