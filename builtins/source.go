package builtins

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Source(args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: source FILE")
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line[0] == '#' {
			continue
		}

		parts := strings.Split(line, " ")
		cmd := parts[0]
		args := parts[1:]

		err := executeCommand(cmd, args...)
		if err != nil {
			return fmt.Errorf("error executing command in file %s: %v", fileName, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}
