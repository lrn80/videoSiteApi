package controllers

import (
	"encoding/json"
	"fmt"
	"fyoukuApi-master/fyoukuapi/models"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/astaxie/beego"
)

type AliyunController struct {
	beego.Controller
}


type JSONS struct {
	RequestId     string
	UploadAddress string
	UploadAuth    string
	VideoId       string
}

var (
	accessKeyId     = "LTAI4GHtdQ1AvJhgZ8Hryi97"
	accessKeySecret = "MZ6uz9ICUBZeohp4XaiPOJcPpRrVz6"
)

func (c *AliyunController) InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {
	// 点播服务接入区域
	regionId := "cn-shanghai"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func (c *AliyunController) MyCreateUploadVideo(client *vod.Client, title string, desc string, fileName string, coverUrl string, tags string) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = title
	request.Description = desc
	request.FileName = fileName
	//request.CateId = "-1"
	request.CoverURL = coverUrl
	request.Tags = tags
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}


// @router /aliyun/create/upload/video [*]
func (c *AliyunController) CreateUploadVideo() {
	title := c.GetString("title")
	desc := c.GetString("desc")
	fileName := c.GetString("fileName")
	coverUrl := c.GetString("coverUrl")
	tags := c.GetString("tags")
	client, err := c.InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	response, err := c.MyCreateUploadVideo(client, title, desc, fileName, coverUrl, tags)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetHttpContentString())
	//fmt.Printf("VideoId: %s\n UploadAddress: %s\n UploadAuth: %s",
	//    response.VideoId, response.UploadAddress, response.UploadAuth)
	data := &JSONS{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *AliyunController) MyRefreshUploadVideo(client *vod.Client, videoId string) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.RefreshUploadVideo(request)
}

// @router /aliyun/refresh/upload/video [*]
func (c *AliyunController) RefreshUploadVideo() {
	videoId := c.GetString("videoId")

	client, err := c.InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	response, err := c.MyRefreshUploadVideo(client, videoId)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetHttpContentString())
	//fmt.Printf("UploadAddress: %s\n UploadAuth: %s", response.UploadAddress, response.UploadAuth)
	data := &JSONS{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.Data["json"] = data
	c.ServeJSON()
}

type PlayJSONS struct {
	PlayAuth string
}


// 获取上传凭证
// @router /aliyun/video/play/auth [*]
func (c *AliyunController) GetPlayAuth() {
	videoId := c.GetString("videoId")
	client, err := c.InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	response, err := c.MyGetPlayAuth(client, videoId)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetHttpContentString())
	//fmt.Printf("%s: %s\n", response.VideoMeta, response.PlayAuth)
	data := &PlayJSONS{
		response.PlayAuth,
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *AliyunController) MyGetPlayAuth(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetVideoPlayAuth(request)
}



type CallbackData struct {
	EventTime   string
	EventType   string
	VideoId     string
	Status      string
	Exteng      string
	StreamInfos []CallbackStreamInfosData
}
type CallbackStreamInfosData struct {
	Status     string
	Bitrate    int
	Definition string
	Duration   int
	Encrypt    bool
	FileUrl    string
	Format     string
	Fps        int
	Height     int
	Size       int
	Width      int
	JobId      string
}

//回调函数
// @router /aliyun/video/callback [*]
func (this *AliyunController) VideoCallback() {
	var ob CallbackData
	r := this.Ctx.Input.RequestBody
	json.Unmarshal(r, &ob)

	models.SaveAliyunVideo(ob.VideoId, string(r))
	this.Ctx.WriteString("success")
}
