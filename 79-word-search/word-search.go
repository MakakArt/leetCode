func exist(board [][]byte, word string) bool {
    n := len(board) - 1
    m := len(board[0]) - 1
    been := make([]bool, encode(n, m, n + 1) + 1)
    var r func(s, x, y int) bool
    r = func(s, x, y int) bool {
        if s == len(word) {
            return true
        }
        if x < 0 || x > n || y < 0 || y > m || been[encode(x, y, n+1)]{
            return false
        }
        if word[s] == board[x][y] {
            been[encode(x, y, n+1)] = true
            s++
            if r(s, x+1, y) || r(s, x, y-1) || r(s, x-1, y) || r(s, x, y+1) {
                been[encode(x, y, n+1)] = false
                return true
            }
            been[encode(x, y, n+1)] = false
        }
        return false
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] {
                if r(0, i, j) { 
                    return true
                }
            }
        }
    }
    return false
}

func encode(x, y, w int) int {
    return y * w + x
}

