package emojies

import "unicode/utf8"

// ToEmojies can be used to convert a
// slice of strings to [Emojies].
func ToEmojies(emojies []string) Emojies {
	em := Emojies{}

	for i := 0; i < len(emojies); i++ {
		current := emojies[i]
		for j := 0; j < len(current); {
			r, width := utf8.DecodeRuneInString(current[j:])
			em = append(em, Emoji(r))

			j += width
		}
	}

	return em
}
