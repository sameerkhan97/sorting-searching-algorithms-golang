package main

import "fmt"

func main() {
	arr := []int{45, 54, 56, 34, 364, -416, -41, 0, 545, 1000}
	fmt.Println("Before Sort : ", arr)
	shellSort(arr)
	fmt.Println("After Sort : ", arr)
}

func shellSort(arr []int) {
	//initially Gap is considered half of size of array
	//function will compare every value with the value present at a particular gap meaning ith element will be compared with (i+gap)th element
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := 0; i+gap < len(arr); i++ { //i will be use to compare elements with their gap element
			//if 0th element is compared with 3rd element and later 3rd element is swaped with 7th element
			// then swapped element at 3rd will be again compared with 0th element using statement i=i-gap
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
}
