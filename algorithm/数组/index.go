package main

import (
	"fmt"
)

func main() {
	fmt.Println(AddOne([]int{9}))
}

// AddOne  数组加 1
/**
 * 示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。
*/
func AddOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] += 1
		nums[i] %= 10
		if nums[i] != 0 {
			return nums
		}
	}
	return append([]int{1}, nums...)
}
