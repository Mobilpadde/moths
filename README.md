# moths 🪰

> e**mo**jicon au**th**entication**s**

[![Go Reference](https://pkg.go.dev/badge/github.com/Mobilpadde/moths/v6.svg)](https://pkg.go.dev/github.com/Mobilpadde/moths/v6)

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
>
> You don't need to use environment variables, this is just my choice.

```sh
echo 'moths' | sha256sum | base64 -w 0 | { echo -n "MOTHS_SECRET=$(cat -)" } > .env
```

2. Run the program 🏃

```sh
go run .
```

As I said, easy-peasy! 💖

### options 👓

To setup a new `code`-generator, you must call [`token.NewGenerator`](token/newGenerator.go#L11-L33) as

```go
gen, err := token.NewGenerator()
```

But I do recommend to add some customization, as it will fail immediately if not.

Like so

```go
gen, err := token.NewGenerator(
  token.OptionWithSecret(secret), // Any string as the secret
  token.OptionWithPeriod(time.Second), // Each code is only valid for a second
  token.OptionWithAmount(6), // Each must is always `6` emojies
  token.OptionWithEmojies(emojies.CATS), // A pure slice of cats
  token.OptionWithTime(time.Now().AddDate(10, 0, 0)), // 10 years into the future
)
```

There are a few options to choose from, these are

- [`OptionWithSecret(secret string)`](token/option/secret.go#L7-L11)
  - The secret to generate from
  - Defaults to a random 32-character string
- [`OptionWithPeriod(period time.Duration)`](token/option/period.go#L11-L15)**\***
  - On which interval should a new `code` be generated
  - A `code` will only be valid during this duration - Until (if) skewed interval are implemented
- [`OptionWithAmount(amount int)`](token/option/amount.go#L9-L13)
  - The amount of emojies to generate in a `moth`
  - Defaults to `6`-charater tokens
- [`OptionWithEmojies(emojies emojies.Emojies)`](token/option/emojies.go#L11-L15)**\***
  - Take a look in the [`emojies`](token/emojies)-package to see your options
  - You can also add your own emojies - [How to](#emojies-)
- [`OptionWithTime(t time.Time)`](token/option/time.go#L14-L18)
  - This will allow you to add a custom time
  - Meaning you can validate towards old `code`s
  - You can even add future dates ⌛
  - Defaults to _now_
- [`OptionFromEncoded(str string)`](token/option/encode.go#L11-L15)
  - Use this if you have [exported](token/encoding.go#L57) a generator earlier

> **Warning**
>
> Options marked with an asterix (\*) are required!

### generating 🖇

Now that you have a sparkly new `code`-generator, you can use it as

```go
code, err := gen.Next()
```

Now that you have a [`code`](token/code), you can use its functions

- [`Validate(emojies string) bool`](token/code/validate.go#L3-L5)
  - Will validate a code (pattern of emojies) directly
- [`ValidateToken(token string) bool`](token/code/validate.go#L8-L10) - **_DEPRECATED_**
  - Will validate a token
  - You'll need to expose the token to your user(s) for this - not recommended
- [`String() string`](token/code/config.go#L16-L18)
  - Returns the code as a string
- [`SpacedString() string`](token/code/config.go#L20-L22)
  - Returns the code as a string with spaces inbetween
- [`Slice() []string`](token/code/config.go#L24-L26)
  - Returns the code as a slice of strings
  - Returns the code as a string
- [`CreatedAt() time.Time`](token/code/config.go#L28-L30)
  - Returns the time the code was generated
- [`ExpiresAt() time.Time`](token/code/config.go#L32-L34)
  - Returns the time when the code will expire
- [`Token() string`](token/code/config.go#L37-L39) - **_DEPRECATED_**
  - Returns the token - for whatever reason that might be needed

### validating 🧑‍🔬

To validate, you'll need both the `code` (or the `token`) and the generator

```go
str := code.String() // Ideally you'd get this from the user
ok := gen.Validate(str)
```

### emojies 😻

To use your own set of **known** emojies, you can reference the [`cat`-emojies](token/emojies/cats.go).

I've chosen the `cats` as it's a great reference, both for creating and re-using.

If we would like to add an easter-egg to a purebred slice of cats, we could do it like:

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
simply make a slice of your desired emojies, and use the [`ToEmojies`-func](token/emojies/helper.go#L7-L21).
Then provide this as an argument in `token.OptionWithEmojies(CATS_HOTDOG)` when calling the `token.NewGenerator()`.

## example 🤷

Check out [`main.go`](main.go) for an example

## show-case 🕺

![Three iterations of moths](./data/sample.png)

## history ✍

- [`v6.0.0`](https://github.com/Mobilpadde/moths/tree/v6.0.0) 💘

- <details>
    <summary>Older</summary>

  - [`v5.0.2`](https://github.com/Mobilpadde/moths/tree/v5.0.2)
  - [`v5.0.1`](https://github.com/Mobilpadde/moths/tree/v5.0.1)
  - [`v5.0.0`](https://github.com/Mobilpadde/moths/tree/v5.0.0)
  - [`v4.0.1`](https://github.com/Mobilpadde/moths/tree/v4.0.1)
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

- Add [Skew Intervals](https://www.ibm.com/docs/en/sva/9.0.6?topic=authentication-configuring-totp-one-time-password-mechanism) 🕰️
- Add better documentation 🫢
- Get into fixing [`geatures`](token/emojies/gestures.go#L28-L60) 🤦
- ~~`Export` / `Import`~~ - [d1bb4cc](https://github.com/Mobilpadde/moths/commit/d1bb4cc92abcf1f00d19088bbeee7013f6c8c302)
- ~~Rename `moth` to something better (`OTP` / `Token`)~~ - [6105848](https://github.com/Mobilpadde/moths/commit/6105848b336d57af5cc60fe53aa60532d2f979a4)
- ~~Rename the `Moths`-struct to `Generator`~~ - [4c973ef](https://github.com/Mobilpadde/moths/commit/4c973ef15c6f6102aaf3741aeb64ea35663b0b9c)
  - ~~This means renaming the `otp`-package as well?~~ - [fcdf295](https://github.com/Mobilpadde/moths/commit/fcdf295111bec0b516db62c3879bf4b7d7fd4436)
- ~~Get [pkg.go.dev](https://pkg.go.dev/github.com/Mobilpadde/moths) up-to-date~~ - Works with [v6](https://pkg.go.dev/github.com/Mobilpadde/moths/v6)

## shoutout 📢💨

I couldn't have done it without these lovely OSS 🦾

- <https://github.com/aidarkhanov/nanoid>
- <https://github.com/enescakir/emoji>
- <https://github.com/tilaklodha/google-authenticator>
- <https://github.com/pquerna/otp/>

In no specific order 🤷
