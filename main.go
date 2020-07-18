package main

import (
	"encoding/json"
	"fmt"
	"hal9k/internal/constants"
	"hal9k/pkg/logger"
	"hal9k/pkg/utils/hashutils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
	Restrict   int       `json:"restrict"`
}

type Illust struct {
	IllustrationInfo `json:"illust"`
}

type Illusts struct {
	I []*IllustrationInfo `json:"illusts"`
}

func main() {
	//pattern := `^([1-9]\d*)\b` //反斜杠要转义
	//str := "12453435353535353"
	//result,_ := regexp.MatchString(pattern,str)
	//fmt.Println(result)
	//searchByWord("FGO")
	//fmt.Println(getUserSign(540386505))
	//searchByID("77558582")
	//	helpMsg := `
	//Bot相关指令:
	//1. /image: 图片相关功能(上传图片等)，详细命令使用/image help查询
	//2. /img $catalog: 随机发送一张分类为$catalog的图片
	//3. /luck,/占卜: 今日浅草寺占卜
	//4. /div,/卜卦: 周易卜卦
	//5. /pixiv: pixiv图片相关功能，详细命令使用/pixiv help查询`
	//	fmt.Println(helpMsg)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rangeRand(1, 10))
	}
}

func rangeRand(min, max int) int {
	result := rand.Intn(max - min + 1)
	return min + result
}

func searchByID(id string) {
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
	logger.Info(nil, string(body))
	var illust = Illust{}
	err = json.Unmarshal(body, &illust)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	logger.Info(nil, "+%v", illust)
}

//func searchByWord(keyword string) {
//	key := url.QueryEscape(keyword)
//	c := http.Client{
//		Transport:     nil,
//		CheckRedirect: nil,
//		Jar:           nil,
//		Timeout:       5 * time.Second,
//	}
//	resp, err := c.Get(fmt.Sprintf(constants.PixivSearchByKeyword, key))
//	if err != nil {
//		logger.Error(nil, "search pixiv  by keywork error, %v", err)
//		return
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logger.Error(nil, err.Error())
//		return
//	}
//	var illust Illusts
//	err = json.Unmarshal(body, &illust)
//	if err != nil {
//		logger.Error(nil, err.Error())
//		return
//	}
//	if len(illust.Illusts) > 0 {
//		index := RangeRand(0, int64(len(illust.Illusts)-1))
//		if illust.Illusts[index].ImageUrl.Large == "" {
//			return
//		}
//		fmt.Println(fmt.Sprintf("[CQ:image,file=%s,cache=0]", illust.Illusts[index].ImageUrl.Large))
//	} else {
//		return
//	}
//
//}

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
