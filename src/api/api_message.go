package api

import (
	"encoding/xml"
)

// 消息类型
const (
	EV_EVENT 		   string = "event"              // 消息推送
	EV_TEXT 		   string = "text"               // 用户文本消息
	EV_IMAGE           string = "image"              // 用户图片消息
	EV_VOICE           string = "voice"              // 用户语音消息    
	EV_VIDEO           string = "video"              // 用户视频消息
	EV_SVIDEO          string = "shortvideo"         // 用户短视频消息
	EV_LOCATION        string = "location"           // 用户地理位置信息
	EV_LINK            string = "link"               // 用户连接类型信息
)

// 操作类型
const (
	OP_SUBSCRIBE 	   string = "subscribe"
	OP_UNSUBSCRIBE 	   string = "unsubscribe"
	OP_SCAN 		   string = "SCAN"                 // 重复扫二维码关注
	OP_LOCATION        string = "LOCATION"             // 上报地理位置
	OP_CLICK           string = "CLICK"                // 自定义菜单点击事件
)

type MessageBase struct {
	ToUserName  	string
	FromUserName 	string
	CreateTime   	int64
	MsgType         string
}

type TextMessage struct {
	XMLName 		xml.Name 	`xml:"xml"`
	MessageBase
	Content  		string
}

type Image struct {
	MediaId         int64
}

type Voice struct {
	MediaId         int64
}

type Video struct {
	MediaId         int64
	Title           string
	Description     string
}

type MusicMessage struct {
	Title           string
	Description     string
	MusicURL        string
	HQMusicUrl      string
	ThumbMediaId    int64
}

type Article struct {
	Title           string
	Description     string
	PicUrl          string
	Url             string
}

type MessageRoute struct {
	MsgType         string
	Event           string
}

type Message struct {
	XMLName 		xml.Name 	`xml:"xml" json:-`
	ToUserName 		string 
	FromUserName 	string
	CreateTime 		int

	MessageRoute      // 路由的关键信息

	// 带参数二维码扫描
	// 1 用户从未关注=>关注: qrscene_<parameter>
	// 2 用户关注=>再次关注: uint32, scene_id
	// 自定义菜单事件中KEY值
	EventKey        string    
	// 二维码的ticket，可用来换取二维码图片
	Ticket          string

	// 地理位置信息
	Latitude        float64             // 纬度
	Longitude       float64             // 经度
	Precision       float64             // 位置精确度


	// 用户消息
	// 所有类型消息公用
	MsgId           int64
	// 视频、图片、短视屏公用
	MediaId         int64       // 多媒体类型资源ID
	// 文本消息
	Content         string      // 文本消息内容
	// 图片类消息
	PicUrl         	string      // 图片消息内容url
	// 语音消息
	Format          string      // 语音消息格式
	Recognition     string      // 语音识别结果
	// 视频、短视屏消息
	ThumbMediaId    string      // 视频消息缩略图
	// 地址位置消息
	Location_X      float64     // 纬度
	Location_Y      float64     // 经度
	Scale           int         // 缩放大小
	Label           string      // 地理位置信息
	// 链接类型消息
	Title           string      // 链接标题
	Description     string      // 链接描述
	Url             string      // 链接URL
}
