package image

import "strconv"

type Image struct {
	URL    string
	Width  int
	Height int
	Alt    string
}

func ImageSizeFromString(s string) int {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0
	}
	return int(v)
}
