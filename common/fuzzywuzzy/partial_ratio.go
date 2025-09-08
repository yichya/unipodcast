package fuzzywuzzy

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// return the smallest string and then the bigger string
func strMinMax(s1, s2 string) (string, string) {
	if len(s1) <= len(s2) {
		return s1, s2
	}
	return s2, s1
}

// process the string to change to lowercase and remove
// things that are not letters...
func processString(s string) string {
	space := false
	var str []rune
	for _, v := range strings.ToLower(s) {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			space = false
			str = append(str, v)
		} else if !space {
			space = true
			str = append(str, ' ')
		}
	}

	return strings.TrimSpace(string(str))
}

func LevenshteinDistance(s1, s2 string) int {
	//length of the string s1, s2
	s1Len, s2Len := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2)

	//if the two strings equals
	if s1 == s2 {
		return 0
	}
	//if a string is length 0
	if s1Len == 0 {
		return s2Len
	}
	if s2Len == 0 {
		return s1Len
	}

	v0 := make([]int, s2Len+1)
	v1 := make([]int, s2Len+1)

	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < s1Len; i++ {

		v1[0] = i + 1

		for j := 0; j < s2Len; j++ {
			cost := 1
			if s1[i] == s2[j] {
				cost = 0
			}
			v1[j+1] = slices.Min([]int{v1[j] + 1, v0[j+1] + 1, v0[j] + cost})
		}

		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}

	}
	return v1[s2Len]
}

func Ratio(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	l := utf8.RuneCountInString(s1) + utf8.RuneCountInString(s2)
	dist := LevenshteinDistance(processString(s1), processString(s2))
	return int((1 - (float32(dist) / float32(l))) * 100)
}

func PartialRatio(s1, s2 string) int {
	minStr, maxStr := strMinMax(s1, s2)
	var bestRatio int
	for i := 0; i < len(maxStr)-len(minStr)+1; i++ {
		Ratio := Ratio(minStr, maxStr[i:i+len(minStr)])
		if Ratio > bestRatio {
			bestRatio = Ratio
		}
	}
	return bestRatio
}
