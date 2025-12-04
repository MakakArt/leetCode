func findPoisonedDuration(timeSeries []int, duration int) int {
    var count, p int
    for _, r := range timeSeries {
        if r < p {
            nP := r + duration
            if nP < p {
                continue
            }
            count += nP - p
            p = nP
        } else {
            count += duration
            p = r + duration
        }
    }
    return count
}   