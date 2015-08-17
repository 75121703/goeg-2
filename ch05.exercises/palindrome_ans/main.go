package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"
)

// IsPalindrome ...
var IsPalindrome func(string) bool

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "-a" || os.Args[1] == "--ascii") {
		os.Args = append(os.Args[:1], os.Args[2:]...) // Strip out arg.
		IsPalindrome = func(s string) bool {          // Simple ASCII-only version
			j := len(s) - 1
			for i := 0; i < len(s)/2; i++ {
				if s[i] != s[j] {
					return false
				}
				j--
			}
			return true
		}
	} else {
		IsPalindrome = func(s string) bool { // UTF-8 version
			for len(s) > 0 {
				first, sizeOfFirst := utf8.DecodeRuneInString(s)
				if sizeOfFirst == len(s) {
					break // s only has one character
				}
				last, sizeOfLast := utf8.DecodeLastRuneInString(s)
				if first != last {
					return false
				}
				s = s[sizeOfFirst : len(s)-sizeOfLast]
			}
			return true
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s [-a|--ascii] word1 [word2 [... wordN]]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	words := os.Args[1:]
	for _, word := range words {
		fmt.Printf("%5t %q\n", IsPalindrome(word), word)
	}
}
