package models

import (
	_ "errors"
	"fmt"
	"strconv"
	_ "time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
type CMN_ORG_TB struct {
	Orgid    string `orm:"pk;column(orgid)"`
	Orgname  string `orm:"column(orgname)"`
	Parentid string `orm:"column(parentid)"`
	Orgtype  string `orm:"column(orgtype);null"`
	Orglevel string `orm:"column(orglevel)"`
	Url      string `orm:"column(url);null"`
	Remark   string `orm:"column(remark);null"`
}
type CMN_ORGLEADER_TB struct {
	Orgid      string
	Userid     string
	Leadertype string
}
type CMN_ORGANDLEADER_TB struct {
	Org       CMN_ORG_TB
	Orgleader []CMN_ORGLEADER_TB
}
type CMN_ORGROLE_TB struct {
	Orgid  string
	Roleid string
}

func (u *CMN_ORG_TB) TableName() string {
	return "cmn_org_tb"
}

func SaveOrg(mt CMN_ORG_TB, orglds []CMN_ORGLEADER_TB) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Delete(&mt)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	m, _, _ := GetOrgbyid(CMN_ORG_TB{Orgid: mt.Parentid})
	orglevel, err := strconv.Atoi(m.Orglevel)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	strorglevel := strconv.Itoa(orglevel + 1)
	mt.Orglevel = strorglevel

	//_, err = o.Insert(&mt)
	insertsql := "insert into cmn_org_tb(orgid,orgname,parentid,orgtype,orglevel,url,remark) values(?,?,?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	_, err = o.Raw(insertsql, mt.Orgid, mt.Orgname, mt.Parentid, mt.Orgtype, mt.Orglevel, mt.Url, mt.Remark).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	dbtype := Getdbtype()
	deletesql := "delete from cmn_orgleader_tb where orgid=?"
	deletesql = ConvertSQL(deletesql, dbtype)
	_, err = o.Raw(deletesql, mt.Orgid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql := "insert into cmn_orgleader_tb(orgid,userid,leadertype) values(?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, orgld := range orglds {
		_, err = o.Raw(sql, orgld.Orgid, orgld.Userid, orgld.Leadertype).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	err = o.Commit()
	return err
}

func GetAllOrg() (admins []CMN_ORG_TB, err error) {
	dbtype := Getdbtype()
	admins = make([]CMN_ORG_TB, 0)
	o := orm.NewOrm()
	//getleadertitle(orgid,orgname) as
	sql := "select a.orgid,a.parentid,a.orgtype,a.orglevel,a.remark,a.url,(case when CONCAT(a.orgname,'(',a.orgid,')',GROUP_CONCAT(DISTINCT c.username order by b.leadertype SEPARATOR ',') ) is null then CONCAT(a.orgname,'(',a.orgid,')')   else CONCAT(a.orgname,'(',a.orgid,')',GROUP_CONCAT(DISTINCT c.username order by b.leadertype SEPARATOR ',')) end) as orgname from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	if dbtype == "postgres" {
		sql = "select a.orgid,a.parentid,a.orgtype,a.orglevel,a.remark,a.url,(case when CONCAT(a.orgname,'(',a.orgid,')',array_to_string(group_concat(c.username),',') ) is null then CONCAT(a.orgname,'(',a.orgid,')')   else CONCAT(a.orgname,'(',a.orgid,')',array_to_string(group_concat(c.username),',')) end) as orgname from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	}
	//GROUP_CONCAT(DISTINCT a ORDER BY a DESC SEPARATOR ‘-‘)
	if dbtype == "sqlite3" {
		sql = "select a.orgid,a.parentid,a.orgtype,a.orglevel,a.remark,a.url,(case when a.orgname||'('||a.orgid||')'||GROUP_CONCAT(DISTINCT c.username)  is null then  a.orgname||'('||a.orgid||')' else a.orgname||'('||a.orgid||')'||GROUP_CONCAT(DISTINCT c.username )  end) as orgname from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	}
	_, err = o.Raw(sql).QueryRows(&admins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admins)
	return admins, err
}

func GetAllOrgoptions() (admins []OPTIONS, err error) {
	dbtype := Getdbtype()
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()
	//getleadertitle(orgid,orgname) as
	sql := "select a.orgid as value,(case when CONCAT(a.orgname,'(',a.orgid,')',GROUP_CONCAT(DISTINCT c.username order by b.leadertype SEPARATOR ',') ) is null then CONCAT(a.orgname,'(',a.orgid,')')   else CONCAT(a.orgname,'(',a.orgid,')',GROUP_CONCAT(DISTINCT c.username order by b.leadertype SEPARATOR ',')) end) as label from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	if dbtype == "postgres" {
		sql = "select a.orgid as value,(case when CONCAT(a.orgname,'(',a.orgid,')',array_to_string(group_concat(c.username),',') ) is null then CONCAT(a.orgname,'(',a.orgid,')')   else CONCAT(a.orgname,'(',a.orgid,')',array_to_string(group_concat(c.username),',')) end) as label from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	}
	//GROUP_CONCAT(DISTINCT a ORDER BY a DESC SEPARATOR ‘-‘)
	if dbtype == "sqlite3" {
		sql = "select a.orgid as value,(case when a.orgname||'('||a.orgid||')'||GROUP_CONCAT(DISTINCT c.username)  is null then  a.orgname||'('||a.orgid||')' else a.orgname||'('||a.orgid||')'||GROUP_CONCAT(DISTINCT c.username )  end) as label from cmn_org_tb a left join cmn_orgleader_tb b on a.orgid=b.orgid  left join cmn_user_tb c on b.userid=c.userid group by a.orgid"
	}
	_, err = o.Raw(sql).QueryRows(&admins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admins)
	return admins, err
}

func GetOrgbyid(mt CMN_ORG_TB) (admin CMN_ORG_TB, admins []CMN_ORGLEADER_TB, err error) {
	admins = make([]CMN_ORGLEADER_TB, 0)
	o := orm.NewOrm()
	dbtype := Getdbtype()
	sql := "select * from cmn_org_tb where orgid=?"
	sql = ConvertSQL(sql, dbtype)
	err = o.Raw(sql, mt.Orgid).QueryRow(&admin)

	sql = "select * from cmn_orgleader_tb where orgid=?"
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, mt.Orgid).QueryRows(&admins)

	return admin, admins, err
}
func DeleteOrgbyid(mt CMN_ORG_TB) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Delete(&mt)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql := "delete from cmn_orgleader_tb where orgid=?"
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err = o.Raw(deletesql, mt.Orgid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}
func AddMultiCMN_ORGROLE_TB(orgid string, u []CMN_ORGROLE_TB) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete  from cmn_orgrole_tb where orgid=?"
	dbtype := Getdbtype()
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, orgid).Exec()

	if err != nil {
		o.Rollback()
		return err
	}
	sql = "insert into cmn_orgrole_tb(orgid,roleid) values(?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, orgrole := range u {
		_, err = o.Raw(sql, orgrole.Orgid, orgrole.Roleid).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}
func GetCMN_ORGROLE_TB(u CMN_ORG_TB) (admins []CMN_ORGROLE_TB, err error) {
	admins = make([]CMN_ORGROLE_TB, 0)
	o := orm.NewOrm()
	sql := "select orgid,roleid from cmn_orgrole_tb where 1=1 "

	if u.Orgid != "" {
		sql = sql + " and orgid='" + u.Orgid + "'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func DeleteCMN_ORGROLE_TB(u CMN_ORG_TB) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete from cmn_orgrole_tb where orgid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Orgid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}
