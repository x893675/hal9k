package command

import (
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"github.com/catsworld/qq-bot-api/cqcode"
	"hal9k/internal/constants"
	"hal9k/pkg/client/qbot"
	"hal9k/pkg/logger"
	"hal9k/pkg/utils/envutils"
	"io"
	"net/http"
	"os"
)

var (
	imagePath = envutils.GetEnvironment("IMAGE_PATH", "/Users/hanamichi/work/github/hal9k/data")
)

func init() {
	qbot.RegistryCommandHandler("image", UploadImageCommand)
}

func UploadImageCommand(update qqbotapi.Update) {
	_, args := update.Message.Command()
	switch args[0] {
	case constants.UploadCommand:
		uploadImage(update, args[2:])
		return
	case constants.HelpCommand:
	default:
		reply(update, "参数错误!\n/不支持的自命令，使用/image help查看用法")
		return
	}
}

func uploadImage(update qqbotapi.Update, args []string) {
	if len(args) != 2 {
		reply(update, "参数错误!\n/image upload $catalog image")
		return
	}
	msg, err := cqcode.ParseMessage(args[1])
	if err != nil {
		logger.Error(nil, "parse message error")
		return
	}
	var imgUrl string
	var fileId string
	for _, item := range msg {
		if sc, ok := item.(*cqcode.Image); ok {
			imgUrl = sc.URL
			fileId = sc.FileID
		}
	}
	logger.Info(nil, "img catalog is %s", args[0])
	logger.Info(nil, "img url is %s", imgUrl)
	go downloadFile(args[0], imgUrl, fileId)
	reply(update, fmt.Sprintf("%s,%s", args[0], imgUrl))
}

func downloadFile(catalog string, url string, fileId string) {
	directory := fmt.Sprintf("%s/%s", imagePath, catalog)
	exists, err := PathExists(directory)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	if !exists {
		if err := os.Mkdir(directory, 755); err != nil {
			logger.Error(nil, err.Error())
			return
		}
	}
	res, err := http.Get(url)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	f, err := os.Create(fmt.Sprintf("%s/%s", directory, fileId))
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(f, res.Body)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
