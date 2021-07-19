func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	item_len := 2*numRows - 2
	res := make([][]string, numRows, numRows)

	for index, v := range s {

		mod := index % item_len

		if mod < numRows {
			res[mod] = append(res[mod], string(v))
		} else {
			i := numRows - (mod - numRows) - 2
			res[i] = append(res[i], string(v))
		}
	}

	var str string

	for _, arr := range res {
		for _, v := range arr {
			str += v
		}
	}

	return str
}
