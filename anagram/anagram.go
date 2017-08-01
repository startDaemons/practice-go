package anagram

import (
	s "strings"
)

func FindAnagrams(dictionary []string, _word string) []string {

	result := make([]string, 0)

	filterWord := func (word string) string {return s.ToLower(s.Trim(word, " "))}

	word := filterWord(_word)

	// check
	if word == "" {
		return result 
	}

	// compare occurances of letters in word with those in each dictionary word.
	// If the count of each letter matches and the number of unique letters match
	// then the words are anagrams of each other.

	wOccurances := getOccurances(word)
	
	for _, _dword := range dictionary {

		dword := filterWord(_dword)

		if dword == "" || word == dword {
			continue
		}

		dOccurances := getOccurances(dword)

		if !occurancesMatch(wOccurances, dOccurances) {
			continue
		}	

		result = append(result, _dword)
	}

	return result
}

func occurancesMatch(wOccurances map[string]int, dOccurances map[string]int) bool {
	if len(wOccurances) != len(dOccurances) {
		return false
	}

	for char, count := range wOccurances {
		if count != dOccurances[char] {
			return false
		}
	}

	return true
}

// Retuns a map where the keys are unique letters in the word and 
// values are the number of occurances of that letter in the word  
func getOccurances(word string) map[string]int {
	occurances := make(map[string]int)

	for _, _char := range word {
		char := s.ToUpper(string(_char)) // this seems to filter out non alpha characters- it's a bit sketchy
		
		if string(char) == "" || occurances[string(char)] != 0 || string(char) == " " {
			continue
		}

		occurances[string(char)] = s.Count(s.ToUpper(word), string(char))
	}

	return occurances
}