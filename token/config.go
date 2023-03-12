package token

import (
	"time"

	"github.com/Mobilpadde/moths/v6/token/emojies"
)

type Generator struct {
	secret  []byte
	amount  int
	emojies emojies.Emojies

	period    time.Duration
	lastToken string

	timing struct {
		time time.Time
		curr time.Time
		last time.Time
	}
}

type Option func(*Generator) error
