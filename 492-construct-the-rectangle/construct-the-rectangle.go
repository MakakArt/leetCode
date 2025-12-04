func constructRectangle(area int) []int {
    result := []int{area, 1}
    for i := 1; i <= int(math.Sqrt(float64(area))); i++ {
        if area%i == 0 {
            result[0] = area/i
            result[1] = i
        }
    }
    return result
}