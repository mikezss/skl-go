package models

import (
	_ "errors"
	"fmt"
	_ "strconv"
	_ "time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
type CMN_MODUAL_TB struct {
	Modualid   string `orm:"pk;column(modualid)"`
	Modualname string `orm:"column(modualname)"`
	Parentid   string `orm:"column(parentid)"`
	Url        string `orm:"column(url);null"`
	Remark     string `orm:"column(remark);null"`
	Displayno  int64  `orm:"column(displayno);default(1)"`
}
type CMN_MODUALTEMPLATE_TB struct {
	Modualid       string `orm:"pk;column(modualid)"`
	Flowtemplateid string `orm:"column(flowtemplateid)"`
	Tablename      string `orm:"column(tablename)"`
}

func (u *CMN_MODUALTEMPLATE_TB) TableName() string {
	return "cmn_modualtemplate_tb"
}

func (u *CMN_MODUAL_TB) TableName() string {
	return "cmn_modual_tb"
}

func SaveModual(mt CMN_MODUAL_TB) (err error) {

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
	//_, err = o.Insert(&mt)
	insertsql := "insert into cmn_modual_tb(modualid,modualname,parentid,url,remark,displayno) values(?,?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	_, err = o.Raw(insertsql, mt.Modualid, mt.Modualname, mt.Parentid, mt.Url, mt.Remark, mt.Displayno).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}

func GetAllModual() (admins []CMN_MODUAL_TB, err error) {
	admins = make([]CMN_MODUAL_TB, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_modual_tb order by displayno"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllModualoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select modualid as value,modualname as label from cmn_modual_tb order by displayno"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetModualbyid(mt CMN_MODUAL_TB) (admin CMN_MODUAL_TB, err error) {

	o := orm.NewOrm()

	sql := "select * from cmn_modual_tb where modualid=? order by displayno"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, mt.Modualid).QueryRow(&admin)

	return admin, err
}
func DeleteModualbyid(mt CMN_MODUAL_TB) (err error) {

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
func SaveModualtemplate(mt CMN_MODUALTEMPLATE_TB) (err error) {

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
	_, err = o.Insert(&mt)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()
	return err
}
func DeleteModualtemplatebyid(mt CMN_MODUALTEMPLATE_TB) (err error) {

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
func GetModualtemplatebyid(mt CMN_MODUALTEMPLATE_TB) (admin CMN_MODUALTEMPLATE_TB, err error) {

	o := orm.NewOrm()

	sql := "select * from cmn_modualtemplate_tb where modualid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, mt.Modualid).QueryRow(&admin)

	return admin, err
}
