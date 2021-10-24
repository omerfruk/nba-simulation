package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nba-simulation/database"
	"github.com/nba-simulation/model"
	"github.com/nba-simulation/service"
	"github.com/nba-simulation/viewmodels"
	"math/rand"
	"sync"
	"time"
)

func MatchStarting(c *fiber.Ctx) error {
	//adding wait group
	var wg sync.WaitGroup

	var match []model.Match
	doMatchTeam := TeamPrepare()

	for _, vm := range doMatchTeam {
		wg.Add(1)
		go func(vm viewmodels.DoMatchTeamVm, group sync.WaitGroup, matchType model.MatchType) {
			temp := DoMatch(vm, &wg, matchType)
			match = append(match, temp)
		}(vm, wg, model.MatchtypeEnd)
	}
	wg.Wait()

	return c.JSON(match)
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

func TakeShot() viewmodels.ShotVM {
	var shot viewmodels.ShotVM
	attackStartingTime := time.Now()
	attack := rand.Intn(10-1) + 1

	//attack fail
	for {
		if attack < 4 && !time.Now().Before(attackStartingTime.Add(time.Second*24)) {
			shot = viewmodels.ShotVM{
				ShotType: 0,
				OnShot:   false,
				TimeOut:  true,
			}
			return shot
		}
		break
	}
	onlyOneShot := rand.Intn(10-1) + 1
	if onlyOneShot > 7 {
		number := rand.Intn(5-1) + 1
		if number > 3 {
			shot = viewmodels.ShotVM{
				ShotType: model.ScoreTypeThreePoint,
				OnShot:   true,
				TimeOut:  false,
			}
			//time.Sleep(time.Second * 5)
			return shot
		}
		shot = viewmodels.ShotVM{
			ShotType: model.ScoreTypeDoublePoin,
			OnShot:   true,
			TimeOut:  false,
		}
		//time.Sleep(time.Second * 3)
		return shot
	}
	number := rand.Intn(5-1) + 1
	if number > 3 {
		shot = viewmodels.ShotVM{
			ShotType: model.ScoreTypeThreePoint,
			OnShot:   false,
			TimeOut:  false,
		}
		//time.Sleep(time.Second * 3)
		return shot
	}
	shot = viewmodels.ShotVM{
		ShotType: model.ScoreTypeDoublePoin,
		OnShot:   false,
		TimeOut:  false,
	}
	// time.Sleep(time.Second * 2)
	return shot
}

func ShouterPlayer(player []model.Player) []viewmodels.ScoreVM {
	var score []viewmodels.ScoreVM
	shot := TakeShot()

	//is player Substitute
	var playerNumber int
	for {
		playerNumber = rand.Intn(8-0) + 0
		if player[playerNumber].PlayerType == model.PlayerTypePlayer {
			break
		}
	}
	// on shot no assister
	if shot.OnShot {
		score = append(score, viewmodels.ScoreVM{
			ScoreType: shot.ShotType,
			ScorerId:  int(player[playerNumber].ID),
			Scorer:    player[playerNumber],
		})
		return score
	}

	score = append(score, viewmodels.ScoreVM{
		ScoreType: shot.ShotType,
		ScorerId:  int(player[playerNumber].ID),
		Scorer:    player[playerNumber],
	})
	//assister player
	var assistNumber int
	for {
		assistNumber = rand.Intn(8-1) + 1
		if assistNumber != playerNumber && player[assistNumber].PlayerType == model.PlayerTypePlayer {
			break
		}
	}
	score = append(score, viewmodels.ScoreVM{
		ScoreType: model.ScoreTypeAssist,
		ScorerId:  int(player[assistNumber].ID),
		Scorer:    player[assistNumber],
	})
	return score
}

func DoMatch(matcher viewmodels.DoMatchTeamVm, wg *sync.WaitGroup, matchType model.MatchType) model.Match {
	var match model.Match
	match.MatchType = matchType
	ss := service.NewScoreService(database.DB())
	ms := service.NewMatchService(database.DB())
	matcher.StartingTime = time.Now()
	ps := service.NewPlayerService(database.DB())
	team1 := matcher.Teams[0]
	team2 := matcher.Teams[1]

	team1members, err := ps.GetByMatchId(int64(team1.ID))
	if err != nil {
		fmt.Println(err)
	}
	team2members, err := ps.GetByMatchId(int64(team2.ID))
	if err != nil {
		fmt.Println(err)
	}

	matcher.EndTime = time.Now().Add(time.Second * 240)
	fmt.Println(time.Now())
	fmt.Println(matcher.EndTime)
	var scores []model.Score
	for {
		if len(scores)%13 == 0 {
			ChangePlayer(*team1members)
			ChangePlayer(*team2members)
			//	time.Sleep(time.Second * 6)
			fmt.Println("Change Player")
		}
		if time.Now().Before(matcher.EndTime) && len(scores) <= 48 {
			scor := ShouterPlayer(*team1members)
			for _, s := range scor {
				scores = append(scores, model.Score{
					ScoreType: s.ScoreType,
					ScorerId:  s.ScorerId,
					TeamId:    team1.ID,
					Team:      team1,
					Scorer:    s.Scorer,
				})

			}
			scor = ShouterPlayer(*team2members)
			for _, s := range scor {
				scores = append(scores, model.Score{
					ScoreType: s.ScoreType,
					ScorerId:  s.ScorerId,
					TeamId:    team2.ID,
					Team:      team2,
					Scorer:    s.Scorer,
				})
			}
		} else {
			break
		}
	}
	// Score put in db
	if err = ss.CreateAll(&scores); err != nil {
		fmt.Println(err)
	}
	match.Scores = scores
	match.Teams = matcher.Teams
	if err = ms.Create(&match); err != nil {
		fmt.Println(err)
	}
	fmt.Println("wg bitti")
	wg.Done()
	return match
}

func ChangePlayer(players []model.Player) {
	ps := service.NewPlayerService(database.DB())
	var substitute model.Player
	var player model.Player

	for _, m := range players {
		if substitute.ID == 0 && player.ID == 0 {
			if m.PlayerType == model.PlayerTypePlayer {
				player = m
				player.PlayerType = model.PlayerTypeSubstitute
			}
			if m.PlayerType == model.PlayerTypeSubstitute {
				substitute = m
				substitute.PlayerType = model.PlayerTypePlayer
			}
		} else {
			if err := ps.Update(&player); err != nil {
				fmt.Println(err)
			}
			if err := ps.Update(&substitute); err != nil {
				fmt.Println(err)
			}
			break
		}
	}

}
