package main

import "math"

func checkAlmostEquivalent(word1 string, word2 string) bool {
    freq1 := make([]int, 26)
    freq2 := make([]int, 26)

    for _, ch := range word1 {
        freq1[ch-'a']++
    }

    for _, ch := range word2 {
        freq2[ch-'a']++
    }

    for i := 0; i < len(freq1); i++ {
        if math.Abs(float64(freq1[i] - freq2[i])) > 3 {
            return false
        }
    }

    return true
}
