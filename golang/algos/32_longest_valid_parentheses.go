func longestValidParentheses(s string) int {
	count, left, right := 0, 0, 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == "(" {
			left++
		}
		if c == ")" {
			right++
		}

		if left == right {
			count = int(math.Max(float64(count), float64(left+right)))
		}

		if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		c := string(s[i])
		if c == "(" {
			left++
		}
		if c == ")" {
			right++
		}

		if left == right {
			count = int(math.Max(float64(count), float64(left+right)))
		}

		if left > right {
			left, right = 0, 0
		}
	}
	return count
}
