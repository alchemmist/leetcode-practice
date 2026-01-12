//go:build ignore

package main

import (
	"fmt"
)

func frequencyArray(s string) [26]int {
	var result [26]int
	for _, chr := range s {
		result[chr-'a'] += 1
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, s := range strs {
		sS := frequencyArray(s)
		groups[sS] = append(groups[sS], s)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
