package main

import "fmt"

func main() {
	arr := []int{31, 46, 462, 0, -46, 341}
	fmt.Println("Before Sort : ", arr)
	insertionSort(arr)
	fmt.Println("After Sort : ", arr)
}
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		//It will take element at ith inedx and compare it with its previous element first and if found smaller then will it swap with previous.
		//then it will be again compared with next previous element following same procedure untill no previous element is greater than it.
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
