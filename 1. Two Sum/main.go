func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if i == j {
                continue
            }
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{0}
}
