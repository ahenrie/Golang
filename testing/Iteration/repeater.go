package iteration

func Repeat(letter string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result = result + letter
	}
	return result
}
