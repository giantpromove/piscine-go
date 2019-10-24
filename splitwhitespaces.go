package piscine

func lent(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}

func SplitWhiteSpaces(str string) []string {
	TextToString := ""
	startvalue := false
	count_M := 0

	for ii, v := range str {
		if string(v) == " " || string(v) == "\t" || string(v) == "\n" {
			if startvalue == true {
				count_M++
				startvalue = false
			}
		} else {
			if startvalue == false {
				startvalue = true
			}
		}
		if ii == lent(str)-1 && startvalue == true {
			count_M++

		}
	}
	t := make([]string, count_M)
	k := 0
	for i, v := range str {
		if i == lent(str)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
			t[k] = TextToString
		} else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString = TextToString + string(v)
		} else {
			if lent(TextToString) >= 1 {
				t[k] = TextToString
				k = k + 1
			}
			TextToString = ""
		}

	}
	return t
}
