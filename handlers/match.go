package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nba-simulation/database"
	"github.com/nba-simulation/model"
	"github.com/nba-simulation/service"
	"github.com/nba-simulation/viewmodels"
	"time"
	//"github.com/nba-simulation/database"
	//"github.com/nba-simulation/service"
)

func MatchStarting(c *fiber.Ctx) error {

	doMatchTeam := TeamPrepare()
	fmt.Println(doMatchTeam)

	return c.Next()
}

func TeamPrepare() []viewmodels.DoMatchTeamVm {

	var doMatchTeams []viewmodels.DoMatchTeamVm
	var doMatchTeam viewmodels.DoMatchTeamVm
	//var doMatchTeam viewmodels.DoMatchTeamVm
	var towTeam model.Teams

	ts := service.NewTeamService(database.DB())
	teams, err := ts.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, team := range *teams {
		t := *teams
		if i%2 != 0 {
			towTeam = append(towTeam, team)
			towTeam = append(towTeam, t[i-1])
		}
		doMatchTeam = viewmodels.DoMatchTeamVm{
			MatchTitle:   fmt.Sprintf("%d. Match", (i/2)+1),
			StartingTime: time.Now(),
			EndTime:      time.Now().Add(time.Second * 240),
			Teams:        towTeam,
		}
		if i%2 != 0 {
			towTeam = nil
			doMatchTeams = append(doMatchTeams, doMatchTeam)
		}
	}

	return doMatchTeams
}
