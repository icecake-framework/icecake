package errors

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

func ConsoleLogf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show == 0 {
		fmt.Println(err.Error())
	}
	return err
}

func ConsoleWarnf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show <= SHOW_WARNERR {
		js.Global().Call("ickWarn", err.Error())
	}
	return err
}

func ConsoleErrorf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	js.Global().Call("ickError", err.Error())
	return err
}

// ConsoleStackf prints the msg message and the panic recovery message if not nil, then the stacktrace.
// If r is nil ConsoleStackf create a panic
func ConsoleStackf(r any, msg string, a ...any) {
	defer ConsoleLogf("> panic stacktrace:\n" + string(debug.Stack()))
	str := fmt.Sprintf(msg, a...)
	if str != "" {
		ConsoleErrorf(str)
	}
	if r == nil {
		panic(str)
	} else {
		ConsoleErrorf("%+v [recovered]\n", r)
	}
}

func ConsoleStack(r any) {
	ConsoleErrorf("%+v [recovered]\n", r)
	ConsoleLogf("> panic stacktrace:\n" + string(debug.Stack()))
}
