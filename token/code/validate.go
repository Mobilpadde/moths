package code

func (code Code) Validate(emojies string) bool {
	return code.String() == emojies
}

func (code Code) ValidateToken(token string) bool {
	return code.Token() == token
}
