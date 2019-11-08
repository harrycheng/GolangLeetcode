package mysolution

import "fmt"

func TwoSumTest() {
	var nums []int = []int{3, 2, 4}
	var target int = 6

	var re []int = twoSum(nums, target)
	fmt.Print(re)

}

func twoSum(nums []int, target int) []int {
	var re []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				re = append(re, i)
				re = append(re, j)
				return re
			}
		}
	}
	return re
}
