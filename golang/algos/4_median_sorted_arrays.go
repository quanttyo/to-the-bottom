func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Index, nums2Index := 0, 0
	target := make([]int, 0)
	for {
		if nums1Index == len(nums1) {
			for ; nums2Index < len(nums2); nums2Index++ {
				target = append(target, nums2[nums2Index])
			}
			break
		}
		if nums2Index == len(nums2) {
			for ; nums1Index < len(nums1); nums1Index++ {
				target = append(target, nums1[nums1Index])
			}
			break
		}
		if nums1[nums1Index] < nums2[nums2Index] {
			target = append(target, nums1[nums1Index])
			nums1Index++
		} else {
			target = append(target, nums2[nums2Index])
			nums2Index++
		}
	}
	if len(target)%2 == 1 {
		return float64(target[len(target)/2])
	} else {
		return float64(target[len(target)/2]+target[len(target)/2-1]) / 2
	}
}
