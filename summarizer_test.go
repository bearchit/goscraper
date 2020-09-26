package goscraper

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestHeader_Run(t *testing.T) {
	urls := []string{
		"https://aws.amazon.com/blogs/opensource/rust-runtime-for-aws-lambda/",
		"https://m.blog.naver.com/we_are_youth/221543976793",
		"https://www.netflix.com/browse/audio",
		"https://extrememanual.net/31549#%EB%84%B7%ED%94%8C%EB%A6%AD%EC%8A%A4_%ED%95%9C%EA%B5%AD%EC%96%B4_%EB%8D%94%EB%B9%99_%EC%9E%91%ED%92%88_%EB%AA%A9%EB%A1%9D",
		"https://sentry.io/pricing/",
	}
	s := DefaultHeader()

	var wait sync.WaitGroup
	wait.Add(len(urls))
	ctx := context.Background()
	for _, url := range urls {
		go func(url string) {
			defer wait.Done()
			r, err := s.Run(ctx, url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(url, spew.Sdump(r))
		}(url)
	}
	wait.Wait()
}
