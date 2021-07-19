func reverse(x int) int {
	var MaxInt int32 = 2147483647

	if x > int(MaxInt) || x < -int(MaxInt) {
		return 0
	}

	temp := 0
	run := x
	if x < 0 {
		run = -run
	}
	for run > 0 {
		temp *= 10
		digit := run % 10
		temp += digit
		run /= 10
	}

	if temp > int(MaxInt) {
		return 0
	}
	if x < 0 {
		return -temp
	}
	return temp
}
