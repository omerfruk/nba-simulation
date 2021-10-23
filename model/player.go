package model

import "gorm.io/gorm"

type PlayerType int

const (
	PlayerTypeSubstitute PlayerType = iota
	PlayerTypePlayer
)

type Player struct {
	gorm.Model
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Number     int        `json:"number"`
	TeamID     *uint      `json:"team_id"`
	Team       *Team      `json:"team"`
	PlayerType PlayerType `json:"player_type"`
}
