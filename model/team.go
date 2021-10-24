package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string   `json:"team_name"`
	Players  []Player `json:"players"`
}

type Teams []Team
