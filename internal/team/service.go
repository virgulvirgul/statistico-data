package team

import (
	"github.com/joesweeny/sportmonks-go-client"
	"github.com/joesweeny/statshub/internal/model"
	"github.com/joesweeny/statshub/internal/season"
	"log"
	"sync"
)

type Service struct {
	Repository
	SeasonRepo season.Repository
	Factory
	Client *sportmonks.Client
	Logger *log.Logger
}

var waitGroup sync.WaitGroup

func (s Service) Process() error {
	ids, err := s.SeasonRepo.Ids()

	if err != nil {
		return err
	}

	teams := make(chan sportmonks.Team, len(ids))
	done := make(chan bool)

	go s.parseTeamsSync(teams, ids)
	go s.persistTeams(teams, done)

	<-done

	return nil
}

func (s Service) ProcessCurrentSeason() error {
	ids, err := s.SeasonRepo.CurrentSeasonIds()

	if err != nil {
		return err
	}

	s.parseTeamsAsync(ids)
	waitGroup.Wait()

	return nil
}

func (s Service) parseTeamsSync(ch chan<- sportmonks.Team, ids []int) {
	for _, id := range ids {
		res, err := s.Client.TeamsBySeasonId(id, []string{}, 5)

		if err != nil {
			log.Printf("Error when calling client. Message: %s", err.Error())
		}

		for _, team := range res.Data {
			ch <- team
		}
	}

	close(ch)
}

func (s Service) parseTeamsAsync(ids []int) {
	for _, id := range ids {
		waitGroup.Add(1)

		go func(id int) {
			res, err := s.Client.TeamsBySeasonId(id, []string{}, 5)

			if err != nil {
				log.Printf("Error when calling client. Message: %s", err.Error())
			}

			for _, team := range res.Data {
				waitGroup.Add(1)
				go func(team sportmonks.Team) {
					s.persistTeam(&team)
					defer waitGroup.Done()
				}(team)
			}

			defer waitGroup.Done()
		}(id)
	}
}

func (s Service) persistTeams(ch <-chan sportmonks.Team, done chan bool) {
	for team := range ch {
		s.persistTeam(&team)
	}

	done <- true
}

func (s Service) persistTeam(t *sportmonks.Team) {
	team, err := s.GetById(t.ID)

	if (err == ErrNotFound && model.Team{} == *team) {
		created := s.createTeam(t)

		if err := s.Insert(created); err != nil {
			log.Printf("Error '%s' occurred when inserting Team struct: %+v\n,", err.Error(), created)
		}

		return
	}

	updated := s.updateTeam(t, team)

	if err := s.Update(updated); err != nil {
		log.Printf("Error '%s' occurred when updating Team struct: %+v\n,", err.Error(), updated)
	}

	return
}