// AVL二叉树,发明者: G. M. Adelson-Velsky和E. M. Landis
// 高度平衡树,任何节点的两子树的高度最大差别为1
// 递归版本
// todo:
// 1: 线可以优化一下，最底层的线和其他线的计算可以合并一下
// 2: pretty print 单位从1增加到2 不使用float64
// 3: 整个AVL代码的结构还可以整理一下，现在的版本有些随意

package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	pos    float64
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
	leftDepth, rightDepth := checkNode.height, checkNode.height
	if checkNode.Left != nil {
		leftDepth = checkNode.Left.depth()
	}
	if checkNode.Right != nil {
		rightDepth = checkNode.Right.depth()
	}
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
	} else {
		checkNode.Left = nil
	}
	if parent != nil {
		if parent.Left == checkNode {
			parent.Left = leftSon
		} else {
			parent.Right = leftSon
		}
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
	} else {
		checkNode.Right = nil
	}
	if parent != nil {
		if parent.Left == checkNode {
			parent.Left = rightSon
		} else {
			parent.Right = rightSon
		}
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
	} else {
		leftSon.Right = nil
	}
	if leftSonRightSonRightSon != nil {
		leftSonRightSonRightSon.Parent = checkNode
		checkNode.Left = leftSonRightSonRightSon
	} else {
		checkNode.Left = nil
	}
	if parent != nil {
		if parent.Left == checkNode {
			parent.Left = leftSonRightSon
		} else {
			parent.Right = leftSonRightSon
		}
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
	} else {
		checkNode.Right = nil
	}
	if rightSonLeftSonRightSon != nil {
		rightSonLeftSonRightSon.Parent = rightSon
		rightSon.Left = rightSonLeftSonRightSon
	} else {
		rightSon.Left = nil
	}
	if parent != nil {
		if parent.Left == checkNode {
			parent.Left = rightSonLeftSon
		} else {
			parent.Right = rightSonLeftSon
		}
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
	fmt.Print("[", avl.data, ",", avl.pos, "] ")
	if avl.Right != nil {
		avl.Right.printSub()
	}
}

// consolePrint 打印二叉树
// 可自行扩展打印格式
func (avl *AVL) consolePrint() {
	fmt.Println()
	fmt.Println("============= AVL ===============")
	s := make([]*AVL, 1)
	s[0] = avl.root()

	h := avl.root().depth() + 1
	for i := 0; i < h; i++ {
		len := len(s)
		printed := 0.0
		for _, node := range s {
			printed = node.show(printed)
		}
		fmt.Println()

		// 划线
		// printed = 0.0
		// for _, node := range s {
		//     printed = node.showLine1(printed)
		// }
		// fmt.Println()

		printed = 0.0
		for _, node := range s {
			printed = node.showLine2(printed)
		}
		fmt.Println()

		for _, node := range s {
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		s = s[len:]
	}

	fmt.Println("=================================")
	fmt.Println()
}

const halfItemSpace string = "   "
const itemSpace string = "      "

// show 显示一个节点
// 一个节点的显示方式是 [abcd] 共6位, 可根据实际调整
func (avl *AVL) show(printed float64) float64 {
	// show space
	for i := 0; i < int(avl.pos-printed); i++ {
		fmt.Print(itemSpace)
	}
	if printed == 0 && avl.height != avl.root().depth() {
		fmt.Print(halfItemSpace)
	}

	// show data, 后面可以优化为居中显示
	{
		switch len(strconv.Itoa(avl.data)) {
		case 1:
			fmt.Print("[ ", avl.data, "  ]")
		case 2:
			fmt.Print("[ ", avl.data, " ]")
		case 3:
			fmt.Print("[ ", avl.data, "]")
		case 4:
			fmt.Print("[", avl.data, "]")
		default:
			fmt.Print("[abcd]")
		}
	}

	return avl.pos + 1
}

const line1 string = "     |"

func (avl *AVL) showLine1(printed float64) float64 {
	// |
	if avl.Left == nil && avl.Right == nil {
		return 0.0
	}

	for i := 0; i < int(avl.pos-printed); i++ {
		fmt.Print(itemSpace)
	}
	fmt.Print(line1)

	return float64(int(avl.pos) + 1)
}

const line2 string = "  -"
const line3 string = "------"
const line4 string = "---"

func (avl *AVL) showLine2(printed float64) float64 {
	// -
	newPrinted := printed
	if avl.height != avl.root().depth()-1 {
		if avl.Left != nil {
			for i := 0; i < int(avl.Left.pos+0.5-newPrinted); i++ {
				fmt.Print(itemSpace)
			}
			if avl.height == avl.root().depth()-1 {
				// if float64(int(avl.Left.pos-newPrinted)) != avl.Left.pos-newPrinted {
				fmt.Print(halfItemSpace)
				// }
			}
			for i := 0; i < int(avl.pos-avl.Left.pos); i++ {
				fmt.Print(line3)
			}
			if avl.height == avl.root().depth()-1 {
				fmt.Print(line4)
			}
			newPrinted = avl.pos + 0.5
		}
		if avl.Right != nil {
			if avl.Left == nil {
				for i := 0; i < int(avl.pos+0.5-newPrinted); i++ {
					fmt.Print(itemSpace)
				}
				if float64(int(avl.pos+0.5-newPrinted)) != avl.pos+0.5-newPrinted {
					fmt.Print(halfItemSpace)
				}
				newPrinted = avl.pos + 0.5
			}
			for i := 0; i < int(avl.Right.pos+0.5-newPrinted); i++ {
				fmt.Print(line3)
			}
			if avl.height == avl.root().depth()-1 {
				fmt.Print(line4)
			}

			newPrinted = avl.Right.pos + 0.5
		}
	} else {
		// show space
		for i := 0; i < int(avl.pos-newPrinted); i++ {
			fmt.Print(itemSpace)
		}
		if printed == 0 && avl.height != avl.root().depth() {
			fmt.Print(halfItemSpace)
		}

		if avl.Left != nil {
			fmt.Print("---")
		} else {
			fmt.Print("   ")
		}
		if avl.Right != nil {
			fmt.Print("---")
		} else {
			fmt.Print("   ")
		}

		newPrinted = avl.pos + 1

	}

	return newPrinted
}

func pow2(y int) (ret int) {
	ret = 1
	for i := 0; i < y; i++ {
		ret *= 2
	}

	return
}

// calcPos 计算打印位置
// 层次遍历
func (avl *AVL) calcPos() {
	interval := pow2(avl.root().depth()-avl.height) / 2
	if avl.Parent == nil {
		avl.pos = float64(interval) - 0.5
	} else {
		if avl == avl.Parent.Left {
			avl.pos = avl.Parent.pos - float64(interval)
			if interval == 0 {
				avl.pos -= 0.5
			}
		} else {
			avl.pos = avl.Parent.pos + float64(interval)
			if interval == 0 {
				avl.pos += 0.5
			}
		}
	}

	if avl.Left != nil {
		avl.Left.calcPos()
	}
	if avl.Right != nil {
		avl.Right.calcPos()
	}
}

func main() {
	var newNode, checkNode *AVL
	node := AVL{data: 1}
	root := &node

	for i := 0; i < 30; i++ {
		rand.Seed(time.Now().UnixNano())
		newNode = root.add(rand.Intn(50))
		checkNode = newNode.check()
		if checkNode != nil {
			root = root.balance(newNode, checkNode)
		}
		fmt.Println()
		fmt.Println("new item: ", newNode.data)
		root.calcPos()
		root.print()
		time.Sleep(time.Millisecond * 10)
	}

	root.consolePrint()
}
