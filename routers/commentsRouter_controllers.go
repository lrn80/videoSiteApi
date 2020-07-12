package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["fyoukuapi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/barrage/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "BarrageWs",
            Router: "/barrage/ws",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:BaseController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: "/channel/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:BaseController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channel/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/comment/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/comment/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:TopController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: "/channel/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:TopController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: "/type/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginDo",
            Router: "/login/do",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveRegister",
            Router: "/register/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMessageDo",
            Router: "/send/message",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: "/channel/advert",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHot",
            Router: "/channel/hot",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendRegionList",
            Router: "/channel/recommend/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetChannelRecomendTypeList",
            Router: "/channel/recommend/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: "/channel/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: "/video/episodes/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["fyoukuapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: "/video/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
