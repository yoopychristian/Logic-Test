//YOOPY CHRISTIAN - Stockbit Golang Developer

package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	kata = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func anagram() {
	list := make(map[string][]string)

	for _, word := range kata {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	anagram()
}
