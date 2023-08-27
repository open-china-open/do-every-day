package main

/*
leetcode: https://leetcode.cn/problems/remove-element/
代码随想录：https://www.bilibili.com/video/BV12A4y1Z7LP
题目分析：
	1. 暴力解法，就是遍历数组，遇到相同的就删除，但是这样的时间复杂度是O(n^2)
	2. 双指针法 （快慢指针）时间复杂度是O(n)
*/

// 暴力解法
func removeElement(nums []int, val int) int {
	size := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val { // 遇到相同的元数，后面的整体向前移动
			for j := i + 1; j < len(nums); j++ { // 这里就是整体向前移动，不过 j = i +1 是为了简单的处理，这样不用加判断条件了
				nums[j-1] = nums[j]
			}
			size--
		}
	}
	return size
}

// 双指针法
// 其实就是快指针遇到相同的元素就跳过，如果没有遇到就替换，慢指针就是记录替换的位置
func removeElement2(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {

}
