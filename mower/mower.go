package main

import (
	"fmt"
	"github.com/ouesfa/go-mower/models"
	"github.com/ouesfa/go-mower/parser"
)

func main() {
	lines := `5 5
1 2 N
LFLFLFLFF
sdf dsf sdcx
FFRFFRFRRF
3 3 E
FFRFFRFRRF
3 3 E
fdgs sfdg
3 3 E dsf
FFRFFRFRRF`

	boundaries, mowersWithIndices, actionsWithIndices := parser.Parse(lines)
	var movedMowers []models.Mower
	for _, mowerWithIndex := range mowersWithIndices {
		for _, actionWithIndex := range actionsWithIndices {
			mowerCorrespondsToAction := actionWithIndex.Index-mowerWithIndex.Index == 1
			if mowerCorrespondsToAction {
				for _, instruction := range actionWithIndex.Instructions {
					mowerWithIndex.Mower = move(instruction, mowerWithIndex.Mower, boundaries)
				}
				movedMowers = append(movedMowers, mowerWithIndex.Mower)
			}
		}
	}
	fmt.Println("input:\n" + lines)
	fmt.Println("\nfinal positions:\n", movedMowers)
}

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
