package models

import "strconv"

type (
	Boundaries struct {
		_           struct{} `^`
		SNorthLimit string   `\d`
		_           struct{} `\s`
		SEastLimit  string   `\d`
		_           struct{} `$`
		NorthLimit  int
		EastLimit   int
	}
	Orientation rune
	Mower       struct {
		_           string `^`
		SX          string `\d`
		_           string `\s`
		SY          string `\d`
		_           string `\s`
		Orientation string `[NSEW]`
		_           string `$`
		X           int
		Y           int
	}
	Action struct {
		Instructions string `[RLF]+`
	}
)

func NewMower(x string, y string, orientation string) *Mower {
	i, _ := strconv.Atoi(x)
	j, _ := strconv.Atoi(y)
	return &Mower{SX: x, SY: y, Orientation: orientation, X: i, Y: j}
}

func NewBoundaries(northLimit string, eastLimit string) *Boundaries {
	n, _ := strconv.Atoi(northLimit)
	e, _ := strconv.Atoi(eastLimit)
	return &Boundaries{SNorthLimit: northLimit, SEastLimit: eastLimit, NorthLimit: n, EastLimit: e}
}
