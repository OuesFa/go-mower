package models

type (
	Boundaries struct {
		NorthLimit int
		EastLimit  int
	}
	Orientation rune
	Mower       struct {
		X           int
		Y           int
		Orientation string
	}
	Action struct {
		Instructions string `[RLF]+`
	}
)
