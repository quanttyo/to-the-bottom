func intToRoman(num int) string {
	m := []string{"", "M", "MM", "MMM"}
	c := []string{"", "C", "CC", "CCC", "CD", "D",
		"DC", "DCC", "DCCC", "CM"}
	x := []string{"", "X", "XX", "XXX", "XL", "L",
		"LX", "LXX", "LXXX", "XC"}
	i := []string{"", "I", "II", "III", "IV", "V",
		"VI", "VII", "VIII", "IX"}
	thousands := m[num/1000]
	hundereds := c[(num%1000)/100]
	tens := x[(num%100)/10]
	ones := i[num%10]

	ans := thousands + hundereds + tens + ones

	return ans
}
