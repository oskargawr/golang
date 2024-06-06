package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	var rows []TableRow

	c.OnHTML(".infobox tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			header := el.ChildText("th")
			value := el.ChildText("td")
			if header != "" && value != "" {
				row := TableRow{
					Header: header,
					Value:  value,
				}
				rows = append(rows, row)
			}
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")

	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Header", "Value"})

	for _, row := range rows {
		writer.Write([]string{row.Header, row.Value})
	}

	fmt.Println("done")
}
