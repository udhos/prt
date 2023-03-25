// Package main implements an example.
package main

import (
	"fmt"

	"github.com/udhos/prt/prt"
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

	vars := prt.Data{
		"name":   func() string { return info.name },
		"age":    func() string { return fmt.Sprint(info.age) },
		"active": func() string { return fmt.Sprint(info.active) },
	}

	msg := vars.Str("Str: O jogador {name} tem {age} anos e continua na ativa? Resposta: {active}\n")
	fmt.Print(msg)

	info.age = 44

	vars.Out("Out: O jogador {name} tem {age} anos e continua na ativa? Resposta: {active}\n")
}
