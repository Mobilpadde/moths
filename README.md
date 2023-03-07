# moths 🪰

> e**mo**jicon au**th**entication**s**

[![Go Reference](https://pkg.go.dev/badge/github.com/Mobilpadde/moths.svg)](https://pkg.go.dev/github.com/Mobilpadde/moths/v4)

## what (is this 💩)

Emojies as [TOTP](https://rublon.com/blog/hotp-totp-difference/), [because](#why-).

## why 🥲

We _all_ **could** use a little more emotion in our lives 🤗 Go show the world how you really feel!

It's a great companion for any app that uses 2FA - as every app should! Make your app's 2FA as star-striking as your app! 🤩

So whether your mascot is:

0. 🦋
1. 🐰
2. 🐶
3. 🐕
4. 🐷
5. 🐥
6. 🐱
7. 🐣
8. 🐻
9. 🐈

**moths** has got you covered! 🪰

> P.S. These emojies wasn't chosen by accident (random), they really are the 10 ut[most frequently used (animal) emojies of 2021](https://home.unicode.org/emoji/emoji-frequency/)

Show that sparkly emojies ✨

## how 🧑‍💼

Running this is quite easy 💨

0. Download the dependencies ⏬

> **Note**
>
> You can skip this step and let the `run`-command handle it - but you knew that 🧠

```sh
go mod download
```

1. Add a secret 🔐

> **Warning**
>
> This is just an one-liner to write the `base64`-encoded `sha256` of `moths` into `.env`.
> Make sure to use something better 🤷

```sh
echo 'moths' | sha256sum | base64 -w 0 | { echo -n "MOTHS_SECRET=$(cat -)" } > .env
```

> **Note**
>
> You don't need to use environment variables, this is just my choice.

2. Run the program 🏃

```sh
go run .
```

As I said, easy-peasy! 💖

### options 👓

To setup a new `moth`-generator, you must call [`moths.NewMoths`](moths/new.go#L9-L43) as

```go
gen, err := moths.NewMoths()
```

But I do recommend to add some customization, as it will fail immediately if not.

Like so

```go
gen, err := moths.NewMoths(
  moths.OptionWithSecret(secret), // any string as the secret
  moths.OptionWithInterval(time.Second), // Each is only valid for a second
  moths.OptionWithAmount(6), // Each must is always `6` emojies
  moths.OptionWithEmojies(emojies.CATS), // A pure slice of cats
  moths.OptionWithTime(time.Now().AddDate(10, 0, 0)), // 10 years into the future
)
```

There are a few options to choose from, these are

- [`OptionWithSecret(secret string)`](moths/options.go#L29-L45)**\***
  - The secret to generate from
  - Must be 32 characters ⚠
- [`OptionWithInterval(interval time.Duration)`](moths/options.go#L47-L56)**\***
  - On which interval should a new `moth` be generated
  - A `moth` will only be valid during this duration
- [`OptionWithAmount(amount int)`](moths/options.go#L58-L67)
  - The amount of emojies to generate in a `moth`
  - Defaults to `6`-charater tokens
- [`OptionWithEmojies(emojies emojies.Emojies)`](moths/options.go#L69-L78)**\***
  - Take a look in the [`emojies`](moths/emojies)-package to see your options
  - You can also add your own emojies - [How to](#emojies-)
- [`OptionWithTime(t time.Time)`](moths/options.go#L80-L85)
  - This will allow you to add a custom time
  - Meaning you can validate towards old `moths`
  - You can even add future dates ⌛
  - Defaults to _now_

> **Warning**
>
> Options marked with an asterix (\*) are required!

### generating 🖇

Now that you have a sparkly new `moth`-generator, you can use it as

```go
otp, err := gen.Next()
```

Now that you have an [`OTP`](moths/otp), you can use its functions

- [`Validate(moth string) bool`](moths/otp/validate.go#L3-L5)
  - Will validate a moth (pattern of emojies) directly
- [`ValidateToken(token string) bool`](moths/otp/validate.go#L7-L9)
  - Will validate a OTP token
  - You'll need to expose the code for your user too, for this - not recommended
- [`Token() string`](moths/otp/config.go#L12-L14)
  - Returns the token - for whatever reason that might be needed
- [`String() string`](moths/otp/config.go#L16-L18)
  - Returns the `moth` as a string
- [`SpacedString() string`](moths/otp/config.go#L20-L22)
  - Returns the `moth` as a string with spaces inbetween
- [`Slice() []string`](moths/otp/config.go#L24-L26)
  - Returns the `moth` as a slice of strings

### validating 🧑‍🔬

To validate, you'll need both the `moth` (or the `token`) and the generator

```go
token := otp.String() // Ideally you'd get this from the user
ok := gen.Validate(token)
```

### emojies 😻

To use your own set of **known** emojies, you can reference the [`cat`-emojies](moths/emojies/cats.go).

I've chosen the `cats` as it's a great reference, both for creating and re-using.

If we would like to add an easter-egg to a pure bred slice of cats, we could do it like:

```go
// A slice of cat emojies and a single hotdog 🌭
var catsHotdog = []string{
	emoji.GrinningCat.String(),
	emoji.GrinningCatWithSmilingEyes.String(),
	emoji.CatWithTearsOfJoy.String(),
	emoji.SmilingCatWithHeartEyes.String(),
	emoji.CatWithWrySmile.String(),
	emoji.KissingCat.String(),
	emoji.WearyCat.String(),
	emoji.CryingCat.String(),
	emoji.PoutingCat.String(),

	emoji.HotDog.String(),
}

var CATS_HOTDOG = ToEmojies(catsHotdog)
```

It doesn't even have to be from the [emoji](https://github.com/enescakir/emoji)-package,
simply make a slice of your desired emojies, and use the [`ToEmojies`-func](moths/emojies/helper.go#L5-L19).
Then provide this as an argument in `moths.OptionWithEmojies(CATS_HOTDOG)` when calling the `moths.NewMoths()`.

## example 🤷

Check out [`main.go`](main.go) for an example

## show-case 🕺

![Three iterations of moths](./data/sample.png)

## history ✍

- [`v4.0.1`](https://github.com/Mobilpadde/moths/tree/v4.0.1) 💘

- <details>
    <summary>Older</summary>

  - [`v4.0.0`](https://github.com/Mobilpadde/moths/tree/v4.0.0)
  - [`v3.0.0`](https://github.com/Mobilpadde/moths/tree/v3.0.0)
  - [`v2.2.2`](https://github.com/Mobilpadde/moths/tree/v2.2.2)
  - [`v2.2.1`](https://github.com/Mobilpadde/moths/tree/v2.2.1)
  - [`v2.2.0`](https://github.com/Mobilpadde/moths/tree/v2.2.0)
  - [`v2.1.0`](https://github.com/Mobilpadde/moths/tree/v2.1.0)
  - [`v2.0.0`](https://github.com/Mobilpadde/moths/tree/v2.0.0)
  - [`v1.0.0`](https://github.com/Mobilpadde/moths/tree/v1.0.0)
  - [`v0.1.0`](https://github.com/Mobilpadde/moths/tree/v0.1)
  </details>

## future 🔮

- ~~Rename `moth` to something better (`OTP` / `Token`)~~ - TODO: Add commit here
- Add better documentation 🫢
- ~~Rename the `Moths`-struct to `Generator`~~ - TODO: Add commit here
  - This means renaming the [`otp`](moths/otp)-package as well (`totp`)?
- Get [pkg.go.dev](https://pkg.go.dev/github.com/Mobilpadde/moths) up-to-date

## shoutout 📢💨

I couldn't have done it without these lovely OSS 🦾

- <https://github.com/aidarkhanov/nanoid>
- <https://github.com/enescakir/emoji>
- <https://github.com/tilaklodha/google-authenticator>
- <https://github.com/pquerna/otp/>

In no specific order 🤷
