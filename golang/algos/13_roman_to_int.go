func romanToInt(s string) int {
	var romans = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var val int
	letters := strings.Split(s, "")
	for i := 0; i <= (len(s) - 1); i++ {
		if i == (len(s) - 1) {
			val = val + romans[letters[i]]
		} else if romans[letters[i]] < romans[letters[i+1]] {
			val = val + romans[(letters[i]+letters[i+1])]
			i += 1
		} else {
			val = val + romans[letters[i]]
		}
	}
	return val
}
