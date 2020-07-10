package controllers

import (
	"fmt"
	"fyoukuapi/models"
	"github.com/astaxie/beego"
	"regexp"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title user register
// @Description save user info
// @Param mobile fromData string true "用户手机号"
// @Param password fromData string true "用户密码"
// @router /register/save [post]
func (this *UserController) SaveRegister()  {
	var (
		mobile string
		password string
		err error
	)

	mobile = this.GetString("mobile")
	password = this.GetString("password")

	// 判断
	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobile)
	fmt.Println(isorno)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机格式本正确")
		this.ServeJSON()
	}

	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不能为空")
		this.ServeJSON()
	}

	status := models.IsUserMobile(mobile)
	fmt.Println(status)
	if status {
		this.Data["json"] = ReturnError(4005, "这个手机号已经注册过了")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			this.Data["json"] = ReturnSuccess(200, "添加成功！", nil, 0)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
		}
	}
}

// @router /login/do [*]
func (c *UserController) LoginDo() {
	mobile := c.GetString("mobile")
	password := c.GetString("password")

	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空！")
		c.ServeJSON()
	}

	if password == "" {
		c.Data["json"] = ReturnError(4002, "密码不能为空！")
		c.ServeJSON()
	}

	isorno, _ := regexp.MatchString(`^1(3|5|7|8|4)[0-9]{4,9}$`, mobile)

	if !isorno {
		c.Data["json"] = ReturnError(4003, "手机号码的格式不正确")
		c.ServeJSON()
	}

	uid, name := models.IsMobileLogin(mobile, MD5V(password))

	if uid != 0 {
		c.Data["json"] = ReturnSuccess(0, "登陆成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "账号或者密码错误")
		c.ServeJSON()
	}
}