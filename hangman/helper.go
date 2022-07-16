package hangman

var Words = [47]string{
	"python", "apple", "orange",
	"existansialism", "disgusting", "christ", "yummy", "ronaldo",
	"bitcoin", "football", "cryptocurrency", "corridor", "execution",
	"whiteboard", "delicious", "exciting", "immortality", "dangrous",
	"cosmon", "homeless", "psycology", "necklace", "zebra", "rhinoceros",
	"japan", "spider", "monnument", "pyramid", "avenue", "antidisestablishmentarianism",
	"lantern", "liberty", "shoulder", "finger", "thumb", "locket", "monsoone", "avadacadabra",
	"ethereum", "pronunciation", "naruto", "neverland", "golang", "france", "hogrider", "bottle", "heavymetal",
}

type WordStruct struct {
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Pronunciation string `json:"pronunciation"`
}