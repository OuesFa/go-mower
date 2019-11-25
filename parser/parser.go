package parser

import (
	"github.com/alexflint/go-restructure"
	"github.com/ouesfa/go-mower/models"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Parse : main parsing function, extracts all information needed from program string input
// returns boundaries
// also returns mowers & actions with indices in order to bind a mower to its own action in main program
func Parse(input string) (models.Boundaries, []models.MowerWithIndex, []models.ActionWithIndex) {
	lines := strings.Split(input, "\n")
	boundaries := parseBoundaries(lines[0])
	var actions []models.ActionWithIndex
	var mowers []models.MowerWithIndex
	for i, v := range lines[1:] {
		if action := parseAction(v); action != nil && i%2 != 0 {
			actions = append(actions, models.ActionWithIndex{Index: i, Action: *action})
		} else if mower := parseMower(v); mower != nil && i%2 == 0 {
			mowers = append(mowers, models.MowerWithIndex{Index: i, Mower: *mower})
		}
	}
	return boundaries, mowers, actions
}
func parseBoundaries(input string) models.Boundaries {
	const boundariesPattern = `^(\d)\s(\d)$`
	re := regexp.MustCompile(boundariesPattern)
	if matchedBoundaries := re.FindStringSubmatch(input); matchedBoundaries == nil {
		panic("failed to Parse boundaries")
	} else {
		matchedNorthLimit, _ := strconv.Atoi(matchedBoundaries[1])
		matchedEastLimit, _ := strconv.Atoi(matchedBoundaries[2])
		return models.Boundaries{
			NorthLimit: matchedNorthLimit,
			EastLimit:  matchedEastLimit,
		}
	}
}
func parseMower(input string) *models.Mower {
	const mowerPattern = `^(\d)\s(\d)\s([NSEW])$`
	re := regexp.MustCompile(mowerPattern)
	matchedMower := re.FindStringSubmatch(input)
	if matchedMower == nil {
		log.Println("failed to Parse mower from:", input)
		return nil
	}
	matchedX, _ := strconv.Atoi(matchedMower[1])
	matchedY, _ := strconv.Atoi(matchedMower[2])
	matchedOrientation := matchedMower[3]
	return &models.Mower{
		Orientation: matchedOrientation,
		X:           matchedX,
		Y:           matchedY,
	}
}
func parseAction(input string) *models.Action {
	var actions models.Action
	match, _ := restructure.Find(&actions, input)
	if !match {
		defer log.Println("failed to Parse actions")
		return nil
	}
	return &models.Action{Instructions: actions.Instructions}
}
