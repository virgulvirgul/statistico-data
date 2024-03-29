package fixture

import (
	"github.com/jonboulle/clockwork"
	"github.com/statistico/sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var t = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)
var clock = clockwork.NewFakeClockAt(t)
var f = Factory{clock}

func TestFactoryCreateFixture(t *testing.T) {
	t.Run("a new domain fixture struct is hydrated", func(t *testing.T) {
		t.Helper()

		s := f.createFixture(newClientFixture())

		a := assert.New(t)
		a.Equal(34, s.ID)
		a.Equal(987, s.SeasonID)
		a.Equal(451, *s.VenueID)
		a.Equal(4, s.HomeTeamID)
		a.Equal(98, s.AwayTeamID)
		a.Nil(s.RoundID)
		a.Nil(s.RefereeID)
		a.Equal("2019-03-12T19:45:00Z", s.Date.Format(time.RFC3339))
		a.Equal("2019-01-14T11:25:00Z", s.CreatedAt.Format(time.RFC3339))
		a.Equal("2019-01-14T11:25:00Z", s.UpdatedAt.Format(time.RFC3339))
	})
}

func TestFactoryUpdateFixture(t *testing.T) {
	t.Run("updates an existing fixture struct", func(t *testing.T) {
		t.Helper()

		clientFixture := newClientFixture()

		c := f.createFixture(clientFixture)

		clock.Advance(10 * time.Minute)

		ref := 32
		round := 2

		clientFixture.Time.StartingAt.Timestamp = 1552420800
		clientFixture.RefereeID = &ref
		clientFixture.RoundID = &round
		clientFixture.VenueID = &ref

		updated := f.updateFixture(clientFixture, c)

		a := assert.New(t)
		a.Equal(34, updated.ID)
		a.Equal(987, updated.SeasonID)
		a.Equal(32, *updated.VenueID)
		a.Equal(4, updated.HomeTeamID)
		a.Equal(98, updated.AwayTeamID)
		a.Equal(2, *updated.RoundID)
		a.Equal(32, *updated.RefereeID)
		a.Equal("2019-03-12T20:00:00Z", updated.Date.Format(time.RFC3339))
		a.Equal("2019-01-14T11:25:00Z", updated.CreatedAt.Format(time.RFC3339))
		a.Equal("2019-01-14T11:35:00Z", updated.UpdatedAt.Format(time.RFC3339))
	})
}

func newClientFixture() *sportmonks.Fixture {
	var venue = 451
	var t = sportmonks.FixtureTime{
		Status: "Live",
		StartingAt: struct {
			DateTime  string `json:"date_time"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Timestamp int64  `json:"timestamp"`
			Timezone  string `json:"timezone"`
		}{
			DateTime:  "",
			Date:      "",
			Time:      "",
			Timestamp: 1552419900,
			Timezone:  "",
		},
	}

	return &sportmonks.Fixture{
		ID:            34,
		LeagueID:      1356,
		SeasonID:      987,
		VenueID:       &venue,
		LocalTeamID:   4,
		VisitorTeamID: 98,
		Time:          t,
	}
}
