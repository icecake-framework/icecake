package console

import (
	"fmt"
	"runtime/debug"
	"syscall/js"
)

type SHOW_LEVEL uint

const (
	SHOW_ALL     SHOW_LEVEL = 0x00
	SHOW_WARNERR SHOW_LEVEL = 0x01
	SHOW_ERRONLY SHOW_LEVEL = 0x02
)

var Show SHOW_LEVEL

// Logf prints a formated log (in white) to the console.
// Logf does nothing if SHOW_LEVEL is not equal to SHOW_ALL.
// Logf returns an error built with the formated message.
func Logf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show == 0 {
		fmt.Println(err.Error())
	}
	return err
}

// Warnf prints a formated warning (in yellow) to the console.
// Warnf does nothing if SHOW_LEVEL equal SHOW_ERRONLY
// Warnf returns an error built with the formated message.
func Warnf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show <= SHOW_WARNERR {
		js.Global().Call("ickWarn", err.Error())
	}
	return err
}

// Errorf prints a formated error (in red) to the console.
// Errorf returns an error built with the formated message.
func Errorf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	js.Global().Call("ickError", err.Error())
	return err
}

// Stackf prints the stack trace to the console
func Stackf() {
	Logf("> stack trace:\n" + string(debug.Stack()))
}

// Panicf is equivalent to Errorf followed by a panic
// If r is nil Stackf create a panic
func Panicf(msg string, a ...any) {
	e := Errorf(msg, a...)
	panic(e.Error())
}
