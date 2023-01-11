package otp

func (otp OTP) Validate(moth string) bool {
	return otp.String() == moth
}

func (otp OTP) ValidateToken(token string) bool {
	return otp.Token() == token
}
