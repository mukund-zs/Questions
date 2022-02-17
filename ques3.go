package main

//pascalTriangle is function to print pascal triangle upto a given number n
func pascalTriangle(n int) [][]int {
	var arr [][]int
	if n == 0 {
		return arr
	}
	dummy := []int{1}
	arr = append(arr, dummy)
	if n == 1 {
		return arr
	}
	dummy = []int{1, 1}
	arr = append(arr, dummy)
	if n == 2 {
		return arr
	}
	dummy2 := []int{}
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			if j == i-1 || j == 0 {
				dummy2 = append(dummy2, 1)
			} else {
				dummy2 = append(dummy2, dummy[j-1]+dummy[j])
			}
		}
		arr = append(arr, dummy2)
		dummy = dummy2
		dummy2 = []int{}
	}
	return arr
}
