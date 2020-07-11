package models

import "github.com/astaxie/beego/orm"

type Region struct {
	Id   int
	Name string
}
type Type struct {
	Id   int
	Name string
}

func GetChannelRegion(channelId int) (int64, []Region, error) {
	o := orm.NewOrm()
	var regions []Region
	num, err := o.Raw("SELECT id,name FROM channel_region WHERE status=1 AND channel_id=? ORDER BY sort DESC", channelId).QueryRows(&regions)
	return num, regions, err
}

func GetChannelType(channelId int) (int64, []Type, error) {
	o := orm.NewOrm()
	var types []Type
	num, err := o.Raw("SELECT id,name FROM channel_type WHERE status=1 AND channel_id=? ORDER BY sort DESC", channelId).QueryRows(&types)
	return num, types, err
}
