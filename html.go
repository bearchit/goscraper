package goscraper

import (
	"context"

	"github.com/gocolly/colly"
)

type html struct{}

func NewHTMLSummarizer() *html {
	return &html{}
}

func (p html) Summarize(ctx context.Context, e *colly.HTMLElement) (*Summary, error) {
	s := new(Summary)

	s.Title = e.ChildText("title")
	s.Description = e.ChildAttr(`meta[name="description"]`, "content")

	return s, nil
}
