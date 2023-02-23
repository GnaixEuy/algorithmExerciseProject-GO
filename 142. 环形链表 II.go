package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]*ListNode)
	for head != nil {
		v, b := mp[head.Next]
		if b {
			return v
		} else {
			mp[head] = head
			head = head.Next
		}
	}
	return nil
}
