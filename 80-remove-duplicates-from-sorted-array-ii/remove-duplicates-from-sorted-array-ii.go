func removeDuplicates(nums []int) int {    
    count := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            if count > 2 {
                nums = append(nums[:i-count+2], nums[i:]...)
                i = i-count+1
            }
            count = 1
        } else {
            count++
        }
    }
    if count > 2 {
        nums = nums[:len(nums)-(count-2)]
    }
    return len(nums)
}