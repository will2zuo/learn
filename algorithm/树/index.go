package main

func main() {

}

type LinkNode struct {
	Val   int
	Left  *LinkNode
	Right *LinkNode
}

func preOrder(node *LinkNode, nums []int) []int {
	if node == nil {
		return []int{}
	}
	nums = append(nums, node.Val)
	preOrder(node.Left, nums)
	preOrder(node.Right, nums)
	return nums
}
