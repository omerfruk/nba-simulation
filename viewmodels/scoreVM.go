package viewmodels

import "github.com/nba-simulation/model"

type ScoreVM struct {
	ScoreType model.ScoreType
	ScorerId  int
	Scorer    model.Player
}
