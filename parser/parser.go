package parser

import (
	"github.com/alexflint/go-restructure"
	"github.com/ouesfa/go-mower/models"
	"log"
	"strings"
)

func parse(input string) (models.Boundaries, []models.Mower, []models.Action) {
	lines := strings.Split(input, "\n")
	boundaries := ParseBoundaries(lines[0])
	var actions []models.Action
	var mowers []models.Mower
	for i, v := range lines[1:] {
		if action := ParseActions(v); action != nil && i%2 != 0 {
			actions = append(actions, *action)
		} else if mower := ParseMower(v); mower != nil && i%2 == 0 {
			mowers = append(mowers, *mower)
		}
	}
	return boundaries, mowers, actions
}
func ParseBoundaries(input string) models.Boundaries {
	var boundaries models.Boundaries
	match, _ := restructure.Find(&boundaries, input)
	if !match {
		panic("failed to parse boundaries")
	}
	return *models.NewBoundaries(boundaries.SNorthLimit, boundaries.SEastLimit)
}
func ParseMower(input string) *models.Mower {
	var mower models.Mower
	match, _ := restructure.Find(&mower, input)
	if !match {
		defer log.Println("failed to parse mower")
		return nil
	}
	return models.NewMower(mower.SX, mower.SY, mower.Orientation)
}
func ParseActions(input string) *models.Action {
	var actions models.Action
	match, _ := restructure.Find(&actions, input)
	if !match {
		defer log.Println("failed to parse actions")
		return nil
	}
	return &models.Action{Instructions: actions.Instructions}
}

//boundariesPattern := regexp.MustCompile(`\d \d`)
//wantPanic := boundariesPattern.FindString(input)
//limits := strings.Split(wantPanic, " ")
//northLimit, err0 := strconv.Atoi(limits[0])
//eastLimit, err1 := strconv.Atoi(limits[1])
//if err0 != nil || err1 != nil {
//	log.Fatal(err1, err0)
//	return nil
//} else {
//	return &mower.Boundaries{northLimit, eastLimit}
//}
