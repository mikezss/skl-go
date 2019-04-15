package controllers

import (
	"encoding/json"
	"github.com/mikezss/skl-go/models"

	_ "fmt"

	"github.com/astaxie/beego"
)

// Operations about Users
type ENUMController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMController) Saveenum() {
	var status = ""

	ob := models.ENUMANDITEM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiENUM(ob.Enum, ob.Enumitem)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
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
func (ctl *ENUMController) Getenum() {
	var status = ""

	ob, err := models.GetAllENUM()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMController) Getenumoptions() {
	var status = ""

	ob, err := models.GetAllENUMoptions()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMController) Getenumbyid() {
	var status = ""
	ob := models.ENUM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err := models.GetENUMBYID(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
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
func (ctl *ENUMController) Getenumitem() {
	var status = ""
	ob := models.ENUM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	expenseitem, err := models.GetAllENUMITEM(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = expenseitem
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMController) Getenumitemoptions() {
	var status = ""
	ob := models.ENUM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	expenseitem, err := models.GetAllENUMITEMoptions(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = expenseitem
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *ENUMController) Deleteenumbyid() {
	var status = ""
	ob := models.ENUM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteENUM(&ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}
