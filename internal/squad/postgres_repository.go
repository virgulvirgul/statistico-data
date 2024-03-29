package squad

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/statistico/statistico-data/internal/model"
	"strconv"
	"time"
)

var ErrNotFound = errors.New("not found")

type PostgresSquadRepository struct {
	Connection *sql.DB
}

func (p *PostgresSquadRepository) Insert(m *model.Squad) error {
	query := `
	INSERT INTO sportmonks_squad (season_id, team_id, player_ids, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := p.Connection.Exec(query, m.SeasonID, m.TeamID, pq.Array(m.PlayerIDs), m.CreatedAt.Unix(), m.UpdatedAt.Unix())

	return err
}

func (p *PostgresSquadRepository) Update(m *model.Squad) error {
	if _, err := p.BySeasonAndTeam(m.SeasonID, m.TeamID); err != nil {
		return err
	}

	query := `
	UPDATE sportmonks_squad set player_ids = $3, updated_at = $4 where season_id = $1 and team_id = $2`

	_, err := p.Connection.Exec(query, m.SeasonID, m.TeamID, pq.Array(m.PlayerIDs), m.UpdatedAt.Unix())

	return err
}

func (p *PostgresSquadRepository) BySeasonAndTeam(seasonId, teamId int) (*model.Squad, error) {
	query := `SELECT * FROM sportmonks_squad where season_id = $1 AND team_id = $2`

	m := model.Squad{}

	var players []string
	var created int64
	var updated int64

	err := p.Connection.QueryRow(query, seasonId, teamId).Scan(&m.SeasonID, &m.TeamID, pq.Array(&players), &created, &updated)

	if err != nil {
		return &m, ErrNotFound
	}

	for _, i := range players {
		text, _ := strconv.Atoi(i)
		m.PlayerIDs = append(m.PlayerIDs, text)
	}

	m.CreatedAt = time.Unix(created, 0)
	m.UpdatedAt = time.Unix(updated, 0)

	return &m, err
}

func (p *PostgresSquadRepository) All() ([]model.Squad, error) {
	query := `SELECT * FROM sportmonks_squad order by season_id ASC, team_id ASC`

	var squads []model.Squad

	rows, err := p.Connection.Query(query)

	if err != nil {
		return squads, err
	}

	return parseRows(rows, squads)
}

func (p *PostgresSquadRepository) CurrentSeason() ([]model.Squad, error) {
	query := `SELECT * FROM sportmonks_squad WHERE season_id in (SELECT id from sportmonks_season WHERE is_current = true)
 	order by season_id ASC, team_id ASC`

	var squads []model.Squad

	rows, err := p.Connection.Query(query)

	if err != nil {
		return squads, err
	}

	return parseRows(rows, squads)
}

func parseRows(r *sql.Rows, m []model.Squad) ([]model.Squad, error) {
	for r.Next() {
		var players []string
		var created int64
		var updated int64
		var squad model.Squad

		if err := r.Scan(&squad.SeasonID, &squad.TeamID, pq.Array(&players), &created, &updated); err != nil {
			return m, err
		}

		for _, i := range players {
			text, _ := strconv.Atoi(i)
			squad.PlayerIDs = append(squad.PlayerIDs, text)
		}

		squad.CreatedAt = time.Unix(created, 0)
		squad.UpdatedAt = time.Unix(updated, 0)

		m = append(m, squad)
	}

	return m, nil
}
