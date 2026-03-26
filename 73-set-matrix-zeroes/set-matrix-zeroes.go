func setZeroes(matrix [][]int)  {
    var row bool
    columns := make([]bool, len(matrix[0]))
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                row = true
                columns[j] = true
                k := i
                for k > -1 {
                    matrix[k][j] = 0
                    k--
                }
                k = j
                for k > -1 {
                    matrix[i][k] = 0
                    k--
                }
                continue
            }
            if columns[j] || row { 
                matrix[i][j] = 0
            }
        }
        row = false
    }
}