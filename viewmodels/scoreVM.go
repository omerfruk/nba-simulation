package viewmodels

import "github.com/nba-simulation/model"

type ScoreVM struct {
	ScoreType model.ScoreType
	TeamId    uint
	Team      model.Team
	ScorerId  int
	Scorer    PlayerVM
}

type ReturnVM struct {
	MatchName   string
	TeamScoreVM []TeamScoreVM
}

type TeamScoreVM struct {
	TeamName      string
	PlayerScoreVM []PlayerScoreVM
}
type PlayerScoreVM struct {
	PlayerName      string
	PlayerType      model.PlayerType
	PlayerScoreType model.ScoreType
}
