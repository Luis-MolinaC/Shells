package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	text := strings.Join(args, "Hello world")
	fmt.Fprintln(w, text)
	return nil
}
