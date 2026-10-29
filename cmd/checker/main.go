package main

import (
	"fmt"
	"os"
	"push-swap/internal/parse"
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
	 fmt.Printf("%v \n", nombres)

	 
}
