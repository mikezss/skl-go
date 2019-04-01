package models

import (
	_ "errors"
	"fmt"
	_ "strconv"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
type ENUM struct {
	Enumcode string `orm:"pk;column(enumcode)"`
	Enumname string `orm:"column(enumname)"`
}
type ENUMITEM struct {
	Id       int64  `orm:"pk;auto"`
	Enumcode string `orm:"column(enumcode)"`
	Value    string `orm:"column(value)"`
	Label    string `orm:"column(label)"`
}
type ENUMANDITEM struct {
	Enum     ENUM
	Enumitem []ENUMITEM
}

func (u *ENUM) TableName() string {
	return "skl_enum_tb"
}

// 多字段唯一键
func (u *ENUMITEM) TableUnique() [][]string {
	return [][]string{
		[]string{"Enumcode", "Value"},
	}
}
func (u *ENUMITEM) TableName() string {
	return "skl_enumitem_tb"
}

func AddMultiENUM(u ENUM, u2 []ENUMITEM) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	_, err = o.Delete(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Insert(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	sql := "delete from  skl_enumitem_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Enumcode).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.InsertMulti(len(u2), &u2)
	sql = "insert into skl_enumitem_tb(enumcode,value,label) values(?,?,?)"
	sql = ConvertSQL(sql, Getdbtype())
	for _, u3 := range u2 {
		_, err = o.Raw(sql, u.Enumcode, u3.Value, u3.Label).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}

func GetAllENUM() (admins []ENUM, err error) {
	admins = make([]ENUM, 0)
	o := orm.NewOrm()

	sql := "select * from skl_enum_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllENUMoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select enumcode as value,enumname as label from skl_enum_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetENUMBYID(e ENUM) (admin ENUM, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_enum_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Enumcode).QueryRow(&admin)

	return admin, err
}
func GetAllENUMITEM(e ENUM) (admins []ENUMITEM, err error) {
	admins = make([]ENUMITEM, 0)
	o := orm.NewOrm()

	sql := "select * from skl_enumitem_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, e.Enumcode).QueryRows(&admins)

	return admins, err
}
func GetAllENUMITEMoptions(e ENUM) (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select * from skl_enumitem_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, e.Enumcode).QueryRows(&admins)

	return admins, err
}
func DeleteENUM(u *ENUM) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	sql := "delete from skl_enumitem_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Enumcode).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}
func GetENUM(e ENUM) (admin ENUM, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_enum_tb where enumcode=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Enumcode).QueryRow(&admin)

	return admin, err
}
