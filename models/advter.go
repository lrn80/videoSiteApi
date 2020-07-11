package models

import "github.com/astaxie/beego/orm"

type Advert struct {
	Id int
	Title string
	SubTitle string
	AddTime int64
	Img string
	Url string
}

func init()  {
	orm.RegisterModel(new(Advert))
}

func GetChannelAdvert(channelId int) (int64, []Advert, error) {
	o := orm.NewOrm()
	var adverts []Advert
	num, err := o.Raw("select id, title, sub_title, img, add_time, url from advert " +
		"where status=1 and channel_id=? order by sort desc limit 1", channelId).QueryRows(&adverts)
	return num, adverts, err
}
