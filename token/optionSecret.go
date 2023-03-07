package token

import (
	"encoding/base64"

	"github.com/Mobilpadde/moths/v4/token/checks"
)

func OptionWithSecret(secret string) Option {
	return func(m *Generator) error {
		if err := checks.CheckSecret(secret); err != nil {
			return err
		}

		// Encode the secret as base64 - for extra security (🤷)
		data := []byte(secret)
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(dst, data)

		m.secret = dst
		return nil
	}
}
