package command

import (
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/pkg/client/qbot"
)

func init() {
	qbot.RegistryCommandHandler("div", DivCommand)
	qbot.RegistryCommandHandler("卜卦", DivCommand)
}

func DivCommand(update qqbotapi.Update) {
	reply(update, "To be continued...")
}
