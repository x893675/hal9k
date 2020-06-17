package main

import (
	"github.com/catsworld/qq-bot-api"
	"log"
	"net/http"
	"sync"
)

func main() {
	bot, err := qqbotapi.NewBotAPI("MyCoolqHttpToken", "http://localhost:5700", "CQHTTP_SECRET")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := qqbotapi.NewWebhook("/webhook_endpoint")
	u.PreloadUserInfo = true

	// Use WebHook as event method
	updates := bot.ListenForWebhook(u)
	// Or if you love WebSocket Reverse
	// updates := bot.ListenForWebSocket(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go messageHandler(1,bot, updates, &wg)
	wg.Add(1)
	go messageHandler(2,bot, updates, &wg)

	wg.Wait()
	return
	//for update := range updates {
	//	log.Println("len of channel: ", len(updates))
	//	if update.Message == nil {
	//		continue
	//	}
	//	log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)
	//
	//	bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
	//	log.Println("len of channel: ", len(updates))
	//	log.Println()
	//}
}

func messageHandler(i int, bot *qqbotapi.BotAPI, updates qqbotapi.UpdatesChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for update := range updates {
		log.Printf("in handler %d len of channel: %d", i,len(updates))
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)
		bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
	}
}
