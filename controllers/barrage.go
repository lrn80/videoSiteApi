package controllers

import (
	"encoding/json"
	"fyoukuapi/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}


type WsData struct {
	CurrentTime int
	EpisodesId  int
}


var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 获取弹幕websocket
// @router /barrage/ws [*]
func (c *BarrageController) BarrageWs() {
	var (
		conn *websocket.Conn
		err error
		data []byte
		barrages []models.BarrageData
	)

	conn, err = upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	if err != nil {
		goto ERR
	}

	for  {
		 _, data, err = conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		// 获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)

		if err == nil {
			err := conn.WriteJSON(barrages)
			if err != nil {
				goto ERR
			}
		}
	}

ERR:
conn.Close()
}

//保存弹幕
// @router /barrage/save [*]
func (this *BarrageController) Save() {
	uid, _ := this.GetInt("uid")
	content := this.GetString("content")
	currentTime, _ := this.GetInt("currentTime")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")

	if content == "" {
		this.Data["json"] = ReturnError(4001, "弹幕不能为空")
		this.ServeJSON()
	}
	if uid == 0 {
		this.Data["json"] = ReturnError(4002, "请先登录")
		this.ServeJSON()
	}
	if episodesId == 0 {
		this.Data["json"] = ReturnError(4003, "必须指定剧集ID")
		this.ServeJSON()
	}
	if videoId == 0 {
		this.Data["json"] = ReturnError(4005, "必须指定视频ID")
		this.ServeJSON()
	}

	if currentTime == 0 {
		this.Data["json"] = ReturnError(4006, "必须指定视频播放时间")
		this.ServeJSON()
	}
	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", "", 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}
}