package main

import "fmt"

func main() {
	arr := []int{50, 78, 34, 65, 21}
	fmt.Println("Before Sort : ", arr)
	arr = countSort(arr)
	fmt.Println("After Sort : ", arr)
}

func countSort(arr []int) []int {
	max := 0
	for i := 0; i < len(arr); i++ { // Find the largest element of the array
		if max < arr[i] {
			max = arr[i]
		}
	}
	count := make([]int, max+1) // Initialize count array with all zeros.
	b := make([]int, len(arr))
	for i := 0; i < len(arr); i++ { // Store the count of each element
		count[arr[i]]++
	}
	for i := 1; i < len(count); i++ { // Store the cummulative count of each array element
		count[i] = count[i] + count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- { // Find the index of each element of the original array in count array, and
		b[count[arr[i]]-1] = arr[i] // place the elements in output array
		count[arr[i]] = count[arr[i]] - 1
	}
	return b
}
