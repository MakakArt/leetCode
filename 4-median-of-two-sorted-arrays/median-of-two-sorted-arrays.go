func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m
    
	totalLeft := (m + n + 1) / 2

	for left <= right {
		cut1 := (left + right) / 2
		cut2 := totalLeft - cut1

		var l1, r1 float64
		if cut1 == 0 { l1 = math.Inf(-1) } else { l1 = float64(nums1[cut1-1]) }
		if cut1 == m { r1 = math.Inf(1) }  else { r1 = float64(nums1[cut1]) }

		var l2, r2 float64
		if cut2 == 0 { l2 = math.Inf(-1) } else { l2 = float64(nums2[cut2-1]) }
		if cut2 == n { r2 = math.Inf(1) }  else { r2 = float64(nums2[cut2]) }

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 0 {
				return (max(l1, l2) + min(r1, r2)) / 2
			}
			return max(l1, l2)
		}

		if l1 > r2 {
			right = cut1 - 1
		} else {
			left = cut1 + 1
		}
	}
	return 0
}
