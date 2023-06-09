package main

import (
	"fmt"
)

var heap []int

func minHeap(arr []int) {
	buildHeap(arr)
}
func buildHeap(arr []int) {
	heap = arr
	for i := parent(len(heap) - 1); i >= 0; i-- {
		shiftDown(i)
	}
}
func shiftDown(currentIdx int) {
	endIdx := len(heap) - 1

	for currentIdx >= 0 {
		leftIdx := leftCHild(currentIdx)
		rightIdx := rightChild(currentIdx)

		idxToShift := leftIdx
		if rightIdx <= endIdx && heap[rightIdx] > heap[leftIdx] {
			idxToShift = rightIdx
		}else{
			idxToShift = leftIdx
		}

		if idxToShift <= endIdx && heap[idxToShift] > heap[currentIdx] {
			swap(heap, idxToShift, currentIdx)
			currentIdx = idxToShift
		} else {
			return
		}
	}
}


func peak() {
	fmt.Println(heap[0])
}



func shiftUp(currentIdx int) {
	parentIdx := parent(currentIdx)
	for currentIdx > 0 && heap[parentIdx] > heap[currentIdx] {
		swap(heap, currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = parent(currentIdx)
	}
}
func remove() {
	swap(heap, 0, len(heap)-1)
	heap = heap[:len(heap)-1]
	shiftDown(0)
}

func insert(value int) {
	heap = append(heap, value)
	shiftUp(len(heap) - 1)
}

func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func parent(i int) int {
	return (i - 1) / 2
}
func leftCHild(i int) int {
	return (i * 2) + 1
}
func rightChild(i int) int {
	return (i * 2) + 2
}

func main() {
	heap = []int{4, 2, 3, 1, 7, 6, 5}
	display(heap)
	fmt.Println()
	buildHeap(heap)
	// insert(1)
	// insert(5)
	// insert(7)
	// display(heap)
	// remove()
	fmt.Println()
	display(heap)
}
