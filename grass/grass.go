package grass

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	githubBaseURL = "https://github.com/"
)

const (
	svgClassName = ".js-calendar-graph-svg"
	grassNodeTag = "rect"
)

const (
	maxCap = 365 + 7
)

type Grass struct {
	x, y      int
	dataCount int
	color     string
	date      string
}

func (g *Grass) String() string {
	return fmt.Sprintf("{x:%d, y:%d, n:%d, c:%s, d:%s}",
		g.x, g.y, g.dataCount, g.color, g.date)
}

func Mow(id string) ([]*Grass, error) {
	url := createURL(id)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return scrape(doc), nil
}

func createURL(id string) string {
	return githubBaseURL + id
}

func grassesNodesSelector() string {
	return fmt.Sprintf("%s %s", svgClassName, grassNodeTag)
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
			g.date = a.Val
		}
	}
	return g
}
