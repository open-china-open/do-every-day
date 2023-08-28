package main

/*
leetcode: https://leetcode.cn/problems/minimum-size-subarray-sum/
代码随想录：https://www.bilibili.com/video/BV1tZ4y1q7XE
题目分析：
1. 暴力解法：O(n^2)
2. 滑动窗口：O(n)

分析：
	因为是【连续】的子数组，所以可以使用滑动窗口来解决这个问题。一般连续的这类问题，可以考虑滑动窗口
*/

// 1. 暴力解法
func minSubArrayLen(target int, nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ { // j = i 表示从 i 开始累计求和，这样就不会重复计算了
			sum += nums[j]
			if sum >= target { // 发现符合条件的就更新 result
				subLength := j - i + 1 // j 是从 i 开始的，所以要加 1
				if result == 0 || subLength < result {
					result = subLength
				}
				break // 发现就可以 break 了，因为是找最小的
			}
		}
	}
	return result
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	result := 0
	sum := 0                         // 滑动窗口之和
	i := 0                           // 滑动窗口的左边界
	for j := 0; j < len(nums); j++ { // j 是滑动窗口的右边界
		sum += nums[j]
		for sum >= target { // 滑动窗口的左边界开始向右移动
			subLength := j - i + 1
			if result == 0 || subLength < result {
				result = subLength
			}
			sum -= nums[i] // 左边界向右移动，所以要减去左边界的值
			i++
		}
	}

	return result
}

func main() {

}
