package emojies

import "unicode/utf8"

func ToEmojies(emojies []string) Emojies {
	em := Emojies{}

	for i := 0; i < len(emojies); i++ {
		current := emojies[i]
		for j := 0; j < len(current); {
			r, width := utf8.DecodeRuneInString(current[j:])
			em = append(em, r)

			j += width
		}
	}

	return em
}
