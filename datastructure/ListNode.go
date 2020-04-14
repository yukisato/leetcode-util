package datastructure

// ListNode is a singly-linked list data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListNode creates chain of ListNode from []int
func CreateListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{nums[0], &ListNode{-1, nil}}
	node := head

	for _, num := range nums[1:] {
		node = node.Next
		node.Val = num
		node.Next = &ListNode{-1, nil}
	}

	node.Next = nil

	return head
}
