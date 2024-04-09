package main

/*
leetcode: https://leetcode.cn/problems/binary-search/
代码随想录：https://www.bilibili.com/video/BV1fA4y1o715/
题目很容易写成 O(n) 的时间复杂度，但是这样就没有意义了，因为题目给出的数组是有序的，所以我们可以使用二分法来解决这个问题。
题目分析：
 1. 二分法查找一定要是有序的
 2. 二分法的时间复杂度是O(logn)

二分法有两种写法：
1. 左闭右闭 (while left <= right)
2. 左闭右开 (while left < right)
*/

/*
左闭右闭
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 注意这里是 <=
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 注意这里是 <
			left = mid + 1
		} else if nums[mid] > target { // 注意这里是 >
			right = mid - 1
		}
	}
	return -1
}

// 左闭右开
func search2(nums []int, target int) int {
	// right 应该是 len(nums) 而不是 len(nums)-1 应为右边已经不包含了。
	left, right := 0, len(nums)
	for left < right { // 注意这里是 <
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 注意这里是 <
			left = mid + 1
		} else if nums[mid] > target { // 注意这里是 >
			right = mid
		}
	}
	return -1
}

func main() {
	println("11")
}
