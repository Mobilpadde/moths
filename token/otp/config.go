package otp

import (
	"github.com/Mobilpadde/moths/v4/token/emojies"
)

type OTP struct {
	token   string
	emojies emojies.Emojies
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
