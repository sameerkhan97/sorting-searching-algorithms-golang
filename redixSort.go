package main

import "fmt"

func main() {
	arr := []int{45, 54, 56, 34, 364, 416, 41, 0, 545, 1000}
	fmt.Println("Before Sort : ", arr)
	redixSort(arr)
	fmt.Println("After Sort : ", arr)
}
func redixSort(arr []int) {
	max := 0
	for _, i := range arr {
		if max < i {
			max = i
		}
	}
	for exp := 1; max/exp > 0; exp *= 10 {
		count := [10]int{0}	//defining count array of size 10 bcoz digit vary from 0-9
		dummy := make([]int, len(arr))2
		for i := 0; i < len(arr); i++ {
			count[(arr[i]/exp)%10]++	//value at a[i] will be devided and particular digit of value will be used as an index in count incrementing it to 1. 
		}
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]	//values at count array will be incremented by adding previous element values.
		}
		for i := len(arr) - 1; i >= 0; i-- {
			//value at a[i] will be devided and particular digit of value will be used as an index in count value at that index of count
			//... will be decremented by one and used as index of array b,value at ith index in a will be copied at that index of b.
			count[(arr[i]/exp)%10]--	
			dummy[count[(arr[i]/exp)%10]] = arr[i]
		}
		for i := 0; i < len(arr); i++ {
			arr[i] = dummy[i]
		}
	}
}
