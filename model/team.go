package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string   `json:"team_name"`
	Match    []*Match `json:"match"gorm:"many2many:match_teams;"`
	Players  []Player `json:"players"`
}

type Teams []Team
