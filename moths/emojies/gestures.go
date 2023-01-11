package emojies

import "github.com/enescakir/emoji"

var GESTURES = ToEmojies([]string{
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
})

// IDEA: Fix?
// Below does not work atm, as they are using "multple" emojies in one

// var GESTURES_MAN = ToEmojies([]string{
// 	emoji.ManFrowning.String(),
// 	emoji.ManPouting.String(),
// 	emoji.ManGesturingNo.String(),
// 	emoji.ManGesturingOk.String(),
// 	emoji.ManTippingHand.String(),
// 	emoji.ManRaisingHand.String(),
// 	emoji.DeafMan.String(),
// 	emoji.ManBowing.String(),
// 	emoji.ManFacepalming.String(),
// 	emoji.ManShrugging.String(),
// })

// var GESTURES_WOMAN = ToEmojies([]string{
// 	emoji.WomanFrowning.String(),
// 	emoji.WomanPouting.String(),
// 	emoji.WomanGesturingNo.String(),
// 	emoji.WomanGesturingOk.String(),
// 	emoji.WomanTippingHand.String(),
// 	emoji.WomanRaisingHand.String(),
// 	emoji.DeafWoman.String(),
// 	emoji.WomanBowing.String(),
// 	emoji.WomanFacepalming.String(),
// 	emoji.WomanShrugging.String(),
// })
