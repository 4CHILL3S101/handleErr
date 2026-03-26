package handleErr

import (
	"fmt"
)

// CheckErr checks if err is not nil.
// Returns a boolean and a user-facing message.
// If devMode is true, prints debug info and stack trace.
func CheckErr(err error, userMsg string, dev bool) (bool, string) {
	if err != nil {
		if dev {
			fmt.Println("DEBUG:", err)
			fmt.Println("Stack Trace:\n", stackTrace(callers()))
		}
		return true, userMsg
	}
	return false, ""
}
