package internal

import (
	"errors"
	"io"
	"os"
)

var (
	// ErrMissingFileParam error when no STDIN nor --from param is passed.
	ErrMissingFileParam = errors.New("missing --from parameter")
)

// GetInputStream returns an io.Reader from a file or a piped STDIN.
func GetInputStream(in io.Reader, from string) (io.Reader, error) {
	if from == "" && isPiped() {
		return in, nil
	}

	if from == "" {
		return nil, ErrMissingFileParam
	}

	_, err := os.Stat(from)

	if err != nil {
		return nil, err
	}

	return os.Open(from)
}

// isPiped detect if there is any input through STDIN.
func isPiped() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	notPipe := info.Mode()&os.ModeNamedPipe == 0

	return !notPipe || info.Size() > 0
}
