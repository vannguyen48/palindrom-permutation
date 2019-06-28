package main

import "fmt"

func IsPalindromeOrPermutation(input string) bool {
	counts := make(map[rune]int)
	inputLen := 0
	for _, r := range input {
		// keeping track of how many times a letter appears.
		counts[r]++
		inputLen++
	}
	// Check to see if the word has odd number letters.
	odd := inputLen%2 == 1
	seenOdd := false
	for _, val := range counts {
		// when a letter count is an odd number.
		if val%2 == 1 {
			if !seenOdd && odd {
				seenOdd = true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	// palindrome, odd number word
	tacocat := IsPalindromeOrPermutation("tacocat")
	//palindrome, even number word
	hannah := IsPalindromeOrPermutation("hannah")
	// not a palindrome
	falala := IsPalindromeOrPermutation("falala")
	// palindrome, even number word
	nanan := IsPalindromeOrPermutation("nanan")
	fmt.Println(tacocat, hannah, falala, nanan)
}