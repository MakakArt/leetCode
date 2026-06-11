func removeDuplicates(nums []int) int {
    k := 0
    for _, x := range nums {
        if k < 2 || nums[k-2] != x {
            nums[k] = x
            k++
        }
    }
    return k
}