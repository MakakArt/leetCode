func subsets(nums []int) [][]int {
    var result [][]int
    buff := make([]int, len(nums))
    var r func(depth, curr int)
    r = func(depth, curr int) {
        res := make([]int, curr)
        copy(res, buff)
        result = append(result, res)
        for i := depth; i < len(nums); i++ {
            buff[curr] = nums[i]
            r(i+1, curr+1)
        }
    }
    r(0, 0)
    return result
}