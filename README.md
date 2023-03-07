# DJIA Constituent Scraper

## Installation
```bash
go get github.com/phoobynet/djia-constituent-scraper
```

## Usage

```go
package main

import djia "github.com/phoobynet/djia-constituent-scraper"

func main() {
	constituents, err := djia.ScrapeDJIA()

	if err != nil {
		panic(err)
	}
	
	println(constituents)
}
```


