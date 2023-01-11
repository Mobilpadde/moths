package otp

import (
	"moths/moths/emojies"
)

type OTP struct {
	token   string
	emojies emojies.Emojies
}

func (otp OTP) Validate(code string) bool {
	return otp.token == code
}

func (otp OTP) Token() string {
	return otp.token
}

func (otp OTP) String() string {
	return otp.emojies.String()
}

func (otp OTP) SpacedString() string {
	return otp.emojies.SpacedString()
}

func (otp OTP) Slice() []string {
	return otp.emojies.Slice()
}
