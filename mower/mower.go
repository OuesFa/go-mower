package mower

import (
	"github.com/ouesfa/go-mower/models"
)

func move(instruction rune, mower models.Mower, boundaries models.Boundaries) models.Mower {
	switch instruction {
	case 'L':
		mower.Orientation = leftOrientation[mower.Orientation]
	case 'R':
		mower.Orientation = rightOrientation[mower.Orientation]
	case 'F':
		mower = moveForward(mower, boundaries)
	}
	return mower
}

func moveForward(mower models.Mower, boundaries models.Boundaries) models.Mower {
	switch mower.Orientation {
	case "N":
		if mower.Y < boundaries.NorthLimit {
			mower.Y += 1
		}
	case "S":
		if mower.Y > 0 {
			mower.Y -= 1
		}
	case "E":
		if mower.X < boundaries.EastLimit {
			mower.X += 1
		}
	case "W":
		if mower.X > 0 {
			mower.X -= 1
		}
	}
	return mower
}

var rightOrientation = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}
var leftOrientation = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}
