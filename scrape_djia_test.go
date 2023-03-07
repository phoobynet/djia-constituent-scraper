package djia_constituent_scraper

import (
	"testing"
)

func TestScrapeDJIA(t *testing.T) {
	constituents, err := ScrapeDJIA()

	if err != nil {
		t.Errorf("ScrapeDJIA() returned an error: %s", err)
	}

	if len(constituents) > 0 {
		t.Logf("Found %d constituents", len(constituents))
		t.Logf("First constituent: %s", constituents[0])
	} else {
		t.Fatal("No constituents found")
	}
}
