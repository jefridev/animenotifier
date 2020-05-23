package animeflv

import "time"

// Anime holds attribute values.
type Anime struct {
	Title            string     `json:"title"`
	Type             string     `json:"type"`
	Cover            string     `json:"cover"`
	Status           string     `json:"status"`
	Synopsis         string     `json:"synopsis"`
	LastAiredEpisode int        `json:"lastAiredEpisode"`
	NextRelease      *time.Time `json:"nextRelease"`
	Genres           []string   `json:"genres"`
}
