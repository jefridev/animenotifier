package animeflv

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
)

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

// GetAnime returns pointer to a new single anime representation.
func GetAnime(sourceURL string) (*Anime, error) {
	doc, err := htmlquery.LoadURL(sourceURL)
	if err != nil {
		return nil, err
	}

	titleNode, err := htmlquery.Query(doc, titleXpath)
	if err != nil {
		return nil, err
	}

	typeNode, err := htmlquery.Query(doc, typeXpath)
	if err != nil {
		return nil, err
	}

	coverNode, err := htmlquery.Query(doc, coverXpath)
	if err != nil {
		return nil, err
	}

	statusNode, err := htmlquery.Query(doc, statusXpath)
	if err != nil {
		return nil, err
	}

	genresNodes, err := htmlquery.QueryAll(doc, genresXpath)
	if err != nil {
		return nil, err
	}

	synopsisNode, err := htmlquery.Query(doc, synopsisXpath)
	if err != nil {
		return nil, err
	}

	contentScriptNode, err := htmlquery.Query(doc, scriptContentXpath)
	if err != nil {
		return nil, err
	}

	anime := &Anime{}
	anime.Title = htmlquery.InnerText(titleNode)
	anime.Type = htmlquery.InnerText(typeNode)
	anime.Cover = htmlquery.InnerText(coverNode)
	anime.Status = htmlquery.InnerText(statusNode)
	for _, genre := range genresNodes {
		anime.Genres = append(anime.Genres, htmlquery.InnerText(genre))
	}

	contentScript := htmlquery.InnerText(contentScriptNode)
	nextRelease, _ := getNextReleaseDateFromString(contentScript)
	anime.NextRelease = nextRelease

	lastEpisode, _ := getLastEpisodeFromString(contentScript)
	anime.LastAiredEpisode = lastEpisode

	anime.Synopsis = htmlquery.InnerText(synopsisNode)
	return anime, nil
}

func getNextReleaseDateFromString(contentScript string) (*time.Time, error) {
	re := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}`)
	results := re.FindAllString(contentScript, 1)
	if len(results) > 0 {
		layout := "2006-01-02"
		releaseDate, err := time.Parse(layout, results[0])
		if err != nil {
			return nil, err
		}
		return &releaseDate, nil
	}

	return nil, errors.New("It couldn't find value to parse within content script")
}

func getLastEpisodeFromString(contentScript string) (int, error) {

	re := regexp.MustCompile(`\[([0-9]+)\,`)
	results := re.FindAllStringSubmatch(contentScript, 1)
	if len(results) > 0 {
		parseValue, err := strconv.Atoi(results[0][1])
		if err != nil {
			return 0, err
		}
		return parseValue, nil
	}
	return 0, errors.New("It couldn't find value to parse within content script")
}
