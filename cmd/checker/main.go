package main

import (
	"fmt"
	"push-swap/internal/stack"
)

func main() {
	teste := []int{2, 1, 3}
	result := stack.NewEtat(teste)
	fmt.Printf("le teste %+v", result)
}
