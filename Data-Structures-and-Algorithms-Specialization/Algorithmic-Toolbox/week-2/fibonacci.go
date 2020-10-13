package course

/**
 * Fibonacci
 * 
 * Recursive (Naive solution)
 * 
 * Table (Efficient solution)
 * - Increasing the space complexity to improve the time complexity
 * 
 * TODO: implement big int
 */

func fibonacciTable (n int) int {
	arr := make([]int, n)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}

	return arr[n - 1]
}