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
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	numberPattern = `^([1-9]\d*)\b`
	numberRe      *regexp.Regexp
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
	Restrict   int       `json:"x-restrict"`
	Tags       []string  `json:"tags"`
}

type Illust struct {
	IllustrationInfo `json:"illust"`
}

func init() {
	var err error
	numberRe, err = regexp.Compile(numberPattern)
	if err != nil {
		panic(err)
	}
	qbot.RegistryCommandHandler("pixiv", PixivCommand)
}

func PixivCommand(update qqbotapi.Update) {
	_, args := update.Message.Command()
	logger.Info(nil, "args is  %v", args)
	if len(args) < 1 {
		reply(update, "参数错误!")
		return
	}
	if numberRe.MatchString(args[0]) {
		searchByID(update, args[0])
	} else {
		searchByWord(update, args[0])
	}
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
	if illust.ImageUrl.Large == "" {
		reply(update, "图片不存在")
		return
	}
	reply(update, revproxy(fmt.Sprintf("[CQ:image,file=%s,cache=0]", illust.ImageUrl.Large)))
}

func searchByWord(update qqbotapi.Update, keyword string) {
	key := url.QueryEscape(keyword)
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}
	resp, err := c.Get(fmt.Sprintf(constants.PixivSearchByKeyword, key))
	if err != nil {
		logger.Error(nil, "search pixiv  by keywork error, %v", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	var illust []*Illust
	err = json.Unmarshal(body, &illust)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	if len(illust) > 0 {
		index := RangeRand(0, int64(len(illust)-1))
		if illust[index].ImageUrl.Large == "" {
			reply(update, "未找到相关图片")
			return
		}
		reply(update, revproxy(fmt.Sprintf("[CQ:image,file=%s,cache=0]", illust[index].ImageUrl.Large)))
	} else {
		reply(update, "未找到相关图片")
		return
	}

}

func revproxy(url string) string {
	u := strings.ReplaceAll(url, constants.PixivOriginalDomain, constants.PixivRevproxyDomain)
	return strings.ReplaceAll(u, "_webp", "")
}
