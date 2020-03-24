package goscraper

import (
	"context"

	"github.com/gocolly/colly"
)

type twitter struct{}

func NewTwitterCardSummarizer() *twitter {
	return &twitter{}
}

func (p twitter) Summarize(ctx context.Context, e *colly.HTMLElement) (*Summary, error) {
	s := new(Summary)

	s.Title = e.ChildAttr(`meta[name="twitter:title"]`, "content")
	s.Description = e.ChildAttr(`meta[name="twitter:description"]`, "content")
	s.Cover = p.parseCover(ctx, e)

	return s, nil
}

func (p twitter) parseCover(_ context.Context, e *colly.HTMLElement) *Image {
	url := e.ChildAttr(`meta[name="twitter:image"]`, "content")
	if url != "" {
		url = e.Request.AbsoluteURL(url)
	}
	return &Image{
		URL: url,
		Alt: e.ChildAttr(`meta[name="twitter:image:alt"]`, "content"),
	}
}
