package djia_constituent_scraper

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

import (
	"fmt"
	"github.com/gocolly/colly"
)

type DJIAConstituent struct {
	Ticker    string    `json:"ticker"`
	Exchange  string    `json:"exchange"`
	Company   string    `json:"company"`
	Industry  string    `json:"industry"`
	DateAdded time.Time `json:"dateAdded"`
	Notes     string    `json:"notes"`
	Weighting float64   `json:"weighting"`
}

func (d DJIAConstituent) String() string {
	return fmt.Sprintf("%s %s %s %s %s %.2f%%\n", d.Ticker, d.Exchange, d.Company, d.Industry, d.DateAdded, d.Weighting*100)
}

func (d DJIAConstituent) JSON() (string, error) {
	j, err := json.MarshalIndent(d, "", "  ")

	if err != nil {
		return "", err
	}

	return string(j), nil
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
			el.ForEach("tr", func(rowIndex int, el *colly.HTMLElement) {
				if rowIndex == 0 {
					// parse header, creating a map of header name to column index
					el.ForEach("th", func(headerIndex int, el *colly.HTMLElement) {
						header := strings.TrimSpace(strings.TrimSuffix(el.Text, "\n"))
						headerMap[header] = headerIndex
					})
				} else {
					// data
					ticker := el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Symbol"]+1))

					if ticker != "" {

						constituent := DJIAConstituent{
							Ticker:   ticker,
							Exchange: el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Exchange"]+1)),
							Company:  el.ChildText(fmt.Sprintf("th:nth-child(%d)", headerMap["Company"]+1)),
							Industry: el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Industry"]+1)),
							Notes:    el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Notes"]+1)),
						}

						dateAdded, err := time.Parse(time.DateOnly, el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Date added"]+1)))

						if err == nil {
							constituent.DateAdded = dateAdded
						}

						weightingRaw := el.ChildText(fmt.Sprintf("td:nth-child(%d)", headerMap["Index weighting"]+1))

						if weightingRaw != "" {
							weighting, err := strconv.ParseFloat(strings.TrimSuffix(weightingRaw, "%"), 64)

							if err == nil {
								constituent.Weighting = weighting / 100
							}
						}

						constituents = append(constituents, constituent)
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
