package main

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	// 存在虚拟节点
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 {
		return -1
	}
	current := this.head
	// 0 也会进循环, 获得第一个节点
	for i := 0; i <= index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.Next = &ListNode{Val: val, Next: this.head.Next}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	current := this.head
	// 最后一个为空
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: val, Next: nil}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		return
	}
	cur := this.head
	// 获得 n - 1 节点
	for ; index > 0; index-- {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val, Next: cur.Next}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	cur := this.head
	// 获得 n - 1 节点
	for ; index > 0; index-- {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
}
