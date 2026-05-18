func sortColors(nums []int)  {
    var col [3]int
    for _, r := range nums {
        switch r{
            case 0:
                col[0]++
            case 1:
                col[1]++
            case 2:
                col[2]++
        }
    }
    for i := 0; i < col[0]; i++ {
        nums[i] = 0
    }
    for i := col[0]; i < col[0] + col[1]; i++ {
        nums[i] = 1
    }
    for i := col[0] + col[1]; i < len(nums); i++ {
        nums[i] = 2
    }
}