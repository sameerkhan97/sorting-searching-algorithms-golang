package main

import "fmt"

func main() {
	arr := []int{45, 54, 56, 34, 364, 416, 41, 0, 545, 1000}
	fmt.Println("Before Sort : ", arr)
	arr = heapSort(arr)
	fmt.Println("After Sort : ", arr)
}

func heapSort(data []int) []int {
	heapify(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		siftDown(data, 0, i)
	}
	return data
}

func heapify(data []int) {
	for i := (len(data) - 1) / 2; i >= 0; i-- {
		siftDown(data, i, len(data))
	}
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}

	}
}
