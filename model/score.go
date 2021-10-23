package model

import "gorm.io/gorm"

type ScoreType int

const (
	ScoreTypeAssist     ScoreType = 0
	ScoreTypeDoublePoin           = 2
	ScoreTypeThreePoint           = 3
)

type Score struct {
	gorm.Model
	ScoreType ScoreType `json:"score_type"`
}
