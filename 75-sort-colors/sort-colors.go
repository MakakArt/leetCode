func sortColors(nums []int)  {
    l, m, h := 0, 0, len(nums) - 1
    for m <= h {
        switch nums[m] {
            case 0:
                nums[l], nums[m] = nums[m], nums[l]
                l++
                m++
            case 2:
                nums[h], nums[m] = nums[m], nums[h]
                h--
            default: 
                m++
        }
    }
}