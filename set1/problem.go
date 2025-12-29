package set1

import (
	"fmt"
	"strings"
)

func ReverseWords() {

	str := "Welcome to Go"

	str1 := strings.Fields(str)
	fmt.Println(str1)
	for i, j := 0, len(str1)-1; i < j; i, j = i+1, j-1 {
		str1[i], str1[j] = str1[j], str1[i]
	}

	fmt.Println(str1)
	fmt.Println(longestWord("lets start programming go"))

}

func longestWord(s string) string {

	s1 := strings.Fields(s)
	result := " "

	for i := 0; i < len(s1); i++ {
		if len(result) < len(s1[i]) {
			result = s1[i]
		}

	}
	return result
}