package controllers

import (
	"encoding/json"
	_ "fmt"
	"github.com/mikezss/skl-go/models"
	"time"

	"github.com/astaxie/beego"
)

// Operations about Users
type LOGINController struct {
	//beego.Controller
	BASEController
}
type Userpassword struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *LOGINController) Login() {
	var status = ""
	var ob Userpassword
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	beego.Debug(ob.UserName)
	beego.Debug(ob.Password)
	ip := ctl.Ctx.Request.RemoteAddr
	companycode := "19000500"
	isok := models.Checklogin("19000500", ob.UserName, ob.Password)

	//isok = true
	if !isok {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {

		ctl.CruSession.Set("PROFILE", models.PROFILE{Userid: ob.UserName, Username: ob.UserName, Companycode: companycode, Loginip: ip, Logintime: time.Now()})

		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}
