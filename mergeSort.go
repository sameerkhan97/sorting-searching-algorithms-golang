package main

import "fmt"

func main() {
	arr := []int{45, 21, 35, 78, 64, 01, 0, 74}
	fmt.Println("Before Sort : ", arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("After Sort : ", arr)
}
func mergeSort(arr []int, l, u int) {
	var mid int
	if l < u {
		mid = l + int((u-l)/2)
		mergeSort(arr, l, mid)
		mergeSort(arr, mid+1, u)
		merge(arr, l, mid, u)

	}
}
func merge(arr []int, l, mid, u int) {
	i := l
	j := mid + 1
	k := l
	b := make([]int, u+1)
	for i <= mid && j <= u {
		if arr[i] <= arr[j] {
			b[k] = arr[i]
			i++
		} else {
			b[k] = arr[j]
			j++
		}
		k++
	}
	if i <= mid {
		for i <= mid {
			b[k] = arr[i]
			k++
			i++
		}
	}

	if j <= u {
		for j <= u {
			b[k] = arr[j]
			k++
			j++
		}
	}

	for k := l; k <= u; k++ {
		arr[k] = b[k]
	}
}
