package goscraper

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/require"

	"github.com/gocolly/colly"
)

func TestColly(t *testing.T) {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildAttr(`meta[property="og:title"]`, "content"))
		fmt.Println(e.ChildAttr(`meta[property="og:image"]`, "content"))

		fmt.Println(e.ChildAttr(`meta[name="twitter:title"]`, "content"))
		fmt.Println(e.ChildAttr(`meta[name="twitter:image"]`, "content"))

		p := new(opengraph)
		s, err := p.Summarize(context.Background(), e)
		require.NoError(t, err)
		t.Log(spew.Sdump(s))
	})

	c.Visit("https://www.clien.net/service/board/park/14598853?type=recommend")
	c.Wait()
}

func TestScrapper_Parse(t *testing.T) {
	urls := []string{
		"https://aws.amazon.com/blogs/opensource/rust-runtime-for-aws-lambda/",
		"https://m.blog.naver.com/we_are_youth/221543976793",
		"https://www.netflix.com/browse/audio",
		"https://extrememanual.net/31549#%EB%84%B7%ED%94%8C%EB%A6%AD%EC%8A%A4_%ED%95%9C%EA%B5%AD%EC%96%B4_%EB%8D%94%EB%B9%99_%EC%9E%91%ED%92%88_%EB%AA%A9%EB%A1%9D",
		"https://sentry.io/pricing/",
	}
	s := Default()

	var wait sync.WaitGroup
	wait.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wait.Done()
			r, err := s.Parse(context.Background(), url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(url, spew.Sdump(r))
		}(url)
	}
	wait.Wait()
}
