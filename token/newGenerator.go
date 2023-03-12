package token

import (
	"math/rand"

	"github.com/Mobilpadde/characters/v2/letters"
)

// NewGenerator creates a new token generator
// with the given [Option]s.
func NewGenerator(opts ...Option) (*Generator, error) {
	g := &Generator{
		amount: 6, // Defaults to `6` as most other TOTP codes are this length
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	if g.secret == nil {
		// set a default secret since none has been specified
		secret := randSecret(32)
		g.SetSecret(secret)
	}

	// now := time.Now().UTC()
	// g.timing.curr = now
	// g.timing.last = now.Add(-g.period)

	return g, nil
}

func randSecret(n int) string {
	chars := letters.Combined
	charsLength := len(chars)

	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(charsLength)]
	}
	return string(b)
}
