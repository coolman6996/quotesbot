package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Quote struct {
	Text   string
	Author string
}

func main() {
	pref := tele.Settings{
		Token:  "7999490418:AAHwVaxL4lF8M_Yx8PkXwy2qSPSpLCYa_ac",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	var quotes = []Quote{
		{"I don't grandfather, You don't grandmother, what?...", "Jason Statham"},
		{"This is your", "Pavlik"},
		{"Remember, there is nothing more precious than a pig.", "Zelya"},
		{"Last Cristmas, I gave you my heart...", "Mark"},
	}
	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Mom is mother, Dad is father \"Unkwown\"")
	})
	b.Handle("/numquotes", func(c tele.Context) error {
		count := len(quotes)
		return c.Send(fmt.Sprintf("Number of quotes: %d", count))
	})
	b.Handle("/quote", func(c tele.Context) error {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := len(quotes)
		randomNum := rand.Intn(max-min+1) + min
		i := randomNum - 1
		text := quotes[i].Text
		author := quotes[i].Author
		c.Send(text)
		c.Send(author)
		return c.Send("")
	})
	fmt.Println("Bot is running")
	b.Start()
}
