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