package builtins

import (
	"errors"
	"fmt"
)

// Alias represents the alias command
func Alias(args ...string) error {
	// Check if arguments are provided
	if len(args) == 0 {
		return errors.New("usage: alias NAME=VALUE")
	}

	// Split the argument into name and value
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid alias format: %s", arg)
		}
		name, value := parts[0], parts[1]

		// Store the alias in a map or perform other actions
		fmt.Printf("Alias created: %s -> %s\n", name, value)
	}

	return nil
}
