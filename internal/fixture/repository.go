package fixture

import (
	"github.com/statistico/statistico-data/internal/model"
	"time"
)

type Repository interface {
	Insert(f *model.Fixture) error
	Update(f *model.Fixture) error
	ById(id uint64) (*model.Fixture, error)
	Ids() ([]int, error)
	IdsBetween(from, to time.Time) ([]int, error)
	Between(from, to time.Time) ([]model.Fixture, error)
	// Id of the Team concerned
	// Limit parameter to limit the number of Fixture structs returned
	// Date constraint returning fixtures from before that date
	ByTeamId(id int64, limit int32, before time.Time) ([]model.Fixture, error)
	// Id of the Season
	// Date constraint returning fixtures from before the given date
	BySeasonId(id int64, before time.Time) ([]model.Fixture, error)
	// ID of the Home Team concerned
	// ID of the Away Team concerned
	// Limit parameter to limit the number of Fixture structs returned
	// Date constraint returning fixtures from before that date
	ByHomeAndAwayTeam(homeTeamId, awayTeamId uint64, limit uint32, before time.Time) ([]model.Fixture, error)
	// Return a struct of Team IDs linked to the seasonId provided
	TeamIdsForSeason(seasonId uint64) ([]int, error)
}
