package summary

import (
	"context"

	"github.com/gocolly/colly"
)

type Summarizer interface {
	Summarize(ctx context.Context, e *colly.HTMLElement) (*Summary, error)
}
