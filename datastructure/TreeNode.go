package datastructure

import (
	"strconv"
)

// TreeNode literally indicates a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type queueData struct {
	node   *TreeNode
	isLeft bool
}

// CreateTreeNodeFromString creates a tructure of TreeNode
// from an array of string(or "null") in LeetCode expression
// ex) ["1", "2", "null", "3"]
func CreateTreeNodeFromString(numStrings []string) *TreeNode {
	if len(numStrings) == 0 {
		return nil
	}

	numString := numStrings[0]
	numStrings = numStrings[1:]

	num, err := strconv.Atoi(numString)

	if err != nil {
		return nil
	}

	root := TreeNode{num, nil, nil}

	if len(numStrings) == 0 {
		return &root
	}

	queue := []queueData{
		{&root, true},
		{&root, false},
	}

	for len(numStrings) > 0 {
		q := queue[0]
		queue = queue[1:]
		numString := numStrings[0]
		numStrings = numStrings[1:]

		num, err := strconv.Atoi(numString)

		if err != nil {
			continue
		}

		if q.isLeft {
			q.node.Left = &TreeNode{num, nil, nil}
			queue = append(
				queue,
				queueData{q.node.Left, true},
				queueData{q.node.Left, false},
			)
		} else {
			q.node.Right = &TreeNode{num, nil, nil}
			queue = append(
				queue,
				queueData{q.node.Right, true},
				queueData{q.node.Right, false},
			)
		}
	}

	return &root
}
