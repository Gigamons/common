package consts

// Playmodes
const (
	STD   = 0
	Taiko = 1
	CTB   = 2
	Mania = 3
)

// ToPlaymodeString convert an int8 into an Readable string. (0, 1, 2, 3) -> ("std", "taiko", "ctb", "mania")
func ToPlaymodeString(p byte) string {
	switch p {
	case STD:
		return "std"
	case Taiko:
		return "taiko"
	case CTB:
		return "ctb"
	case Mania:
		return "mania"
	}
	return "std"
}
