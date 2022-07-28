package main

import (
	"github.com/xvbnm48/learn-hexagonal/internal/adapters/core/arithmetic"
	"github.com/xvbnm48/learn-hexagonal/internal/ports"
)

func main() {
	var core = ports.ArithmaticPort{}

	core = arithmetic.NewAdapter()

	// arithAdapter := arithmetic.NewAdapter()
	// result, err := arithAdapter.Subtraction(1, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result)
	// arithAdapter.Addition(1, 2)
	// arithAdapter.Multiplication(1, 2)
	// arithAdapter.Division(1, 2)
}
