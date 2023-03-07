package djia_constituent_scraper

import "testing"

func TestScrapeDJIA(t *testing.T) {
	actual, err := ScrapeDJIA()

	if err != nil {
		t.Errorf("ScrapeDJIA() returned an error: %s", err)
	}

	t.Log(actual)
}
