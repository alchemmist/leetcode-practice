//go:build ignore

package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		sS := SortString(s)
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
