package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	Name string
	Mobile string
	Password string
	Status int
	Avatar string
	AddTime int64
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}


func init() {
	orm.RegisterModel(new(User))
}

func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile:mobile}
	fmt.Println(user)
	err := o.Read(&user, "Mobile")
	fmt.Println(err)
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

func UserSave(mobile, password string) error {
	o := orm.NewOrm()
	var user User
	user.Name = ""
	user.Password = password
	user.Mobile = mobile
	user.Status = 1
	user.AddTime = time.Now().Unix()
	_, err := o.Insert(&user)
	return err
}

func IsMobileLogin(mobile, password string) (int, string) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("password", password).Filter("mobile", mobile).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	}

	return user.Id, user.Name
}

//根据用户ID获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	o := orm.NewOrm()
	var user UserInfo
	err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
	return user, err
}