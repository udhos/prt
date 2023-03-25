// Package main implements an example.
package main

import (
	"fmt"

	"github.com/udhos/prt/prt"
	"golang.org/x/exp/slices"
)

type state struct {
	name   string
	age    int
	active bool
}

func main() {

	info := state{
		name:   "john",
		age:    33,
		active: true,
	}

	weights := []int{12, 31, 64, 22, 77, 99, 12, 81}

	vars := prt.Data{
		"name":   func() string { return info.name },
		"age":    func() string { return fmt.Sprint(info.age) },
		"active": func() string { return fmt.Sprint(info.active) },
		"heaviest": func() string {
			if len(weights) == 0 {
				return "<empty_list>"
			}
			w := append(weights[:0:0], weights...)
			slices.Sort(w)
			return fmt.Sprint(w[len(w)-1])
		},
	}

	msg := vars.Str("Str: O jogador {name} tem {age} anos e continua na ativa? Resposta: {active}\n")
	fmt.Print(msg)

	info.age = 44

	vars.Out("Out: O jogador {name} tem {age} anos e continua na ativa? Resposta: {active}\n")

	vars.Out("The heaviest stone weighs {heaviest} kg.\n")

	weights = append(weights, 222)

	vars.Out("The heaviest stone weighs {heaviest} kg.\n")

	weights = nil

	vars.Out("The heaviest stone weighs {heaviest} kg.\n")

}
