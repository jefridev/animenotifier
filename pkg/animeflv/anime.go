package animeflv

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// AnimeUtils contains encode definition method for
// fill up anime information once obtained.
type AnimeUtils interface {
	DecodeNode(node *html.Node) error
}

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

// DecodeNode  returbs
func (a *Anime) DecodeNode(node *html.Node) error {
	title, err := getTextFromNode(node, titleXpath)
	if err != nil {
		return err
	}
	a.Title = title

	typeValue, err := getTextFromNode(node, typeXpath)
	if err != nil {
		return err
	}
	a.Type = typeValue

	cover, err := getTextFromNode(node, coverXpath)
	if err != nil {
		return err
	}
	a.Cover = cover

	status, err := getTextFromNode(node, statusXpath)
	if err != nil {
		return err
	}
	a.Status = status

	geners, err := htmlquery.QueryAll(node, genresXpath)
	if err != nil {
		return err
	}

	for _, itemNode := range geners {
		if itemNode == nil {
			return nil
		}
		a.Genres = append(a.Genres, htmlquery.InnerText(itemNode))
	}

	synopsis, err := getTextFromNode(node, synopsisXpath)
	if err != nil {
		return err
	}
	a.Synopsis = synopsis

	scriptContent, err := getTextFromNode(node, scriptContentXpath)
	if err != nil {
		return err
	}

	nextRelease, err := getNextReleaseDateFromString(scriptContent)
	if err != nil {
		return err
	}
	a.NextRelease = nextRelease

	lastEpisode, err := getLastEpisodeFromString(scriptContent)
	if err != nil {
		return err
	}
	a.LastAiredEpisode = lastEpisode
	return nil
}

// GetAnime returns pointer to a new single anime representation.
func GetAnime(sourceURL string) (*Anime, error) {
	doc, err := htmlquery.LoadURL(sourceURL)
	if err != nil {
		return nil, err
	}
	anime := &Anime{}
	err = anime.DecodeNode(doc)
	if err != nil {
		return nil, err
	}
	return anime, nil
}

// GetAllCurrentlyAiringShows returns a list of all the anime shows
// that has been airing in current season.
func GetAllCurrentlyAiringShows() ([]Anime, error) {
	sources, err := GetSources()
	if err != nil {
		return nil, err
	}

	if len(sources) == 0 {
		return nil, errors.New("There are no sources available to load new anime data")
	}

	var animes []Anime
	for _, source := range sources {
		anime, err := GetAnime(source)
		if err != nil {
			return nil, err
		}
		animes = append(animes, *anime)
	}

	return animes, nil
}

// getTextFromNode returns inner text from a specific node.
func getTextFromNode(doc *html.Node, xpath string) (innerText string, err error) {
	node, err := htmlquery.Query(doc, xpath)
	if err != nil {
		return "", err
	}
	if node == nil {
		return "", fmt.Errorf("There is an error trying to get node with specified xpath: %v", xpath)
	}
	nodeValue := htmlquery.InnerText(node)
	return nodeValue, nil
}

// getNextReleaseDateFromString returns time from current scriptContent.
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

// getLastEpisodeFromString returns the last episode on air.
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
