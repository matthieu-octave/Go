package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		str := fmt.Sprintf("%v", y)
		fmt.Printf("Erreur : %q n'exite pas dans le tableau", str)
		return strError("-1")
	}
	fmt.Println("Index :", Index(x, y), "->", "Valeur :", x[Index(x, y)])
	return nil
}

func main() {
	type T interface{}
	var argNumber int
	var isNumberSet bool
	var err error
	var x []T = make([]T, 0)
	var y T

	// faire 2 tableaux, 1 de int et 1 de strings
	// en fonction de la saisie, remplir l'un ou l'autre

	// var a T
	// var b T

	// Remplacer la saisie du tableau par une suite de prompt
	/*
		fmt.Println("Entrez les éléments du tableau :")
		for {
			_, err = fmt.Scan(a)
			if err != nil {
				break
			}
			x = append(x, a)
		}

		fmt.Println("Entrez la valeur à rechercher :")
		fmt.Scanln(b)
		fmt.Println("valeur à rechercher :", b)
	*/

	// parser les arguments pour pouvoir passer le fichier en paramètres à l'éxécution

	progName := os.Args[0]
	os.Args = os.Args[1:]

	progName = filepath.Base(progName)

	for len(os.Args) != 0 {
		arg := os.Args[0]
		os.Args = os.Args[1:]

		switch arg {
		case "-h", "--help":
			fmt.Println("Usage: ", progName, " -i <int> <int> [int...]")
			fmt.Println("OU")
			fmt.Println("Usage: ", progName, " -s <int> <string> [string...]")
			os.Exit(1)
		default:
			// convertir arg(string) en tableau x

			// a revoir faire 2 cases "-i" et "-s"
			fields := strings.Fields(arg)
			x = make([]T, len(fields))
			for i, field := range fields {
				// Convert `field` to `T` and store it in `x[i]`
				x[i] = T(field)
			}

			y = os.Args

			//
			isNumberSet = true
			if err != nil {
				fmt.Println("Usage: ", progName, " <int> <int> [int...]")
				fmt.Println("OU")
				fmt.Println("Usage: ", progName, " <int> <string> [string...]")
				panic(err)
			}
		}
	}

	if isNumberSet == false {
		fmt.Println("Veuillez entrer une valeur à rechercher !")
		os.Exit(1)
	}

	// Index works on a slice of ints or strings

	fmt.Printf("Nombre d'éléments dans le tableau : %d", argNumber)
	printIndex(x, y)
}

// Output: 0Erreur : "[]" n'exite pas dans le tableau
