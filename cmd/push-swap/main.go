package main

import (
	"fmt"
	"os"
	"push-swap/internal/parse"
	"push-swap/internal/stack"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	nombres, err := parse.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	if len(nombres) == 0 {
		return
	}
	Newetat := stack.NewEtat(nombres)
	Newetat.Trier()

	for _, v := range stack.Nettoyer(Newetat.Coups()) {
		fmt.Println(v)
	}

}
