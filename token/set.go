package token

import (
	"encoding/base64"
	"time"

	"github.com/Mobilpadde/moths/v6/token/checks"
	"github.com/Mobilpadde/moths/v6/token/emojies"
)

// SetSecret is used to specify
// the secret to generate codes from.
func (g *Generator) SetSecret(secret string) error {
	if err := checks.CheckSecret(secret); err != nil {
		return err
	}

	// Encode the secret as base64 - is this even extra secure? 🤷
	data := []byte(secret)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)

	g.secret = dst
	return nil
}

// SetAmount is used to specify
// the amount of emojies in a code.
func (g *Generator) SetAmount(amount int) error {
	if err := checks.CheckAmount(amount); err != nil {
		return err
	}

	g.amount = amount
	return nil
}

// SetEmojies is used to specify
// which emojies that can be used in
// any given code.
func (g *Generator) SetEmojies(emojies emojies.Emojies) error {
	if err := checks.CheckEmojies(emojies); err != nil {
		return err
	}

	g.emojies = emojies
	return nil
}

// SetPeriod is used to specify
// how long a code will remain valid.
func (g *Generator) SetPeriod(period time.Duration) error {
	if err := checks.CheckPeriod(period); err != nil {
		return err
	}

	g.period = period
	g.syncTiming()
	return nil
}

// OptionWithTime is used to specify
// a custom time to generate code from.
//
// If none is specified, the current time
// will be used.
func (g *Generator) SetTime(t time.Time) error {
	g.timing.time = t
	if g.timing.time == (time.Time{}) {
		g.timing.time = time.Now().UTC()
	}

	g.syncTiming()
	return nil
}

// Since we're updating time-related
// options, we need to synchronize
// the current and last time.
func (g *Generator) syncTiming() {
	now := time.Now().UTC()
	g.timing.curr = now
	g.timing.last = now.Add(-g.period)
}
