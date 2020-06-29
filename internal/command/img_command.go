package command

import (
	"crypto/rand"
	"fmt"
	qqbotapi "github.com/catsworld/qq-bot-api"
	"hal9k/pkg/client/qbot"
	"hal9k/pkg/logger"
	"hal9k/pkg/utils/envutils"
	"io/ioutil"
	"math"
	"math/big"
	"os"
)

var (
	graphBedAddr = envutils.GetEnvironment("GRAPH_ADDR", "http://localhost:8080")
)

func init() {
	qbot.RegistryCommandHandler("img", ShowImage)
}

func ShowImage(update qqbotapi.Update) {
	_, args := update.Message.Command()
	if len(args) != 1 {
		reply(update, "参数错误!\n/img $catalog")
		return
	}
	directory := fmt.Sprintf("%s/%s", imagePath, args[0])
	exists, err := PathExists(directory)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	if !exists {
		reply(update, fmt.Sprintf("不存在标签为%s的图片!", args[0]))
		return
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		logger.Error(nil, err.Error())
		return
	}
	imgs := make([]os.FileInfo, 0)
	for _, item := range files {
		if !item.IsDir() && isHidden(item.Name()) {
			imgs = append(imgs, item)
		}
	}
	imgNum := len(imgs)
	if imgNum == 0 {
		reply(update, fmt.Sprintf("不存在标签为%s的图片!", args[0]))
		return
	}
	index := RangeRand(0, int64(imgNum)-1)
	imgname := imgs[index].Name()
	msg := fmt.Sprintf("[CQ:image,file=%s/%s/%s]", graphBedAddr, args[0], imgname)
	logger.Info(nil, "msg is %v", msg)
	reply(update, msg)
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

func isHidden(filename string) bool {
	// linux/unix only
	if filename[0:1] == "." {
		return true
	}
	return false
}
