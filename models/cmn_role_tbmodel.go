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
type CMN_ROLE_TB struct {
	Roleid    string `orm:"pk;column(roleid)"`
	Rolename  string `orm:"column(rolename)"`
	Parentid  string `orm:"column(parentid)"`
	Rolelevel string `orm:"column(rolelevel)"`
	Remark    string `orm:"column(remark);null"`
}
type CMN_ROLEPRIVILEGE_TB struct {
	Roleid   string
	Modualid string
}

func (u *CMN_ROLE_TB) TableName() string {
	return "cmn_role_tb"
}

func SaveRole(mt CMN_ROLE_TB) (err error) {

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

	m, _ := GetRolebyid(CMN_ROLE_TB{Roleid: mt.Parentid})
	rolelevel, err := strconv.Atoi(m.Rolelevel)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	strrolelevel := strconv.Itoa(rolelevel + 1)
	mt.Rolelevel = strrolelevel

	//_, err = o.Insert(&mt)
	insertsql := "insert into cmn_role_tb(roleid,rolename,parentid,rolelevel,remark) values(?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	_, err = o.Raw(insertsql, mt.Roleid, mt.Rolename, mt.Parentid, mt.Rolelevel, mt.Remark).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

func GetAllRole() (admins []CMN_ROLE_TB, err error) {
	admins = make([]CMN_ROLE_TB, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_role_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllRoleoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select roleid as value,rolename as label from cmn_role_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetRolebyid(mt CMN_ROLE_TB) (admin CMN_ROLE_TB, err error) {

	o := orm.NewOrm()

	sql := "select * from cmn_role_tb where roleid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, mt.Roleid).QueryRow(&admin)

	return admin, err
}
func DeleteRolebyid(mt CMN_ROLE_TB) (err error) {

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

	err = o.Commit()

	return err
}
func AddMultiCMN_ROLEPRIVILEGE_TB(roleid string, u []CMN_ROLEPRIVILEGE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	dbtype := Getdbtype()
	sql := "delete  from cmn_roleprivilege_tb where roleid=?"
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, roleid).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	sql = "insert into cmn_roleprivilege_tb values(?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, rolerole := range u {
		_, err = o.Raw(sql, roleid, rolerole.Modualid).Exec()
		if err != nil {
			err = o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}
func GetCMN_ROLEPRIVILEGE_TB(roleid string) (admins []CMN_ROLEPRIVILEGE_TB, err error) {
	admins = make([]CMN_ROLEPRIVILEGE_TB, 0)
	o := orm.NewOrm()
	sql := "select roleid,modualid from cmn_roleprivilege_tb where 1=1 "

	if roleid != "" {
		sql = sql + " and roleid='" + roleid + "'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func DeleteCMN_ROLEPRIVILEGE_TB(roleid string) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql := "delete  from cmn_roleprivilege_tb where roleid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, roleid).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}
func AddMultiCMN_USERROLE_TBbyrole(u []CMN_USERROLE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	sql := "delete from cmn_userrole_tb where roleid=?"
	dbtype := Getdbtype()
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, u[0].Roleid).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	if u[0].Userid != "" {
		sql = "insert into cmn_userrole_tb(userid,roleid) values(?,?)"
		sql = ConvertSQL(sql, dbtype)
		for _, orgrole := range u {
			_, err = o.Raw(sql, orgrole.Userid, orgrole.Roleid).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}
	err = o.Commit()

	return err
}
func GetCMN_USERROLE_TBbyroleid(u CMN_USERROLE_TB) (admins []CMN_USERROLE_TB, err error) {
	admins = make([]CMN_USERROLE_TB, 0)
	o := orm.NewOrm()
	sql := "select * from cmn_userrole_tb where roleid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Roleid).QueryRows(&admins)

	return admins, err
}
