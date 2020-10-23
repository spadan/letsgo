package leetcode

import (
	"strconv"
	"strings"
)

type node struct {
	val   int
	left  *node
	right *node
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res string
	queen := make([]*TreeNode, 0)
	queen = append(queen, root)
	for len(queen) > 0 {
		node := queen[0]
		queen = queen[1:]
		if node == nil {
			res += "null,"
		} else {
			res += strconv.Itoa(node.Val) + ","
			queen = append(queen, node.Left)
			queen = append(queen, node.Right)
		}
	}
	return res[:len(res)-1]
}

// 谦虚遍历
func (this *Codec) serialize1(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return buildStr(root)
}

func buildStr(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return strconv.Itoa(root.Val) + "," + buildStr(root.Left) + "," + buildStr(root.Right)
}

func (this *Codec) deserialize1(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data, ",")
	return buildTree(&nodes)
}

func buildTree(nodes *[]string) *TreeNode {
	if (*nodes)[0] == "nil" {
		return nil
	}
	v := (*nodes)[0]
	newNodes:=make([]string,len(*nodes)-1)
	copy(newNodes,(*nodes)[1:])
	root := TreeNode{Val: str2int(v)}
	root.Left = buildTree(&newNodes)
	root.Right = buildTree(nodes)
	return &root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data, ",")
	queen := make([]*TreeNode, 0)
	root := &TreeNode{Val: str2int(nodes[0])}
	queen = append(queen, root)
	i := 0
	for len(queen) > 0 {
		parent := queen[0]
		queen = queen[1:]
		i++
		if i < len(data) && nodes[i] != "null" {
			parent.Left = &TreeNode{Val: str2int(nodes[i])}
			queen = append(queen, parent.Left)
		}
		i++
		if i < len(data) && nodes[i] != "null" {
			parent.Right = &TreeNode{Val: str2int(nodes[i])}
			queen = append(queen, parent.Right)
		}
	}
	return root
}

func str2int(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
