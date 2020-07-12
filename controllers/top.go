package controllers

import (
	"fyoukuapi/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

//根据频道获取排行榜
// @router /channel/top [*]
func (this *TopController) ChannelTop() {
	//获取频道ID
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelTop(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}