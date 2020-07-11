package models

import "github.com/astaxie/beego/orm"

type VideoData struct {
	Id int
	Title string
	SubTitle string
	AddTime int64
	Img string
	Img1 string
	EpisodesCount int
	IsEnd int
}

func init()  {
	orm.RegisterModel(new(VideoData))
}

func GetChannelHotList(channelId int) (int64, []VideoData, error)  {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end " +
		"from video where status=1 and is_hot=1 and channel_id=? " +
		"order by episodes_update_time desc limit 9", channelId).QueryRows(&videos)
	return num, videos, err
}
