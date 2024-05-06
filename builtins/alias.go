package builtins

import (
	"errors"
	"fmt"
)

func Alias(args ...string) error {
	if len(args) == 0 {
		return errors.New("usage: alias NAME=VALUE")
	}
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid alias format: %s", arg)
		}
		name, value := parts[0], parts[1]
		fmt.Printf("Alias created: %s -> %s\n", name, value)
	}

	return nil
}
