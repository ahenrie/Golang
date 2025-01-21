package main

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		// Check if the current numeral is smaller than the next
		if i > 0 && romanMap[rune(s[i])] > romanMap[rune(s[i-1])] {
			total += romanMap[rune(s[i])] - 2*romanMap[rune(s[i-1])]
		} else {
			total += romanMap[rune(s[i])]
		}
	}

	return total
}
