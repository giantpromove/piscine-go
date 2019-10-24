package piscine

func ConcatParams(args []string) string {
	result := ""
	counter := 0
	for range args {
		counter++
	}
	for i := 0; i < counter; i++ {
		result += args[i]
		if i != counter-1 {
			result = result + "\n"
		}
	}
	return result
}
