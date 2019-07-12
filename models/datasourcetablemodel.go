package models

import (
	_ "errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//{{Unescapedjs .uppercomponentname}}
type DATASOURCETABLE struct {
	Id         int64  `orm:"pk;auto"` //主键，自动增长
	Datasource string `orm:"column(datasource)"`
	Tablename  string `orm:"column(tablename)"`
}

func (u *DATASOURCETABLE) TableName() string {
	return "skl_datasourcetable_tb"
}

// 多字段唯一键
func (u *DATASOURCETABLE) TableUnique() [][]string {
	return [][]string{
		[]string{"Datasource", "Tablename"},
	}
}

func AddDATASOURCETABLE(u *DATASOURCETABLE) error {
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

func AddMultiDATASOURCETABLE(u []DATASOURCETABLE) error {
	o := orm.NewOrm()
	err := o.Begin()

	sql := "select count(1) as ncount from skl_datasourcetable_tb where datasource=? and tablename=?"
	updatesql := "update skl_datasourcetable_tb set datasource=?,tablename=? where datasource=? and tablename=?"

	insertsql := "insert into skl_datasourcetable_tb(datasource,tablename) values(?,?)"

	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		ncount := 0
		err = o.Raw(sql, u1.Datasource, u1.Tablename).QueryRow(&ncount)
		if ncount > 0 {
			_, err = o.Raw(updatesql, u1.Datasource, u1.Tablename, u1.Datasource, u1.Tablename).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		} else {
			_, err = o.Raw(insertsql, u1.Datasource, u1.Tablename).Exec()
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

func GetDATASOURCETABLEBYID(e DATASOURCETABLE) (admin DATASOURCETABLE, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_datasourcetable_tb where datasource=? and tablename=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Datasource, e.Tablename).QueryRow(&admin)

	return admin, err
}
func GetDATASOURCETABLEBYdatasource(datasource string) (admins []DATASOURCETABLE, err error) {
	admins = make([]DATASOURCETABLE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetable_tb where datasource=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, datasource).QueryRows(&admins)

	return admins, err
}

func GetAllDATASOURCETABLE() (admins []DATASOURCETABLE, err error) {
	admins = make([]DATASOURCETABLE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetable_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetsourceDATASOURCETABLE() (admins []DATASOURCETABLE, err error) {
	admins = make([]DATASOURCETABLE, 0)
	o := orm.NewOrm()

	sql := "select a.* from skl_datasourcetable_tb a inner join skl_datasource_tb b on a.datasource=b.datasource where sourcetargettype='source'"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

//获得数据条数
func Getdatasourcetablecount() (page PAGE, err error) {
	o := orm.NewOrm()
	sql := "SELECT count(1) as total  from skl_datasourcetable_tb a"
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)
	return page, err
}

//获得分页数据
func Getdatasourcetablebypageindex(l PAGE) (admins []DATASOURCETABLE, err error) {
	dbtype := Getdbtype()
	admins = make([]DATASOURCETABLE, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetable_tb where 1=1 "

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

func GetAllDATASOURCETABLEoptions(u *DATASOURCE) (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select tablename as value,tablename as label from skl_datasourcetable_tb where datasource=?"

	_, err = o.Raw(sql, u.Datasource).QueryRows(&admins)

	return admins, err
}

func GetDATASOURCETABLE(u *DATASOURCETABLE) (admins []DATASOURCETABLE, err error) {
	admins = make([]DATASOURCETABLE, 0)
	o := orm.NewOrm()
	sql := "select * from skl_datasourcetable_tb where 1=1 "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func DeleteDATASOURCETABLE(u *DATASOURCETABLE) error {

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

func UpdateDATASOURCETABLE(u *DATASOURCETABLE) error {

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
func GetmetaDATASOURCETABLE() (err error) {
	reg := regexp.MustCompile(`.*copy$`)
	reg2 := regexp.MustCompile(`.*[0-9]{8}.*`)
	admins := make([]DATASOURCETABLE, 0)
	dsarr, err := GetAllDATASOURCE()
	if err != nil {
		fmt.Println(err)
		Getlog().Error("GetAllDATASOURCE==>" + err.Error())
		return err
	}
	for _, ds := range dsarr {
		switch ds.Dbtype {
		case "mongodb":
			Getlog().Debug("ds.Dbtype==>" + ds.Dbtype)
			sourcemongoconn, err := GetMongoConn(ds)
			var tablenamesarr []string
			if err == nil {
				tablenamesarr, err = sourcemongoconn.Getcollectionnames(ds.Schema)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("sourcemongoconn.Getcollectionnames(ds.Schema)==>" + err.Error())
					return err
				}
				Getlog().Debug("sourcemongoconn.Getcollectionnames==>" + orm.ToStr(tablenamesarr))
				for _, tablename := range tablenamesarr {
					if reg.MatchString(tablename) || reg2.MatchString(tablename) {
						continue
					}
					admins = append(admins, DATASOURCETABLE{Datasource: ds.Datasource, Tablename: tablename})
				}
				err = AddMultiDATASOURCETABLE(admins)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("AddMultiDATASOURCETABLE(admins)==>" + err.Error())
					return err
				}
			} else {
				fmt.Println(err)
				Getlog().Error("GetMongoConn(ds)==>" + err.Error())
				return err
			}
			break
		case "mysql":
			admins, err = Getmysqltablenames(ds)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("Getmysqltablenames(ds)==>" + err.Error())
				return err
			}
			Getlog().Debug("Getmysqltablenames==>" + orm.ToStr(admins))
			err = AddMultiDATASOURCETABLE(admins)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("AddMultiDATASOURCETABLE(admins)==>" + err.Error())
				return err
			}
			break
		}
	}

	return err
}
func GetmetaDATASOURCETABLEbydatasource(ds2 DATASOURCE) (err error) {
	reg := regexp.MustCompile(`.*copy$`)
	reg2 := regexp.MustCompile(`.*[0-9]{8}.*`)
	admins := make([]DATASOURCETABLE, 0)
	ds, err := GetDATASOURCEBYID(ds2)
	if err != nil {
		fmt.Println(err)
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	switch ds.Dbtype {
	case "mongodb":
		sourcemongoconn, err := GetMongoConn(ds)
		var tablenamesarr []string
		if err == nil {
			tablenamesarr, err = sourcemongoconn.Getcollectionnames(ds.Schema)
			if err != nil {
				fmt.Println(err)
				return err
			}
			Getlog().Debug("sourcemongoconn.Getcollectionnames==>" + orm.ToStr(tablenamesarr))
			for _, tablename := range tablenamesarr {
				if reg.MatchString(tablename) || reg2.MatchString(tablename) {
					continue
				}
				admins = append(admins, DATASOURCETABLE{Datasource: ds.Datasource, Tablename: tablename})
			}
			err = AddMultiDATASOURCETABLE(admins)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("AddMultiDATASOURCETABLE==>" + err.Error())
				return err
			}
		} else {
			fmt.Println(err)
			return err
		}
		break
	case "mysql":
		admins, err = Getmysqltablenames(ds)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("Getmysqltablenames==>" + err.Error())
			return err
		}
		Getlog().Debug("Getmysqltablenames==>" + orm.ToStr(admins))
		err = AddMultiDATASOURCETABLE(admins)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("AddMultiDATASOURCETABLE==>" + err.Error())
			return err
		}
		break
	case "oracle":
		admins, err = Getoracletablenames2(ds)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("Getoracletablenames==>" + err.Error())
			return err
		}
		Getlog().Debug("Getoracletablenames==>" + orm.ToStr(admins))
		err = AddMultiDATASOURCETABLE(admins)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("AddMultiDATASOURCETABLE==>" + err.Error())
			return err
		}
		break
	case "sqlserver":
		admins, err = Getmssqltablenames(ds)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("Getmssqltablenames==>" + err.Error())
			return err
		}
		Getlog().Debug("Getmssqltablenames==>" + orm.ToStr(admins))
		err = AddMultiDATASOURCETABLE(admins)
		if err != nil {
			fmt.Println(err)
			Getlog().Error("AddMultiDATASOURCETABLE==>" + err.Error())
			return err
		}
		break
	}

	return err
}
