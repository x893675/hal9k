package qbot

import (
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/pkg/logger"
)

var _commandHandler = map[string]func(update qqbotapi.Update){}

func RegistryCommandHandler(command string, fn func(update qqbotapi.Update)) {
	if _, ok := _commandHandler[command]; ok {
		panic(fmt.Sprintf("command: %s already exist", command))
	}
	_commandHandler[command] = fn
}

func Log(update qqbotapi.Update) {
	logger.Debug(nil, "[%s] %s", update.Message.From.String(), update.Message.Text)
}

// 群组普通消息
func HandlerGroupNormalMessage(update qqbotapi.Update) {
	HandlerFriendMessage(update)
}

// 群组匿名消息
func HandlerGroupAnonymousMessage(update qqbotapi.Update) {

}

// 群组系统通知消息
func HandlerGroupNoticeMessage(update qqbotapi.Update) {

}

// 好友信息
func HandlerFriendMessage(update qqbotapi.Update) {
	if update.Message.IsCommand() {
		cmd, _ := update.Message.Command()
		if _, ok := _commandHandler[cmd]; ok {
			logger.Info(nil, "command is %s", cmd)
			go _commandHandler[cmd](update)
		} else {
			_, err := bot.SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, "不支持的命令")
			if err != nil {
				logger.Error(nil, "send message error, the update is [%v]", update)
			}
		}
	}
}
