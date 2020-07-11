package controllers

import (
	"fyoukuapi/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//获取频道地区列表
// @router /channel/region [*]
func (this *BaseController) ChannelRegion() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", regions, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

//获取频道类型列表
// @router /channel/type [*]
func (this *BaseController) ChannelType() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	num, types, err := models.GetChannelType(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", types, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}
