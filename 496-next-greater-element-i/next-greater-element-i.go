func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ge := make(map[int]int, len(nums2))
    stack := []int{nums2[0]}
    for i := 1; i < len(nums2); i++ {
        for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
            ge[stack[len(stack)-1]] = nums2[i]
            stack = stack[0:len(stack)-1]
        }
        stack = append(stack, nums2[i])
    }
    var result []int
    for _, r := range nums1 {
        e, ok := ge[r]
        if !ok {
            e = -1
        }
        result = append(result, e)
    }
    return result
}