package controllers

import (
	"encoding/json"

	"github.com/mikezss/skl-go/models"

	_ "fmt"

	"github.com/astaxie/beego"
)

// Operations about Users
type ENUMSEARCHController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMSEARCHController) Getenumsearchbypageindex() {
	var status = ""
	ob := models.ENUMSEARCH{}
	ob2 := make([]models.ENUMSEARCHITEM, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Getenumsearchbypageindex(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob2
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMSEARCHController) Getenumsearchcount() {
	var err error
	fif := models.ENUMSEARCH{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &fif)
	if err != nil {
		beego.Error(err)
	}
	ob := models.PAGE{}
	ob, err = models.Getenumsearchcount(fif)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}
}
