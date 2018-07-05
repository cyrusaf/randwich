package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_American_sandwiches")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find sandwiches
	sandwiches := make([]string, 0)
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			return
		}
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				if i > 0 {
					return
				}
				sandwiches = append(sandwiches, s.Text())
			})
		})
	})

	rand.Seed(time.Now().UnixNano())
	sandwich := sandwiches[rand.Intn(len(sandwiches))]
	fmt.Println(sandwich)
}
