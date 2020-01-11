package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/guuid"
	"github.com/gogf/gf/util/gvalid"
	"jwkauthserver/app/model/pgdb"
	"time"
)

var mydb, err = pgdb.Newdb()

// 用户注册
func SignUp(data map[string]string) error {
	// 数据校验
	rules := []string{
		"passport @required|length:6,16#账号不能为空|账号长度应当在:min到:max之间",
		"password2@required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间",
		"password @required|length:6,16|same:password2#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等",
	}
	if e := gvalid.CheckMap(data, rules); e != nil {
		return errors.New(e.String())
	}
	if _, ok := data["id"]; !ok {
		data["id"] = guuid.New().String()
	}

	if data["password"], err = gmd5.EncryptString(data["id"] + data["password"]); err != nil {
		return err
	}
	if _, ok := data["nickname"]; !ok {
		data["nickname"] = data["passport"]
	}

	// 唯一性数据检查
	if !CheckPassport(data["passport"]) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data["passport"]))
	}
	if !CheckNickName(data["nickname"]) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data["nickname"]))
	}
	// 记录账号创建/注册时间
	if _, ok := data["create_time"]; !ok {
		data["create_time"] = gtime.Now().String()
	}
	if _, err := mydb.Insert("go.users", gdb.Map{
		"id":data["id"],
		"passport":data["passport"],
		"password":data["password"],
		"nickname":data["nickname"],
		"create_time":data["create_time"],
	}); err != nil {
		return err
	}
	return nil
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(passport, password string) error {
	record, err := mydb.GetOne("select id,passport,password from go.users where passport=?",passport)
	if err != nil {
		return err
	}
	id := record["id"].String()
	passwordmd5 := record["password"].String()
	if passwordmd5!=gmd5.MustEncryptString(id+password)  {
		return errors.New("账号或密码错误")
	}
	return nil
}

func Usertoken(passport string) (gin.H,error){
	role := [2]string{"role_a", "role_b"}
	now := time.Now()
	var nowafter1hour,nowafter24hour time.Time
	usertoken:=gin.H{}
	if After1Hour,err := time.ParseDuration("1h"); err==nil{
		nowafter1hour = now.Add(After1Hour)
	} else {
		return usertoken,err
	}
	if After24Hour,err := time.ParseDuration("24h");err == nil {
		nowafter24hour = now.Add(After24Hour)
	} else {
		return usertoken,err
	}
	access_token:=gin.H{
		"aud": "localhost",
		"iss": "AuthCenter",
		"sub": passport,
		"roles": role,
		"exp": nowafter1hour.Unix(),
	}
	refresh_token:=gin.H{
		"aud": "localhost",
		"iss": "AuthCenter",
		"sub": passport,
		"exp": nowafter24hour.Unix(),
	}
	usertoken = gin.H{
		"access_token":access_token,
		"refresh_token":refresh_token,
	}
	return usertoken,nil
}

/*// 用户注销
func SignOut(session *ghttp.Session) {
	session.Remove(USER_SESSION_MARK)
}*/

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckPassport(passport string) bool {
	i, err := mydb.GetValue("select passport from go.users where passport=?", passport)
	if err == nil {
		if i.String() == passport {
			return false
		}
	}
	return true
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	i, err := mydb.GetValue("select nickname from go.users where nickname=?", nickname)
	if err == nil {
		if i.String() == nickname {
			return false
		}
	}
	return true
}
