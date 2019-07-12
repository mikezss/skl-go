package models

import (
	_ "errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//{{Unescapedjs .uppercomponentname}}
type DATASOURCETABLEFIELDCHILD struct {
	Id             int64  `orm:"pk;auto"` //主键，自动增长
	Datasource     string `orm:"column(datasource);size(150)"`
	Tablename      string `orm:"column(tablename);size(150)"`
	Fieldname      string `orm:"column(fieldname);size(150)"`
	Childfieldname string `orm:"column(childfieldname);size(150)"`
	Fieldtype      string `orm:"column(fieldtype)"`
	Fieldlength    int    `orm:"column(fieldlength)"`
	Isprimary      string `orm:"column(isprimary);default(0)"`
	Isnull         string `orm:"column(isnull);default(0)"`
	Defaultvalue   string `orm:"column(defaultvalue)"`
	Isauto         string `orm:"column(isauto);default(0)"`
	Isparent       string `orm:"column(isparent);default(0)"`
	Comment        string `orm:"column(comment);"` //备注
}

func (u *DATASOURCETABLEFIELDCHILD) TableName() string {
	return "skl_datasourcetablefieldchild_tb"
}

// 多字段唯一键
func (u *DATASOURCETABLEFIELDCHILD) TableUnique() [][]string {
	return [][]string{
		[]string{"Datasource", "Tablename", "Fieldname", "Childfieldname"},
	}
}
func AddDATASOURCETABLEFIELDCHILD(u *DATASOURCETABLEFIELDCHILD) error {
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

func AddMultiDATASOURCETABLEFIELDCHILD(u []DATASOURCETABLEFIELDCHILD) error {
	o := orm.NewOrm()
	err := o.Begin()

	sql := "select count(1) as ncount from skl_datasourcetablefieldchild_tb where datasource=? and tablename=? and fieldname=? and childfieldname=?"
	updatesql := "update skl_datasourcetablefieldchild_tb set tablename=?,fieldname=?,childfieldname=?,fieldtype=?,fieldlength=?,isprimary=?,isnull=?,defaultvalue=?,isauto=?,isparent=? where datasource=? and tablename=? and fieldname=? and childfieldname=?"

	insertsql := "insert into skl_datasourcetablefieldchild_tb(datasource,tablename,fieldname,childfieldname,fieldtype,fieldlength,isprimary,isnull,defaultvalue,isauto,isparent) values(?,?,?,?,?,?,?,?,?,?,?)"

	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		ncount := 0
		err = o.Raw(sql, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Childfieldname).QueryRow(&ncount)
		if ncount > 0 {
			_, err = o.Raw(updatesql, u1.Tablename, u1.Fieldname, u1.Childfieldname, u1.Fieldtype, u1.Fieldlength, u1.Isprimary, u1.Isnull, u1.Defaultvalue, u1.Isauto, u1.Isparent, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Childfieldname).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		} else {
			_, err = o.Raw(insertsql, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Childfieldname, u1.Fieldtype, u1.Fieldlength, u1.Isprimary, u1.Isnull, u1.Defaultvalue, u1.Isauto, u1.Isparent).Exec()
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

func GetAllDATASOURCETABLEFIELDCHILD() (admins []DATASOURCETABLEFIELDCHILD, err error) {
	admins = make([]DATASOURCETABLEFIELDCHILD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefieldchild_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得数据条数
func Getdatasourcetablefieldchildcount() (page PAGE, err error) {
	o := orm.NewOrm()
	sql := "SELECT count(1) as total  from skl_datasourcetablefieldchild_tb a"
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)
	return page, err
}

//获得分页数据
func Getdatasourcetablefieldchildbypageindex(l PAGE) (admins []DATASOURCETABLEFIELDCHILD, err error) {
	dbtype := Getdbtype()
	admins = make([]DATASOURCETABLEFIELDCHILD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefieldchild_tb where 1=1 "

	var limitstr string = " limit "
	if dbtype == "postgres" {
		limitstr = limitstr + strconv.Itoa(l.Pagesize) + " offset " + strconv.Itoa((l.Pageindex-1)*l.Pagesize)

	} else if dbtype == "mysql" {
		limitstr = limitstr + strconv.Itoa((l.Pageindex-1)*l.Pagesize) + "," + strconv.Itoa(l.Pagesize)

	} else {
		limitstr = limitstr + strconv.Itoa((l.Pageindex-1)*l.Pagesize) + "," + strconv.Itoa(l.Pagesize)

	}
	sql = sql + limitstr

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetAllDATASOURCETABLEFIELDCHILDoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefieldchild_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetDATASOURCETABLEFIELDCHILD(u *DATASOURCETABLEFIELDCHILD) (admins []DATASOURCETABLEFIELDCHILD, err error) {
	admins = make([]DATASOURCETABLEFIELDCHILD, 0)
	o := orm.NewOrm()
	sql := "select * from skl_datasourcetablefieldchild_tb where 1=1 "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetDATASOURCETABLEFIELDCHILDbyid(u *DATASOURCETABLEFIELDCHILD) (admins []DATASOURCETABLEFIELDCHILD, err error) {
	admins = make([]DATASOURCETABLEFIELDCHILD, 0)
	o := orm.NewOrm()
	sql := "select * from skl_datasourcetablefieldchild_tb where id=? "

	_, err = o.Raw(sql, u.Id).QueryRows(&admins)

	return admins, err
}
func GetDATASOURCETABLEFIELDCHILDbyfield(u *DATASOURCETABLEFIELDCHILD) (admins []DATASOURCETABLEFIELDCHILD, err error) {
	admins = make([]DATASOURCETABLEFIELDCHILD, 0)
	o := orm.NewOrm()
	sql := "select * from skl_datasourcetablefieldchild_tb where datasource=? and tablename=? and fieldname=? "

	_, err = o.Raw(sql, u.Datasource, u.Tablename, u.Fieldname).QueryRows(&admins)

	return admins, err
}

func DeleteDATASOURCETABLEFIELDCHILD(u *DATASOURCETABLEFIELDCHILD) error {

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

func UpdateDATASOURCETABLEFIELDCHILD(u *DATASOURCETABLEFIELDCHILD) error {

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
func GetmetaDATASOURCETABLEFIELDCHILD(dstf DATASOURCETABLEFIELDCHILDSEARCH) (err error) {
	admins := make([]DATASOURCETABLEFIELDCHILD, 0)

	ds, _ := GetDATASOURCEBYID(DATASOURCE{Datasource: dstf.Datasource})

	switch ds.Dbtype {
	case "mongodb":
		sourcemongoconn, err := GetMongoConn(ds)

		if err == nil {
			fieldinfomap, err := sourcemongoconn.Getcollectionfieldinfochild(ds.Schema, dstf.Tablename, dstf.Fieldname, dstf.Maptype)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("sourcemongoconn.Getcollectionfieldinfochild()==>" + err.Error())
				return err
			}
			//fmt.Println(fieldinfomap)
			for key, value := range fieldinfomap {
				admins = append(admins, DATASOURCETABLEFIELDCHILD{Datasource: dstf.Datasource, Tablename: dstf.Tablename, Fieldname: dstf.Fieldname, Childfieldname: key, Fieldtype: value.Fieldtype, Fieldlength: value.Fieldlength, Isprimary: value.Isprimary, Isauto: value.Isauto, Isnull: value.Isnull, Isparent: value.Isparent, Defaultvalue: value.Defaultvalue})
			}

			err = AddMultiDATASOURCETABLEFIELDCHILD(admins)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("AddDATASOURCETABLEFIELDCHILD()==>" + err.Error())
				return err
			}

		} else {

			//defer sourcemongoconn.Session.Close()
			fmt.Println(err)
			Getlog().Error("GetMongoConn()==>" + err.Error())
			return err
		}
		break

	}

	return nil
}
