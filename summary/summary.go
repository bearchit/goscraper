package summary

import "github.com/bearchit/goscraper/image"

type Summary struct {
	Title       string
	Description string
	Cover       *image.Image
}
