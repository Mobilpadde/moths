package emojies

import "github.com/enescakir/emoji"

var (
	FACES_SLICE = []string{
		emoji.PileOfPoo.String(),
		emoji.ClownFace.String(),
		emoji.Ogre.String(),
		emoji.Goblin.String(),
		emoji.Ghost.String(),
		emoji.Alien.String(),
		emoji.AlienMonster.String(),
		emoji.Robot.String(),
	}

	FACES = ToEmojies(FACES_SLICE)
)
