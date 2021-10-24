package model

import "gorm.io/gorm"

type ScoreType int

const (
	ScoreTypeAssist ScoreType = iota
	ScoreTypeNotYet
	ScoreTypeDoublePoin
	ScoreTypeThreePoint
)

type Score struct {
	gorm.Model
	ScoreType ScoreType `json:"score_type"`
	ScorerId  int       `json:"scorer_id"`
	TeamId    uint      `json:"team_id"`
	Team      Team      `json:"match"`
	Scorer    Player    `json:"scorer"`
}
