func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)

    for _, num := range nums {
        m[num] = !m[num]    // all unique num will be true
    } 

    for _, val := range m{
        if !val {
            return true
        }
    }

    return false
}