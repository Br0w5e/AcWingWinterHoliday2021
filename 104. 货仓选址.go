package main

import (
	"fmt"
	"sort"
)

func main()  {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	sort.Ints(arr)
	res := 0
	for i := 0; i < n; i++ {
		res += abs(arr[i]-arr[n/2])
	}
	fmt.Println(res)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
