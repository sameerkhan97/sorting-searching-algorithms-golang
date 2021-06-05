package main

import "fmt"

func main() {
	arr := []int{31, 46, 462, 0, -46, 341}
	fmt.Println("Before Sort : ", arr)
	selectionSort(arr)
	fmt.Println("After Sort : ", arr)
}
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min, pos := arr[i], i //min-store smallest value in array ,pos-store location of minimum value
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min { //finding smallest element on right side of comparing element in Array
				min, pos = arr[j], j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i] //Smallest element in right side of array will be swapped with the comparing element
	}
}
