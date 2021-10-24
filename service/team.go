package service

import (
	"github.com/nba-simulation/model"
	"gorm.io/gorm"
)

type TeamService struct {
	db *gorm.DB
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{db: db}
}

func (s *TeamService) Create(m *model.Team) error {
	err := s.db.Create(m).Error
	return err
}

func (s *TeamService) GetAll() (*[]model.Team, error) {
	m := new([]model.Team)
	err := s.db.Limit(6).Find(&m).Error
	return m, err
}

func (s *TeamService) GetById(id int64) (*model.Team, error) {
	m := new(model.Team)
	err := s.db.Where("id = ?", id).First(&m).Error
	return m, err
}

func (s *TeamService) Update(m *model.Team) error {
	err := s.db.Save(m).Error
	return err
}

func (s *TeamService) Delete(m *model.Team) error {
	err := s.db.Delete(m).Error
	return err
}
