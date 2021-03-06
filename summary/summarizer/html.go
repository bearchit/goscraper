package summarizer

import (
	"context"

	"github.com/bearchit/goscraper/summary"

	"github.com/gocolly/colly"
)

type html struct{}

func NewHTML() summary.Summarizer {
	return &html{}
}

func (p html) Summarize(ctx context.Context, e *colly.HTMLElement) (*summary.Summary, error) {
	return &summary.Summary{
		Title:       e.ChildText("title"),
		Description: e.ChildAttr(`meta[name="description"]`, "content"),
	}, nil
}
