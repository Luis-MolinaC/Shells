package builtins

import (
	"fmt"
)

func Unset(args ...string) error {
	if len(args) == 0 {
		return errors.New("usage: unset NAME")
	}

	for _, arg := range args {
		os.Unsetenv(arg)
		fmt.Printf("Unset variable: %s\n", arg)
	}

	return nil
}

