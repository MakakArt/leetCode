func searchMatrix(matrix [][]int, target int) bool {
    l := 0
    h := len(matrix) - 1
    for l <= h {
        m := l+(h-l)/2
        if matrix[m][0] > target {
            h = m-1
        } else if matrix[m][len(matrix[m])-1] < target {
            l = m+1
        } else {
            return classicBinarySearch(matrix[m], target)
        }
    }
    return false
}

func classicBinarySearch(arr []int, t int) bool {
    l := 0
    h := len(arr) - 1
    for l <= h {
        m := l+(h-l)/2
        if arr[m] > t {
            h = m - 1
        } else if arr[m] < t {
            l = m + 1
        } else {
            return true
        }
    }
    return false
}
