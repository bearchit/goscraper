package goscraper

import (
	"context"
	"net/url"

	"github.com/bearchit/goscraper/summary"

	"github.com/imdario/mergo"

	"github.com/bearchit/goscraper/summarizer"

	"github.com/gocolly/colly"
)

type Summarizer interface {
	Summarize(ctx context.Context, e *colly.HTMLElement) (*summary.Summary, error)
}

type Header struct {
	summarizers []Summarizer
}

func (h Header) Run(ctx context.Context, rawurl string) (*summary.Summary, error) {
	_, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	c := colly.NewCollector()
	var summary summary.Summary
	c.OnHTML("html", func(e *colly.HTMLElement) {
		for _, s := range h.summarizers {
			r, err := s.Summarize(ctx, e)
			if err != nil {

			} else {
				mergo.Merge(&summary, r)
			}
		}
	})

	if err := c.Visit(rawurl); err != nil {
		return nil, err
	}

	return &summary, nil
}

func DefaultHeader() *Header {
	return &Header{summarizers: []Summarizer{
		summarizer.NewOpenGraph(),
		summarizer.NewTwitterCard(),
		summarizer.NewHTML(),
	}}
}
