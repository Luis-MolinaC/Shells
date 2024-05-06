package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	// Join the arguments into a single string
	text := strings.Join(args, "Hello world")
	// Print the text
	fmt.Fprintln(w, text)
	return nil
}
