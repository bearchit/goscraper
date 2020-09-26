package summarizer

import (
	"context"

	"github.com/bearchit/goscraper/summary"

	"github.com/bearchit/goscraper/image"

	"github.com/bearchit/goscraper"

	"github.com/gocolly/colly"
)

type openGraph struct{}

func NewOpenGraph() goscraper.Summarizer {
	return &openGraph{}
}

func (p openGraph) Summarize(ctx context.Context, e *colly.HTMLElement) (*summary.Summary, error) {
	return &summary.Summary{
		Title:       e.ChildAttr(`meta[property="og:title"]`, "content"),
		Description: e.ChildAttr(`meta[property="og:description"`, "content"),
		Cover:       p.parseCover(ctx, e),
	}, nil
}

func (p openGraph) parseCover(_ context.Context, e *colly.HTMLElement) *image.Image {
	url := e.ChildAttr(`meta[property="og:image"]`, "content")
	if url == "" {
		url = e.ChildAttr(`meta[property="og:image:url"]`, "content")
		if url == "" {
			return nil
		}
	}

	if url != "" {
		url = e.Request.AbsoluteURL(url)
	}

	return &image.Image{
		URL:    url,
		Width:  image.ImageSizeFromString(e.ChildAttr(`meta[property="og:image:width]"`, "content")),
		Height: image.ImageSizeFromString(e.ChildAttr(`meta[property="og:image:height]"`, "content")),
		Alt:    e.ChildAttr(`meta[property="og:image:alt]"`, "content"),
	}
}

func (p openGraph) parseImages(_ context.Context, e *colly.HTMLElement) ([]*image.Image, error) {
	images := make([]*image.Image, 0)
	var prevImage *image.Image = nil
	e.ForEach("meta", func(i int, element *colly.HTMLElement) {
		if element.Attr("property") == "og:image" {
			if prevImage != nil {
				images = append(images, prevImage)
				prevImage = nil
			} else {
				prevImage = &image.Image{
					URL: e.Request.AbsoluteURL(element.Attr("content")),
				}
			}
		}

		if element.Attr("property") == "og:image:width" {
			if prevImage != nil {
				prevImage.Width = image.ImageSizeFromString(element.Attr("content"))
			}
		}

		if element.Attr("property") == "og:image:height" {
			if prevImage != nil {
				prevImage.Height = image.ImageSizeFromString(element.Attr("content"))
			}
		}

		if element.Attr("property") == "og:image:alt" {
			prevImage.Alt = element.Attr("content")
		}
	})
	if prevImage != nil {
		images = append(images, prevImage)
	}
	return images, nil
}
