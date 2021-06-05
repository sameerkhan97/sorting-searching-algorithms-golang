package main

import "fmt"

func main() {
	arr := []int{50, 78, 34, 65, 21}
	fmt.Println("Before Sort : ", arr)
	bubbleSort(arr)
	fmt.Println("After Sort : ", arr)
}
func bubbleSort(arr []int) {
	for c := 1; c < len(arr); c++ { //c will count the number of time array is checked and while loop will operate untill c is smaller than array size.
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] { //Element at ith index will be compared with next element and if found smaller then will be swapped with next one.
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
