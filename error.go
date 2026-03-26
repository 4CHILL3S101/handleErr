package handleErr

import (
	"fmt"
	"runtime"
)

type Error struct {
	Msg     string
	Code    string
	Context map[string]any
	Inner   error
	stack   []uintptr
}

// New creates a new error
func New(msg, code string) *Error {
	return &Error{
		Msg:   msg,
		Code:  code,
		stack: callers(),
	}
}

// Wrap wraps an existing error
func Wrap(err error, msg, code string, ctx map[string]any) *Error {
	return &Error{
		Msg:     msg,
		Code:    code,
		Context: ctx,
		Inner:   err,
		stack:   callers(),
	}
}

func (e *Error) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	if e.Inner != nil {
		return e.Inner.Error()
	}
	return "an unexpected error occurred"
}

func (e *Error) Debug() string {
	out := fmt.Sprintf("Error Code: %s\nMessage: %s\n", e.Code, e.Msg)
	if e.Context != nil {
		out += fmt.Sprintf("Context: %+v\n", e.Context)
	}
	if e.Inner != nil {
		out += fmt.Sprintf("Inner: %v\n", e.Inner)
	}
	out += fmt.Sprintf("Stack Trace:\n%s\n", stackTrace(e.stack))
	return out
}

func callers() []uintptr {
	pcs := make([]uintptr, 32)
	n := runtime.Callers(3, pcs)
	return pcs[:n]
}

func stackTrace(pcs []uintptr) string {
	frames := runtime.CallersFrames(pcs)
	out := ""
	for {
		frame, more := frames.Next()
		out += fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	return out
}
