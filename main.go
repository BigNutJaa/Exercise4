package main

import (
	"fmt"
	"sort"
	"strconv"
)

var number string

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	for {
		fmt.Println("Enter some number:")
		fmt.Scanln(&number)
		if _, err := strconv.ParseFloat(number, 64); err != nil {
			fmt.Printf("%q is not a number.\n", number)
		} else {
			break

		}
	}

	newNumber := []rune(number)
	sort.Slice(newNumber, func(i int, j int) bool { return newNumber[i] < newNumber[j] })
	fmt.Println("Input number =", number)
	fmt.Println("New number(sort asc) =", string(newNumber))

}
