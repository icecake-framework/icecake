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

func Logf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show == 0 {
		fmt.Println(err.Error())
	}
	return err
}

func Warnf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	if Show <= SHOW_WARNERR {
		js.Global().Call("ickWarn", err.Error())
	}
	return err
}

func Errorf(msg string, a ...any) error {
	err := fmt.Errorf(msg, a...)
	js.Global().Call("ickError", err.Error())
	return err
}

// Stackf prints the msg message, the r recovery message and the stacktrace.
func Stackf(r any, msg string, a ...any) {
	defer Logf("> panic stacktrace:\n" + string(debug.Stack()))
	str := fmt.Sprintf(msg, a...)
	if str != "" {
		Errorf(str)
	}
	Errorf("%+v [recovered]\n", r)
}

// Panicf prints the msg message and the panic recovery message if not nil, then the stacktrace.
// If r is nil Stackf create a panic
func Panicf(msg string, a ...any) {
	defer Logf("> panic stacktrace:\n" + string(debug.Stack()))
	str := fmt.Sprintf(msg, a...)
	panic(str)
}

// func ConsoleStack(r any) {
// 	Errorf("%+v [recovered]\n", r)
// 	Logf("> panic stacktrace:\n" + string(debug.Stack()))
// }
