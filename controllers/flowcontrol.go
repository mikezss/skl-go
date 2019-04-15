package controllers

import (
	"encoding/json"
	"github.com/mikezss/skl-go/models"

	_ "fmt"

	"github.com/astaxie/beego"
)

// Operations about Users
type FLOWController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Saveflowtemplate() {
	var status = ""
	ob := models.FLOWTEMPLATEANDITEM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiFLOWTEMPLATE(ob.Flowtemplate, ob.Flowtemplateitem)

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
func (ctl *FLOWController) Copyflowtemplate() {
	var status = ""
	ob := models.COPYFLOWTEMPLATE{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.Copyflowtemplate(ob.Flowtemplateid, ob.COPYFlowtemplateid, ob.COPYFlowtemplatename)

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
func (ctl *FLOWController) Getflowtemplate() {
	var status = ""

	ob, err := models.GetAllFLOWTEMPLATE()

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
func (ctl *FLOWController) Getflowtemplateoptions() {
	var status = ""

	ob, err := models.GetAllFLOWTEMPLATEoptions()

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
func (ctl *FLOWController) Getflowtemplateitem() {
	var status = ""
	ob := models.FLOWTEMPLATE{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	flowtemplateitem, err := models.GetAllFLOWTEMPLATEITEM(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = flowtemplateitem
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getflowtemplatebyid() {
	var status = ""
	ob := models.FLOWTEMPLATE{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	flowtemplate, err := models.GetFLOWTEMPLATE(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = flowtemplate
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Deleteflowtemplate() {
	var status = ""
	ob := models.FLOWTEMPLATE{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteFLOWTEMPLATE(ob)

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
func (ctl *FLOWController) Saveflowtask() {
	var status = ""
	ob := models.FLOWTASKANDACTIONS{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddFLOWDEFINE(ob.Flowtask, ob.Flowmantaskaction, ob.Flowmantaskexecuter, ob.Flowswitchtaskaction)

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
func (ctl *FLOWController) Getflowtask() {
	var status = ""
	ob := models.FLOWTASK{}
	flowtask := make([]models.FLOWTASK, 0)
	mantaskactions := make([]models.FLOWMANTASKACTION, 0)
	mantaskexecuters := make([]models.FLOWTASKEXECUTER, 0)
	switchtaskactions := make([]models.FLOWSWITCHTASKACTION, 0)

	ob3 := make([]map[string]interface{}, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	flowtask, mantaskactions, mantaskexecuters, switchtaskactions, _ = models.GetFLOWDEFINE(ob.Flowtemplateid)
	for _, ft := range flowtask {
		mp := make(map[string]interface{}, 0)
		mp["flowtask"] = ft

		manswitchtaskactions := make([]interface{}, 0)
		for _, mt := range mantaskactions {
			if mt.Taskid == ft.Taskid {
				manswitchtaskactions = append(manswitchtaskactions, mt)
			}
		}
		for _, mt := range switchtaskactions {
			if mt.Taskid == ft.Taskid {
				manswitchtaskactions = append(manswitchtaskactions, mt)
			}
		}
		mp["actions"] = manswitchtaskactions
		mantaskexecuters2 := make([]models.FLOWTASKEXECUTER, 0)

		for _, mt := range mantaskexecuters {
			if mt.Taskid == ft.Taskid {
				mantaskexecuters2 = append(mantaskexecuters2, mt)
			}
		}
		mp["executers"] = mantaskexecuters2
		ob3 = append(ob3, mp)
	}
	//ob2 :={"Flowtask": flowtask, "Flowmantaskaction": mantaskactions, "Flowmantaskexecuter": mantaskexecuters, "Flowswitchtaskaction": switchtaskactions}

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob3
		ctl.ServeJSON()
	}

}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Deleteflowtask() {
	var status = ""
	ob := models.FLOWTASK{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteTaskid(ob)

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
func (ctl *FLOWController) Gettaskinfo() {
	var status = ""
	ob := models.TASKINFO{}
	ob2 := make([]models.TASKINFO, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.GetTaskinfo(ob)

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
func (ctl *FLOWController) Gettodotask() {
	var status = ""
	ob := models.FIFLOW{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Gettodotask(ob)

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
func (ctl *FLOWController) Gettodotasklist() {
	var status = ""
	ob := models.FITASK{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Gettodotasklist(ob)

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
func (ctl *FLOWController) Getdonetask() {
	var status = ""
	ob := models.FIFLOW{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Getdonetask(ob)

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
func (ctl *FLOWController) Getdonetasklist() {
	var status = ""
	ob := models.FITASK{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Getdonetasklist(ob)

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
func (ctl *FLOWController) Getflowmonitorbypageindex() {
	var status = ""
	ob := models.FIFLOWPAGEINDEX{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Getflowmonitorbypageindex(ob)

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
func (ctl *FLOWController) Getmyflow() {
	var status = ""
	ob := models.FIFLOW{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Getmyflow(ob)

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
func (ctl *FLOWController) Cancelflow() {
	var status = ""
	ob := models.FIFLOW{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.Cancelflow(ob)

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
func (ctl *FLOWController) Skiptask() {
	var status = ""
	ob := models.TODOTASKLIST{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	err = models.Skiptask(ob)

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
func (ctl *FLOWController) Getfloworgvarybypageindex() {
	var err error
	ob := models.PAGE{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	floworgvarys := make([]models.FLOWORGVARY, 0)

	floworgvarys, err = models.GetAllFLOWORGVARYBYPAGEINDEX(ob)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = floworgvarys
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getfloworgvary() {
	var err error

	floworgvarys := make([]models.FLOWORGVARY, 0)

	floworgvarys, err = models.GetAllFLOWORGVARY()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = floworgvarys
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getfloworgvarycount() {
	var err error
	ob := models.PAGE{}
	ob, err = models.GetAllFLOWORGVARYCOUNT()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

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
func (ctl *FLOWController) Savefloworgvary() {

	floworgvarys := make([]models.FLOWORGVARY, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &floworgvarys)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiFLOWORGVARY(floworgvarys)

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
func (ctl *FLOWController) Getflowstatus() {
	var err error
	flowstatus := make([]models.FISTATUS, 0)

	flowstatus, err = models.GetAllFLOWSTATUS()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = flowstatus
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Saveflowstatus() {

	flowstatus := make([]models.FISTATUS, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &flowstatus)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiFLOWSTATUS(flowstatus)

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
func (ctl *FLOWController) Saveagent() {

	agents := make([]models.AGENT, 0)
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &agents)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddMultiagent(agents)

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
func (ctl *FLOWController) Getagent() {
	var err error
	agents := make([]models.AGENT, 0)

	agents, err = models.GetAllagent()

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = agents
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getuserforagent() {
	var err error
	userforagent := models.USERFORAGENT{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &userforagent)
	if err != nil {
		beego.Error(err)
	}
	agents := make([]models.PROFILE, 0)

	agents, err = models.GetUserforagent(userforagent)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = agents
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getuserforagentoptions() {
	var err error
	userforagent := models.USERFORAGENT{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &userforagent)
	if err != nil {
		beego.Error(err)
	}

	agents, err1 := models.GetUserforagentoptions(userforagent)

	if err1 != nil {

		ctl.Data["json"] = map[string]string{"status": err1.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = agents
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Getflowstatusbyfiid() {
	var err error
	ff := models.FIFLOW{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &ff)
	if err != nil {
		beego.Error(err)
	}
	fs := models.Getflowstatus(ff.Fiid)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {
		ctl.Data["json"] = fs
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *FLOWController) Transfersign() {
	var err error
	ff := models.TRANSFER{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &ff)
	if err != nil {
		beego.Error(err)
	}
	err = models.Transfersign(ff)

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
func (ctl *FLOWController) Transferpost() {
	var err error
	ff := models.TRANSFER{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &ff)
	if err != nil {
		beego.Error(err)
	}
	err = models.Transferpost(ff)

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
func (ctl *FLOWController) Leave() {
	var err error
	ff := models.TRANSFER{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &ff)
	if err != nil {
		beego.Error(err)
	}
	err = models.Leave(ff)

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
func (ctl *FLOWController) Gettodotaskfortransfer() {
	var status = ""
	ob := models.TRANSFER{}
	ob2 := make([]models.TODOTASKLIST, 0)

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err = models.Gettodotaskfortransfer(ob)

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
func (ctl *FLOWController) Getflowmonitorcount() {
	var err error
	fif := models.FIFLOW{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &fif)
	if err != nil {
		beego.Error(err)
	}
	ob := models.PAGE{}
	ob, err = models.Getflowmonitorcount(fif)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

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
func (ctl *FLOWController) Gettaskidbytiidforng() {
	var err error
	fif := models.FITASK{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &fif)
	if err != nil {
		beego.Error(err)
	}

	fif.Taskid = models.Gettaskidbytiidforng(fif.Tiid)

	if err != nil {

		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()
		return
	} else {

		ctl.Data["json"] = fif
		ctl.ServeJSON()
	}
}
