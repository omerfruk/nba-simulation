package service

import (
	"github.com/nba-simulation/model"
	"gorm.io/gorm"
)

type ScoreService struct {
	db *gorm.DB
}

func NewScoreService(db *gorm.DB) *ScoreService {
	return &ScoreService{db: db}
}

func (s *ScoreService) Create(m *model.Score) error {
	err := s.db.Create(m).Error
	return err
}

func (s *ScoreService) GetById(id int64) (*model.Score, error) {
	m := new(model.Score)
	err := s.db.Where("id = ?", id).First(&m).Error
	return m, err
}

func (s *ScoreService) Update(m *model.Score) error {
	err := s.db.Save(m).Error
	return err
}

func (s *ScoreService) Delete(m *model.Score) error {
	err := s.db.Delete(m).Error
	return err
}
