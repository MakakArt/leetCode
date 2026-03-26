func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1 
        }
    }
    
    return calculateR(word1, word2, m-1, n-1, memo)
}

func calculateR(word1, word2 string, i, j int, memo [][]int) int {
    if i < 0 {
        return j + 1
    }
    if j < 0 {
        return i + 1
    }
    
    if memo[i][j] != -1 {
        return memo[i][j]
    }
    
    if word1[i] == word2[j] {
        memo[i][j] = calculateR(word1, word2, i-1, j-1, memo)
        return memo[i][j]
    }
    
    insertCost := calculateR(word1, word2, i, j-1, memo)
    deleteCost := calculateR(word1, word2, i-1, j, memo)
    replaceCost := calculateR(word1, word2, i-1, j-1, memo)

    memo[i][j] = 1 + min(insertCost, min(deleteCost, replaceCost))
    
    return memo[i][j]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}