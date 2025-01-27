package main

func mergeAlternately(word1 string, word2 string) string {
	// Find the minimum length of the two words
	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}

	var merged []byte

	// Append alternating characters from both words
	for i := 0; i < minLen; i++ {
		merged = append(merged, word1[i], word2[i])
	}

	// Append the remaining characters from the longer word
	if len(word1) > minLen {
		merged = append(merged, word1[minLen:]...)
	} else if len(word2) > minLen {
		merged = append(merged, word2[minLen:]...)
	}

	return string(merged)
}
