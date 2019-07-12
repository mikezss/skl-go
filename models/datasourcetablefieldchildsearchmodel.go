package models

import (
	_ "errors"
	_ "fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type DATASOURCETABLEFIELDCHILDSEARCH struct {
	Sourcetargettype string
	Datasource       string
	Tablename        string
	Fieldname        string
	Pageindex        int
	Pagesize         int
	Maptype          string
}

type DATASOURCETABLEFIELDCHILDSEARCHITEM struct {
	Id               int64
	Sourcetargettype string
	Datasource       string
	Tablename        string
	Fieldname        string
	Childfieldname   string
	Edit             string
	Delete           string
	Pageindex        int
	Pagesize         int
}

//获得数据条数
func Getdatasourcetablefieldchildsearchcount(u DATASOURCETABLEFIELDCHILDSEARCH) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total from skl_datasourcetablefieldchild_tb a inner join skl_datasource_tb b on a.datasource=b.datasource where 1=1 "

	if u.Sourcetargettype != "" {
		sql = sql + " and b.sourcetargettype='" + u.Sourcetargettype + "'"
	}
	if u.Datasource != "" {
		sql = sql + " and a.datasource='" + u.Datasource + "'"
	}
	if u.Tablename != "" {
		sql = sql + " and a.tablename='" + u.Tablename + "'"
	}
	if u.Fieldname != "" {
		sql = sql + " and a.fieldname='" + u.Fieldname + "'"
	}

	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}

//获得分页数据
func Getdatasourcetablefieldchildsearchbypageindex(u DATASOURCETABLEFIELDCHILDSEARCH) (admins []DATASOURCETABLEFIELDCHILDSEARCHITEM, err error) {
	dbtype := Getdbtype()
	admins = make([]DATASOURCETABLEFIELDCHILDSEARCHITEM, 0)
	o := orm.NewOrm()

	sql := "SELECT a.* from skl_datasourcetablefieldchild_tb a inner join skl_datasource_tb b on a.datasource=b.datasource where 1=1 "

	if u.Sourcetargettype != "" {
		sql = sql + " and b.sourcetargettype='" + u.Sourcetargettype + "'"
	}
	if u.Datasource != "" {
		sql = sql + " and a.datasource='" + u.Datasource + "'"
	}
	if u.Tablename != "" {
		sql = sql + " and a.tablename='" + u.Tablename + "'"
	}
	if u.Fieldname != "" {
		sql = sql + " and a.fieldname='" + u.Fieldname + "'"
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
