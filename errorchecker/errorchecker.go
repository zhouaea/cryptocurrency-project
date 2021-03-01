package errorchecker

import (
	"fmt"
	"os"
)

func CheckError(error error) {
	if error != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", error.Error())
		os.Exit(2)
	}
}