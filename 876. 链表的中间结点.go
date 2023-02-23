package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	hMap := make(map[int]*ListNode)
	for i := 0; head != nil; i++ {
		hMap[i] = head
		head = head.Next
	}
	return hMap[len(hMap)/2]
}
