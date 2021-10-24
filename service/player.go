package service

import (
	"github.com/nba-simulation/model"
	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	return &PlayerService{db: db}
}

func (s *PlayerService) Create(m *model.Player) error {
	err := s.db.Create(m).Error
	return err
}

func (s *PlayerService) GetAll() (*[]model.Player, error) {
	m := new([]model.Player)
	err := s.db.Limit(48).Find(&m).Error
	return m, err
}

func (s *PlayerService) GetById(id int64) (*model.Player, error) {
	m := new(model.Player)
	err := s.db.Where("id = ?", id).First(&m).Error
	return m, err
}
func (s *PlayerService) GetByMatchId(id int64) (*[]model.Player, error) {
	m := new([]model.Player)
	err := s.db.Limit(8).Where("team_id = ?", id).Find(&m).Error
	return m, err
}

func (s *PlayerService) Update(m *model.Player) error {
	err := s.db.Save(m).Error
	return err
}

func (s *PlayerService) Delete(m *model.Player) error {
	err := s.db.Delete(m).Error
	return err
}
