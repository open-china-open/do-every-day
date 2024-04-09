package main

/*
递归法:
f(n) = f(n-1) + f(n-2)
f(0) = 0
f(1) = 1
*/
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
迭代法: 循环遍历
时间复杂度: O(n)
*/
func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}
	f0, f1, s := 0, 1, 0
	for i := 2; i <= n; i++ {
		s = f0 + f1
		f0 = f1
		f1 = s
	}
	return s
}

func main() {
	println(fibonacci(10))
	println(fibonacci2(10))
}
