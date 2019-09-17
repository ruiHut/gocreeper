package engine

import (
	"crepper/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+
				"fetching url: %s: %v"+r.Url, err)
			continue
		}

		parseRes := r.ParserFunc(body)
		for _, res := range parseRes.Requests {
			if res.Url != "" {
				requests = append(requests, res)
			}
		}

		for _, item := range parseRes.Items {
			fmt.Printf("Got item %v", item)
		}

	}
}
