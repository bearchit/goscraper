package goscraper

import (
	"context"

	"github.com/gocolly/colly"
)

type opengraph struct{}

func NewOpenGraphSummarizer() *opengraph {
	return &opengraph{}
}

func (p opengraph) Summarize(ctx context.Context, e *colly.HTMLElement) (*Summary, error) {
	s := new(Summary)

	s.Title = e.ChildAttr(`meta[property="og:title"]`, "content")
	s.Description = e.ChildAttr(`meta[property="og:description"`, "content")
	s.Cover = p.parseCover(ctx, e)

	return s, nil
}

func (p opengraph) parseCover(_ context.Context, e *colly.HTMLElement) *Image {
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

	return &Image{
		URL:    url,
		Width:  imageSizeFromString(e.ChildAttr(`meta[property="og:image:width]"`, "content")),
		Height: imageSizeFromString(e.ChildAttr(`meta[property="og:image:height]"`, "content")),
		Alt:    e.ChildAttr(`meta[property="og:image:alt]"`, "content"),
	}
}

func (p opengraph) parseImages(_ context.Context, e *colly.HTMLElement) ([]*Image, error) {
	images := make([]*Image, 0)
	var prevImage *Image = nil
	e.ForEach("meta", func(i int, element *colly.HTMLElement) {
		if element.Attr("property") == "og:image" {
			if prevImage != nil {
				images = append(images, prevImage)
				prevImage = nil
			} else {
				prevImage = &Image{
					URL: e.Request.AbsoluteURL(element.Attr("content")),
				}
			}
		}

		if element.Attr("property") == "og:image:width" {
			if prevImage != nil {
				prevImage.Width = imageSizeFromString(element.Attr("content"))
			}
		}

		if element.Attr("property") == "og:image:height" {
			if prevImage != nil {
				prevImage.Height = imageSizeFromString(element.Attr("content"))
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
