package models

type (
	// Boundaries : lawn limits
	Boundaries struct {
		NorthLimit int
		EastLimit  int
	}
	// Orientation : mower orientation
	Orientation rune
	// Mower : position & orientation of a mower
	Mower struct {
		X           int
		Y           int
		Orientation string
	}
	// Action : set of instructions to apply to a mower
	Action struct {
		Instructions string `[RLF]+`
	}
	// MowerWithIndex : add index in order to bind a mower to its own action
	MowerWithIndex struct {
		Index int
		Mower
	}
	// ActionWithIndex : add index in order to bind an action its own mower
	ActionWithIndex struct {
		Index int
		Action
	}
)
