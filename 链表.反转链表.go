package main

/*
*
leetcode: https://leetcode.cn/problems/reverse-linked-list/description/
bilibili: https://www.bilibili.com/video/BV1nB4y1i7eL/?spm_id_from=333.788&vd_source=4c86ec124a4d3d914ab42f01ff9d4d4c
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	current := head

	for current != nil {
		// 缓存数据
		next := current.Next
		// 修改指向
		current.Next = pre
		// 移动
		pre = current
		current = next

	}

	return pre
}
func main() {
}
