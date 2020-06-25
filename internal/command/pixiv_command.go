package command

import (
	"encoding/json"
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/internal/constants"
	"hal9k/pkg/client/qbot"
	"hal9k/pkg/logger"
	"io/ioutil"
	"net/http"
	"time"
)

type ImageUrls struct {
	Large        string `json:"large"`
	Medium       string `json:"medium"`
	SquareMedium string `json:"square_medium"`
}

type IllustrationInfo struct {
	CreateDate string    `json:"create_date"`
	Height     int       `json:"height"`
	Id         int       `json:"id"`
	ImageUrl   ImageUrls `json:"image_urls"`
}

type Illust struct {
	IllustrationInfo `json:"illust"`
}

func init() {
	qbot.RegistryCommandHandler("pixiv", PixivCommand)
}

func PixivCommand(update qqbotapi.Update) {
	_, args := update.Message.Command()
	logger.Info(nil, "args is  %v", args)
	searchByID(update, args[0])
}

func searchByID(update qqbotapi.Update, id string) {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}
	resp, err := c.Get(fmt.Sprintf("%s%s", constants.PixivSearchByIDURL, id))
	if err != nil {
		logger.Error(nil, "search pixiv  by id error, %v", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	var illust = Illust{}
	err = json.Unmarshal(body, &illust)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	reply(update, fmt.Sprintf("[CQ:image,file=%s,cache=0]", illust.ImageUrl.Large))
}
