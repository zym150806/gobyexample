package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b", "d"}
	sort.Strings(strs)
	fmt.Println("strings: ", strs)

	ints := []int{8, 2, 5, 7, 6}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)
}
