package command

import (
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	"hal9k/pkg/client/qbot"
	"hal9k/pkg/constants"
	"hal9k/pkg/logger"
	"log"
)

func init() {
	qbot.RegistryCommandHandler("test", TestCommand)
}

func TestCommand(update qqbotapi.Update) {
	_, args := update.Message.Command()
	argsMedia, _ := cqcode.ParseMessage(args)
	for _, v := range argsMedia {
		switch r := v.(type) {
		case *cqcode.At:
			log.Println("The command includes an At!, qq is ", r.QQ)
		case *cqcode.Face:
			fmt.Print("The command includes a Face!")
		case *cqcode.Image:
			log.Println("the command includes a image ", r.URL)
		}
	}
	reply(update, update.Message.Text)
	//qbot.Bot().SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, fmt.Sprintf("[CQ:image,file=b26939971854291f0f13d488350a8f51.jpg]"))
	//qbot.Bot().SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
}

func reply(update qqbotapi.Update, msg string) {
	if update.MessageType == constants.Group {
		msg = fmt.Sprintf("[CQ:at,qq=%d]\n%s", update.UserID, msg)
	}
	_, err := qbot.Bot().SendMessage(update.Message.Chat.ID, update.Message.Chat.Type, msg)
	if err != nil {
		logger.Error(nil, "send message error [%v]", err)
	}
}
