package controllers

import (
	"encoding/json"
	"github.com/mikezss/skl-go/models"

	_ "fmt"

	"github.com/astaxie/beego"
)

// Operations about Users
type COMPANYController struct {
	//beego.Controller
	BASEController
}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *COMPANYController) Savecompany() {
	var status = ""
	ob := make([]models.COMPANY, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCOMPANY(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *COMPANYController) Getcompany() {
	ob := models.COMPANY{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	var status = ""

	ob2, err := models.GetCOMPANY(&ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *COMPANYController) Getallcompany() {
	//var status = ""

	users, err := models.GetAllCOMPANY()

	if err != nil {

		//status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		//status = "ok"
		ctl.Data["json"] = users
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *COMPANYController) Getallcompanyoptions() {
	//var status = ""

	users, err := models.GetAllCOMPANYoptions()

	if err != nil {

		//status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		//status = "ok"
		ctl.Data["json"] = users
		ctl.ServeJSON()
	}

}
func (ctl *COMPANYController) Deletecompany() {
	var status = ""
	ob := models.COMPANY{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteCOMPANY(&ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}
