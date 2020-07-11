package controllers

import (
	"fyoukuapi/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// 获取顶部广告
// @router /channel/advert [*]
func (c *VideoController) ChannelAdvert() {
	channelId, _ := c.GetInt("channelId")

	if channelId == 0 {
		c.Data["json"] =  ReturnError(4001, "未指定频道")
		c.ServeJSON()
	}

	num, videos, err := models.GetChannelAdvert(channelId)

	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "为请求到数据，请稍后再试")
		c.ServeJSON()
	}
}

// 正在热播功能
// @router /channel/hot [get]
func (c *VideoController) ChannelHot() {
	channelId, _ := c.GetInt("channelId")

	if channelId == 0 {
		c.Data["json"] =  ReturnError(4001, "未指定频道")
		c.ServeJSON()
	}

	num, videos, err := models.GetChannelHotList(channelId)

	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "并未找到相应的热播功能，请稍后再试！")
		c.ServeJSON()
	}
}
