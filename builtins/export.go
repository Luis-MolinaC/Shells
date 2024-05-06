package builtins

import (
	"fmt"
)

func Export(args ...string) error {
	if len(args) == 0 {
		return errors.New("usage: export NAME=VALUE")
	}

	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid export format: %s", arg)
		}
		name, value := parts[0], parts[1]

		// Set the environment variable
		os.Setenv(name, value)
		fmt.Printf("Exported variable: %s=%s\n", name, value)
	}

	return nil
}

