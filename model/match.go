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
	Team       []Team
	TotalScore []Score
}
