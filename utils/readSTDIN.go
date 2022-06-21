package utils

import (
	"bufio"
	"errors"
	"os"
)

// ReadSTDIN provides input reader from STDIN
func ReadSTDIN() (*bufio.Reader, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, errors.New("failed to read input from STDIN")
	}

	return bufio.NewReader(os.Stdin), nil
}
