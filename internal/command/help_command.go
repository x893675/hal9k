package command

import (
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/pkg/client/qbot"
)

func init() {
	qbot.RegistryCommandHandler("help", HelpCommand)
	qbot.RegistryCommandHandler("帮助", HelpCommand)
}

func HelpCommand(update qqbotapi.Update) {
	helpMsg := `
Bot相关指令:
1. /image: 图片相关功能(上传图片等)，详细命令使用/image help查询
2. /img $catalog: 随机发送一张分类为$catalog的图片
3. /luck,/占卜: 今日浅草寺占卜
4. /div,/卜卦: 周易卜卦
5. /pixiv: pixiv图片相关功能，详细命令使用/pixiv help查询`

	reply(update, helpMsg)
}
