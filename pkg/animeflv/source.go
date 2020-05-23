package animeflv

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

// GetSources returns all urls that could be use to extract information
// from anime that's been currently airing.
func GetSources() ([]string, error) {
	doc, err := htmlquery.LoadURL(hostname)
	if err != nil {
		return nil, err
	}

	endpointNodes, err := htmlquery.QueryAll(doc, endpointXpath)
	if err != nil {
		return nil, err
	}

	var sources []string
	for _, node := range endpointNodes {
		completeURL := fmt.Sprintf("%s%s", hostname, htmlquery.SelectAttr(node, "href"))
		sources = append(sources, completeURL)
	}

	return sources, nil
}
