package main

/*
leetcode: https://leetcode.cn/problems/spiral-matrix-ii/
代码随想录：https://www.bilibili.com/video/BV1SL4y1N7mV
题目分析：
	循环不变量：每一条边的第一个位置处理，最后一个位置留个下一个边处理。
	这样保证了所有的逻辑都是一样的。左闭右开的处理逻辑。这样一圈一圈的去处理。
	圈数：n/2  n是奇数的时候，最后一个位置是中心点，不用处理。直接赋值就行了。
	时间复杂度：O(n^2)
*/

func generateMatrix(n int) [][]int {
	count := 1   // 计数器，到 n*n 的时候就结束
	x, y := 0, 0 // x, y 是坐标
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	loop := n / 2            // 圈数
	offset := n              // 每一圈的边长
	for ; loop > 0; loop-- { // 处理一圈的逻辑
		// 一圈一圈的处理
		for i := 0; i < offset-1; i++ {
			matrix[x][y] = count
			count++
			y++ // 这里是y++ , 应为二维数组，x是行
		}
		for j := 0; j < offset-1; j++ {
			matrix[x][y] = count
			count++
			x++
		}
		for i := 0; i < offset-1; i++ {
			matrix[x][y] = count
			count++
			y--
		}
		for j := 0; j < offset-1; j++ {
			matrix[x][y] = count
			count++
			x--
		}
		offset -= 2 // 每一圈的边长减少2
		y++         // 每一圈的起始位置向右下角移动一位
		x++

	}

	if n%2 == 1 { // 如果 n 是奇数，最后一个位置是中心点，直接赋值就行了
		matrix[x][y] = count
	}

	return matrix
}

func main() {
	generateMatrix(3)
}
