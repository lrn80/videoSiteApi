package controllers

import (
	"fmt"
	"fyoukuapi/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

//根据频道获取排行榜
// @router /channel/top [get]
func (c *TopController) ChannelTop() {
	//获取频道ID
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}
}

//根据类型获取排行榜
// @router /type/top [get]
func (c *TopController) TypeTop() {
	typeId, _ := c.GetInt("typeId")
	if typeId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定类型")
		c.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(typeId)
	fmt.Println(videos)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}
}
