// Provides convience methods to print information to color-aware terminal
package fmtc

import (
	"fmt"

	"github.com/alexsaveliev/go-colorable-wrapper/streamc"
)

// Mimics fmt.Print
func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(streamc.Stdout, a...)
}

// Mimics fmt.Printf
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(streamc.Stdout, format, a...)
}

// Mimics fmt.Println
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(streamc.Stdout, a...)
}
