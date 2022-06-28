package bilibili

import (
	"encoding/json"
	"fmt"

	"github.com/FloatTech/zbputils/binary"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	// ArticleInfoURL 查看专栏信息
	ArticleInfoURL = "https://api.bilibili.com/x/article/viewinfo?id=%v"
	// CURL b站专栏前缀
	CURL = "https://www.bilibili.com/read/cv"
)

// ArticleCard2msg 专栏转消息
func ArticleCard2msg(str string, defaultID string) (msg []message.MessageSegment, err error) {
	var (
		card Card
	)
	msg = make([]message.MessageSegment, 0, 16)
	err = json.Unmarshal(binary.StringToBytes(str), &card)
	if err != nil {
		return
	}
	for i := 0; i < len(card.OriginImageUrls); i++ {
		msg = append(msg, message.Image(card.OriginImageUrls[i]))
	}
	msg = append(msg, message.Text(card.Title, "\n"))
	msg = append(msg, message.Text("UP主: ", card.AuthorName, "\n"))
	msg = append(msg, message.Text(humanNum(card.Stats.View), "阅读 ", humanNum(card.Stats.Reply), "评论", "\n"))
	msg = append(msg, message.Text(CURL, defaultID))
	return
}

// ArticleInfo 用id查专栏信息
func ArticleInfo(id string) (msg []message.MessageSegment, err error) {
	var data []byte
	data, err = web.GetData(fmt.Sprintf(ArticleInfoURL, id))
	if err != nil {
		return
	}
	fmt.Println(string(data))
	return ArticleCard2msg(gjson.ParseBytes(data).Get("data").Raw, id)
}
