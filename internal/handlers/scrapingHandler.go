package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/Zucke/CodeTracker/internal/data"
	"github.com/Zucke/CodeTracker/pkg/response"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

//MakeScrapingRequest meje scrap Request
func MakeScrapingRequest(w http.ResponseWriter, r *http.Request) {
	var scrapData data.ScrapData
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	_, m, err := ws.ReadMessage()
	json.Unmarshal([]byte(m), &scrapData)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if scrapData.UseRegex {
			r := regexp.MustCompile(scrapData.ToFind)
			if r.MatchString(e.Text) {
				ws.WriteMessage(1, []byte(link))
			}

		} else {
			if strings.Contains(e.Text, scrapData.ToFind) {
				ws.WriteMessage(1, []byte(link))
			}

		}

	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(link)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	c.Visit(scrapData.CurrendURL)

}

//MainHandler the main handler
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./public/"))
	fs.ServeHTTP(w, r)

}
