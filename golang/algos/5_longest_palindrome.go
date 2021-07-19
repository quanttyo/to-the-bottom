func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	orig := strings.Split(s, "")
	length := len(orig)
	dest := make([]string, 0)
	lastCenter := 2*length - 1
	destLg := 0
	for counter := 1; counter <= lastCenter; counter++ {
		odd := counter%2 == 1
		tmp := make([]string, 0)
		startRight := (counter + 1) / 2
		startLeft := counter / 2
		if !odd {
			startRight = counter/2 + 1
			startLeft = counter/2 - 1
		}
		for idxL, idxR := startLeft, startRight; idxL >= 0 && idxR < length; idxL, idxR = idxL-1, idxR+1 {
			if orig[idxL] == orig[idxR] {
				tmp = append(tmp, orig[idxR])
			} else {
				break
			}
		}
		if odd {
			if len(tmp)*2 > destLg {
				dest = tmp
				destLg = len(tmp) * 2
			}
		} else {
			if len(tmp) > 0 && len(tmp)*21 > destLg {
				dest = append([]string{orig[counter/2]}, tmp...)
				destLg = len(tmp)*2 + 1
			}
		}
	}
	add := ""
	if len(dest) > 0 {
		for i := 0; i < len(dest); i++ {
			add = dest[i] + add
		}
		if destLg%2 == 1 {
			add = add[:len(add)-1]
		}
	}
	if len(dest) == 1 && destLg%2 == 1 {
		return ""
	}
	ret := add + strings.Join(dest, "")
	if len(ret) == 0 {
		return orig[0]
	}
	return ret
}+
