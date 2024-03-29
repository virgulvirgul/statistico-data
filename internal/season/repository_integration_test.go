package season

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/statistico/statistico-data/internal/config"
	"github.com/statistico/statistico-data/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresSeasonRepository{Connection: conn}

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i < 4; i++ {
			s := newSeason(i, true)

			if err := repo.Insert(s); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from sportmonks_season")

			var count int

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, i, count)
		}
	})

	t.Run("returns error when ID primary key violates unique constraint", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newSeason(10, true)

		if err := repo.Insert(c); err != nil {
			t.Errorf("Test failed, expected nil, got %s", err)
		}

		if e := repo.Insert(c); e == nil {
			t.Fatalf("Test failed, expected %s, got nil", e)
		}
	})

	conn.Close()
}

func TestUpdate(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresSeasonRepository{Connection: conn}

	t.Run("modifies existing record", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		s := newSeason(50, true)

		if err := repo.Insert(s); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		var d = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)

		s.IsCurrent = false
		s.LeagueID = 2
		s.UpdatedAt = d

		if err := repo.Update(s); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.Id(50)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(50, r.ID)
		a.Equal("2018-2019", r.Name)
		a.Equal(false, r.IsCurrent)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns an error if record does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newSeason(146, true)

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}
	})

	conn.Close()
}

func TestGetById(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresSeasonRepository{Connection: conn}

	t.Run("season can be retrieved by ID", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		s := newSeason(146, true)

		err := repo.Update(s)

		if err == nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}

		if err := repo.Insert(s); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		r, err := repo.Id(146)

		if err != nil {
			t.Errorf("Error when retrieving a record from the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(146, r.ID)
		a.Equal("2018-2019", r.Name)
		a.Equal(560, r.LeagueID)
		a.True(r.IsCurrent)
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns error if season does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		_, err := repo.Id(4)

		if err == nil {
			t.Errorf("Test failed, expected %v, got nil", err)
		}

		if err != ErrNotFound {
			t.Fatalf("Test failed, expected %v, got %s", ErrNotFound, err)
		}
	})

	conn.Close()
}

func TestGetIds(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresSeasonRepository{Connection: conn}

	t.Run("test returns a slice of int ids", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i <= 4; i++ {
			s := newSeason(i, true)

			if err := repo.Insert(s); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}
		}

		ids, err := repo.Ids()

		want := []int{1, 2, 3, 4}

		if err != nil {
			t.Fatalf("Test failed, expected %v, got %s", want, err.Error())
		}

		assert.Equal(t, want, ids)
	})
}

func TestCurrentSeasonIds(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresSeasonRepository{Connection: conn}

	t.Run("returns records with is current season set to true", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		seasons := []model.Season{}

		for i := 1; i <= 4; i++ {
			s := newSeason(i, true)

			seasons = append(seasons, *s)

			if err := repo.Insert(s); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}
		}

		if err := repo.Insert(newSeason(10, false)); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		retrieved, err := repo.CurrentSeasonIds()

		if err != nil {
			t.Fatalf("Test failed, expected %v, got %s", seasons, err.Error())
		}

		assert.Equal(t, []int{1, 2, 3, 4}, retrieved)
	})
}

var db = config.GetConfig().Database

func getConnection(t *testing.T) (*sql.DB, func()) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)

	db, err := sql.Open(db.Driver, psqlInfo)

	if err != nil {
		panic(err)
	}

	return db, func() {
		_, err := db.Exec("delete from sportmonks_season")
		if err != nil {
			t.Fatalf("Failed to clear database. %s", err.Error())
		}
	}
}

func newSeason(id int, current bool) *model.Season {
	return &model.Season{
		ID:        id,
		Name:      "2018-2019",
		LeagueID:  560,
		IsCurrent: current,
		CreatedAt: time.Unix(1546965200, 0),
		UpdatedAt: time.Unix(1546965200, 0),
	}
}
