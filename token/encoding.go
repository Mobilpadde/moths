// Probably very overengineered right now.
// Will think up a better solution later.

package token

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Mobilpadde/moths/v5/token/emojies"
)

const (
	seperator = ";;;"
	delimiter = ":"

	letterSecret  = "s"
	letterAmount  = "a"
	letterEmojies = "e"
	letterPeriod  = "p"
	letterTime    = "t"

	typeSecret  = "%s"
	typeAmount  = "%d"
	typeEmojies = "%s"
	typePeriod  = "%d"
	typeTime    = "%d"

	patternField = "%s%s%s"
)

var (
	prefixSecret  = letterSecret + delimiter
	prefixAmount  = letterAmount + delimiter
	prefixEmojies = letterEmojies + delimiter
	prefixPeriod  = letterPeriod + delimiter
	prefixTime    = letterTime + delimiter

	patternSecret  = fmt.Sprintf(patternField, prefixSecret, typeSecret, seperator)
	patternAmount  = fmt.Sprintf(patternField, prefixAmount, typeAmount, seperator)
	patternEmojies = fmt.Sprintf(patternField, prefixEmojies, typeEmojies, seperator)
	patternPeriod  = fmt.Sprintf(patternField, prefixPeriod, typePeriod, seperator)
	patternTime    = fmt.Sprintf(patternField, prefixTime, typeTime, seperator)

	patternEncoded = "%s%s%s%s%s"
)

// Encode generators fields as a
// base64 string.
//
// Should not be shared, as this
// shows the secret as a string
// if this is decoded.
func (g *Generator) Export() string {
	secret := fmt.Sprintf(patternSecret, g.secret)
	amount := fmt.Sprintf(patternAmount, g.amount)
	emojies := fmt.Sprintf(patternEmojies, g.emojies)
	period := fmt.Sprintf(patternPeriod, g.period)
	time := fmt.Sprintf(patternTime, g.timing.time.Unix())

	str := fmt.Sprintf(patternEncoded, secret, amount, emojies, period, time)
	str = str[0 : len(str)-3]

	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Try parsing a given base64
// encoded string into Generator.
func (g *Generator) Import(encoded string) error {
	str, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	updatable := []struct {
		prefix string
		fn     func(string) error
	}{
		{prefixSecret, g.parseSecret},
		{prefixAmount, g.parseAmount},
		{prefixEmojies, g.parseEmojies},
		{prefixPeriod, g.parsePeriod},
		{prefixTime, g.parseTime},
	}
	split := strings.Split(string(str), seperator)
	for _, x := range split { // loop over each field
		for _, y := range updatable { // loop over each updatable field
			if strings.HasPrefix(x, y.prefix) { // if current field is updatable
				if err := y.fn(x); err != nil { // udpate
					return err
				}

				continue
			}
		}
	}

	g.Check()
	return nil
}

func (g *Generator) parseSecret(x string) error {
	str := x[len(prefixSecret):]

	// We don't encode secret again,
	// so we update the secret (byte-slice)
	// manually.
	g.secret = []byte(str)

	return nil
}

func (g *Generator) parseAmount(x string) error {
	str := x[len(prefixAmount):]
	amount, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	return g.SetAmount(amount)
}

func (g *Generator) parseEmojies(x string) error {
	str := x[len(prefixEmojies):]
	split := strings.Split(str, "")

	return g.SetEmojies(emojies.ToEmojies(split))
}

func (g *Generator) parsePeriod(x string) error {
	str := x[len(prefixPeriod):]
	period, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	duration := time.Duration(period)
	return g.SetPeriod(duration)
}

func (g *Generator) parseTime(x string) error {
	str := x[len(prefixTime):]
	unix, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	parsed := time.Unix(unix, 0).UTC()
	return g.SetTime(parsed)
}
