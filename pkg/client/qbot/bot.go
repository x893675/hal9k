package qbot

import (
	qqbotapi "github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	"hal9k/pkg/config"
)

var bot *qqbotapi.BotAPI

func NewQbot(cfg *config.QBotConfig, stopCh <-chan struct{}) error {
	var err error
	//bot, err = qqbotapi.NewBotAPI(cfg.Token, cfg.QbotHttpEndpoint, cfg.Secret)
	//if err != nil {
	//	return err
	//}
	//bot.Debug = true
	//
	//cqcode.StrictCommand = true
	//// Set command prefix
	//cqcode.CommandPrefix = "/"
	//
	//u := qqbotapi.NewWebhook(cfg.WebhookPattern)
	//u.PreloadUserInfo = true
	//
	//updates := bot.ListenForWebhook(u)

	//Test
	bot, err = qqbotapi.NewBotAPIWithWSClient(cfg.Token, cfg.QbotHttpEndpoint)
	if err != nil {
		return err
	}
	bot.Debug = true
	cqcode.StrictCommand = true
	cqcode.CommandPrefix = "/"
	u := qqbotapi.NewUpdate(0)
	u.PreloadUserInfo = true
	updates, err := bot.GetUpdatesChan(u)

	ev := qqbotapi.NewEv(updates)
	// Function Echo will get triggered on receiving an update with
	// PostType `message`, MessageType `group` and SubType `normal`
	ev.On("message.group.normal")(HandlerGroupNormalMessage)
	//ev.On("message.group.anonymous")(HandlerGroupAnonymousMessage)
	//ev.On("message.group.notice")(HandlerGroupNoticeMessage)
	ev.On("message.private.friend")(HandlerFriendMessage)
	// Function Log will get triggered on receiving an update with
	// PostType `message`
	ev.On("message")(Log)

	return nil
}

func Bot() *qqbotapi.BotAPI {
	return bot
}
