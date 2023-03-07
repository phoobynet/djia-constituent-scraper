package djia_constituent_scraper

import "strings"

import (
	"fmt"
	"github.com/gocolly/colly"
)

type DJIAConstituent struct {
	Ticker          string `json:"ticker"`
	Company         string `json:"company"`
	GICSSector      string `json:"gicsSector"`
	GICSSubIndustry string `json:"gicsSubIndustry"`
}

// ScrapeDJIA scrapes the DJIA from https://en.wikipedia.org/wiki/Dow_Jones_Industrial_Average (probably not the best source)
func ScrapeDJIA() ([]DJIAConstituent, error) {
	c := colly.NewCollector()

	headerMap := make(map[string]int)
	constituents := make([]DJIAConstituent, 0)

	// Find and visit all links
	c.OnHTML("table#constituents", func(e *colly.HTMLElement) {
		// parse the columns headers to figure out indexes of the columns we want
		e.ForEach("tbody", func(i int, el *colly.HTMLElement) {
			// header
			el.ForEach("tr", func(rowIndex int, el *colly.HTMLElement) {
				if rowIndex == 0 {
					el.ForEach("th", func(headerIndex int, el *colly.HTMLElement) {
						header := strings.TrimSpace(strings.TrimSuffix(el.Text, "\n"))
						headerMap[header] = headerIndex
					})
				} else {
					// data
					ticker := el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Symbol"]+1))

					if ticker != "" {
						indexConstituent := DJIAConstituent{
							Ticker:     ticker,
							Company:    el.ChildText(fmt.Sprintf("th:nth-child(%d)", headerMap["Company"]+1)),
							GICSSector: el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Industry"]+1)),
						}
						constituents = append(constituents, indexConstituent)
					}
				}
			})
		})
	})

	err := c.Visit("https://en.wikipedia.org/wiki/Dow_Jones_Industrial_Average")
	if err != nil {
		return nil, err
	}

	return constituents, nil
}
