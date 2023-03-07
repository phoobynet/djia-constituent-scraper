# DJIA Constituent Scraper

## Installation
```bash
go get github.com/phoobynet/djia-constituent-scraper
```

## Usage

```go
package main

import (
    "fmt"
    djia "github.com/phoobynet/djia-constituent-scraper"
)

func main() {
    constituents, err := djia.ScrapeDJIA()

    if err != nil {
        panic(err)
    }

    if len(constituents) > 0 {
        fmt.Printf("Found %d", len(constituents))
        fmt.Printf("First constituent: %s", constituents[0])
    } else {
        panic("No constituents found")
    }
}
```


