package main

import (
    "fmt"
    "log"
    "net/url"
    "time"

    "github.com/gocolly/colly"
)

func main() {
    keyword := "amFoss club"
    searchURL := "https://www.google.com/search?q=" + url.QueryEscape(keyword)

    c := colly.NewCollector(
        colly.AllowedDomains("www.google.com", "google.com"),
        colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"),
    )
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*google.*",
        RandomDelay: 5 * time.Second,
    })
    var results []map[string]string
    c.OnHTML("div.yuRUbf a", func(e *colly.HTMLElement) {
        title := e.Text
        link := e.Attr("href")

        if title != "" && link != "" {
            results = append(results, map[string]string{
                "title": title,
                "link":  link,
            })
        }
        
        if len(results) >= 10 {
            e.Request.Abort()
        }
    })
    fmt.Println("Starting search for:", keyword)
    err := c.Visit(searchURL)
    if err != nil {
        log.Fatal(err)
    }
    for i, result := range results {
        fmt.Printf("%d: %s\n%s\n\n", i+1, result["title"], result["link"])
    }
}
