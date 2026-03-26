package handleErr

import (
	"fmt"
)

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
