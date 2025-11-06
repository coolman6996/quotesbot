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
		{"The only way to do great work is to love what you do.", "Steve Jobs"},
		{"Be yourself; everyone else is already taken.", "Oscar Wilde"},
		{"I have not failed. I've just found 10,000 ways that won't work.", "Thomas Edison"},
		{"In the middle of difficulty lies opportunity.", "Albert Einstein"},
		{"The greatest glory in living lies not in never falling, but in rising every time we fall.", "Nelson Mandela"},
		{"The future belongs to those who believe in the beauty of their dreams.", "Eleanor Roosevelt"},
		{"It is during our darkest moments that we must focus to see the light.", "Aristotle"},
		{"Life is what happens to you while you're busy making other plans.", "John Lennon"},
		{"The only thing we have to fear is fear itself.", "Franklin D. Roosevelt"},
		{"Success is not final, failure is not fatal: it is the courage to continue that counts.", "Winston Churchill"},
		{"To be, or not to be, that is the question.", "William Shakespeare"},
		{"The journey of a thousand miles begins with one step.", "Lao Tzu"},
		{"If you tell the truth, you don't have to remember anything.", "Mark Twain"},
		{"Happiness depends upon ourselves.", "Aristotle"},
		{"There are decades where nothing happens; and there are weeks where decades happen.", "Vladimir Lenin"},
		{"I'm not the guy you want to mess with. I'm the guy you want to pay to mess with the guy you don't like.", "Jason Statham"},
		{"Training is a daily thing. You don't take days off. If you're dedicated, you don't need a day off.", "Jason Statham"},
		{"I'm a great believer in the philosophy of hard work. There are no shortcuts.", "Jason Statham"},
		{"You either get busy living, or you get busy dying. That's how I see life.", "Jason Statham"},
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
		return c.Send("Great people")
	})
	fmt.Println("Bot is running")
	b.Start()
}
