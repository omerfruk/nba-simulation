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

	match := make([]model.Match, 0)
	doMatchTeam := TeamPrepare()

	for _, vm := range doMatchTeam {
		wg.Add(1)
		go func(vm viewmodels.DoMatchTeamVm, group *sync.WaitGroup, matchType model.MatchType) {
			temp := DoMatch(vm, matchType)
			match = append(match, temp)
			wg.Done()
		}(vm, &wg, model.MatchtypeEnd)
	}
	wg.Wait()

	mep := make(map[string]map[string][]viewmodels.PlayerScoreVM)
	for _, m := range match {
		mep[fmt.Sprintf("%s Vs %s ", m.Teams[0].TeamName, m.Teams[1].TeamName)] = returnTypeTranslater(m)
	}

	var player []viewmodels.PlayerScoreVM
	var team []viewmodels.TeamScoreVM
	var retur []viewmodels.ReturnVM
	for teamsName, m := range mep {
		for teamName, vms := range m {
			for _, vm := range vms {
				player = append(player, vm)
			}
			team = append(team, viewmodels.TeamScoreVM{
				TeamName:      teamName,
				PlayerScoreVM: player,
			})

		}
		retur = append(retur, viewmodels.ReturnVM{
			MatchName:   teamsName,
			TeamScoreVM: team,
		})
		player = make([]viewmodels.PlayerScoreVM, 0)
		team = make([]viewmodels.TeamScoreVM, 0)
	}
	return c.JSON(retur)
}

func returnTypeTranslater(match model.Match) map[string][]viewmodels.PlayerScoreVM {
	m := make(map[string][]viewmodels.PlayerScoreVM)
	for _, score := range match.Scores {
		m[score.Team.TeamName] = append(m[score.Team.TeamName], viewmodels.PlayerScoreVM{
			PlayerId:        score.ScorerId,
			PlayerName:      score.Scorer.Name + " " + score.Scorer.Surname,
			PlayerType:      score.Scorer.PlayerType,
			PlayerScoreType: score.ScoreType,
		})
	}
	return m
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
		playerNumber = rand.Intn(7-0) + 0
		if player[playerNumber].PlayerType == model.PlayerTypePlayer {
			break
		}
	}
	// on shot no assister
	if shot.OnShot {
		score = append(score, viewmodels.ScoreVM{
			ScoreType: shot.ShotType,
			ScorerId:  int(player[playerNumber].ID),
			Scorer: viewmodels.PlayerVM{
				Name:       player[playerNumber].Name,
				Surname:    player[playerNumber].Surname,
				Number:     player[playerNumber].Number,
				TeamID:     player[playerNumber].TeamID,
				Team:       player[playerNumber].Team,
				PlayerType: player[playerNumber].PlayerType,
			},
		})
		return score
	}

	score = append(score, viewmodels.ScoreVM{
		ScoreType: shot.ShotType,
		ScorerId:  int(player[playerNumber].ID),
		Scorer: viewmodels.PlayerVM{
			Name:       player[playerNumber].Name,
			Surname:    player[playerNumber].Surname,
			Number:     player[playerNumber].Number,
			TeamID:     player[playerNumber].TeamID,
			Team:       player[playerNumber].Team,
			PlayerType: player[playerNumber].PlayerType,
		},
	})
	//assister player
	var assistNumber int
	for {
		assistNumber = rand.Intn(7-0) + 0
		if assistNumber != playerNumber && player[assistNumber].PlayerType == model.PlayerTypePlayer {
			break
		}
	}
	score = append(score, viewmodels.ScoreVM{
		ScoreType: model.ScoreTypeAssist,
		ScorerId:  int(player[assistNumber].ID),
		Scorer: viewmodels.PlayerVM{
			Name:       player[playerNumber].Name,
			Surname:    player[playerNumber].Surname,
			Number:     player[playerNumber].Number,
			TeamID:     player[playerNumber].TeamID,
			Team:       player[playerNumber].Team,
			PlayerType: player[playerNumber].PlayerType,
		},
	})
	return score
}

func DoMatch(matcher viewmodels.DoMatchTeamVm, matchType model.MatchType) model.Match {
	var match model.Match
	match.MatchType = matchType
	//ss := service.NewScoreService(database.DB())
	ms := service.NewMatchService(database.DB())
	ps := service.NewPlayerService(database.DB())
	matcher.StartingTime = time.Now()

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
		/*	if len(scores)%13 == 0 {
			ChangePlayer(*team1members)
			ChangePlayer(*team2members)
			//	time.Sleep(time.Second * 6)
			fmt.Println("Change Player")
		}*/
		if time.Now().Before(matcher.EndTime) && len(scores) <= 48 {
			scor := ShouterPlayer(*team1members)
			for _, s := range scor {
				scores = append(scores, model.Score{
					ScoreType: s.ScoreType,
					ScorerId:  s.ScorerId,
					TeamId:    team1.ID,
					Team:      team1,
					Scorer: model.Player{
						Name:       s.Scorer.Name,
						Surname:    s.Scorer.Surname,
						Number:     s.Scorer.Number,
						Team:       &s.Team,
						PlayerType: s.Scorer.PlayerType,
					},
				})

			}
			scor = ShouterPlayer(*team2members)
			for _, s := range scor {
				scores = append(scores, model.Score{
					ScoreType: s.ScoreType,
					ScorerId:  s.ScorerId,
					TeamId:    team2.ID,
					Team:      team2,
					Scorer: model.Player{
						Name:       s.Scorer.Name,
						Surname:    s.Scorer.Surname,
						Number:     s.Scorer.Number,
						Team:       &s.Team,
						PlayerType: s.Scorer.PlayerType,
					},
				})
			}
		} else {
			break
		}
	}
	// Score put in db

	/*

		acikca sotlemek gerekirse gorm'un azametine ugradim
		normalde score olustururken team ve player'in id si olmasina
		ragmen db de gereksiz yere bos player ve team lar olusturuyor
		1 gun ful buna ugrastim cozemedim zamanim da azaldigindan dolayi boyle bi
		yontem yapmak zorunda kaldim hatamin farkindayim

	*/

	/*if err = ss.CreateAll(&scores); err != nil {
		fmt.Println(err)
	}*/
	match.Teams = matcher.Teams
	if err = ms.Create(&match); err != nil {
		fmt.Println(err)
	}
	match.Scores = scores

	fmt.Println("wg bitti")
	return match
}

/*

gorm burda da id si olan veriyi olusturmakta israr ediyor
yeniden kaydetmek yerine tekrardan bos veri olusturuyor

*/

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
