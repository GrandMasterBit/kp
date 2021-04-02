package main

import (
	"flag"
	"fmt"
	film "kinopoiskraiting/filmlist"
	"log"
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

var (
	user_id int  = 14036738
	debug   bool = true
)

func init() {
	flag.IntVar(&user_id, "id", user_id, "Имя пользователя на кинопоиске.")
	flag.BoolVar(&debug, "d", debug, "Включить отладку.")
}

func main() {
	flag.Parse()
	site := colly.NewCollector(
		colly.Async(false),
		colly.MaxDepth(1),
	)
	site.Limit(&colly.LimitRule{
		Delay: 5 * time.Second,
	})
	site.OnHTML(".profileFilmsList > div:not(.top)", func(e *colly.HTMLElement) {
		name := e.ChildText(".nameRus")
		name_eng := e.ChildText(".nameEng")
		date := e.ChildText(".date")
		vote := e.ChildText(".vote")
		if vote == "" {
			vote = "watched"
		}
		f := film.New(vote, name, name_eng, date)
		fmt.Printf("%s\n\n", f.String())
	})

	// Простиет за это :(
	arrSelector := "#list > tbody:nth-child(1) > tr:nth-child(3) > td:nth-child(1) > div:nth-child(1) > div:nth-child(2) > ul:nth-child(2) > li:nth-child(5) > a:nth-child(1)"
	site.OnHTML(arrSelector, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		site.Visit(e.Request.AbsoluteURL(link))
		log.Println(link)
	})

	link_site := generateKinoLink(user_id)
	if debug {
		log.Println(link_site)
	}
	site.Visit(link_site)
	site.Wait()
}

func generateKinoLink(ID int) string {
	return fmt.Sprintf("https://www.kinopoisk.ru/user/%d/votes/list/ord/date/page/1/", ID)
}

func returnSomfingInt() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}
