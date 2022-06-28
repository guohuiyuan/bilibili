package bilibili

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/FloatTech/zbputils/binary"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	// VideoInfoURL 查看视频信息
	VideoInfoURL = "https://api.bilibili.com/x/web-interface/view?aid=%v&bvid=%v"
	// SearchVideoInfoURL 搜索视频信息
	SearchVideoInfoURL = "https://api.bilibili.com/x/web-interface/search/all/v2?%v"
	// VURL 视频网址前缀
	VURL = "https://www.bilibili.com/video/"
)

// VideoCard2msg 视频卡片转消息
func VideoCard2msg(str string) (msg []message.MessageSegment, err error) {
	var (
		card Card
	)
	msg = make([]message.MessageSegment, 0, 16)
	err = json.Unmarshal(binary.StringToBytes(str), &card)
	if err != nil {
		return
	}
	msg = append(msg, message.Text("标题: ", card.Title, "\n"))
	msg = append(msg, message.Text("UP主: ", card.Owner.Name, "\n"))
	msg = append(msg, message.Text("播放: ", humanNum(card.Stat.View), ", 弹幕: ", humanNum(card.Stat.Danmaku)))
	msg = append(msg, message.Image(card.Pic))
	msg = append(msg, message.Text("点赞: ", humanNum(card.Stat.Like), ", 投币: ", humanNum(card.Stat.Coin), "\n"))
	msg = append(msg, message.Text("收藏: ", humanNum(card.Stat.Favorite), ", 分享: ", humanNum(card.Stat.Share), "\n"))
	msg = append(msg, message.Text(VURL, card.BvID))
	return
}

// VideoInfo 用av或bv查视频信息
func VideoInfo(id string) (msg []message.MessageSegment, err error) {
	var data []byte
	_, err = strconv.Atoi(id)
	if err == nil {
		data, err = web.GetData(fmt.Sprintf(VideoInfoURL, id, ""))
	} else {
		data, err = web.GetData(fmt.Sprintf(VideoInfoURL, "", id))
	}
	if err != nil {
		return
	}
	return VideoCard2msg(gjson.ParseBytes(data).Get("data").Raw)
}
