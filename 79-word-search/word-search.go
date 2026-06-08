func exist(board [][]byte, word string) bool {
    n := len(board) - 1
    m := len(board[0]) - 1
    var r func(s, x, y int, been []bool) bool
    r = func(s, x, y int, been []bool) bool {
        if s == len(word) {
            return true
        }
        if x < 0 || x > n || y < 0 || y > m || been[encode(x, y, n+1)]{
            return false
        }
        if word[s] == board[x][y] {
            been[encode(x, y, n+1)] = true
            s++
            if r(s, x+1, y, been) || r(s, x, y-1, been) || r(s, x-1, y, been) || r(s, x, y+1, been) {
                return true
            }
            been[encode(x, y, n+1)] = false
        }
        return false
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] {
                been := make([]bool, encode(n, m, n + 1) + 1)
                if r(0, i, j, been) { 
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

// func decode(code, w int) (int, int) {
//     return code % w, code / w
// }
