package database

import (
	"encoding/json"
	"fmt"
	"github.com/nba-simulation/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
)

const (
	HOST     = "localhost"
	DATABASE = "nba-simulation"
	USER     = "nba-simulation"
	PASSWORD = "guzelsifre"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}
func ConnectDB() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=5432", HOST, USER, PASSWORD, DATABASE)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return errors.Wrapf(err, "DB'ye bağlanılamadı: "+dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(2)
	return nil
}
func AutoMigrate() {
	db.AutoMigrate(model.Player{})
	db.AutoMigrate(model.Score{})
	db.AutoMigrate(model.Team{})
	db.AutoMigrate(model.Match{})
}

func CreateRandomTeamAndMember() {

	var playerModel []*model.Player
	var teams []model.Team

	type playerJson struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	jsonFile, err := os.Open("players.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonPlayers []playerJson
	json.Unmarshal(byteValue, &jsonPlayers)

	for i := 0; i < 8; i++ {
		teams = append(teams, model.Team{
			TeamName: fmt.Sprintf("Team %d", i+1),
		})
	}
	if err = db.Create(&teams).Error; err != nil {
		fmt.Println(err)
	}
	temp := 1
	for i, player := range jsonPlayers {
		var typePlayer model.PlayerType

		if temp <= 5 {
			typePlayer = model.PlayerTypePlayer
		} else {
			typePlayer = model.PlayerTypeSubstitute
		}
		var teamId uint
		teamId = uint((i / 8) + 1)
		playerModel = append(playerModel, &model.Player{
			Model:      gorm.Model{},
			Name:       player.FirstName,
			Surname:    player.LastName,
			Number:     i + 1,
			TeamID:     &teamId,
			PlayerType: typePlayer,
		})
		temp++
		if temp == 9 {
			temp = 0
		}
	}
	if err = db.Create(&playerModel).Error; err != nil {
		fmt.Println(err)
	}

}
