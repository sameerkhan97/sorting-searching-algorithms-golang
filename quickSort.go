package main

import "fmt"

func main() {
	arr := []int{45, 54, 56, 34, 364, -416, -41, 0, 545, 1000}
	fmt.Println("Before Sort : ", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After Sort : ", arr)
}

func quickSort(arr []int, l, u int) {
	if l < u {
		loc := partition(arr, l, u) //variable storing location of pivot element
		quickSort(arr, l, loc)      //calls for leftside array from pivot element
		quickSort(arr, loc+1, u)    //calls for rightside array from pivot element
	}
}
func partition(arr []int, l, u int) (end int) {
	//making 1st element as pivot element of array
	pivot := arr[l]
	start, end := l, u
	//keep iterating array and swaping elements untill start and end cross eachother
	for start < end {
		for arr[start] <= pivot { //start pointer will stop at element > pivot element
			start++
		}
		for arr[end] > pivot { //end pointer will stop at element < pivot element
			end--
		}
		if start < end { //swap element at start and end only if start and end didnt crossed each other.
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	arr[l], arr[end] = arr[end], arr[l] //pivot element will be swaped with element pointed by end
	return                              //return location of pivot element
}
