package main

func main() {

}

// 最小栈

var (
	data []int
	minD []int
	size int
)

func top() int {
	return data[size-1]
}

func push(num int) {
	data[size] = num
	if len(minD) == 0 || num <= getMin() {
		minD[size] = num
	} else {
		minD[size] = minD[size-1]
	}
	size++
}

func peek() int {
	return data[size-1]
}

func pop() {
	size--
}

func getMin() int {
	return minD[size-1]
}
