package code

func (code Code) Validate(emojies string) bool {
	return code.String() == emojies
}

// Deprecated: Insecure. Use `Validate` instead.
func (code Code) ValidateToken(token string) bool {
	return code.Token() == token
}
