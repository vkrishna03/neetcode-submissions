func twoSum(nums []int, target int) []int {
    i, j := 0,1
	sol := make([]int, 2) 

	// single-loop sliding window
	for i < len(nums)-1 {
		if nums[i] + nums[j] == target {
			sol[0] = i
			sol[1] = j
			break
		}

		if j == len(nums)-1 {
			i++
			j = i + 1
		} else {
            j++
        }
	}
	return sol
}
