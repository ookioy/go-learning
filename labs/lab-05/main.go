package main

import "fmt"

func Shift(arr []int, k int, n int) []int {
	if k <= 0 {
		return arr
	}

	if k >= n {
		k %= n
	}

	shifted := make([]int, n)
	for i := 0; i < n; i++ {
		ix := i + k
		if ix >= n {
			ix -= n
		}
		shifted[ix] = arr[i]
	}

	return shifted
}

func main() {
	var n, k int

	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&k)

	arr = Shift(arr, k, n)
	fmt.Println(arr)
}
