package notifierlib

import (
	"log"

	"github.com/jefridev/animenotifier/pkg/animeflv"
)

// AnimeService handles logic operations.
type AnimeService interface {
	LoadAnimeToFirebase() error
}

type service struct {
	r AnimeRepository
	l *log.Logger
}

func (s *service) LoadAnimeToFirebase() error {
	animes, err := animeflv.GetAllCurrentlyAiringShows()
	if err != nil {
		s.l.Printf("It was not possible to extract anime information %v\n", err)
		return err
	}

	err = s.r.AddRange(animes)
	if err != nil {
		s.l.Printf("It was not possible to save extracted animes %v\n", err)
		return err
	}
	return nil
}

// NewAnimeService returns a new service.
func NewAnimeService(log *log.Logger, repository AnimeRepository) AnimeService {
	return &service{
		l: log,
		r: repository,
	}
}
