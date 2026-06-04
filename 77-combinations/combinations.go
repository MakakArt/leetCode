func combine(n int, k int) [][]int {
    var result [][]int
    buff := make([]int, k)
    var r func (start, depth int)
    r = func(start, depth int) {
        if depth == k {
            res := make([]int, k)
            copy(res, buff)
            result = append(result, res)
        } else {
            for i := start; i <= n-(k-depth)+1; i++ {
                buff[depth] = i
                r(i+1, depth+1)
            }
        }
    }
    r(1, 0)
    return result
}