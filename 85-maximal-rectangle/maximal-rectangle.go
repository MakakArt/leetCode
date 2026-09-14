func maximalRectangle(matrix [][]byte) int {
    var max int
    hist := make([]int, len(matrix[0]) + 1)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == '0' {
                hist[j] = 0
            } else {
                hist[j]++
            }
        }
        fmt.Println("hist", hist)
        // it stores only indeces
        var stack []int
        for n, h := range hist {
            for len(stack) > 0 && h < hist[stack[len(stack) - 1]] {
                index := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]

                height := hist[index]
                for index > 0 && height <= hist[index - 1] {
                    index--
                }
                width := n - index
                
                sq := height * width
                if sq > max {
                    max = sq
                }
            }
            stack = append(stack, n)
        }
    }
    return max
}