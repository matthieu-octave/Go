package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci(x int) int {

	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {

	// parser les arguments pour pouvoir passer le fichier en paramètres à l'éxécution
	progName := os.Args[0]
	os.Args = os.Args[1:]

	progName = filepath.Base(progName)
	var argNumber int
	var isNumberSet bool
	var err error

	// fmt.Printf("args=%s\n\n", os.Args)

	for len(os.Args) != 0 {
		arg := os.Args[0]
		os.Args = os.Args[1:]

		switch arg {
		case "-h", "--help":
			fmt.Println("Usage: ", progName, " <int>")
			os.Exit(1)
		default:
			argNumber, err = strconv.Atoi(arg)
			isNumberSet = true
			if err != nil {
				fmt.Println("Usage: ", progName, " <int>")
				panic(err)
			}
		}
	}
	if isNumberSet == false {
		fmt.Println("Veuillez entrer un nombre !")
		os.Exit(1)
	}

	f := fibonacci(argNumber)
	fmt.Println(f)
}
