package model

import "gorm.io/gorm"

type PlayerType int

const (
	PlayerTypesubstitute PlayerType = iota
	PlayerTypePlayer
)

type Player struct {
	gorm.Model
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	PlayerType PlayerType `json:"player_type"`
}
