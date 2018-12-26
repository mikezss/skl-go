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
//{{Unescapedjs .uppercomponentname}}
type COMPANY struct {
	Companycode      string `orm:"pk;column(companycode)"`
	Companyname      string `orm:"column(companyname)"`
	Companyshortname string `orm:"column(companyshortname)"`
	Manager          string `orm:"column(manager)"`
	Telphone         string `orm:"column(telphone)"`
	Fax              string `orm:"column(fax)"`
	Email            string `orm:"column(email)"`
	Address          string `orm:"column(address)"`
	Postcode         string `orm:"column(postcode)"`
	Companytype      string `orm:"column(companytype)"`
	Exportflag       string `orm:"column(exportflag)"`
}

func (u *COMPANY) TableName() string {
	return "skl_company_tb"
}

func AddCOMPANY(u *COMPANY) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Insert(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func AddMultiCOMPANY(u []COMPANY) error {
	o := orm.NewOrm()
	err := o.Begin()
	//_, err = o.Delete(&u)
	deletesql := "delete from skl_company_tb"
	_, err = o.Raw(deletesql).Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.InsertMulti(len(u), &u)
	insertsql := "insert into skl_company_tb(companycode,companyname,companyshortname,manager,telphone,fax,email,address,postcode,companytype,exportflag) values(?,?,?,?,?,?,?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		_, err = o.Raw(insertsql, u1.Companycode, u1.Companyname, u1.Companyshortname, u1.Manager, u1.Telphone, u1.Fax, u1.Email, u1.Address, u1.Postcode, u1.Companytype, u1.Exportflag).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}

func GetAllCOMPANY() (admins []COMPANY, err error) {
	admins = make([]COMPANY, 0)
	o := orm.NewOrm()

	sql := "select * from skl_company_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllCOMPANYoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select companycode as value,companyname as label from skl_company_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetCOMPANY(u *COMPANY) (admins []COMPANY, err error) {
	admins = make([]COMPANY, 0)
	o := orm.NewOrm()
	sql := "select * from skl_company_tb where 1=1 "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func DeleteCOMPANY(u *COMPANY) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}

func UpdateCOMPANY(u *COMPANY) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Update(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}
