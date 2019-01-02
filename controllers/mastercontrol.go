package controllers

import (
	"encoding/json"
	"github.com/mikezss/skl-go/models"

	_ "fmt"

	"github.com/astaxie/beego"
)

// Operations about Users
type MASTERController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Savemodual() {
	var status = ""
	ob := models.CMN_MODUAL_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.SaveModual(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getallmodual() {
	var status = ""

	ob, err := models.GetAllModual()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getallmodualoptions() {
	var status = ""

	ob, err := models.GetAllModualoptions()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getmodual() {
	var status = ""
	ob := models.CMN_MODUAL_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err := models.GetModualbyid(ob)

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
func (ctl *MASTERController) Getmodualtreejson() {

	modualtreeson := models.CreateModualTreeJson()

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(modualtreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deletemodual() {
	var status = ""
	ob := models.CMN_MODUAL_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteModualbyid(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Saveorg() {
	var status = ""
	ob := models.CMN_ORGANDLEADER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.SaveOrg(ob.Org, ob.Orgleader)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getallorg() {
	var status = ""

	ob, err := models.GetAllOrg()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}
}

func (ctl *MASTERController) Getallorgoptions() {
	var status = ""

	ob, err := models.GetAllOrgoptions()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getorg() {
	var status = ""
	ob := models.CMN_ORG_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, ob3, err := models.GetOrgbyid(ob)
	mp := make(map[string]interface{}, 0)
	mp["org"] = ob2
	mp["orgleader"] = ob3

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = mp
		ctl.ServeJSON()
	}

}
func (ctl *MASTERController) Getorgtreejson() {

	orgtreeson := models.CreateOrgTreeJson()

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(orgtreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deleteorg() {
	var status = ""
	ob := models.CMN_ORG_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteOrgbyid(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}
func (ctl *MASTERController) Getleaders() {
	var status = ""

	leaders, err := models.GetLeaders()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = leaders
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Saveuser() {
	var status = ""
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob.Password = "666666"
	ob.Expired = "0"
	err = models.AddCMN_USER_TB(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Deleteuser() {
	var status = ""
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteCMN_USER_TB(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getuser() {
	//var status = ""
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	users, err := models.GetCMN_USER_TB(ob)

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
func (ctl *MASTERController) Getuseroptionsbyorgid() {
	//var status = ""
	ob := models.CMN_ORG_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	users, err := models.Getuseroptionsbyorgid(ob)

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
func (ctl *MASTERController) Getalluser() {
	//var status = ""

	users, err := models.GetALLCMN_USER_TB()

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
func (ctl *MASTERController) Getalluseroptions() {
	//var status = ""

	users, err := models.GetALLCMN_USER_TBoptions()

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
func (ctl *MASTERController) Saverole() {
	var status = ""
	ob := models.CMN_ROLE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.SaveRole(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getallrole() {
	var status = ""

	ob, err := models.GetAllRole()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getallroleoptions() {
	var status = ""

	ob, err := models.GetAllRoleoptions()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getrole() {
	var status = ""
	ob := models.CMN_ROLE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err := models.GetRolebyid(ob)

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
func (ctl *MASTERController) Getroletreejson() {

	roletreeson := models.CreateRoleTreeJson()

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(roletreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deleterole() {
	var status = ""
	ob := models.CMN_ROLE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteRolebyid(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Saveroleprivileges() {
	status := ""
	ob := make([]models.CMN_ROLEPRIVILEGE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCMN_ROLEPRIVILEGE_TB(ob[0].Roleid, ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()
}
func (ctl *MASTERController) Getroleprivilegetreejson() {
	ob := models.CMN_ROLE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	roleprivilege, _ := models.GetCMN_ROLEPRIVILEGE_TB(ob.Roleid)
	roletreeson := models.CreateRolePrivilegeTreeJson(roleprivilege)

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(roletreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deleteroleprivilege() {
	status := ""
	ob := models.CMN_ROLE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteCMN_ROLEPRIVILEGE_TB(ob.Roleid)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Saveorgprivileges() {
	status := ""
	ob := make([]models.CMN_ORGROLE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCMN_ORGROLE_TB(ob[0].Orgid, ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()
}
func (ctl *MASTERController) Getorgroletreejson() {
	ob := models.CMN_ORG_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	orgprivilege, _ := models.GetCMN_ORGROLE_TB(ob)
	roletreeson := models.CreateOrgRoleTreeJson(orgprivilege)

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(roletreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deleteorgprivilege() {
	status := ""
	ob := models.CMN_ORG_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteCMN_ORGROLE_TB(ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Savegroupprivileges() {
	status := ""
	ob := make([]models.CMN_GROUPROLE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCMN_GROUPROLE_TB(ob[0].Groupid, ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()
}

func (ctl *MASTERController) Deletegroupprivilege() {
	status := ""
	ob := models.CMN_GROUP_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteCMN_GROUPROLE_TB(ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Savegroup() {
	var status = ""
	ob := models.CMN_GROUP_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.AddCMN_GROUP_TB(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getallgroup() {
	var status = ""

	ob, err := models.GetAllCMN_GROUP_TB()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getgroup() {
	var status = ""
	ob := models.CMN_GROUP_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err := models.GetGroupbyid(ob)

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

func (ctl *MASTERController) Getgrouptreejson() {

	grouptreeson := models.CreateGroupTreeJson()

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(grouptreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Getgrouproletreejson() {
	ob := models.CMN_GROUP_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	groupprivilege, _ := models.GetCMN_GROUPROLE_TB(ob)
	roletreeson := models.CreateGroupRoleTreeJson(groupprivilege)

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(roletreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Deletegroup() {
	var status = ""
	ob := models.CMN_GROUP_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteGroupbyid(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Saveusergroup() {
	var status = ""
	ob := make([]models.CMN_USERGROUP_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.AddMultiCMN_USERGROUP_TB(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getusergroup() {

	ob := models.CMN_GROUP_TB{}
	obs := make([]models.CMN_USERGROUP_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	obs, err = models.GetCMN_USERGROUP_TB(ob)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = obs
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Saveuserprivileges() {
	status := ""
	ob := make([]models.CMN_USERROLE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCMN_USERROLE_TB(ob[0].Userid, ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()
}

func (ctl *MASTERController) Deleteuserprivilege() {
	status := ""
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteCMN_USERROLE_TB(ob)
	if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	ctl.Data["json"] = map[string]string{"status": status}
	ctl.ServeJSON()
}

func (ctl *MASTERController) Getuserroletreejson() {
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	userprivilege, _ := models.GetCMN_USERROLE_TB(ob)
	roletreeson := models.CreateUserRoleTreeJson(userprivilege)

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(roletreeson))
	ctl.Ctx.Output.Body(content)

}
func (ctl *MASTERController) Getusermodualtreejson() {
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	modualtreeson := models.CreateModualTreeJsonForuser(ob)

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(modualtreeson))
	ctl.Ctx.Output.Body(content)

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Savemodualtemplate() {
	var status = ""
	ob := models.CMN_MODUALTEMPLATE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.SaveModualtemplate(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Deletemodualtemplate() {
	var status = ""
	ob := models.CMN_MODUALTEMPLATE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteModualtemplatebyid(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getmodualtemplate() {

	ob := models.CMN_MODUALTEMPLATE_TB{}
	ob2 := models.CMN_MODUALTEMPLATE_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.GetModualtemplatebyid(ob)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

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
func (ctl *MASTERController) Getusersbyorgid() {

	org := models.CMN_ORG_TB{}
	users := make([]models.CMN_USER_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &org)
	if err != nil {
		beego.Error(err)
	}
	users, err = models.GetUsersbyorgid(org)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

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
func (ctl *MASTERController) Getorgvary() {
	var err error
	orgvarys := make([]models.ORGVARY, 0)

	orgvarys, err = models.GetAllORGVARY()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = orgvarys
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getorgvaryoptions() {

	orgvarys, err := models.GetAllORGVARYoptions()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = orgvarys
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Saveorgvary() {

	orgvarys := make([]models.ORGVARY, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &orgvarys)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiORGVARY(orgvarys)

	if err != nil {
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = map[string]string{"status": "ok"}
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Passwordchange() {

	user := models.CMN_USER_TB{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &user)
	if err != nil {
		beego.Error(err)
	}

	err = models.PASSWORDCHANGE(user)

	if err != nil {
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = map[string]string{"status": "ok"}
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Passwordreset() {

	user := models.CMN_USER_TB{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &user)
	if err != nil {
		beego.Error(err)
	}

	err = models.PASSWORDRESET(user)

	if err != nil {
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = map[string]string{"status": "ok"}
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Saveuserrole() {
	var status = ""
	ob := make([]models.CMN_USERROLE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.AddMultiCMN_USERROLE_TBbyrole(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getuserrole() {

	ob := models.CMN_USERROLE_TB{}
	obs := make([]models.CMN_USERROLE_TB, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	obs, err = models.GetCMN_USERROLE_TBbyroleid(ob)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = obs
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Updateuserinfo() {
	var status = ""
	ob := models.CMN_USER_TB{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateCMN_USER_TB(ob)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Uploadusers() {
	var status = ""
	Filepath := models.CMN_FILEINFO_TB{}
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &Filepath)
	if err != nil {
		beego.Error(err)
	}

	err = models.Uploadusers(Filepath)

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

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Savelang() {
	var status = ""
	ob := make([]models.CMN_LANG_TB, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiCMN_LANG_TB(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
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
func (ctl *MASTERController) Getlang() {
	var status = ""

	ob, err := models.GetAllCMN_LANG_TB()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
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
func (ctl *MASTERController) Getlangcount() {
	var status = ""
	ob := models.LANG{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err1 := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err1 != nil {
		beego.Error(err1)
	}
	ob2, err := models.GetAllCMN_LANG_TBcount(ob)

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
func (ctl *MASTERController) Getlangbypageindex() {
	var status = ""
	ob := models.LANGPAGEINDEX{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err1 := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err1 != nil {
		beego.Error(err1)
	}
	ob2, err := models.GetAllCMN_LANG_TBbypageindex(ob)

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

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getnavigatormodualbyuser() {
	var status = ""
	ob := models.CMN_USER_TB{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	navigatormodual, err := models.GetNAVIGATORMODUALBYUSER(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = navigatormodual
		ctl.ServeJSON()
	}
}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getmenumodualbyparent() {
	var status = ""
	ob := models.CMN_USERMODUAL_TB{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	menumodual, err := models.GetMENUMODUALBYPARENT(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = menumodual
		ctl.ServeJSON()
	}
}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Loadlangjson() {
	var status = ""

	err := models.Loadlangjson()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}
}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Updatelangjson() {
	var status = ""
	ob := make([]models.CMN_LANG_TB, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.Updatelangjson(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Updateprojectpath() {
	var status = ""
	ob := models.PROJECTPATH{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.SAVEPROJECTPATH(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}

}

// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *MASTERController) Getprojectpath() {
	var status = ""

	ob, err := models.GETPROJECTPATH()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
		return
	} else {
		status = "ok"
		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}

}
