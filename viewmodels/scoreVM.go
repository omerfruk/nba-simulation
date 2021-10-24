package viewmodels

import "github.com/nba-simulation/model"

type ScoreVM struct {
	ScoreType model.ScoreType
	TeamId    int
	Team      model.Team
	ScorerId  int
	Scorer    PlayerVM
}
