package viewmodels

import (
	"github.com/nba-simulation/model"
	"time"
)

type DoMatchTeamVm struct {
	MatchTitle   string
	StartingTime time.Time
	EndTime      time.Time
	Teams        model.Teams
}
