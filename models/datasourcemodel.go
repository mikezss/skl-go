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
type DATASOURCE struct {
	Datasource       string `orm:"pk;column(datasource)"`
	Datasourcename   string `orm:"column(datasourcename)"`
	Dbtype           string `orm:"column(dbtype)"`
	Ip               string `orm:"column(ip)"`
	Port             int    `orm:"column(port)"`
	Schema           string `orm:"column(schema)"`
	Username         string `orm:"column(username)"`
	Password         string `orm:"column(password)"`
	Sourcetargettype string `orm:"column(sourcetargettype)"`
}

func (u *DATASOURCE) TableName() string {
	return "skl_datasource_tb"
}

func AddMultiDATASOURCE(u DATASOURCE) (err error) {
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

	err = o.Commit()

	return err
}

func GetAllDATASOURCE() (admins []DATASOURCE, err error) {
	admins = make([]DATASOURCE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasource_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllMYSQLDATASOURCE() (admins []DATASOURCE, err error) {
	admins = make([]DATASOURCE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasource_tb where dbtype='mysql'"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllORACLEDATASOURCE() (admins []DATASOURCE, err error) {
	admins = make([]DATASOURCE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasource_tb where dbtype='oracle'"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetDATASOURCEBYID(e DATASOURCE) (admin DATASOURCE, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_datasource_tb where datasource=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Datasource).QueryRow(&admin)

	return admin, err
}
func DeleteDATASOURCE(u *DATASOURCE) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	err = o.Commit()

	return err
}
func GetAllDATASOURCEoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select datasource as value,datasourcename as label from skl_datasource_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetDATASOURCEoptionsbytype(ds DATASOURCE) (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select datasource as value,datasourcename as label from skl_datasource_tb where sourcetargettype=?"

	_, err = o.Raw(sql, ds.Sourcetargettype).QueryRows(&admins)

	return admins, err
}
