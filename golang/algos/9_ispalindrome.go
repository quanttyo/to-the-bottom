func isPalindrome(x int) bool {
	var n int
	tmp := x
	for tmp > 0 {
		n = n*10 + tmp%10
		tmp /= 10
	}
	return n == x
}
