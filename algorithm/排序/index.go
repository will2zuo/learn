package main

func main() {

}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// QuickSort 快速排序（Lomuto分区方案）
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortLomuto(arr, 0, len(arr)-1)
}

func quickSortLomuto(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partitionLomuto(arr, low, high)
	quickSortLomuto(arr, low, p-1)
	quickSortLomuto(arr, p+1, high)
}

func partitionLomuto(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// MergeSort 归并排序（原地合并，但需额外空间）
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	merge(arr, mid)
}

func merge(arr []int, mid int) {
	left := make([]int, mid)
	copy(left, arr[:mid])
	right := arr[mid:]
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
