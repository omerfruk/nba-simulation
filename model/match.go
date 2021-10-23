package model

import "gorm.io/gorm"

type MatchType int

const (
	MatchTypeNotStart MatchType = iota
	MatchTypeStart
	MatchtypeEnd
	MatchTypeNextWeek
)

type Match struct {
	gorm.Model
	MatchType MatchType `json:"match_type"`
	Teams     Teams     `json:"teams" gorm:"many2many:match_teams;"`
	Scores    []Score   `json:"scores" gorm:"many2many:match_scores;"`
}
