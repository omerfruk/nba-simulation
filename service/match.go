package service

import (
	"github.com/nba-simulation/model"
	"gorm.io/gorm"
)

type MatchService struct {
	db *gorm.DB
}

func NewMatchService(db *gorm.DB) *MatchService {
	return &MatchService{db: db}
}

func (s *MatchService) Create(m *model.Match) error {
	err := s.db.Create(m).Error
	return err
}

func (s *MatchService) GetAll() (*[]model.Match, error) {
	m := new([]model.Match)
	err := s.db.Find(&m).Error
	return m, err
}

func (s *MatchService) GetById(id int64) (*model.Match, error) {
	m := new(model.Match)
	err := s.db.Where("id = ?", id).First(&m).Error
	return m, err
}

func (s *MatchService) Update(m *model.Match) error {
	err := s.db.Save(m).Error
	return err
}

func (s *MatchService) Delete(m *model.Match) error {
	err := s.db.Delete(m).Error
	return err
}
