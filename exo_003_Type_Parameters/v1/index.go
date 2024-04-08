package main

import (
	"fmt"
)

// Index returns the index of x in s, or e if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

type strError string

func (e strError) Error() string {
	return string(e)
}

func printIndex[T comparable](x []T, y T) error {
	if Index(x, y) == -1 {
		fmt.Printf("Erreur : %v n'exite pas dans le tableau", y)
		return strError("-1")
	}
	fmt.Println("Index :", Index(x, y), "->", "Valeur :", x[Index(x, y)])
	return nil
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	printIndex(si, 15)

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	printIndex(ss, "hello")
}
