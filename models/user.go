package models

import (
	"fmt"
	redisClient "fyoukuapi/service/redis"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"strconv"
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


//增加redis缓存 - 根据用户ID获取用户信息
func RedisGetUserInfo(uid int) (UserInfo, error) {
	var user UserInfo
	conn := redisClient.PoolConnect()
	defer conn.Close()

	redisKey := "user:id:" + strconv.Itoa(uid)
	//判断redis是否存在
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &user)
	} else {
		o := orm.NewOrm()
		err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
		if err == nil {
			//保存redis
			_, err = conn.Do("hmset", redis.Args{redisKey}.AddFlat(user)...)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return user, err
}
