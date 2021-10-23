package viewmodels

import "github.com/nba-simulation/model"

type ShotVM struct {
	ShotType model.ScoreType
	OnShot   bool
	TimeOut  bool
}
