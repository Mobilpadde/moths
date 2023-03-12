package token_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v6/token"
	"github.com/Mobilpadde/moths/v6/token/emojies"
	"github.com/Mobilpadde/moths/v6/token/errs"
	"github.com/Mobilpadde/moths/v6/token/option"
)

func TestNewGenerator(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	g, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}

	err = g.Check()
	if err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}
}

func TestNewGeneratorNoSecret(t *testing.T) {
	amount := 6

	_, err := token.NewGenerator(
		option.OptionWithSecret(""),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrSecretLength) {
		t.Error("Expected to return an error when creating new generator without a secret")
	}
}

func TestNewGeneratorNoInterval(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	g, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}

	err = g.Check()
	if !errors.Is(err, errs.ErrPeriod1) {
		t.Error("Expected to return an error when creating new generator without an interval")
	}
}

func TestNewGeneratorNoAmount(t *testing.T) {
	secret := strings.Repeat("a", 32)

	_, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithAmount(0),
		option.OptionWithPeriod(time.Second),
		option.OptionWithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when creating new generator without an amount")
	}
}

func TestNewGeneratorNoEmojies(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	g, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}

	err = g.Check()
	if !errors.Is(err, errs.ErrEmojies1) {
		t.Error("Expected to return an error when creating new generator without any emojies")
	}
}

// Instantiate a new generator with
// a secret, period, amount, and emojies.
//
// Use this generator to generate new codes
// from the options provided.
func ExampleNewGenerator() {
	amount := 6
	secret := strings.Repeat("a", 32)

	gen, _ := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
	)

	// It's good practice to use the run `g.Check()`
	// method to check if everything is working.
	// This is not requried though.
	if err := gen.Check(); err != nil {
		panic(err)
	}

	gen.Next()
}

// Instantiates a generator as above,
// but also with a specified time.
//
// This will **always** generate the same
// codes *virtually forever*.
//
// If we chose to generate **five** codes,
// these would be as follows:
//
// 🙀 😾 😹 🙀 😼 😹
//
// 😻 😹 😽 😹 😽 😿
//
// 😹 😽 😻 😸 😻 🙀
//
// 😼 😼 😾 😾 😿 😹
//
// 😿 😽 😿 🙀 😼 😻
func ExampleNewGenerator_with_time() {
	secret := strings.Repeat("a", 32)

	gen, _ := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(6),
		option.OptionWithEmojies(emojies.CATS),
		option.OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	)

	gen.Next()
}

// Instantiate a new generator with
// a secret, amount, and emojies.
//
// We omit the period option and
// set it later.
//
// This can be done with the other
// options as well.
func ExampleNewGenerator_set_period() {
	secret := strings.Repeat("a", 32)

	gen, _ := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithAmount(6),
		option.OptionWithEmojies(emojies.CATS),
	)

	// If we run `gen.Check()` here,
	// it'll return an error, because
	// the period is not set.
	gen.Check() // not nil

	// set the period to 1 second
	gen.SetPeriod(time.Second)

	// If we run `gen.Check()` again,
	// everything works, as expected
	// because the period is now set.
	gen.Check() // nil
}
