package moths

import "errors"

func NewMoths(secret string, interval int) (*Moths, error) {
	if len(secret) < 32 {
		return nil, errors.New("secret must be at least 32 characters")
	}

	m := &Moths{
		secret:   secret,
		interval: interval,
	}

	if _, err := m.getHOTPToken(); err != nil {
		return nil, err
	}
	return m, nil
}
