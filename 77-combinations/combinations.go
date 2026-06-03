func combine(n int, k int) [][]int {
    var result [][]int
    var r func (start int, body []int)
    r = func(start int, body []int) {
        if start > n {
            return
        }
        for i := start; i <= n; i++ {
            r(i+1, append(body, i))
            if len(body) + 1 == k {
                res := make([]int, k)
                copy(res, append(body, i))
                result = append(result, res)
            }
        }
    }
    r(1, []int{})
    return result
}