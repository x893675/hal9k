package command

import (
	"encoding/json"
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/pkg/client/qbot"
	"hal9k/pkg/logger"
	"hal9k/pkg/utils/hashutils"
	"os"
	"strconv"
	"strings"
	"time"
)

type LuckData struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
	ImgUrl string `json:"img_url"`
}

var (
	_luckMap = map[int]*LuckData{}
	luckFile string
)

func init() {
	v := os.Getenv("LUCK_FILE")
	if v == "" {
		luckFile = "/Users/hanamichi/work/github/hal9k/config/luck_data.json"
	} else {
		luckFile = v
	}
	data, err := parseLuckData()
	if err != nil {
		panic(err)
	}
	for _, item := range data {
		_luckMap[item.Number] = item
	}
	qbot.RegistryCommandHandler("luck", LuckCommand)
}

func parseLuckData() ([]*LuckData, error) {
	file, err := os.Open(luckFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data []*LuckData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func LuckCommand(update qqbotapi.Update) {
	index := getUserSign(update.UserID)
	if sign, ok := _luckMap[index]; ok {
		logger.Info(nil, "user %d has sign %v", update.UserID, sign)
		reply(update, fmt.Sprintf("[CQ:image,file=%s]", sign.ImgUrl))
	} else {
		reply(update, fmt.Sprintf("[CQ:image,file=%s]", fmt.Sprintf("[CQ:at,qq=%d]\n好像出了点问题，明天再试试吧~", update.UserID)))
	}
}

func getUserSign(userId int64) int {
	ts := time.Now()
	today := fmt.Sprintf("%d%d%d", ts.Year(), int(ts.Month()), ts.Day())
	formatToday, _ := strconv.ParseInt(today, 10, 64)
	formatUserId, _ := strconv.ParseInt(strconv.FormatInt(userId, 10)[:6], 10, 64)
	strnum := strconv.FormatInt(formatToday*formatUserId, 10)
	str := hashutils.MD5Hash([]byte(strnum))
	num, _ := strconv.ParseUint(strings.ToUpper(str)[:8], 16, 64)
	return int(num%100 + 1)
}
