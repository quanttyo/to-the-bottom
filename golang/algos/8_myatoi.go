func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	sig, l := 1, 0
	switch str[l] {
	case '-':
		sig = -1
		str = str[1:]
	case '+':
		str = str[1:]
		sig = 1
	default:
		sig = 1
	}

	ret := 0
	for _, v := range str {
		if v < '0' || v > '9' {
			break
		}
		ret = ret*10 + int(v-'0')
		if ret > math.MaxInt32 {
			if sig == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return ret * sig
}
