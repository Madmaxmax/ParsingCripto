package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strconv"
)

type Auditor struct {
	Rank           int      `json:"row-cell rank"`
	InfluencerNik  string   `json:"influencer_nik"`
	InfluencerName string   `json:"influencer_name"`
	Image          string   `json:"image"`
	Categories     []string `json:"category"`
	Followers      string   `json:"followers"`
	Country        string   `json:"country"`
	EngAuth        string   `json:"eng_auth"`
	EngAvg         string   `json:"eng_avg"`
}

func arrayToString(arr []string) string {
	return "\"" + fmt.Sprintf("%v", arr) + "\""
}

func saveData(data []Auditor, name string) {
	file, e := os.Create(name + ".csv")
	if e != nil {
		panic(e)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Rank", "InfluencerNik", "InfluencerName", "Image", "Categories", "Followers", "Country", "EngAuth", "EngAvg"}
	if err := writer.Write(headers); err != nil {
		panic(err)
		return
	}

	for _, d := range data {
		row := []string{
			fmt.Sprintf("%d", d.Rank),
			d.InfluencerNik,
			d.InfluencerName,
			d.Image,
			arrayToString(d.Categories),
			d.Followers,
			d.Country,
			d.EngAuth,
			d.EngAvg,
		}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}

func main() {
	var url = "https://hypeauditor.com/top-instagram-all-russia/"
	collector := colly.NewCollector()
	var data []Auditor
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		panic(e)
		fmt.Println("Got this error:", e)
	})
	collector.OnHTML(".row", func(element *colly.HTMLElement) {
		news := &Auditor{}
		element.ForEach(".row__top", func(_ int, h *colly.HTMLElement) {
			rankStr := h.ChildText(".row-cell.rank > span:first-child")
			rank, err := strconv.Atoi(rankStr)
			if err != nil {
				panic(err)
			}
			news.Rank = rank
			news.InfluencerNik = h.ChildText(".contributor__name-content")
			news.InfluencerName = h.ChildText(".contributor__title")
			imgElement := h.DOM.Find(".avatar.contributor__avatar img")
			if imgElement.Length() > 0 {
				imgUrl, exists := imgElement.Attr("src")
				if exists {
					news.Image = imgUrl
				}
			}
			var categories []string
			element.ForEach(".tag__content", func(_ int, h *colly.HTMLElement) {
				category := h.Text
				categories = append(categories, category)

			})
			news.Categories = categories
			news.Followers = h.ChildText(".row-cell.subscribers")
			news.Country = h.ChildText(".row-cell.audience")
			news.EngAuth = h.ChildText(".row-cell.authentic")
			news.EngAvg = h.ChildText(".row-cell.engagement")
			data = append(data, *news)
		})
	})

	err := collector.Visit(url)
	if err != nil {
		return
	}
	saveData(data, "Top_50_influencers")
}
