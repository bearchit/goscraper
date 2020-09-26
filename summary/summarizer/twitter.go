package summarizer

import (
	"context"

	"github.com/gocolly/colly"

	"github.com/bearchit/goscraper/image"
	"github.com/bearchit/goscraper/summary"
)

type twitterCard struct{}

func NewTwitterCard() summary.Summarizer {
	return &twitterCard{}
}

func (p twitterCard) Summarize(ctx context.Context, e *colly.HTMLElement) (*summary.Summary, error) {
	return &summary.Summary{
		Title:       e.ChildAttr(`meta[name="twitter:title"]`, "content"),
		Description: e.ChildAttr(`meta[name="twitter:description"]`, "content"),
		Cover:       p.parseCover(ctx, e),
	}, nil
}

func (p twitterCard) parseCover(_ context.Context, e *colly.HTMLElement) *image.Image {
	url := e.ChildAttr(`meta[name="twitter:image"]`, "content")
	if url != "" {
		url = e.Request.AbsoluteURL(url)
	}
	return &image.Image{
		URL: url,
		Alt: e.ChildAttr(`meta[name="twitter:image:alt"]`, "content"),
	}
}
