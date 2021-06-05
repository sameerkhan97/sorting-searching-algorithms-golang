package main

import "fmt"

func main() {
	arr := []int{45, 54, 56, 34, 364, 416, 41, 0, 545, 1000}
	fmt.Println("Before Sort : ", arr)
	arr = bucketSort(arr)
	fmt.Println("After Sort : ", arr)
}

func bucketSort(arr []int) []int {
	//finding min-max to decide size of the buckets
	var min, max int
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	//total no. of buckets
	bucketSize := int(max-min)/len(arr) + 1
	buckets := make([][]int, bucketSize)
	for i := 0; i < bucketSize; i++ {
		buckets[i] = make([]int, 0)
	}
	//storing elements in each bucket according to their value
	for _, n := range arr {
		idx := int(n-min) / len(arr)
		buckets[idx] = append(buckets[idx], n)
	}
	sorted := make([]int, 0)
	//sorting each bucket and appending them into sorted array
	for _, buck := range buckets {
		if len(buck) > 0 {
			insertionSort(buck)
			sorted = append(sorted, buck...)
		}
	}
	return sorted
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
