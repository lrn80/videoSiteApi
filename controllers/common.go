package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"time"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64 `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg, Items: items, Count: count}
	return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg}
	return
}

func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

//格式化时间
func DateFormat(times int64) string {
	videoTime := time.Unix(times, 0)
	return videoTime.Format("2006-01-02")
}
