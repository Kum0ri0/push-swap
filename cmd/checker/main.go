package main

import (
	"bufio"
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
	scan := bufio.NewScanner(os.Stdin)
	sliceScan := []string{}
	for scan.Scan() {
		sliceScan = append(sliceScan, scan.Text())
	}
	err = scan.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	etat := stack.NewEtat(nombres)
	for _, v := range sliceScan {
		if v == "" {
			continue
		}
		if !etat.Executer(v) {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}
	if etat.EstResolu() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
