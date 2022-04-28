# moths

> e**mo**ji a**ths**

# what

Emojies as TOTP

# why

Becasue why not?

# how

Runnig this is quite easy, just run the command:

Download the dependencies:

```sh
go mod download
```

Next-up start the program

```sh
go run .
```

As I said, easy-peasy!

# sample

A sample out-put might be `😻🙀😺🙀🙀` which would equal the `920811` TOTP-token. Using the super secret secret specified in [`./secret.pem`](./secret.pem) - genereated at `2022/04/28 02:05:42` and a `5`-second interval.

# shoutout

I couldn't have done it without the lovely OSS, listed below:

- https://github.com/aidarkhanov/nanoid
- https://github.com/enescakir/emoji
- https://github.com/tilaklodha/google-authenticator
