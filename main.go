package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"hal9k/internal/constants"
	"hal9k/pkg/logger"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"net/url"
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
	searchByWord("FGO")
	//searchByID("77558582")
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func searchByID(id string) {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5*time.Second,
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


func searchByWord(keyword string) {
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
	var illust Illusts
	err = json.Unmarshal(body, &illust)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	if len(illust.Illusts) > 0 {
		index := RangeRand(0, int64(len(illust.Illusts)-1))
		if illust.Illusts[index].ImageUrl.Large == "" {
			return
		}
		fmt.Println(fmt.Sprintf("[CQ:image,file=%s,cache=0]", illust.Illusts[index].ImageUrl.Large))
	} else {
		return
	}

}