package errs

import (
	"errors"
)

const (
	// Secret must have at least 1 characters
	secretLength string = "secret must be at least 1 character"

	// Key must be of 20 in length
	// key20 string = "key could not be parsed from secret"

	// Emojies must have at least one character
	emojies1 string = "emojies must be at least 1 character"

	// Interval must be at least 1s
	interval1 string = "interval must be at least 1000ms"

	// Amount must be at least 1 emoji
	amount1 string = "interval must be at least 1 emoji"
)

var (
	ErrSecretLength error = errors.New(secretLength)
	ErrEmojies1     error = errors.New(emojies1)
	ErrInterval1    error = errors.New(interval1)
	ErrAmount1      error = errors.New(amount1)
)
