package notifierlib

import (
	"context"

	"firebase.google.com/go/db"
	"github.com/jefridev/animenotifier/pkg/animeflv"
)

// AnimeRepository interfaces that defines common behaviour for handling data.
type AnimeRepository interface {
	GetAll() ([]animeflv.Anime, error)
	AddRange(anime []animeflv.Anime) error
}

type repo struct {
	ctx context.Context
	db  *db.Client
}

func (r *repo) GetAll() ([]animeflv.Anime, error) {
	var animes []animeflv.Anime
	err := r.db.NewRef("animes").Get(r.ctx, &animes)
	if err != nil {
		return animes, err
	}
	return animes, nil
}

func (r *repo) AddRange(animes []animeflv.Anime) error {
	err := r.db.NewRef("animes").Set(r.ctx, animes)
	if err != nil {
		return err
	}
	return nil
}

//NewAnimeRepository returns a new reference type for repository.
func NewAnimeRepository(ctx context.Context, client *db.Client) AnimeRepository {
	return &repo{
		ctx: ctx,
		db:  client,
	}
}
