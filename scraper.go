package goscraper

import (
	"context"

	"github.com/imdario/mergo"

	"github.com/gocolly/colly"
)

type Scrapper struct {
	summarizers []Summarizer
}

func Default() *Scrapper {
	return &Scrapper{
		summarizers: []Summarizer{
			NewOpenGraphSummarizer(),
			NewTwitterCardSummarizer(),
			NewHTMLSummarizer(),
		},
	}
}

func (s Scrapper) Parse(ctx context.Context, url string) (*Summary, error) {
	c := colly.NewCollector()

	scrap := new(Summary)
	c.OnHTML("html", func(e *colly.HTMLElement) {
		for _, p := range s.summarizers {
			r, _ := p.Summarize(ctx, e)
			mergo.Merge(scrap, r)
		}
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return scrap, nil
}

type WellKnownServices int

const (
	YouTube WellKnownServices = iota + 1
	Instagram
	Facebook
	Airbnb
	Twitter
)

type Summary struct {
	Title       string
	Description string
	Cover       *Image
}

type Images struct {
	Cover  *Image
	Images []*Image
}

type Detail struct {
	Coordinate *Coordinate
}

type Coordinate struct {
	Lon int
	Lat int
}

type Summarizer interface {
	Summarize(ctx context.Context, e *colly.HTMLElement) (*Summary, error)
}

type ImageExtractor interface {
	Execute(ctx context.Context, e *colly.HTMLElement) (*Images, error)
}

type DetailExtractor interface {
	Execute(ctx context.Context, e *colly.HTMLElement) (*Detail, error)
}
