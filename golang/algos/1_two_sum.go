func twoSum(nums []int, target int) []int {
	res := []int{0, 0}
	for i := range nums {
		for j := range nums {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
				break
			}
		}
	}
	return res

}
