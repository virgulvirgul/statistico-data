package fixture

import (
	pb "github.com/joesweeny/statistico-data/proto/fixture"
	"time"
	"log"
	"errors"
)

var ErrTimeParse = errors.New("unable to parse date provided in Request")

type Service struct {
	Repository
	Handler
	Logger *log.Logger
}

func (s *Service) ListFixtures(r *pb.Request, stream pb.FixtureService_ListFixturesServer) error {
	from, err := time.Parse(time.RFC3339, r.DateFrom)

	if err != nil {
		return ErrTimeParse
	}

	to, err := time.Parse(time.RFC3339, r.DateTo)

	if err != nil {
		return ErrTimeParse
	}

	fixtures, err := s.Repository.Between(from, to)

	if err != nil {
		return err
	}

	for _, fix := range fixtures {
		f, err := s.HandleFixture(&fix)

		if err != nil {
			return err
		}

		if err := stream.Send(f); err != nil {
			s.Logger.Printf("Error hydrating Fixture. ID: %d. Error: %s", fix.ID, err.Error())
			return err
		}
	}

	return nil
}
