package viewmodels

import "github.com/nba-simulation/model"

type PlayerVM struct {
	Name       string
	Surname    string
	Number     int
	TeamID     *uint
	Team       *model.Team
	PlayerType model.PlayerType
}
