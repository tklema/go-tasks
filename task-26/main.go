package main

import (
	"fmt"
	"strings"
)

func isAllUniqueByte(str string) bool {
	letters := make(map[byte]bool)

	strLower := strings.ToLower(str)

	for i := range len(strLower) {
		letter := strLower[i]
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}

	return true
}

func isAllUniqueRune(str string) bool {
	letters := make(map[rune]bool)

	strLower := strings.ToLower(str)

	for _, elem := range strLower {
		if letters[elem] {
			return false
		}
		letters[elem] = true
	}

	return true
}

func main() {

	s1 := "abcd"
	s2 := "abCdefA"
	s3 := "aabcd"

	fmt.Printf("(byte) %s: %v \n", s1, isAllUniqueByte(s1))
	fmt.Printf("(byte) %s: %v \n", s2, isAllUniqueByte(s2))
	fmt.Printf("(byte) %s: %v \n", s3, isAllUniqueByte(s3))

	fmt.Printf("(rune) %s: %v \n", s1, isAllUniqueRune(s1))
	fmt.Printf("(rune) %s: %v \n", s2, isAllUniqueRune(s2))
	fmt.Printf("(rune) %s: %v \n", s3, isAllUniqueRune(s3))

}
