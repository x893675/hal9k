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

func main() {
	searchByID("77558582")
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
