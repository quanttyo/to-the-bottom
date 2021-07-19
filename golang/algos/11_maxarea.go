func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}
func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
