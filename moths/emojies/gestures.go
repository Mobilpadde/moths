package emojies

import "github.com/enescakir/emoji"

var (
	GESTURES_SLICE = []string{
		emoji.PersonFrowning.String(),
		emoji.PersonPouting.String(),
		emoji.PersonGesturingNo.String(),
		emoji.PersonGesturingOk.String(),
		emoji.PersonTippingHand.String(),
		emoji.PersonRaisingHand.String(),
		emoji.DeafPerson.String(),
		emoji.PersonBowing.String(),
		emoji.PersonFacepalming.String(),
		emoji.PersonShrugging.String(),
	}

	GESTURES = ToEmojies(GESTURES_SLICE)
)

// IDEA: Fix?
// Below does not work right now,
// they seem to be using "multple" emojies in one.
//
// Thus now exposed for usage

// var (
// 	GESTURES_MAN_SLICE = []string{
// 		emoji.ManFrowning.String(),
// 		emoji.ManPouting.String(),
// 		emoji.ManGesturingNo.String(),
// 		emoji.ManGesturingOk.String(),
// 		emoji.ManTippingHand.String(),
// 		emoji.ManRaisingHand.String(),
// 		emoji.DeafMan.String(),
// 		emoji.ManBowing.String(),
// 		emoji.ManFacepalming.String(),
// 		emoji.ManShrugging.String(),
// 	}

// 	GESTURES_MAN = ToEmojies(GESTURES_MAN_SLICE)
// )

// var (
// 	GESTURES_WOMAN_SLICE = []string{
// 		emoji.WomanFrowning.String(),
// 		emoji.WomanPouting.String(),
// 		emoji.WomanGesturingNo.String(),
// 		emoji.WomanGesturingOk.String(),
// 		emoji.WomanTippingHand.String(),
// 		emoji.WomanRaisingHand.String(),
// 		emoji.DeafWoman.String(),
// 		emoji.WomanBowing.String(),
// 		emoji.WomanFacepalming.String(),
// 		emoji.WomanShrugging.String(),
// 	}

// 	GESTURES_WOMAN = ToEmojies(GESTURES_WOMAN_SLICE)
// )
