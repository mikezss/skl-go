package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//流程定义fi_template_tb表
type COPYFLOWTEMPLATE struct {
	Flowtemplateid       string
	COPYFlowtemplateid   string
	COPYFlowtemplatename string
}

type FLOWTEMPLATE struct {
	Flowtemplateid   string `orm:"pk;column(flowtemplateid)"`
	Flowtemplatename string `orm:"column(flowtemplatename)"`
	Flowcontent      string `orm:"column(flowcontent);null"`
	Flowinstidcol    string `orm:"column(flowinstidcol);null"`
	Flowstatuscol    string `orm:"column(flowstatuscol);null"`
}

//流程定义fi_templateitem_tb表
type FLOWTEMPLATEITEM struct {
	Id             int    `orm:"auto;column(id)"`
	Flowtemplateid string `orm:"column(flowtemplateid)"`
	Vary           string `orm:"column(vary)"`
	Varyname       string `orm:"column(varyname)"`
	Varytype       string `orm:"column(varytype)"`
	Varyvalue      string `orm:"column(varyvalue);null"`
}

//
type FLOWTEMPLATEANDITEM struct {
	Flowtemplate     FLOWTEMPLATE
	Flowtemplateitem []FLOWTEMPLATEITEM
}

//流程定义fi_flowtask_tb表
type FLOWTASK struct {
	Flowtemplateid string
	Tasktype       string
	Taskid         string
	Taskname       string
	Supportskip    bool
	Sendmessage    bool
	Concurrent     string
	Samepersontask string
	Nopersontask   string
	Skiptotaskid   string
}

//流程定义人工任务action表fi_flowmantaskaction_tb
type FLOWMANTASKACTION struct {
	Flowtemplateid string
	Tasktype       string
	Taskid         string
	Action         string
	Jump           string
	Status         string
	Nexttask       string
	Backtask       string
}

//流程定义分支任务action表fi_flowswitchtaskaction_tb
type FLOWSWITCHTASKACTION struct {
	Flowtemplateid string
	Tasktype       string
	Taskid         string
	Nos            string
	Conditions     string
	Functions      string
	Valuee         string
	Jump           string
	Statuss        string
	Nexttask       string
	Backtask       string
}

//流程定义人工任务执行人表fi_flowtaskexecuter_tb
type FLOWTASKEXECUTER struct {
	Flowtemplateid string
	Tasktype       string
	Taskid         string
	No             string
	Taskexecuter   string
	Expression     string
}

//流程定义分支任务action表
type FLOWTASKANDACTIONS struct {
	Flowtask             FLOWTASK
	Flowmantaskaction    []FLOWMANTASKACTION
	Flowmantaskexecuter  []FLOWTASKEXECUTER
	Flowswitchtaskaction []FLOWSWITCHTASKACTION
}

//fi_flow表
type FIFLOW struct {
	Fiid           int       `orm:"pk;column(fiid)"`
	Modualid       string    `orm:"column(modualid)"`
	Flowtemplateid string    `orm:"column(flowtemplateid)"`
	Flowcontent    string    `orm:"column(flowcontent)"`
	Caller         string    `orm:"column(caller)"`
	Flowstarttime  time.Time `orm:"column(flowstarttime)"`
	Flowfinishtime time.Time `orm:"column(flowfinishtime);null"`
	Flowstatus     string    `orm:"column(flowstatus);default(0)"`
	State          string    `orm:"column(state);null"`
}

//fi_flow表
type FIFLOWPAGEINDEX struct {
	Fiid           int
	Modualid       string
	Flowtemplateid string
	Flowcontent    string
	Caller         string
	Flowstarttime  time.Time
	Flowfinishtime time.Time
	Flowstatus     string
	State          string
	Pageindex      int
	Pagesize       int
}

//fi_task表
type FITASK struct {
	Tiid           int       `orm:"pk;column(tiid)"`
	Fiid           int       `orm:"column(fiid)"`
	Taskid         string    `orm:"column(taskid)"`
	Taskstarttime  time.Time `orm:"column(taskstarttime)"`
	Taskfinishtime time.Time `orm:"column(taskfinishtime);null"`
	Editor         string    `orm:"column(editor);null"`
	Actionid       string    `orm:"column(actionid)"`
	Opinion        string    `orm:"column(opinion);null"`
	Direction      string    `orm:"column(direction);null"`
	Skiptotaskid   string    `orm:"column(skiptotaskid);null"`
	Islast         string    `orm:"column(islast);null"`
	Taskstatus     string    `orm:"column(taskstatus);null;default(0)"`
}

//任务预览表fi_task_prev
type FITASKPREVIEW struct {
	Id            int32  `orm:"auto;column(id)"`
	Fiid          int    `orm:"column(fiid)"`
	Prevtiid      int    `orm:"column(prevtiid)"`
	Prevdirection string `orm:"column(prevdirection);null"`
	Prevtaskid    string `orm:"column(prevtaskid)"`
	Taskid        string `orm:"column(taskid)"`
	Tiid          int    `orm:"column(tiid)"`
}

//fi_owner表
type FIOWNER struct {
	Tiid  int
	Owner string
}

//fi_var表
type FIVAR struct {
	Fiid   int
	Vid    string
	Vvalue string
}

//共通
type MODUALCNTANDMNY struct {
	Opinion   string
	Submitter string
	Content   string
	Amount    float64
}

//主管
type LEADERS struct {
	Leaders []string
}

//任务信息
type TASKINFO struct {
	Modualid    string
	Currentfiid int
	Currenttiid int
	Action      string
}

//待办任务列表
type TODOTASKLIST struct {
	Fiid           int
	Tiid           int
	Caller         string
	Flowcontent    string
	Taskid         string
	Flowstarttime  time.Time
	Flowfinishtime time.Time
	Flowstatus     string
	Url            string
	Routerlink     string
	Opinion        string
	Flowtemplateid string
	Editor         string
	Taskstarttime  time.Time
	Taskfinishtime time.Time
	Taskstatus     string
	Supportskip    bool
	Flowstatusname string
	Taskname       string
	Checked        bool
	Pageindex      int
	Pagesize       int
}

//流程状态
type FISTATUS struct {
	Flowstatus     string `orm:"pk;column(flowstatus)"`
	Flowstatusname string `orm:"column(flowstatusname)"`
}

//变量
type ORGVARY struct {
	Vid          string `orm:"pk;column(vid)"`
	Vname        string `orm:"column(vname)"`
	Defaultvalue string `orm:"column(defaultvalue)"`
}

//流程变量
type FLOWORGVARY struct {
	Orgid  string
	Vid    string
	Vvalue string
}

//代理人
type AGENT struct {
	Userid    string
	Agent     string
	Startdate time.Time
	Enddate   time.Time
}

//代理人
type USERFORAGENT struct {
	Submitter       string
	Userid          string
	Ownerdepartment bool
	Isleader        bool
}
type PORG struct {
	Porg []CMN_ORG_TB
}

//转签转岗离职
type TRANSFER struct {
	Usertype       string
	Submitter      string
	Userid         string
	Cancel         bool
	Transferuserid string
	Listdata       []TODOTASKLIST
}
type PAGE struct {
	Pageindex int
	Pagesize  int
	Total     int
	Orgid     string
	Vid       string
}

//会签
type COUNTERSIGN struct {
	Tiid       int
	Userid     string
	Fiid       int
	Taskid     string
	Taskstatus string
}

//用户更新自定义变量
type Updatefi_var interface {
	Updatevar(fiid int) (int64, error)
}

//机构变量表
func (u *ORGVARY) TableName() string {
	return "cmn_orgvary_tb"
}

//流程状态表
func (u *FISTATUS) TableName() string {
	return "fi_flowstatus"
}

//流程预览表
func (u *FITASKPREVIEW) TableName() string {
	return "fi_task_prev"
}

//任务表
func (u *FITASK) TableName() string {
	return "fi_task"
}

//流程表
func (u *FIFLOW) TableName() string {
	return "fi_flow"
}

//模板主表
func (u *FLOWTEMPLATE) TableName() string {
	return "fi_template_tb"
}

//模板子表
func (u *FLOWTEMPLATEITEM) TableName() string {
	return "fi_templateitem_tb"
}

//流程模板追加
func AddMultiFLOWTEMPLATE(u FLOWTEMPLATE, u2 []FLOWTEMPLATEITEM) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	deletesql := "delete from fi_template_tb where flowtemplateid=?"
	deletesql = ConvertSQL(deletesql, dbtype)
	_, err = o.Raw(deletesql, u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql2 := "delete from fi_templateitem_tb where flowtemplateid=?"
	deletesql2 = ConvertSQL(deletesql2, dbtype)
	_, err = o.Raw(deletesql2, u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	//_, err = o.Insert(&u)
	sql := "insert into fi_template_tb(flowtemplateid,flowtemplatename,flowcontent,flowinstidcol,flowstatuscol) values(?,?,?,?,?)"
	_, err = o.Raw(ConvertSQL(sql, dbtype), u.Flowtemplateid, u.Flowtemplatename, u.Flowcontent, u.Flowinstidcol, u.Flowstatuscol).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i, _ := range u2 {
		u2[i].Flowtemplateid = u.Flowtemplateid
	}
	//_, err = o.InsertMulti(len(u2), &u2)
	sql = "insert into fi_templateitem_tb(flowtemplateid,vary,varyname,varytype,varyvalue) values(?,?,?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, u3 := range u2 {
		_, err = o.Raw(sql, u3.Flowtemplateid, u3.Vary, u3.Varyname, u3.Varytype, u3.Varyvalue).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	selectsql := "select * from fi_flowtask_tb where flowtemplateid=?"
	selectsql = ConvertSQL(selectsql, dbtype)
	ft := make([]FLOWTASK, 0)
	rows, _ := o.Raw(selectsql, u.Flowtemplateid).QueryRows(&ft)

	if rows == 0 {
		sql = "insert into fi_flowtask_tb(flowtemplateid,tasktype,taskid,taskname,supportskip,sendmessage,concurrent,samepersontask,nopersontask) values(?,?,?,?,?,?,?,?,?)"
		sql = ConvertSQL(sql, dbtype)
		_, err = o.Raw(sql, u.Flowtemplateid, "man", "1", "启动", 0, 0, 0, "", "").Exec()

		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		_, err = o.Raw(sql, u.Flowtemplateid, "man", "999", "结束", 0, 0, 0, "", "").Exec()

		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}

//获得所有的流程模板
func GetAllFLOWTEMPLATE() (admins []FLOWTEMPLATE, err error) {
	admins = make([]FLOWTEMPLATE, 0)
	o := orm.NewOrm()

	sql := "select * from fi_template_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得所有的流程模板
func GetAllFLOWTEMPLATEoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select flowtemplateid as value,flowtemplatename as label from fi_template_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得所有的流程模板明细
func GetAllFLOWTEMPLATEITEM(e FLOWTEMPLATE) (admins []FLOWTEMPLATEITEM, err error) {
	admins = make([]FLOWTEMPLATEITEM, 0)
	o := orm.NewOrm()

	sql := "select * from fi_templateitem_tb where flowtemplateid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, e.Flowtemplateid).QueryRows(&admins)

	return admins, err
}

//根据流程模板ID获得流程模板数据
func GetFLOWTEMPLATE(e FLOWTEMPLATE) (admin FLOWTEMPLATE, err error) {

	o := orm.NewOrm()

	sql := "select * from fi_template_tb where flowtemplateid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Flowtemplateid).QueryRow(&admin)

	return admin, err
}

//删除流程模板，同时删除关联表数据
func DeleteFLOWTEMPLATE(u FLOWTEMPLATE) (err error) {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql := "delete from fi_template_tb where flowtemplateid=?"
	deletesql = ConvertSQL(deletesql, dbtype)
	_, err = o.Raw(deletesql, u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql_1 := "delete from fi_templateitem_tb where flowtemplateid=?"
	deletesql_1 = ConvertSQL(deletesql_1, dbtype)
	_, err = o.Raw(deletesql_1, u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql1 := "delete from fi_flowtask_tb where flowtemplateid=?"
	deletesql2 := "delete from fi_flowmantaskaction_tb where flowtemplateid=?"
	deletesql3 := "delete from fi_flowtaskexecuter_tb where flowtemplateid=?"
	deletesql4 := "delete from fi_flowswitchtaskaction_tb where flowtemplateid=?"
	_, err = o.Raw(ConvertSQL(deletesql1, dbtype), u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql2, dbtype), u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql3, dbtype), u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql4, dbtype), u.Flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

//追加流程定义任务
func AddFLOWDEFINE(u FLOWTASK, u2 []FLOWMANTASKACTION, u4 []FLOWTASKEXECUTER, u3 []FLOWSWITCHTASKACTION) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	deletesql1 := "delete from fi_flowtask_tb where flowtemplateid=? and taskid=?"
	deletesql2 := "delete from fi_flowmantaskaction_tb where flowtemplateid=? and taskid=?"
	deletesql3 := "delete from fi_flowtaskexecuter_tb where flowtemplateid=? and taskid=?"
	deletesql4 := "delete from fi_flowswitchtaskaction_tb where flowtemplateid=? and taskid=?"
	_, err = o.Raw(ConvertSQL(deletesql1, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql2, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql3, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql4, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	sql := "insert into fi_flowtask_tb(flowtemplateid,tasktype,taskid,taskname,supportskip,sendmessage,concurrent,samepersontask,nopersontask) values(?,?,?,?,?,?,?,?,?)"

	_, err = o.Raw(ConvertSQL(sql, dbtype), u.Flowtemplateid, u.Tasktype, u.Taskid, u.Taskname, u.Supportskip, u.Sendmessage, u.Concurrent, u.Samepersontask, u.Nopersontask).Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	if u.Tasktype == "man" {

		for i, _ := range u2 {
			u2[i].Flowtemplateid = u.Flowtemplateid
			u2[i].Tasktype = u.Tasktype
			u2[i].Taskid = u.Taskid
		}

		sql := "insert into fi_flowmantaskaction_tb(flowtemplateid,tasktype,taskid,action,jump,status,nexttask,backtask) values(?,?,?,?,?,?,?,?)"
		for _, u3 := range u2 {
			_, err = o.Raw(ConvertSQL(sql, dbtype), u3.Flowtemplateid, u3.Tasktype, u3.Taskid, u3.Action, u3.Jump, u3.Status, u3.Nexttask, u3.Backtask).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
		for i, _ := range u4 {
			u4[i].Flowtemplateid = u.Flowtemplateid
			u4[i].Tasktype = u.Tasktype
			u4[i].Taskid = u.Taskid
		}

		sql = "insert into fi_flowtaskexecuter_tb(flowtemplateid,tasktype,taskid,no,taskexecuter,expression) values(?,?,?,?,?,?)"
		for _, u5 := range u4 {
			_, err = o.Raw(ConvertSQL(sql, dbtype), u5.Flowtemplateid, u5.Tasktype, u5.Taskid, u5.No, u5.Taskexecuter, u5.Expression).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	} else {
		for i, _ := range u3 {
			u3[i].Flowtemplateid = u.Flowtemplateid
			u3[i].Tasktype = u.Tasktype
			u3[i].Taskid = u.Taskid
		}
		sql := "insert into fi_flowswitchtaskaction_tb(flowtemplateid,tasktype,taskid,nos,conditions,functions,valuee,jump,statuss,nexttask,backtask) values(?,?,?,?,?,?,?,?,?,?,?)"

		for _, u4 := range u3 {
			_, err = o.Raw(ConvertSQL(sql, dbtype), u4.Flowtemplateid, u4.Tasktype, u4.Taskid, u4.Nos, u4.Conditions, u4.Functions, u4.Valuee, u4.Jump, u4.Statuss, u4.Nexttask, u4.Backtask).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}

	}

	err = o.Commit()

	return err
}

//根据流程编号获得流程定义
func GetFLOWDEFINE(fiid string) (ft []FLOWTASK, mantaskactions []FLOWMANTASKACTION, mantaskexecuters []FLOWTASKEXECUTER, switchtaskactions []FLOWSWITCHTASKACTION, err error) {
	dbtype := Getdbtype()
	ft = make([]FLOWTASK, 0)
	mantaskactions = make([]FLOWMANTASKACTION, 0)
	mantaskexecuters = make([]FLOWTASKEXECUTER, 0)
	switchtaskactions = make([]FLOWSWITCHTASKACTION, 0)
	o := orm.NewOrm()

	sql := "select * from fi_flowtask_tb where flowtemplateid=? order by taskid"

	_, err = o.Raw(ConvertSQL(sql, dbtype), fiid).QueryRows(&ft)

	sql = "select * from fi_flowmantaskaction_tb where flowtemplateid=? order by taskid"

	_, err = o.Raw(ConvertSQL(sql, dbtype), fiid).QueryRows(&mantaskactions)

	sql = "select * from fi_flowtaskexecuter_tb where flowtemplateid=? order by taskid"

	_, err = o.Raw(ConvertSQL(sql, dbtype), fiid).QueryRows(&mantaskexecuters)

	sql = "select * from fi_flowswitchtaskaction_tb where flowtemplateid=? order by taskid"

	_, err = o.Raw(ConvertSQL(sql, dbtype), fiid).QueryRows(&switchtaskactions)

	return ft, mantaskactions, mantaskexecuters, switchtaskactions, err

}

//根据流程模板ID和任务ID，删除任务定义
func DeleteTaskid(u FLOWTASK) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	deletesql1 := "delete from fi_flowtask_tb where flowtemplateid=? and taskid=?"
	deletesql2 := "delete from fi_flowmantaskaction_tb where flowtemplateid=? and taskid=?"
	deletesql3 := "delete from fi_flowtaskexecuter_tb where flowtemplateid=? and taskid=?"
	deletesql4 := "delete from fi_flowswitchtaskaction_tb where flowtemplateid=? and taskid=?"
	_, err = o.Raw(ConvertSQL(deletesql1, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql2, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql3, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Raw(ConvertSQL(deletesql4, dbtype), u.Flowtemplateid, u.Taskid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}

//停止流程
func stopflow(o orm.Ormer, modualid string, currentfiid int, currenttiid int, editor string) (err error) {
	dbtype := Getdbtype()
	var row int64 = 0
	sql := "update fi_flow set flowstatus='4' where fiid=? and flowstatus='0'"
	rows, err1 := o.Raw(ConvertSQL(sql, dbtype), currentfiid).Exec()
	if err1 != nil {
		fmt.Println(err1)
		o.Rollback()
		return err1
	}
	row, err = rows.RowsAffected()
	if err == nil {
		if row != 1 {
			o.Rollback()
			return errors.New("data has been update by others!")
		}
	}

	sql = "update fi_task set editor=?,direction='9',taskstatus='1',islast='1' where tiid=? and taskstatus='0' "
	rows, err = o.Raw(ConvertSQL(sql, dbtype), editor, currenttiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	row, err = rows.RowsAffected()
	if err == nil {
		if row != 1 {
			o.Rollback()
			return errors.New("data has been update by others!")
		}
	}
	sql = "SELECT a.tablename,b.flowinstidcol,b.flowstatuscol FROM cmn_modualtemplate_tb a "
	sql = sql + " inner join fxmodual.fi_template_tb b on a.flowtemplateid=b.flowtemplateid where a.modualid=?"
	var tablename, flowinstidcol, flowstatuscol string
	err = o.Raw(ConvertSQL(sql, dbtype), modualid).QueryRow(&tablename, &flowinstidcol, &flowstatuscol)

	sql = "update " + tablename + " set " + flowstatuscol + "='4' where " + flowinstidcol + " =?"
	_, err = o.Raw(ConvertSQL(sql, dbtype), currentfiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	return nil
}

//驳回流程
func backflow(o orm.Ormer, modualid string, currentfiid int, currenttiid int, actionid string, editor string) (err error) {
	dbtype := Getdbtype()
	sql := "select * from  fi_flow where fiid=? "
	ff := FIFLOW{}
	err = o.Raw(sql, currentfiid).QueryRow(&ff)

	sql = "insert into fi_owner(tiid,owner) values(?,?)"
	var tiid int
	tiid, err = getnextval(o, "tiid_sequence")
	_, err = o.Raw(ConvertSQL(sql, dbtype), tiid, ff.Caller).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "select a.taskid,b.flowtemplateid from fi_task a "
	sql = sql + " inner join fi_flow b on a.fiid=b.fiid where a.tiid=?"
	var taskid, flowtemplateid string
	err = o.Raw(ConvertSQL(sql, dbtype), currenttiid).QueryRow(&taskid, &flowtemplateid)

	sql = "select * from fi_flowmantaskaction_tb where flowtemplateid=? and taskid=? and action='return'"
	flowmantaskaction := FLOWMANTASKACTION{}
	err = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, taskid).QueryRow(&flowmantaskaction)
	ft := FITASK{Tiid: tiid, Fiid: currentfiid, Taskid: "1", Taskstarttime: time.Now(), Islast: "1", Taskstatus: "0", Skiptotaskid: flowmantaskaction.Backtask}
	_, err = o.Insert(&ft)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	ftp := FITASKPREVIEW{Fiid: currentfiid, Tiid: tiid, Taskid: flowmantaskaction.Backtask, Prevtiid: currenttiid, Prevtaskid: taskid}
	_, err = o.Insert(&ftp)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	sql = "update fi_flow set state=? where fiid=? "
	_, err = o.Raw(ConvertSQL(sql, dbtype), flowmantaskaction.Status, currentfiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "update fi_task set actionid=?,editor=?, taskfinishtime=?,direction='3',taskstatus='1',islast='1' where tiid=? and taskstatus='0'"
	rows, err1 := o.Raw(ConvertSQL(sql, dbtype), actionid, editor, time.Now(), currenttiid).Exec()
	if err1 != nil {
		fmt.Println(err1)
		o.Rollback()
		return err1
	}
	var row int64 = 0
	row, err = rows.RowsAffected()
	if err == nil {
		if row != 1 {
			o.Rollback()
			return errors.New("data has been updated by others!")
		}
	}

	return nil
}

//流程引擎入口
func Doflow(o orm.Ormer, modualid string, currentfiid int, currenttiid int, actionid string, u MODUALCNTANDMNY, varmap []map[string]string) (fiid int, err error, isend bool) {
	dbtype := Getdbtype()
	//中止或取消流程
	if actionid == "stop" || actionid == "cancel" {
		err = stopflow(o, modualid, currentfiid, currenttiid, u.Submitter)
		if err != nil {

			return 0, err, false
		}
		return currentfiid, nil, true
	}
	//驳回流程
	if actionid == "return" {
		err = backflow(o, modualid, currentfiid, currenttiid, actionid, u.Submitter)
		if err != nil {
			return 0, err, false
		}
		return currentfiid, nil, false
	}
	fv := make([]FIVAR, 0)
	sql := "SELECT b.vary as vid FROM  cmn_modualtemplate_tb a "
	sql = sql + "inner join  fi_templateitem_tb b on a.flowtemplateid=b.flowtemplateid "
	sql = sql + "where a.modualid=?"
	_, err = o.Raw(ConvertSQL(sql, dbtype), modualid).QueryRows(&fv)
	//当前tiid为0时，表示是第一次保存或提交申请
	if currenttiid == 0 {
		isend := false
		fiid, err = insertfirstflow(o, modualid, u, actionid, fv, varmap)
		if err != nil {
			return 0, err, false
		}
		//提交时
		if actionid != "save" {
			sql = "select b.* from cmn_modualtemplate_tb a "
			sql = sql + " inner join fi_flowtask_tb b on a.flowtemplateid=b.flowtemplateid and b.taskid='1' "
			sql = sql + " where a.modualid=? "
			flt := FLOWTASK{}
			err = o.Raw(ConvertSQL(sql, dbtype), modualid).QueryRow(&flt)
			err, isend = insertsecondflow(o, flt, actionid, currenttiid, u.Submitter, fiid, modualid)
			if err != nil {
				return 0, err, false
			}
		}
		return fiid, nil, isend
	} else { //保存、审批
		isend := false
		sql = "select c.flowtemplateid,c.taskid,c.taskname,c.tasktype,c.supportskip,c.sendmessage,"
		sql = sql + " c.concurrent,c.samepersontask,c.nopersontask,a.skiptotaskid from  fi_task a "
		sql = sql + " inner join  fi_flow b on a.fiid=b.fiid "
		sql = sql + " inner join  fi_flowtask_tb c on a.taskid=c.taskid and b.flowtemplateid=c.flowtemplateid "
		sql = sql + " where a.tiid=? "
		flt := FLOWTASK{}
		err = o.Raw(ConvertSQL(sql, dbtype), currenttiid).QueryRow(&flt)
		err = updateflow(o, currentfiid, currenttiid, actionid, fv, u, flt)
		if err != nil {
			return 0, err, false
		}
		//提交或审批同意时，则插入第二个任务
		if actionid != "save" {
			var ncount int = 0
			if flt.Concurrent == "1" {
				//判断是否还有未完成的会签
				sql = "select count(1) as ncount from fi_countersign where tiid=? and taskstatus='0'"
				err = o.Raw(ConvertSQL(sql, dbtype), currenttiid).QueryRow(&ncount)
			}
			if ncount > 0 {
				return currentfiid, nil, false
			}
			err, isend = insertsecondflow(o, flt, actionid, currenttiid, u.Submitter, currentfiid, modualid)
			if err != nil {
				return 0, err, false
			}
		}

		return currentfiid, err, isend
	}
	return -1, nil, false
}

//第一次保存或提交申请
func insertfirstflow(o orm.Ormer, modualid string, u MODUALCNTANDMNY, actionid string, fv []FIVAR, varmap []map[string]string) (fiid int, err error) {
	dbtype := Getdbtype()
	var flowintid int
	var tiid int
	flowintid, err = getnextval(o, "fiid_sequence")
	//把所有流程变量插入fi_var表
	sql := "insert into fi_var(fiid,vid) values(?,?)"
	for _, fvc := range fv {
		_, err = o.Raw(ConvertSQL(sql, dbtype), flowintid, fvc.Vid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return 0, err
		}
	}
	//根据申请人获得所在机构级别的变量并更新到fi_var表中
	floworgvarys, _ := GetFLOWORGVARYBYUSERID(u.Submitter)
	sql = "update fi_var set vvalue=? where fiid=? and vid=?"
	for _, floworgvary := range floworgvarys {
		_, err = o.Raw(ConvertSQL(sql, dbtype), floworgvary.Vvalue, flowintid, floworgvary.Vid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return 0, err
		}
	}
	//OA引擎负责更新申请人的orglevel到fi_var表中
	//如果流程模板中的流程变量为orglevel时，OA引擎负责更新orglevel。否则需要在业务模块中负责更新。
	var orglevel string
	sql = "select orglevel from cmn_org_tb a inner join cmn_user_tb b on a.orgid=b.orgid where b.userid=?"
	_ = o.Raw(ConvertSQL(sql, dbtype), u.Submitter).QueryRow(&orglevel)
	sql = "update fi_var set vvalue=? where fiid=? and vid='orglevel'"
	_, err = o.Raw(ConvertSQL(sql, dbtype), orglevel, flowintid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return 0, err
	}
	modualtemplate := CMN_MODUALTEMPLATE_TB{Modualid: modualid}
	err = o.Read(&modualtemplate)
	if err != nil {
		return 0, errors.New(modualid + " not bind template!")
	}
	sql = "select * from fi_flowmantaskaction_tb where flowtemplateid=? and taskid='1' and action=?"
	flowmantaskaction := FLOWMANTASKACTION{}
	err = o.Raw(ConvertSQL(sql, dbtype), modualtemplate.Flowtemplateid, actionid).QueryRow(&flowmantaskaction)
	if err != nil {
		return 0, err
	}
	var flowstatus string
	//是否需要第一个节点的限制条件？
	flowstatus = "0"

	//插入fi_flow表
	ff := FIFLOW{
		Fiid:           flowintid,
		Modualid:       modualtemplate.Modualid,
		Flowtemplateid: modualtemplate.Flowtemplateid,
		Flowcontent:    u.Content,
		Caller:         u.Submitter,
		Flowstarttime:  time.Now(),
		Flowstatus:     flowstatus,
		State:          flowmantaskaction.Status,
	}
	_, err = o.Insert(&ff)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return 0, err
	}
	//获得下个任务的tiid
	tiid, err = getnextval(o, "tiid_sequence")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var taskstatus string
	var direction string
	var editor string
	if actionid == "save" {
		taskstatus = "0"
		direction = ""
		editor = ""
	}
	if actionid != "save" {
		taskstatus = "1"
		direction = "1"
		editor = u.Submitter
	}
	//插入fi_task表
	ft := FITASK{Editor: editor, Direction: direction, Actionid: actionid, Tiid: tiid, Fiid: flowintid, Taskid: "1", Taskstarttime: time.Now(), Islast: "1", Taskstatus: taskstatus}
	_, err = o.Insert(&ft)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return 0, err
	}
	//插入fi_owner表
	sql = "insert into fi_owner(tiid,owner) values(?,?)"
	_, err = o.Raw(ConvertSQL(sql, dbtype), tiid, u.Submitter).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return 0, err
	}
	//if actionid == "save" || actionid == "submit" {
	sql = "update fi_var set vvalue=? where fiid=? and vid=?"
	sql = ConvertSQL(sql, Getdbtype())
	if varmap != nil {
		for _, varmap1 := range varmap {
			for key, value := range varmap1 {
				_, err = o.Raw(sql, value, flowintid, key).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return -1, err
				}
			}
		}
	}

	//}
	return flowintid, nil
}

//插入下一个任务。提交或审批时。
func insertnexttask(modualid string, currenttiid int, o orm.Ormer, submitter string, caller string, fiid int, flowtemplateid string, nexttaskid string) (tiid int, err error, isend bool) {
	dbtype := Getdbtype()
	//下一个任务ID为999，则结束流程。
	if nexttaskid == "999" {
		fmt.Println("end task")
		err = endflow(o, submitter, modualid, fiid, currenttiid)
		if err != nil {
			fmt.Println("end task err")
			return 0, err, false
		} else {
			return -1, nil, true
		}
	}

	ft := FLOWTASK{}
	sql := "select * from fi_flowtask_tb where flowtemplateid=? and taskid=? "
	err = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid).QueryRow(&ft)

	if ft.Tasktype == "man" { //人工任务
		//获得任务的执行人
		sql = "select * from fi_flowtaskexecuter_tb where flowtemplateid=? and taskid=?  "
		flte := make([]FLOWTASKEXECUTER, 0)
		_, err = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid).QueryRows(&flte)

		editors := geteditors(o, caller, flte)
		fmt.Println("editors:")
		fmt.Println(editors)
		if len(editors) > 0 { //能找到执行人信息
			//同人转向
			if len(editors) == 1 && submitter == editors[0] && ft.Samepersontask != "" {
				return insertnexttask(modualid, currenttiid, o, submitter, caller, fiid, flowtemplateid, ft.Samepersontask)
			}
			//获得下一个任务的TIID
			tiid, err = getnextval(o, "tiid_sequence")
			//插入fi_owner表
			sql := "insert into fi_owner(tiid,owner) values(?,?)"
			for _, editor := range editors {
				//下个节点再次包含该审批者时，过滤掉
				//				if submitter == editor {
				//					continue
				//
				//代理处理				}
				insertcountersignsql := "insert into fi_countersign(tiid,userid,fiid,taskid,taskstatus) values(?,?,?,?,?)"

				dsql := "select agent from fi_agent_tb where userid=? and startdate<=? and enddate>=?"
				var agents []string
				_, _ = o.Raw(ConvertSQL(dsql, dbtype), editor, time.Now(), time.Now()).QueryRows(&agents)
				if len(agents) > 0 {
					for _, agent := range agents {
						var ncount int = 0
						existssql := "select count(1) as ncount from fi_owner where tiid=? and owner=?"
						err = o.Raw(ConvertSQL(existssql, dbtype), tiid, agent).QueryRow(&ncount)
						if ncount == 0 {
							_, err = o.Raw(ConvertSQL(sql, dbtype), tiid, agent).Exec()
							if err != nil {
								fmt.Println(err)
								o.Rollback()
								return 0, err, false
							}
						}

						//会签处理
						ncount = 0
						existssql = "select count(1) as ncount from fi_countersign where tiid=? and userid=?"
						err = o.Raw(ConvertSQL(existssql, dbtype), tiid, agent).QueryRow(&ncount)
						if ncount == 0 {
							_, err = o.Raw(ConvertSQL(insertcountersignsql, dbtype), tiid, agent, fiid, nexttaskid, '0').Exec()
							if err != nil {
								fmt.Println(err)
								o.Rollback()
								return 0, err, false
							}
						}

					}

				} else {
					var ncount int = 0
					existssql := "select count(1) as ncount from fi_owner where tiid=? and owner=?"
					err = o.Raw(ConvertSQL(existssql, dbtype), tiid, editor).QueryRow(&ncount)
					if ncount == 0 {
						_, err = o.Raw(ConvertSQL(sql, dbtype), tiid, editor).Exec()
						if err != nil {
							fmt.Println(err)
							o.Rollback()
							return 0, err, false
						}
					}
					//会签处理
					ncount = 0
					existssql = "select count(1) as ncount from fi_countersign where tiid=? and userid=?"
					err = o.Raw(ConvertSQL(existssql, dbtype), tiid, editor).QueryRow(&ncount)
					if ncount == 0 {
						_, err = o.Raw(ConvertSQL(insertcountersignsql, dbtype), tiid, editor, fiid, nexttaskid, '0').Exec()
						if err != nil {
							fmt.Println(err)
							o.Rollback()
							return 0, err, false
						}
					}
				}
			}
			//插入fi_task表
			fitask := FITASK{Tiid: tiid, Fiid: fiid, Taskid: nexttaskid, Taskstarttime: time.Now(), Islast: "1", Taskstatus: "0"}
			_, err = o.Insert(&fitask)

			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return 0, err, false
			}

			return tiid, nil, false
		} else { //找不到执行人信息时
			fmt.Println(flowtemplateid + "," + nexttaskid + " not founded editors")
			//无人转向
			if ft.Nopersontask != "" {
				fmt.Println("find nopersontask:" + ft.Nopersontask)
				return insertnexttask(modualid, currenttiid, o, submitter, caller, fiid, flowtemplateid, ft.Nopersontask)
			} else { //找不到执行人信息，同时无人转向未配置，则报错“流程配置错误”
				fmt.Println(flowtemplateid + "," + nexttaskid + " Nopersontask is also null and flow define is wrong")
				o.Rollback()
				return 0, errors.New(flowtemplateid + "," + nexttaskid + " flow define is wrong"), false
			}
		}
	} else { //分支时
		sql = "select * from fi_flowswitchtaskaction_tb where flowtemplateid=? and taskid=? "
		flst := make([]FLOWSWITCHTASKACTION, 0)
		_, err = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid).QueryRows(&flst)
		//根据流程定义节点的变量，从fi_var表中获得该变量的值
		sql = "select * from fi_var where fiid=? and vid=?"
		fivar := FIVAR{}
		err = o.Raw(ConvertSQL(sql, dbtype), fiid, flst[0].Conditions).QueryRow(&fivar)
		var findswitchaction bool = false
		for _, flst1 := range flst {

			functions := flst1.Functions
			valuee := flst1.Valuee
			nexttask := flst1.Nexttask
			findswitchaction, err := comparefunction(fivar.Vvalue, functions, valuee)
			if err != nil {
				return 0, err, false
			}
			if findswitchaction {
				return insertnexttask(modualid, currenttiid, o, submitter, caller, fiid, flowtemplateid, nexttask)
				break
			}
		}
		if !findswitchaction { //找不到分支任务的匹配项，则报错“流程配置错误”
			fmt.Println(flowtemplateid + "," + nexttaskid + " flow define is wrong ")
			o.Rollback()
			return 0, errors.New(flowtemplateid + "," + nexttaskid + " flow define is wrong"), false
		}
	}
	return -1, nil, false

}

//结束流程
func endflow(o orm.Ormer, submitter string, modualid string, fiid int, currenttiid int) (err error) {
	dbtype := Getdbtype()
	sql := "update fi_flow set flowfinishtime=?,flowstatus='1' where fiid=? and flowstatus='0'"

	rows, err1 := o.Raw(ConvertSQL(sql, dbtype), time.Now(), fiid).Exec()
	if err1 != nil {
		fmt.Println(err1)
		o.Rollback()
		return err1
	}
	var row int64 = 0
	row, err = rows.RowsAffected()
	if err == nil {

		if row != 1 {
			o.Rollback()
			return errors.New("data has been updated by others!")
		}
	}
	sql = "update fi_task set taskfinishtime=?,editor=?,direction='1',taskstatus='1' where tiid=? and taskstatus='0'"
	rows, err = o.Raw(ConvertSQL(sql, dbtype), time.Now(), submitter, currenttiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	if err == nil {

		if row != 1 {
			o.Rollback()
			return errors.New("data has been updated by others!")
		}
	}
	//根据流程模板定义，更新业务表的流程状态
	sql = "SELECT a.tablename,b.flowinstidcol,b.flowstatuscol FROM cmn_modualtemplate_tb a "
	sql = sql + " inner join  fi_template_tb b on a.flowtemplateid=b.flowtemplateid where a.modualid=?"
	var tablename, flowinstidcol, flowstatuscol string
	err = o.Raw(ConvertSQL(sql, dbtype), modualid).QueryRow(&tablename, &flowinstidcol, &flowstatuscol)

	sql = "update " + tablename + " set " + flowstatuscol + "='1' where " + flowinstidcol + " =?"
	_, err = o.Raw(ConvertSQL(sql, dbtype), fiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	return nil

}

//插入第二个流程。提交或审批时。
func insertsecondflow(o orm.Ormer, flt FLOWTASK, actionid string, currenttiid int, submitter string, fiid int, modualid string) (err error, isend bool) {
	dbtype := Getdbtype()
	var nexttaskid string = ""
	var sql string = ""
	//如果fi_task表中的Skiptotaskid不为空，则为驳回后再提交。提交后需要直接跳到驳回节点上。
	if flt.Skiptotaskid != "" {
		nexttaskid = flt.Skiptotaskid
	} else {
		sql = "select * from fi_flowmantaskaction_tb where flowtemplateid=? and taskid=? and action=?  "
		flmt := FLOWMANTASKACTION{}
		err = o.Raw(ConvertSQL(sql, dbtype), flt.Flowtemplateid, flt.Taskid, actionid).QueryRow(&flmt)
		nexttaskid = flmt.Nexttask
		if nexttaskid == "" {
			return errors.New(modualid + "|" + flt.Taskid + " task not find nexttask ,flow define is wrong. "), false
		}
	}
	//o orm.Ormer, caller string, fiid int, flowtemplateid string, taskid string
	var caller string
	if fiid == 0 {
		caller = submitter
	} else {
		sql = "select caller from fi_flow where fiid=?"
		err = o.Raw(ConvertSQL(sql, dbtype), fiid).QueryRow(&caller)
	}
	if actionid != "save" && (currenttiid == 0 || Gettaskidbytiid(o, currenttiid) == "1") {
		//级别高的人员提交时跳过级别低的人员审批，但用户、用户组、角色、自动分支不跳过
		var orglevel string
		sql = "select orglevel from cmn_org_tb a inner join cmn_user_tb b on a.orgid=b.orgid where b.userid=?"
		_ = o.Raw(ConvertSQL(sql, dbtype), caller).QueryRow(&orglevel)
		sql = "select nexttask from fi_flowmantaskaction_tb where flowtemplateid=? and taskid='1' and action=?"
		_ = o.Raw(ConvertSQL(sql, dbtype), flt.Flowtemplateid, actionid).QueryRow(&nexttaskid)
		nexttaskid, err = Getnexttaskbycaller(orglevel, nexttaskid, o, caller, flt.Flowtemplateid, actionid)
		if err != nil {
			return err, false
		}
		if nexttaskid == "999" {
			err = endflow(o, submitter, modualid, fiid, currenttiid)
			if err != nil {
				fmt.Println("end task err")
				return err, false
			} else {
				return nil, true
			}
		}
	}
	//支持分流
	//此节点同时提交给多个下节点
	nexttaskids := strings.Split(nexttaskid, ",")
	istotalend := false
	for _, nexttaskid2 := range nexttaskids {
		var tiid int
		isend := false
		tiid, err, isend = insertnexttask(modualid, currenttiid, o, submitter, caller, fiid, flt.Flowtemplateid, nexttaskid2)
		if err != nil {
			fmt.Println(err)
			return err, false
		} else {
			if isend {
				istotalend = true
			}
		}
		if tiid != -1 {
			ftp := FITASKPREVIEW{Fiid: fiid, Tiid: tiid, Taskid: nexttaskid2, Prevtiid: currenttiid, Prevtaskid: flt.Taskid}
			_, err = o.Insert(&ftp)
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err, false
			}
		}
	}

	return nil, istotalend
}

//更新流程。保存后提交或者审批时。
//第一个节点保存或提交，需要重新更新流程相关变量
func updateflow(o orm.Ormer, currentfiid int, currenttiid int, actionid string, fv []FIVAR, u MODUALCNTANDMNY, flt FLOWTASK) (err error) {
	var sql string
	dbtype := Getdbtype()
	//支持第一次申请和驳回后再次申请
	if currenttiid == 0 || Gettaskidbytiid(o, currenttiid) == "1" {
		var caller string
		sql = "select caller from fi_flow where fiid=?"
		err = o.Raw(ConvertSQL(sql, dbtype), currentfiid).QueryRow(&caller)
		floworgvarys, _ := GetFLOWORGVARYBYUSERID(caller)
		sql = "update fi_var set vvalue=? where fiid=? and vid=?"
		for _, floworgvary := range floworgvarys {
			_, err = o.Raw(ConvertSQL(sql, dbtype), floworgvary.Vvalue, currentfiid, floworgvary.Vid).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	}
	//会签处理
	//如果有未完成的会签则该任务节点未完成
	var ncount int = 0
	if flt.Concurrent == "1" {
		//更新该用户的会签任务为完成状态
		sql = "update fi_countersign set taskstatus='1' where tiid=? and userid=?"
		_, err = o.Raw(ConvertSQL(sql, dbtype), currenttiid, u.Submitter).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		//删除该用户的待办
		sql = "delete from  fi_owner where tiid=? and owner=?"
		_, err = o.Raw(ConvertSQL(sql, dbtype), currenttiid, u.Submitter).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		//判断是否还有未完成的会签
		sql = "select count(1) as ncount from fi_countersign where tiid=? and taskstatus='0'"
		err = o.Raw(ConvertSQL(sql, dbtype), currentfiid).QueryRow(&ncount)
	}
	//如果有未完成的会签则返回
	if ncount > 0 {
		return nil
	}
	sql = "update fi_flow set state=?,flowcontent=? where fiid=?"
	_, err = o.Raw(ConvertSQL(sql, dbtype), flt.Taskname, u.Content, currentfiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	if actionid != "save" {
		sql = "update fi_task set taskfinishtime=?,editor=?,actionid=?,opinion=?,direction='1',taskstatus='1' where tiid=? and taskstatus='0'"

		rows, err1 := o.Raw(ConvertSQL(sql, dbtype), time.Now(), u.Submitter, actionid, u.Opinion, currenttiid).Exec()
		if err != nil {
			fmt.Println(err1)
			o.Rollback()
			return err1
		}
		var row int64 = 0
		row, err = rows.RowsAffected()
		if err == nil {
			if row != 1 {
				o.Rollback()
				return errors.New("data has been updated by others!")
			}
		}
	}
	return nil
}

//获得下一个seq
func getnextval(o orm.Ormer, sequencename string) (returnvalue int, err error) {
	dbtype := Getdbtype()
	//o := orm.NewOrm()
	var nextvalue int
	switch dbtype {
	case "mysql":
		sql := "select nextval('" + sequencename + "')"
		err = o.Raw(sql).QueryRow(&nextvalue)
		if err != nil {
			fmt.Println(err)
			return -1, err
		}
	default:
		var currentValue, increment int
		sql := "select currentValue,increment from sequence where seqname=?"
		err = o.Raw(ConvertSQL(sql, dbtype), sequencename).QueryRow(&currentValue, &increment)
		if err != nil {
			fmt.Println(err)
			return -1, err
		}
		nextvalue = currentValue + increment
		sql = "update sequence set currentValue=? where seqname=?"
		_, err = o.Raw(ConvertSQL(sql, dbtype), nextvalue, sequencename).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return -1, err
		}
	}

	return nextvalue, nil
}

// 变量functions :>= 0
// <= 1
// > 2
// < 3
// = 4
func comparefunction(orgivalue string, functions string, valuee string) (isequal bool, err error) {

	forgivalue, err1 := strconv.ParseFloat(orgivalue, 32/64)
	if err1 != nil {
		return false, err1
	}
	fvaluee, err2 := strconv.ParseFloat(valuee, 32/64)
	if err2 != nil {
		return false, err2
	}
	if functions == "0" {
		return forgivalue >= fvaluee, nil
	}
	if functions == "1" {
		return forgivalue <= fvaluee, nil
	}
	if functions == "2" {
		return forgivalue > fvaluee, nil
	}
	if functions == "3" {
		return forgivalue < fvaluee, nil
	}
	if functions == "4" {
		return forgivalue == fvaluee, nil
	}
	return false, nil
}

//taskexecuter
//用户 1
//角色 2
//用户组 3
//任务执行人 4
//发起人所属指定级别机构主管/副主管 5
//发起人所属指定级别机构主管 6
//发起人所属指定级别机构副主管 7
//发起人直属指定级别机构主管/副主管 8
//发起人直属指定级别机构主管 9
//发起人直属指定级别机构副主管 10
//操作人所属指定级别机构主管/副主管 11
//操作人所属指定级别机构主管 12
//操作人所属指定级别机构副主管 13
//操作人直属指定级别机构副主管/副主管 14
//操作人直属指定级别机构主管 15
//操作人直属指定级别机构副主管 16
//目前实现了最常用的 1 2 3 5 6 7，其它todo
func geteditors(o orm.Ormer, caller string, flte []FLOWTASKEXECUTER) []string {
	editors := make([]string, 0)
	for _, flte1 := range flte {
		taskexecuter := flte1.Taskexecuter
		expression := flte1.Expression
		if taskexecuter == "1" {
			editors = append(editors, expression)
		}
		if taskexecuter == "3" {
			ug, _ := GetCMN_USERGROUP_TB(CMN_GROUP_TB{Groupid: expression})
			for _, ug1 := range ug {
				editors = append(editors, ug1.Userid)
			}
		}
		if taskexecuter == "2" {
			ug, _ := GetUSERSBYROLE(o, CMN_USERROLE_TB{Roleid: expression})
			for _, ug1 := range ug {
				editors = append(editors, ug1.Userid)
			}
		}
		if taskexecuter == "5" || taskexecuter == "6" || taskexecuter == "7" {
			ugl := GetUSERBYORGLEVEL(o, caller, taskexecuter, expression)
			for _, ugl1 := range ugl {
				editors = append(editors, ugl1.Userid)
			}
			fmt.Println("geteditors:")
			fmt.Println(editors)
		}

	}
	return editors
}

//根据角色ID获得用户角色
func GetUSERSBYROLE(o orm.Ormer, u CMN_USERROLE_TB) (admins []CMN_USERROLE_TB, err error) {
	admins = make([]CMN_USERROLE_TB, 0)
	//o := orm.NewOrm()
	sql := "select userid,roleid from cmn_userrole_tb where roleid=? "

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Roleid).QueryRows(&admins)

	return admins, err
}

//获得用户的机构级别
func GetUSERBYORGLEVEL(o orm.Ormer, caller string, taskexecuter string, expressionorglevel string) []CMN_ORGLEADER_TB {

	cmnorg := CMN_ORG_TB{}
	//o := orm.NewOrm()
	sql := "select a.orgid,orglevel from cmn_user_tb a inner join cmn_org_tb b on a.orgid=b.orgid where a.userid=?"
	_ = o.Raw(ConvertSQL(sql, Getdbtype()), caller).QueryRow(&cmnorg)

	users, err := Getorgleaders(o, cmnorg.Orgid, taskexecuter, expressionorglevel)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("users==>")
	fmt.Println(users)
	return users
}

//获得机构主管
func Getorgleaders(o orm.Ormer, orgid string, taskexecuter string, expressionorglevel string) (leaders []CMN_ORGLEADER_TB, err error) {
	dbtype := Getdbtype()
	cmnorg := CMN_ORG_TB{}
	//o := orm.NewOrm()
	//lds := LEADERS{}
	var userorglevelint int
	var expressionorglevelint int
	expressionorglevelint, err = strconv.Atoi(expressionorglevel)
	sql := "select orgid,parentid,orglevel from cmn_org_tb where orgid=?"
	err = o.Raw(ConvertSQL(sql, dbtype), orgid).QueryRow(&cmnorg)
	userorglevelint, err = strconv.Atoi(cmnorg.Orglevel)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("userorglevelint==>")
	fmt.Println(userorglevelint)
	fmt.Println("expressionorglevelint==>")
	fmt.Println(expressionorglevelint)
	if userorglevelint > expressionorglevelint {
		//此处必须有return，否则返回值为空。坑
		return Getorgleaders(o, cmnorg.Parentid, taskexecuter, expressionorglevel)

	}
	if userorglevelint < expressionorglevelint {
		return nil, nil
	}
	if userorglevelint == expressionorglevelint {
		sql := "select * from cmn_orgleader_tb where orgid=?"
		if taskexecuter == "6" {
			sql = sql + " and leadertype='0'"
		}
		if taskexecuter == "7" {
			sql = sql + " and leadertype='1'"
		}
		orglds := make([]CMN_ORGLEADER_TB, 0)
		_, err = o.Raw(ConvertSQL(sql, dbtype), orgid).QueryRows(&orglds)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Print("orglds==>")
		fmt.Println(orglds)
		return orglds, nil
	}
	return nil, nil
}

//获得任务信息
func GetTaskinfo(u TASKINFO) (admins []TASKINFO, err error) {
	dbtype := Getdbtype()
	admins = make([]TASKINFO, 0)
	o := orm.NewOrm()

	sql := "SELECT  action FROM  fi_flowmantaskaction_tb a "
	sql = sql + " inner join  cmn_modualtemplate_tb b on b.flowtemplateid=a.flowtemplateid "
	sql = sql + " where b.modualid=? and a.taskid='1' "

	sql2 := "select  action from   fi_task a "
	sql2 = sql2 + " inner join   fi_flow b on a.fiid=b.fiid "
	sql2 = sql2 + " inner join  cmn_modualtemplate_tb c on c.flowtemplateid=b.flowtemplateid and c.modualid=? "
	sql2 = sql2 + " inner join  fi_flowmantaskaction_tb d on d.flowtemplateid=b.flowtemplateid and d.taskid=a.taskid "
	sql2 = sql2 + " where a.tiid=?"
	fmt.Println(u)
	if u.Currentfiid == 0 {
		_, err = o.Raw(ConvertSQL(sql, dbtype), u.Modualid).QueryRows(&admins)
	} else {
		_, err = o.Raw(ConvertSQL(sql2, dbtype), u.Modualid, u.Currenttiid).QueryRows(&admins)
	}

	return admins, err
}

//获得待办任务
func Gettodotask(u FIFLOW) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,b.taskid,a.flowstarttime,a.flowstatus,c.tiid,a.flowtemplateid, "
	sql = sql + "  e.url,f.flowstatusname,g.taskname from fi_flow a "
	sql = sql + "  inner join fi_task b on a.fiid=b.fiid and taskstatus='0' "
	sql = sql + "  inner join fi_owner c on b.tiid=c.tiid and c.owner=? "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb g on a.flowtemplateid=g.flowtemplateid and b.taskid=g.taskid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where  a.flowstatus='0' "

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Caller).QueryRows(&admins)

	return admins, err
}

//获得待办任务列表数据
func Gettodotasklist(u FITASK) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT b.taskid,b.taskstarttime,  "
	sql = sql + "  c.owner as editor,d.taskname   from fi_flow a "
	sql = sql + "  inner join fi_task b on a.fiid=b.fiid  "
	sql = sql + "  inner join fi_owner c on b.tiid=c.tiid "
	sql = sql + "  inner join fi_flowtask_tb d on a.flowtemplateid=d.flowtemplateid and b.taskid=d.taskid "
	sql = sql + "  where a.fiid=? and a.flowstatus='0' and b.taskstatus='0' "

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Fiid).QueryRows(&admins)

	return admins, err
}

//获得已办任务
func Getdonetask(u FIFLOW) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,b.taskid,a.flowstarttime,a.flowfinishtime,a.flowstatus,b.tiid,a.flowtemplateid, "
	sql = sql + "  e.url,f.flowstatusname,g.taskname  from fi_flow a "
	sql = sql + "  inner join fi_task b on a.fiid=b.fiid and taskstatus='1' and b.editor='" + u.Caller + "' "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb g on a.flowtemplateid=g.flowtemplateid and b.taskid=g.taskid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where a.caller<>b.editor and 1=1 "
	if u.Fiid != 0 {
		sql = sql + " and a.fiid=" + strconv.Itoa(u.Fiid)

	}
	//	if u.Flowstarttime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowstarttime>=" + u.Flowstarttime.Format("2006-01-02")
	//	}
	//	if u.Flowfinishtime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowfinishtime<=" + u.Flowfinishtime.Format("2006-01-02")
	//	}
	if u.Flowtemplateid != "" {
		sql = sql + " and a.flowtemplateid='" + u.Flowtemplateid + "'"
	}
	if u.Flowcontent != "" {
		sql = sql + " and a.flowcontent like '%" + u.Flowcontent + "%' "
	}
	if u.Flowstatus != "" {
		sql = sql + " and a.flowstatus = '" + u.Flowstatus + "' "
	}
	_, err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRows(&admins)

	return admins, err
}

//获得已办任务列表数据
func Getdonetasklist(u FITASK) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT a.fiid,a.caller,a.flowcontent,b.taskid,a.flowstarttime,a.flowfinishtime,a.flowstatus,b.tiid,a.flowtemplateid, "
	sql = sql + "  e.url,b.editor,b.opinion,b.taskstarttime,b.taskfinishtime,b.taskstatus,f.flowstatusname,g.taskname  from fi_flow a "
	sql = sql + "  inner join fi_task b on a.fiid=b.fiid and taskstatus='1' "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb g on a.flowtemplateid=g.flowtemplateid and b.taskid=g.taskid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where 1=1 and a.fiid=? order by b.tiid"

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Fiid).QueryRows(&admins)

	return admins, err
}

//获得流程化监控数据条数
func Getflowmonitorcount(u FIFLOW) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total  "
	sql = sql + "     from fi_flow a "
	sql = sql + "  inner join (select max(tiid) as tiid,fiid from fi_task group by fiid) b on a.fiid=b.fiid "
	sql = sql + "  inner join fi_task c on b.tiid=c.tiid and b.fiid=c.fiid "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb g on g.flowtemplateid=a.flowtemplateid and g.taskid=c.taskid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where 1=1 "
	if u.Fiid != 0 {
		sql = sql + " and a.fiid=" + strconv.Itoa(u.Fiid)

	}
	//	if u.Flowstarttime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowstarttime>=" + u.Flowstarttime.Format("2006-01-02")
	//	}
	//	if u.Flowfinishtime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowfinishtime<=" + u.Flowfinishtime.Format("2006-01-02")
	//	}
	if u.Flowtemplateid != "" {
		sql = sql + " and a.flowtemplateid='" + u.Flowtemplateid + "'"
	}
	if u.Flowcontent != "" {
		sql = sql + " and a.flowcontent like '%" + u.Flowcontent + "%' "
	}
	if u.Flowstatus != "" {
		sql = sql + " and a.flowstatus in(" + u.Flowstatus + ") "
	}
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}

//获得流程监控分页数据
func Getflowmonitorbypageindex(u FIFLOWPAGEINDEX) (admins []TODOTASKLIST, err error) {
	dbtype := Getdbtype()
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,c.taskid,a.flowstarttime,a.flowfinishtime,a.flowstatus,b.tiid,a.flowtemplateid, "
	sql = sql + "  e.url,f.flowstatusname,g.taskname  from fi_flow a "
	sql = sql + "  inner join (select max(tiid) as tiid,fiid from fi_task group by fiid) b on a.fiid=b.fiid "
	sql = sql + "  inner join fi_task c on b.tiid=c.tiid and b.fiid=c.fiid "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb g on g.flowtemplateid=a.flowtemplateid and g.taskid=c.taskid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where 1=1 "
	if u.Fiid != 0 {
		sql = sql + " and a.fiid=" + strconv.Itoa(u.Fiid)

	}
	//	if u.Flowstarttime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowstarttime>=" + u.Flowstarttime.Format("2006-01-02")
	//	}
	//	if u.Flowfinishtime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowfinishtime<=" + u.Flowfinishtime.Format("2006-01-02")
	//	}
	if u.Flowtemplateid != "" {
		sql = sql + " and a.flowtemplateid='" + u.Flowtemplateid + "'"
	}
	if u.Flowcontent != "" {
		sql = sql + " and a.flowcontent like '%" + u.Flowcontent + "%' "
	}
	if u.Flowstatus != "" {
		sql = sql + " and a.flowstatus in(" + u.Flowstatus + ") "
	}
	var limitstr string = " limit "
	if dbtype == "postgres" {
		limitstr = limitstr + strconv.Itoa(u.Pagesize) + " offset " + strconv.Itoa((u.Pageindex-1)*u.Pagesize)

	} else if dbtype == "mysql" {
		limitstr = limitstr + strconv.Itoa((u.Pageindex-1)*u.Pagesize) + "," + strconv.Itoa(u.Pagesize)

	} else {
		limitstr = limitstr + strconv.Itoa((u.Pageindex-1)*u.Pagesize) + "," + strconv.Itoa(u.Pagesize)

	}
	sql = sql + limitstr
	_, err = o.Raw(ConvertSQL(sql, dbtype)).QueryRows(&admins)

	return admins, err
}

//获得我的流程
func Getmyflow(u FIFLOW) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()

	sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,c.taskid,a.flowstarttime,a.flowfinishtime,a.flowstatus,b.tiid,a.flowtemplateid, "
	sql = sql + "  e.url,f.supportskip,g.flowstatusname,f.taskname  from fi_flow a "
	sql = sql + "  inner join (select max(tiid) as tiid,fiid from fi_task group by fiid) b on a.fiid=b.fiid "
	sql = sql + "  inner join fi_task c on a.fiid=c.fiid and c.tiid=b.tiid "
	sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
	sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid "
	sql = sql + "  inner join fi_flowstatus g on a.flowstatus=g.flowstatus "
	sql = sql + "  inner join fi_flowtask_tb f on f.flowtemplateid=a.flowtemplateid and f.taskid=c.taskid where a.caller='" + u.Caller + "' "

	if u.Fiid != 0 {
		sql = sql + " and a.fiid=" + strconv.Itoa(u.Fiid)

	}
	if u.Flowstarttime.Format("2006-01-02") != "9999-01-01" && u.Flowstarttime.Format("2006-01-02") != "0001-01-01" {
		sql = sql + " and DATE_FORMAT(a.flowstarttime,'%Y-%m-%d')>='" + u.Flowstarttime.Format("2006-01-02") + "'"
	}
	if u.Flowfinishtime.Format("2006-01-02") != "9999-01-01" && u.Flowfinishtime.Format("2006-01-02") != "0001-01-01" {
		sql = sql + " and DATE_FORMAT(a.flowfinishtime,'%Y-%m-%d')<='" + u.Flowfinishtime.Format("2006-01-02") + "'"
	}
	//	if u.Flowstarttime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowstarttime>=" + u.Flowstarttime.Format("2006-01-02")
	//	}
	//	if u.Flowfinishtime.Format("2006-01-02") != "" {
	//		sql = sql + " and a.flowfinishtime<=" + u.Flowfinishtime.Format("2006-01-02")
	//	}
	if u.Flowtemplateid != "" {
		sql = sql + " and a.flowtemplateid='" + u.Flowtemplateid + "'"
	}
	if u.Flowcontent != "" {
		sql = sql + " and a.flowcontent like '%" + u.Flowcontent + "%' "
	}
	if u.Flowstatus != "" {
		sql = sql + " and a.flowstatus in(" + u.Flowstatus + ") "
	}
	_, err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRows(&admins)

	return admins, err
}

//流程取消
func Cancelflow(u FIFLOW) (err error) {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err = o.Begin()

	sql := "update fi_flow set flowstatus='4' where fiid=? and (flowstatus='0' or flowstatus='1')"
	rows, err1 := o.Raw(ConvertSQL(sql, dbtype), u.Fiid).Exec()
	if err != nil {
		fmt.Println(err1)
		o.Rollback()
		return err1
	}
	var row int64 = 0
	row, err = rows.RowsAffected()
	if err == nil {
		if row != 1 {
			o.Rollback()
			return errors.New("data has been updated by others!")
		}
	}
	var tablename, flowinstidcol, flowstatuscol string

	sql = "SELECT a.tablename,b.flowinstidcol,b.flowstatuscol FROM cmn_modualtemplate_tb a "
	sql = sql + " inner join  fi_template_tb b on a.flowtemplateid=b.flowtemplateid where a.flowtemplateid=? and a.modualid=? "

	err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Flowtemplateid, u.Modualid).QueryRow(&tablename, &flowinstidcol, &flowstatuscol)

	sql = "update " + tablename + " set " + flowstatuscol + "='4' where " + flowinstidcol + " =?"
	_, err = o.Raw(ConvertSQL(sql, dbtype), u.Fiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "update fi_task set taskstatus='1',direction='9' where fiid =? and taskstatus='0'"

	rows, err = o.Raw(ConvertSQL(sql, dbtype), u.Fiid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	row, err = rows.RowsAffected()
	if err == nil {
		if row != 1 {
			o.Rollback()
			return errors.New("data has been updated by others!")
		}
	}

	err = o.Commit()
	return err
}

//跳过任务节点
//支持跳转的节点必须有actionid="next"配置，否则报错
//todo： 有唯一的forward任务也可
func Skiptask(u TODOTASKLIST) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	sql := "select c.flowtemplateid,c.taskid,c.taskname,c.tasktype,c.supportskip,c.sendmessage,"
	sql = sql + " c.concurrent,c.samepersontask,c.nopersontask from  fi_task a "
	sql = sql + " inner join  fi_flow b on a.fiid=b.fiid "
	sql = sql + " inner join  fi_flowtask_tb c on a.taskid=c.taskid and b.flowtemplateid=c.flowtemplateid "
	sql = sql + " where a.tiid=? "
	flt := FLOWTASK{}
	err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Tiid).QueryRow(&flt)

	err, _ = insertsecondflow(o, flt, "next", u.Tiid, u.Caller, u.Fiid, u.Flowtemplateid)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

//获得机构变量
func GetAllORGVARY() (admins []ORGVARY, err error) {
	admins = make([]ORGVARY, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_orgvary_tb  "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得机构变量
func GetAllORGVARYoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select vid as value,vname as label from cmn_orgvary_tb  "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//机构变量追加
func AddMultiORGVARY(u []ORGVARY) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	//_, err = o.Delete(&u)
	sql := "delete from cmn_orgvary_tb"
	_, err = o.Raw(sql).Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.InsertMulti(len(u), &u)

	sql = "insert into cmn_orgvary_tb(vid,vname,defaultvalue) values(?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, u1 := range u {
		_, err = o.Raw(sql, u1.Vid, u1.Vname, u1.Defaultvalue).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	admins, _ := GetAllOrg()
	sql = "delete from fi_org_vary where vid=?"
	sql = ConvertSQL(sql, dbtype)
	for _, u1 := range u {
		_, err = o.Raw(sql, u1.Vid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	sql = "insert into fi_org_vary(orgid,vid,vvalue) values(?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, u1 := range u {
		for _, admin := range admins {
			_, err = o.Raw(sql, admin.Orgid, u1.Vid, u1.Defaultvalue).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	}
	err = o.Commit()

	return err
}

//获得流程机构变量分页数据
func GetAllFLOWORGVARYBYPAGEINDEX(p PAGE) (admins []FLOWORGVARY, err error) {
	dbtype := Getdbtype()
	admins = make([]FLOWORGVARY, 0)
	o := orm.NewOrm()
	var limitstr string = ""

	//limitstr = strconv.Itoa((p.Pageindex-1)*p.Pagesize) + "," + strconv.Itoa(p.Pagesize)
	if dbtype == "postgres" {
		limitstr = limitstr + strconv.Itoa(p.Pagesize) + " offset " + strconv.Itoa((p.Pageindex-1)*p.Pagesize)

	} else if dbtype == "mysql" {
		limitstr = limitstr + strconv.Itoa((p.Pageindex-1)*p.Pagesize) + "," + strconv.Itoa(p.Pagesize)

	} else {
		limitstr = limitstr + strconv.Itoa((p.Pageindex-1)*p.Pagesize) + "," + strconv.Itoa(p.Pagesize)
	}
	sql := "select * from fi_org_vary where 1=1 "
	if p.Orgid != "" {
		sql = sql + " and orgid='" + p.Orgid + "' "
	}
	if p.Vid != "" {
		sql = sql + " and vid='" + p.Vid + "' "
	}
	sql = sql + " limit " + limitstr
	_, err = o.Raw(ConvertSQL(sql, dbtype)).QueryRows(&admins)

	return admins, err
}

//获得所有的流程机构变量
func GetAllFLOWORGVARY() (admins []FLOWORGVARY, err error) {
	admins = make([]FLOWORGVARY, 0)
	o := orm.NewOrm()

	sql := "select * from fi_org_vary"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得流程机构变量记录数
func GetAllFLOWORGVARYCOUNT() (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "select count(1) as total from fi_org_vary"

	err = o.Raw(sql).QueryRow(&page)

	return page, err
}
func GetFLOWORGVARYBYORGID(orgid string) (admins []FLOWORGVARY, err error) {
	admins = make([]FLOWORGVARY, 0)
	o := orm.NewOrm()

	sql := "select * from fi_org_vary where orgid=? "

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), orgid).QueryRows(&admins)

	return admins, err
}

//根据用户获得流程机构变量
func GetFLOWORGVARYBYUSERID(userid string) (admins []FLOWORGVARY, err error) {
	admins = make([]FLOWORGVARY, 0)
	o := orm.NewOrm()

	sql := "select b.* from cmn_user_tb a inner join fi_org_vary b on a.orgid=b.orgid where a.userid=? "

	_, err = o.Raw(ConvertSQL(sql, Getdbtype()), userid).QueryRows(&admins)

	return admins, err
}

//流程机构变量追加
func AddMultiFLOWORGVARY(u []FLOWORGVARY) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()

	deletesql := "delete from fi_org_vary where orgid=? and vid=?"

	insertsql := "insert into fi_org_vary(orgid,vid,vvalue) values(?,?,?)"
	for _, u1 := range u {
		_, err = o.Raw(ConvertSQL(deletesql, dbtype), u1.Orgid, u1.Vid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		_, err = o.Raw(ConvertSQL(insertsql, dbtype), u1.Orgid, u1.Vid, u1.Vvalue).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}

//流程状态查询
func GetAllFLOWSTATUS() (admins []FISTATUS, err error) {
	admins = make([]FISTATUS, 0)
	o := orm.NewOrm()

	sql := "select * from fi_flowstatus  "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//流程状态追加
func AddMultiFLOWSTATUS(u []FISTATUS) error {
	o := orm.NewOrm()
	err := o.Begin()
	//_, err = o.Delete(&u)
	sql := "delete from fi_flowstatus"
	_, err = o.Raw(sql).Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	}
	_, err = o.InsertMulti(len(u), &u)

	err = o.Commit()

	return err
}

//taskexecuter
//用户 1
//角色 2
//用户组 3
//任务执行人 4
//发起人所属指定级别机构主管/副主管 5
//发起人所属指定级别机构主管 6
//发起人所属指定级别机构副主管 7
//发起人直属指定级别机构主管/副主管 8
//发起人直属指定级别机构主管 9
//发起人直属指定级别机构副主管 10
//操作人所属指定级别机构主管/副主管 11
//操作人所属指定级别机构主管 12
//操作人所属指定级别机构副主管 13
//操作人直属指定级别机构副主管/副主管 14
//操作人直属指定级别机构主管 15
//操作人直属指定级别机构副主管 16

func Getnexttaskbycaller(orglevel string, nexttaskid string, o orm.Ormer, caller string, flowtemplateid string, actionid string) (nexttaskid1 string, err error) {
	dbtype := Getdbtype()
	var tasktype string
	sql := "select tasktype from fi_flowtask_tb where flowtemplateid=? and taskid=?"
	_ = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid).QueryRow(&tasktype)
	//此任务为分支任务不跳过
	if tasktype == "switch" {
		return nexttaskid, nil
	} else {
		sql = "select * from fi_flowtaskexecuter_tb where flowtemplateid=? and taskid=?"
		ftes := make([]FLOWTASKEXECUTER, 0)
		_, _ = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid).QueryRows(&ftes)
		//包含用户、用户组、角色时不跳过
		for _, fte := range ftes {
			if fte.Taskexecuter == "1" || fte.Taskexecuter == "2" || fte.Taskexecuter == "3" {
				return nexttaskid, nil
				break
			}
		}
		//此任务包含了级别高的人员审批则不跳过
		for _, fte := range ftes {
			expression1, _ := strconv.Atoi(fte.Expression)
			orglevel1, _ := strconv.Atoi(orglevel)
			//级别相同，但分主、副主管
			if expression1 < orglevel1 {
				return nexttaskid, nil
				break
			} else {
				if expression1 == orglevel1 {
					var isleader bool
					sql = "select isleader from cmn_user_tb where userid=?"
					_ = o.Raw(ConvertSQL(sql, dbtype), caller).QueryRow(&isleader)
					if !isleader {
						return nexttaskid, nil
						break
					} else {
						var leadertype string
						sql = "select b.leadertype from cmn_user_tb a inner join cmn_orgleader_tb b on a.orgid=b.orgid and b.userid=a.userid where a.userid=?"
						_ = o.Raw(ConvertSQL(sql, dbtype), caller).QueryRow(&leadertype)
						if leadertype != "" {
							//leadertype1, _ := strconv.Atoi(fte.Expression)
							//申请者是副主管，审批者是正主管，则不能跳过
							if fte.Taskexecuter == "6" || fte.Taskexecuter == "9" || fte.Taskexecuter == "12" || fte.Taskexecuter == "15" {
								if leadertype == "1" {
									return nexttaskid, nil
									break
								}
							}

						}
					}

				}
			}
		}
		var nexttaskid2 string
		sql = "select nexttask from fi_flowmantaskaction_tb where flowtemplateid=? and taskid=? and action=?"
		_ = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid, nexttaskid, actionid).QueryRow(&nexttaskid2)
		if nexttaskid2 == "999" {
			return "999", nil
		}
		if nexttaskid2 == "" {
			return "-1", errors.New("can't find nexttaskid by " + nexttaskid)
		}
		return Getnexttaskbycaller(orglevel, nexttaskid2, o, caller, flowtemplateid, actionid)

	}
}

//流程模板追加
func AddMultiagent(u2 []AGENT) error {
	o := orm.NewOrm()
	err := o.Begin()
	deletesql := "delete from fi_agent_tb"
	_, err = o.Raw(deletesql).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	sql := "insert into fi_agent_tb(userid,agent,startdate,enddate) values(?,?,?,?)"
	for _, u3 := range u2 {
		_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u3.Userid, u3.Agent, u3.Startdate, u3.Enddate).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}

//获得所有的流程模板
func GetAllagent() (admins []AGENT, err error) {
	admins = make([]AGENT, 0)
	o := orm.NewOrm()

	sql := "select * from fi_agent_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得本部门或全体用户
func GetUserforagent(ufa USERFORAGENT) (admins []PROFILE, err error) {
	dbtype := Getdbtype()
	admins = make([]PROFILE, 0)
	o := orm.NewOrm()

	sql := "select userid,username from cmn_user_tb where expired='0' and 1=1 "
	if ufa.Userid != "" {
		sql = sql + " and userid like '%" + ufa.Userid + "%'"
	}

	if ufa.Isleader == true {

		if dbtype == "postgres" {
			sql = sql + " and isleader=true"
		} else {
			sql = sql + " and isleader=1"
		}
	} else {

		if dbtype == "postgres" {
			sql = sql + " and isleader=false"
		} else {
			sql = sql + " and isleader=0"
		}
	}
	if ufa.Ownerdepartment == true {
		ownerorgis := getownerorgid(ufa.Submitter)
		if ownerorgis != "all" {
			sql = sql + " and orgid in(" + getownerorgid(ufa.Submitter) + ")"
		}
	}
	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得本部门或全体用户
func GetUserforagentoptions(ufa USERFORAGENT) (admins []OPTIONS, err error) {
	dbtype := Getdbtype()
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select userid as value,username as label from cmn_user_tb where expired='0' and 1=1 "
	if ufa.Userid != "" {
		sql = sql + " and userid like '%" + ufa.Userid + "%'"
	}

	if ufa.Isleader == true {

		if dbtype == "postgres" {
			sql = sql + " and isleader=true"
		} else {
			sql = sql + " and isleader=1"
		}
	} else {

		if dbtype == "postgres" {
			sql = sql + " and isleader=false"
		} else {
			sql = sql + " and isleader=0"
		}
	}
	if ufa.Ownerdepartment == true {
		ownerorgis := getownerorgid(ufa.Submitter)
		if ownerorgis != "all" {
			sql = sql + " and orgid in(" + getownerorgid(ufa.Submitter) + ")"
		}
	}
	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func getownerorgid(userid string) string {
	sql := "select a.orglevel,a.orgid from cmn_org_tb a inner join cmn_user_tb b on a.orgid=b.orgid where b.userid=?"
	o := orm.NewOrm()

	var orglevel, orgid string
	_ = o.Raw(ConvertSQL(sql, Getdbtype()), userid).QueryRow(&orglevel, &orgid)
	var orglevel1 int
	orglevel1, _ = strconv.Atoi(orglevel)
	if orglevel1 == 0 {
		return "all"
	}
	if orglevel1 >= 1 && orglevel1 <= 3 {
		var returnstr string = ""
		orgids := PORG{}
		getallorgidbyparentid(o, &orgids, orgid)
		for index, orgid := range orgids.Porg {
			if index == 0 {
				returnstr = returnstr + "'" + orgid.Orgid + "'"
			} else {
				returnstr = returnstr + ",'" + orgid.Orgid + "'"
			}
		}
		returnstr = returnstr + ",'" + orgid + "'"
		return returnstr
	} else {
		parentorgid := getparentorgidbyorgid(o, orgid)
		var returnstr string = ""
		orgids := PORG{}
		getallorgidbyparentid(o, &orgids, parentorgid)
		for index, orgid := range orgids.Porg {
			if index == 0 {
				returnstr = returnstr + "'" + orgid.Orgid + "'"
			} else {
				returnstr = returnstr + ",'" + orgid.Orgid + "'"
			}
		}
		if returnstr != "" {
			returnstr = returnstr + ",'" + orgid + "'"
		} else {
			returnstr = returnstr + "'" + orgid + "'"
		}
		return returnstr
	}
	return ""
}
func getallorgidbyparentid(o orm.Ormer, allorgid *PORG, parentid string) {

	sql := "select orgid from cmn_org_tb a where parentid=?"
	//o := orm.NewOrm()
	orgids := make([]CMN_ORG_TB, 0)
	_, _ = o.Raw(ConvertSQL(sql, Getdbtype()), parentid).QueryRows(&orgids)
	for _, orgid := range orgids {
		allorgid.Porg = append(allorgid.Porg, orgid)
		getallorgidbyparentid(o, allorgid, orgid.Orgid)
	}
}

//
func getparentorgidbyorgid(o orm.Ormer, orgid string) string {

	sql := "select a.parentid,b.orglevel from cmn_org_tb a inner join cmn_org_tb b on a.parentid=b.orgid where a.orgid=?"
	//o := orm.NewOrm()
	var parentorgid, orglevel string
	_ = o.Raw(ConvertSQL(sql, Getdbtype()), orgid).QueryRow(&parentorgid, &orglevel)
	if orglevel == "3" {
		return parentorgid

	} else {
		return getparentorgidbyorgid(o, parentorgid)
	}
}

//获得流程状态
func Getflowstatus(fiid int) (ff FISTATUS) {

	sql := "select a.flowstatus,b.flowstatusname from fi_flow a inner join fi_flowstatus b on a.flowstatus=b.flowstatus where a.fiid=?"
	o := orm.NewOrm()
	_ = o.Raw(ConvertSQL(sql, Getdbtype()), fiid).QueryRow(&ff)
	return ff
}

//流程模板复制，同时复制流程定义
func Copyflowtemplate(flowtemplateid string, copyflowtemplateid string, copyflowtemplatename string) (err error) {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql := "insert into fi_template_tb(select '"
	sql = sql + copyflowtemplateid + "','" + copyflowtemplatename + "',a.flowcontent,a.flowinstidcol,a.flowstatuscol  from fi_template_tb a where a.flowtemplateid=?)"
	sql = SQLBRACKET2SPACE(sql, dbtype)
	_, err = o.Raw(ConvertSQL(sql, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "insert into fi_templateitem_tb(flowtemplateid,vary,varyname,varytype,varyvalue)"
	sql2 := "(select '" + copyflowtemplateid + "',a.vary,a.varyname,a.varytype,a.varyvalue  from fi_templateitem_tb a where a.flowtemplateid=?)"
	sql2 = SQLBRACKET2SPACE(sql2, dbtype)
	_, err = o.Raw(ConvertSQL(sql+sql2, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "insert into fi_flowtask_tb(flowtemplateid,tasktype,taskid,taskname,supportskip,sendmessage,concurrent,samepersontask,nopersontask)"
	sql2 = "(select '" + copyflowtemplateid + "',tasktype,taskid,taskname,supportskip,sendmessage,concurrent,samepersontask,nopersontask  from fi_flowtask_tb  where flowtemplateid=?)"
	sql2 = SQLBRACKET2SPACE(sql2, dbtype)
	_, err = o.Raw(ConvertSQL(sql+sql2, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "insert into fi_flowmantaskaction_tb(flowtemplateid,tasktype,taskid,action,jump,status,nexttask,backtask)"
	sql2 = "(select '" + copyflowtemplateid + "',tasktype,taskid,action,jump,status,nexttask,backtask  from fi_flowmantaskaction_tb  where flowtemplateid=?)"
	sql2 = SQLBRACKET2SPACE(sql2, dbtype)
	_, err = o.Raw(ConvertSQL(sql+sql2, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "insert into fi_flowtaskexecuter_tb(flowtemplateid,tasktype,taskid,no,taskexecuter,expression)"
	sql2 = "(select '" + copyflowtemplateid + "',tasktype,taskid,no,taskexecuter,expression  from fi_flowtaskexecuter_tb  where flowtemplateid=?)"
	sql2 = SQLBRACKET2SPACE(sql2, dbtype)
	_, err = o.Raw(ConvertSQL(sql+sql2, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql = "insert into fi_flowswitchtaskaction_tb(flowtemplateid,tasktype,taskid,nos,conditions,functions,valuee,jump,statuss,nexttask,backtask)"
	sql2 = "(select '" + copyflowtemplateid + "',tasktype,taskid,nos,conditions,functions,valuee,jump,statuss,nexttask,backtask  from fi_flowswitchtaskaction_tb  where flowtemplateid=?)"
	sql2 = SQLBRACKET2SPACE(sql2, dbtype)
	_, err = o.Raw(ConvertSQL(sql+sql2, dbtype), flowtemplateid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	err = o.Commit()
	return nil
}

//
//转签,待办转给其它审批人审批。
func Transfersign(u TRANSFER) error {
	o := orm.NewOrm()
	err := o.Begin()
	sql := "update fi_owner set owner=? where tiid=?"
	for _, trs := range u.Listdata {
		_, err = o.Raw(ConvertSQL(sql, Getdbtype()), u.Transferuserid, trs.Tiid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	err = o.Commit()

	return err
}

//
//申请人转岗，待办转给新部门领导
//审批人转岗，待办转给部门新领导
func Transferpost(u TRANSFER) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	//geteditors(o orm.Ormer, caller string, flte []FLOWTASKEXECUTER)
	//fi_flowtaskexecuter_tb
	sql := "select * from  fi_flowtaskexecuter_tb where flowtemplateid=? and taskid=?"
	for _, trs := range u.Listdata {
		ftes := make([]FLOWTASKEXECUTER, 0)
		_, err = o.Raw(ConvertSQL(sql, dbtype), trs.Flowtemplateid, trs.Taskid).QueryRows(&ftes)
		editors := geteditors(o, u.Userid, ftes)
		if len(editors) < 1 {
			continue
		}
		deletesql := "delete from fi_owner where tiid=?"
		_, err = o.Raw(ConvertSQL(deletesql, dbtype), trs.Tiid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		for _, editor := range editors {
			insertsql := "insert into fi_owner(tiid,owner) values(?,?)"
			_, err = o.Raw(ConvertSQL(insertsql, dbtype), trs.Tiid, editor).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}

	}
	err = o.Commit()

	return err
}

//
//申请人离职,待办流程转给他人，更新申请人。
//审批者离职,转给他人审批，更新待办的审批人。
//也可取消流程。
func Leave(u TRANSFER) error {
	o := orm.NewOrm()
	err := o.Begin()
	//func stopflow(o orm.Ormer, modualid string, currentfiid int, currenttiid int, editor string) (err error)
	dbtype := Getdbtype()
	if u.Cancel {
		for _, trs := range u.Listdata {
			err = stopflow(o, trs.Flowtemplateid, trs.Fiid, trs.Tiid, u.Submitter)
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	} else {
		if u.Usertype == "0" {
			sql := "update fi_flow set caller=? where fiid=?"
			for _, trs := range u.Listdata {
				_, err = o.Raw(ConvertSQL(sql, dbtype), u.Transferuserid, trs.Fiid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}
		} else {
			sql := "update fi_owner set owner=? where tiid=?"
			for _, trs := range u.Listdata {
				_, err = o.Raw(ConvertSQL(sql, dbtype), u.Transferuserid, trs.Tiid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}
		}
	}
	err = o.Commit()

	return err
}

//为转签转岗业务获得待办
func Gettodotaskfortransfer(u TRANSFER) (admins []TODOTASKLIST, err error) {
	admins = make([]TODOTASKLIST, 0)
	o := orm.NewOrm()
	dbtype := Getdbtype()
	if u.Usertype == "0" {
		sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,b.taskid,a.flowstarttime,a.flowstatus,c.tiid,a.flowtemplateid, "
		sql = sql + "  e.url,f.flowstatusname,c.owner as editor from fi_flow a "
		sql = sql + "  inner join fi_task b on a.fiid=b.fiid and taskstatus='0' "
		sql = sql + "  inner join fi_owner c on b.tiid=c.tiid "
		sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
		sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
		sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where  a.flowstatus='0' and a.caller=? "

		_, err = o.Raw(ConvertSQL(sql, dbtype), u.Userid).QueryRows(&admins)
	}
	if u.Usertype == "1" {
		sql := "SELECT a.fiid,a.caller,CONCAT_WS('/',e.modualname,a.flowcontent) as flowcontent,b.taskid,a.flowstarttime,a.flowstatus,c.tiid,a.flowtemplateid, "
		sql = sql + "  e.url,f.flowstatusname,c.owner as editor from fi_flow a "
		sql = sql + "  inner join fi_task b on a.fiid=b.fiid and taskstatus='0' "
		sql = sql + "  inner join fi_owner c on b.tiid=c.tiid and c.owner=? "
		sql = sql + "  inner join cmn_modualtemplate_tb d on a.flowtemplateid=d.flowtemplateid and a.modualid=d.modualid "
		sql = sql + "  inner join fi_flowstatus f on a.flowstatus=f.flowstatus "
		sql = sql + "  inner join cmn_modual_tb e on d.modualid=e.modualid where  a.flowstatus='0' "

		_, err = o.Raw(ConvertSQL(sql, dbtype), u.Userid).QueryRows(&admins)
	}
	return admins, err
}

//根据tiid获得当前的任务ID
func Gettaskidbytiid(o orm.Ormer, tiid int) string {
	dbtype := Getdbtype()
	taskid := "1"
	sql := "select taskid from fi_task where tiid=?"
	err := o.Raw(ConvertSQL(sql, dbtype), tiid).QueryRow(&taskid)
	if err != nil {
		taskid = "-1"
	}
	return taskid
}

//根据tiid获得当前的任务ID
func Gettaskidbytiidforng(tiid int) string {
	o := orm.NewOrm()
	dbtype := Getdbtype()
	taskid := "1"
	sql := "select taskid from fi_task where tiid=?"
	err := o.Raw(ConvertSQL(sql, dbtype), tiid).QueryRow(&taskid)
	if err != nil {
		taskid = "-1"
	}
	return taskid
}
