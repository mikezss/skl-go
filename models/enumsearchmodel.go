package models

import (
	_ "errors"
	_ "fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ENUMSEARCH struct {
	Enumcode  string
	Enumname  string
	Pageindex int
	Pagesize  int
}

type ENUMSEARCHITEM struct {
	Enumcode  string
	Enumname  string
	Pageindex int
	Pagesize  int
}

//获得数据条数
func Getenumsearchcount(u ENUMSEARCH) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total  from skl_enum_tb a  where 1=1 "
	if u.Enumcode != "" {
		sql = sql + " and enumcode like '%" + u.Enumcode + "%'"
	}
	if u.Enumname != "" {
		sql = sql + " and enumname like '%" + u.Enumname + "%'"
	}

	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}

//获得分页数据
func Getenumsearchbypageindex(u ENUMSEARCH) (admins []ENUMSEARCHITEM, err error) {
	dbtype := Getdbtype()
	admins = make([]ENUMSEARCHITEM, 0)
	o := orm.NewOrm()

	sql := "SELECT a.*  from skl_enum_tb a where 1=1 "

	if u.Enumcode != "" {
		sql = sql + " and enumcode like '%" + u.Enumcode + "%'"
	}
	if u.Enumname != "" {
		sql = sql + " and enumname like '%" + u.Enumname + "%'"
	}
	sql = sql + " order by enumcode "
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
