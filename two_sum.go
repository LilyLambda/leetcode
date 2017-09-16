package leetcode

func twoSum(nums []int, target int) []int {
	return twoSumQuadratic(nums, target)
}

func twoSumQuadratic(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if i != j {
                if nums[i] + nums[j] == target {
                    return []int{i, j}
                }
            }
        }
    }
    return []int{}
}
