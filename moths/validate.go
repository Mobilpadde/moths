package moths

import (
	"time"

	"github.com/Mobilpadde/moths/moths/otp"
)

func (m *Moths) Validate(moth string) bool {
	if len(moth)/otp.EmojiBytes != m.amount {
		return false
	}

	token, err := m.getToken(false)
	if err != nil {
		return false
	}

	same, err := otp.NewOTP(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	// Bump the time for
	// m.getToken(true)
	m.time = m.time.Add(time.Since(m.time))

	return same.Validate(moth)
}

func (m *Moths) ValidateToken(code string) bool {
	token, err := m.getToken(false)
	if err != nil {
		return false
	}

	return token == code
}

// Helpful?
// newOTP, err := otp.NewOTP(token, m.emojies)
// return newOTP.Validate(code), err
// IDEA: Reverse from emojies and validate towards `getToken()`
// ---
// func (m *Moths) Validate(emojies string) (string, error) {
// 	size := emojiBytes * m.amount

// 	// https://github.com/tilaklodha/google-authenticator
// 	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(m.emojies)-1|1))) - 1
// 	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(m.emojies))))

// 	id := new(strings.Builder)
// 	id.Grow(size)

// 	tkn, err := m.getToken()
// 	if err != nil {
// 		return "", err
// 	}

// 	buffer := bufferFromToken(tkn)
// 	for {
// 		randomBuffer, err := buffer(step)
// 		if err != nil {
// 			return "", err
// 		}

// 		for i := 0; i < step; i++ {
// 			currentIndex := int(randomBuffer[i]) & mask

// 			if currentIndex < len(m.emojies) {
// 				if _, err := id.WriteString(m.emojies[currentIndex]); err != nil {
// 					return "", err
// 				} else if id.Len() == size {
// 					return id.String(), nil
// 				}
// 			}
// 		}
// 	}
// }

// func ValidateCustom(passcode string, counter uint64, secret string, opts ValidateOpts) (bool, error) {
// 	passcode = strings.TrimSpace(passcode)

// 	if len(passcode) != opts.Digits.Length() {
// 		return false, otp.ErrValidateInputInvalidLength
// 	}

// 	otpstr, err := GenerateCodeCustom(secret, counter, opts)
// 	if err != nil {
// 		return false, err
// 	}

// 	if subtle.ConstantTimeCompare([]byte(otpstr), []byte(passcode)) == 1 {
// 		return true, nil
// 	}

// 	return false, nil
// }

// func GenerateCodeCustom(secret string, counter uint64, opts ValidateOpts) (passcode string, err error) {
// 	//Set default value
// 	if opts.Digits == 0 {
// 		opts.Digits = otp.DigitsSix
// 	}
// 	// As noted in issue #10 and #17 this adds support for TOTP secrets that are
// 	// missing their padding.
// 	secret = strings.TrimSpace(secret)
// 	if n := len(secret) % 8; n != 0 {
// 		secret = secret + strings.Repeat("=", 8-n)
// 	}

// 	// As noted in issue #24 Google has started producing base32 in lower case,
// 	// but the StdEncoding (and the RFC), expect a dictionary of only upper case letters.
// 	secret = strings.ToUpper(secret)

// 	secretBytes, err := base32.StdEncoding.DecodeString(secret)
// 	if err != nil {
// 		return "", otp.ErrValidateSecretInvalidBase32
// 	}

// 	buf := make([]byte, 8)
// 	mac := hmac.New(opts.Algorithm.Hash, secretBytes)
// 	binary.BigEndian.PutUint64(buf, counter)
// 	if debug {
// 		fmt.Printf("counter=%v\n", counter)
// 		fmt.Printf("buf=%v\n", buf)
// 	}

// 	mac.Write(buf)
// 	sum := mac.Sum(nil)

// 	// "Dynamic truncation" in RFC 4226
// 	// http://tools.ietf.org/html/rfc4226#section-5.4
// 	offset := sum[len(sum)-1] & 0xf
// 	value := int64(((int(sum[offset]) & 0x7f) << 24) |
// 		((int(sum[offset+1] & 0xff)) << 16) |
// 		((int(sum[offset+2] & 0xff)) << 8) |
// 		(int(sum[offset+3]) & 0xff))

// 	l := opts.Digits.Length()
// 	mod := int32(value % int64(math.Pow10(l)))

// 	if debug {
// 		fmt.Printf("offset=%v\n", offset)
// 		fmt.Printf("value=%v\n", value)
// 		fmt.Printf("mod'ed=%v\n", mod)
// 	}

// 	return opts.Digits.Format(mod), nil
// }
