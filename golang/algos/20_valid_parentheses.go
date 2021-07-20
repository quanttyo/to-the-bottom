func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if brackets[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
