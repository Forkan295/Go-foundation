package main

import (
	"fmt"
)

func main() {
	strs := []string{"caater", "caatern", "caar", "caa", "caa"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			//fmt.Println(fmt.Sprintf("%c", int(strs[i][j])))
			//fmt.Println(prefix[:j], prefix, j)
			if prefix[j] == strs[i][j] {
				//prefix = prefix[:j]
			} else {
				prefix = prefix[:j]
			}
			//fmt.Println(prefix)
		}
		if len(prefix) > len(strs[i]) {
			fmt.Println(prefix, prefix[:len(strs[i])], len(strs[i]))
			prefix = prefix[:len(strs[i])]
		}
	}

	return prefix
}
