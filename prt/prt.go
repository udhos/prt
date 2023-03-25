// Package prt expands variables.
package prt

import (
	"fmt"
	"strings"
)

// FuncGet is a utility type for extracting variables as strings.
type FuncGet func() string

// Data maps variables names to function to extract variables values as strings.
type Data map[string]FuncGet

// Str expands variables.
func (d Data) Str(format string) string {
	for name, f := range d {
		format = strings.ReplaceAll(format, "{"+name+"}", f())
	}
	return format
}

// Out expands variables and prints the result to stdout.
func (d Data) Out(format string) {
	fmt.Print(d.Str(format))
}
