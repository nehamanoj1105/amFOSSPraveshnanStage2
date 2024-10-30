# Summary
For this project, I developed a web scraper using Go and the Colly library to extract the top 10 Google search results for a given keyword.
Initially I setup the Colly library to send requests to Google's search result page.
Colly was used for parsing the html and the relevant elements like the title and url were extracted.
In order to handle Google's anti-scraping policies I implemented a delay between the requests.
The scarper was optimised by ensuring that only the top 10 search results were taken into consideration.


The Go code is in the main.go file.
Screenshots of the parsed results have also been uploaded.
