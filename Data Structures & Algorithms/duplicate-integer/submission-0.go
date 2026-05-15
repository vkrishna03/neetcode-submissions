import "slices"

func hasDuplicate(nums []int) bool {
    slices.Sort(nums)
    
    for idx:= 1; idx < len(nums); idx++ {
        if nums[idx - 1] == nums[idx] {
            return true
        }
    } 

    return false
}
