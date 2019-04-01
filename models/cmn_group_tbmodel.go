package models

import (
	_ "errors"
	"fmt"
	_ "strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`

type CMN_GROUP_TB struct {
	Groupid    string `orm:"pk;column(groupid)"`
	Parentid   string `orm:"column(parentid)"`
	Groupname  string `orm:"column(groupname)"`
	Grouplevel string `orm:"column(grouplevel)"`
	Remark     string `orm:"column(remark);null"`
}
type CMN_GROUPROLE_TB struct {
	Groupid string
	Roleid  string
}
type group struct {
	Groupid string
}
type Parentgroupid struct {
	Groupid   string `json:"Parentid"`
	Groupname string `json:"Parentname"`
}

func (u *CMN_GROUP_TB) TableName() string {
	return "cmn_group_tb"
}

type CMN_USERGROUP_TB struct {
	Userid      string
	Groupid     string
	Expiredtime time.Time
	Username    string
	Groupname   string
}

func AddMultiCMN_GROUPROLE_TB(groupid string, u []CMN_GROUPROLE_TB) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete  from cmn_grouprole_tb where groupid=?"
	dbtype := Getdbtype()
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, groupid).Exec()

	if err != nil {
		o.Rollback()
		return err
	}
	sql = "insert into cmn_grouprole_tb(groupid,roleid) values(?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, grouprole := range u {
		_, err = o.Raw(sql, groupid, grouprole.Roleid).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}
func AddCMN_GROUP_TB(u CMN_GROUP_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(&u)
	if err != nil {
		//fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.Insert(&u)
	insertsql := "insert into cmn_group_tb(groupid,parentid,groupname,grouplevel,remark) values(?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	_, err = o.Raw(insertsql, u.Groupid, u.Parentid, u.Groupname, u.Grouplevel, u.Remark).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}

func AddMultiCMN_GROUP_TB(u []CMN_GROUP_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	//_, err = o.InsertMulti(len(u), &u)
	insertsql := "insert into cmn_group_tb(groupid,parentid,groupname,grouplevel,remark) values(?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		_, err = o.Raw(insertsql, u1.Groupid, u1.Parentid, u1.Groupname, u1.Grouplevel, u1.Remark).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}

func GetAllCMN_GROUP_TB() (admins []CMN_GROUP_TB, err error) {
	admins = make([]CMN_GROUP_TB, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_group_tb order by parentid "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetCMN_GROUP_TB(u *CMN_GROUP_TB) (admins []CMN_GROUP_TB, err error) {
	admins = make([]CMN_GROUP_TB, 0)
	o := orm.NewOrm()
	sql := "select * from cmn_group_tb where 1=1 "

	if u.Groupid != "" {
		sql = sql + " and groupid='" + u.Groupid + "'"
	}

	if u.Parentid != "" {
		sql = sql + " and parentid='" + u.Parentid + "'"
	}

	if u.Groupname != "" {
		sql = sql + " and groupname='" + u.Groupname + "'"
	}

	if u.Grouplevel != "" {
		sql = sql + " and grouplevel='" + u.Grouplevel + "'"
	}

	if u.Remark != "" {
		sql = sql + " and remark='" + u.Remark + "'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetCMN_GROUPROLE_TB(u CMN_GROUP_TB) (admins []CMN_GROUPROLE_TB, err error) {
	admins = make([]CMN_GROUPROLE_TB, 0)
	o := orm.NewOrm()
	sql := "select groupid,roleid from cmn_grouprole_tb where groupid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Groupid).QueryRows(&admins)

	return admins, err
}

func DeleteCMN_GROUP_TB(u CMN_GROUP_TB) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(&u)

	if err != nil {
		fmt.Println(err)
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}

func UpdateCMN_GROUP_TB(u CMN_GROUP_TB) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Update(&u)

	if err != nil {
		fmt.Println(err)
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}
func Getparentgroupids() (parentids []Parentgroupid, err error) {
	parentids = make([]Parentgroupid, 0)
	o := orm.NewOrm()

	sql := "select groupid,groupname from cmn_group_tb order by parentid "

	_, err = o.Raw(sql).QueryRows(&parentids)

	return parentids, err
}
func DeleteCMN_GROUPROLE_TB(u CMN_GROUP_TB) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete from cmn_grouprole_tb where groupid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Groupid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}
func GetGroupbyid(mt CMN_GROUP_TB) (admin CMN_GROUP_TB, err error) {

	o := orm.NewOrm()

	sql := "select * from cmn_group_tb where groupid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, mt.Groupid).QueryRow(&admin)

	return admin, err
}
func DeleteGroupbyid(mt CMN_GROUP_TB) (err error) {

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
func AddMultiCMN_USERGROUP_TB(u []CMN_USERGROUP_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	sql := "delete from cmn_usergroup_tb where groupid=?"
	dbtype := Getdbtype()
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, u[0].Groupid).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	sql = "insert into cmn_usergroup_tb(userid,groupid,expireddate) values(?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, orgrole := range u {
		_, err = o.Raw(sql, orgrole.Userid, orgrole.Groupid, time.Now()).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}
func GetCMN_USERGROUP_TB(u CMN_GROUP_TB) (admins []CMN_USERGROUP_TB, err error) {
	admins = make([]CMN_USERGROUP_TB, 0)
	o := orm.NewOrm()
	sql := "select * from cmn_usergroup_tb where groupid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Groupid).QueryRows(&admins)

	return admins, err
}
func GetAllGroupoptions() (admins []OPTIONS, err error) {

	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select a.groupid as value,a.groupname as label from cmn_group_tb a  order by grouplevel"

	_, err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRows(&admins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admins)
	return admins, err
}
