package main

import (
	"fmt"
	"push-swap/internal/stack"
)

func main() {
	 New := stack.NewEtat([]int{2, 1, 3, 6, 8, 5})
	 New.Executer("sa")
	 New.Executer("pb")
	 New.Executer("pb")
	 New.Executer("pb")
	 New.Executer("rb")
	 New.Executer("rrr")
	 New.Executer("pa")
	 New.Executer("pa")
	 New.Executer("pa")
	 fmt.Println(New.EstResolu()) 
}
