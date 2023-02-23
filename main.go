package main

func main() {
	head := ListNode{1, nil}
	tmp := &head
	for i := 2; i < 6; i++ {
		tmpNode := ListNode{i, nil}
		tmp.Next = &tmpNode
		tmp = &tmpNode
	}
	println(reverseList(&head).Val)
}
