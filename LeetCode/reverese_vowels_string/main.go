package main

import (
	"fmt"
	"unicode"
)

func reverseVowels(s string) string {

    chars := []rune(s)
    vowels := []rune{}

    vowelMap := make(map[rune]bool)
    vowelMap['a'] = true
    vowelMap['e'] = true
    vowelMap['i'] = true
    vowelMap['o'] = true
    vowelMap['u'] = true

    for i := 0; i < len(chars); i++ {
        if vowelMap[unicode.ToLower(chars[i])] {
            vowels = append(vowels, (chars[i]))
        }
    }

    for i := 0; i < len(chars); i++ {
        if vowelMap[unicode.ToLower(chars[i])] {
            chars[i] = vowels[len(vowels)-1]
            vowels = vowels[:len(vowels)-1]
        }
    }

    s = string(chars)

    fmt.Println(vowels)
    return s
}
