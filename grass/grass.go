package grass

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	githubBaseURL = "https://github.com/"
)

const (
	svgClassName = ".js-calendar-graph-svg"
	grassNodeTag = "rect"

	dateFormat = "2006-01-02"
)

const (
	maxCap = 365 + 7
)

// Grass represents GitHub daily contributions information.
type Grass struct {
	x, y      int
	dataCount int
	color     string
	date      time.Time
}

func (g *Grass) String() string {
	return fmt.Sprintf("{x:%d, y:%d, n:%d, c:%s, d:%s}",
		g.x, g.y, g.dataCount, g.color, g.date)
}

// Growth returns concentration of grass. (Less 0 1 2 3 4 More)
func (g *Grass) Growth() int {
	switch g.color {
	case "#196127":
		return 4
	case "#239a3b":
		return 3
	case "#7bc96f":
		return 2
	case "#c6e48b":
		return 1
	case "#ebedf0":
		return 0
	}
	return 0
}

// GetMonth returns the month of the grassing date.
func (g *Grass) GetMonth() int {
	return int(g.date.Month())
}

// GetDay returns the day of the grassing date.
func (g *Grass) GetDay() int {
	return g.date.Day()
}

// Mow scrape contributions information from GitHub.
func Mow(id string) ([]*Grass, error) {
	url := createURL(id)
	res, err := get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	// TODO: explicit sort grasses by date asc
	return scrape(doc), nil
}

func createURL(id string) string {
	return githubBaseURL + id
}

func grassesNodesSelector() string {
	return fmt.Sprintf("%s %s", svgClassName, grassNodeTag)
}

func get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to request: status code = %d, url = %s", res.StatusCode, url)
	}
	return res, nil
}

func scrape(doc *goquery.Document) []*Grass {
	grasses := make([]*Grass, 0, maxCap)
	svg := doc.Find(grassesNodesSelector())
	for _, n := range svg.Nodes {
		grasses = append(grasses, rectToGrass(n))
	}
	return grasses
}

func rectToGrass(n *html.Node) *Grass {
	g := &Grass{}
	for _, a := range n.Attr {
		switch a.Key {
		case "x":
			g.x, _ = strconv.Atoi(a.Val)
		case "y":
			g.y, _ = strconv.Atoi(a.Val)
		case "data-count":
			g.dataCount, _ = strconv.Atoi(a.Val)
		case "fill":
			g.color = a.Val
		case "data-date":
			g.date, _ = time.Parse(dateFormat, a.Val)
		}
	}
	return g
}
