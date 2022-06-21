package utils

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
)

// ReadSTDIN provides input reader from STDIN
func ReadSTDIN() ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		return nil, errors.New("failed to read input from STDIN")
	}

	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}
