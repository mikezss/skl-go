package models

import (
	"database/sql"
	_ "errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/config"

	"database/sql/driver"

	"io/ioutil"
	"reflect"

	"github.com/tealeg/xlsx"

	"log"

	"encoding/json"

	"bytes"

	"math/rand"

	"crypto/sha256"
	"encoding/hex"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mozillazg/go-pinyin"
	_ "github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//流程定义fi_template_tb表
const (
	formatTime     = "15:04:05"
	formatDate     = "2006-01-02"
	formatDateTime = "2006-01-02 15:04:05"
)
const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

type COMPONENT struct {
	Componentname  string
	Parentid       string
	Title          string
	Buttons        string
	Style          string
	Gutter         string
	Colcount       string
	Componentlevel string
	Godirectory    string
	Ngdirectory    string
}

type CMN_FILEINFO_TB struct {
	Filename       string    `orm:"pk;column(filename)"`
	Filesize       int64     `orm:"column(filesize)"`
	Fileext        string    `orm:"column(fileext)"`
	Filepath       string    `orm:"column(filepath);null"`
	Filerights     string    `orm:"column(filerights);null"`
	Expired        time.Time `orm:"column(expired);null"`
	Downloadstatus string    `orm:"column(downloadstatus);default(0)"`
	Createuser     string    `orm:"column(createuser);null"`
	Createtime     time.Time `orm:"column(createtime);null"`
	Updateuser     string    `orm:"column(updateuser);null"`
	Updatetime     time.Time `orm:"column(updatetime);null"`
}
type FILELIST struct {
	Uid      int
	Name     string
	Type     string
	Size     int64
	Status   string
	Response string
	Url      string
}
type OPTIONS struct {
	Value   string `json:"value"`
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}

func (u *CMN_FILEINFO_TB) TableName() string {
	return "cmn_fileinfo_tb"
}

func AddCMN_FILEINFO_TB(u *CMN_FILEINFO_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Insert(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func DeleteCMN_FILEINFO_TB(u *CMN_FILEINFO_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func GetAllCMN_FILEINFO_TB() (admins []CMN_FILEINFO_TB, err error) {
	admins = make([]CMN_FILEINFO_TB, 0)
	o := orm.NewOrm()

	sql := "select filerights,expired,downloadstatus from cmn_fileinfo_tb "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetCMN_FILEINFO_TB(u *CMN_FILEINFO_TB) (admins []CMN_FILEINFO_TB, err error) {
	admins = make([]CMN_FILEINFO_TB, 0)
	o := orm.NewOrm()
	sql := "select * from cmn_fileinfo_tb where 1=1 "

	if u.Filerights != "" {
		sql = sql + " and filerights='" + u.Filerights + "'"
	}

	//if u.Expired != "" {
	//sql = sql + " and expired='" + u.Expired + "'"
	//}

	if u.Downloadstatus != "" {
		sql = sql + " and downloadstatus='" + u.Downloadstatus + "'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

type CMN_TEMPLATE_TB struct {
	Templateid   string `orm:"pk;column(templateid)"`
	Templatename string `orm:"column(templatename)"`
	Formate      string `orm:"column(formate)"`
}

func (u *CMN_TEMPLATE_TB) TableName() string {
	return "cmn_template_tb"
}

type CMN_EXPORTTEMPLATE_TB struct {
	Exporttemplateid   string `orm:"pk;column(exporttemplateid)"`
	Exporttemplatename string `orm:"column(exporttemplatename)"`
	Templateid         string `orm:"column(templateid)"`
	Exporttitle        string `orm:"column(exporttitle)"`
	Exporttype         string `orm:"column(exporttype)"`
	Exportsql          string `orm:"column(exportsql)"`
	Exportfilepath     string `orm:"column(exportfilepath)"`
	Exportfilename     string `orm:"column(exportfilename)"`
	Accessmethod       string `orm:"column(accessmethod)"`
	Emailtitle         string `orm:"column(emailtitle)"`
}

func (u *CMN_EXPORTTEMPLATE_TB) TableName() string {
	return "cmn_exporttemplate_tb"
}

type CMN_TEMPLATEITEM_TB struct {
	Templateid   string `orm:"pk;column(templateid)"`
	Colid        string `orm:"pk;column(colid)"`
	Colname      string `orm:"column(colname)"`
	Coltype      string `orm:"column(coltype)"`
	Required     string `orm:"column(required)"`
	Length       string `orm:"column(length)"`
	Accuracy     string `orm:"column(accuracy)"`
	Defaultvalue string `orm:"column(defaultvalue)"`
	Pretype      string `orm:"column(pretype)"`
	Sep          string `orm:"column(sep)"`
}

func (u *CMN_TEMPLATEITEM_TB) TableName() string {
	return "cmn_templateitem_tb"
}
func AddCMN_TEMPLATE_TB(template1 *CMN_TEMPLATE_TB, templateitems []CMN_TEMPLATEITEM_TB) error {
	rows := 0
	db, _ := orm.GetDB("default")
	tr, _ := db.Begin()
	querysql := "select *  from cmn_template_tb where templateid=? "
	querysql = ConvertSQL(querysql, Getdbtype())
	result, err := tr.Query(querysql, template1.Templateid)
	if result.Next() {
		rows = 1
		result.Close()
	}

	fmt.Println(rows)
	if rows > 0 {
		deletesql := "delete  from cmn_template_tb where templateid=? "
		deletesql = ConvertSQL(deletesql, Getdbtype())
		_, err = tr.Exec(deletesql, template1.Templateid)

		if err != nil {
			fmt.Println("delete cmn_template_tb fail:==>")
			fmt.Println(err)
			err = tr.Rollback()
		}
	}

	sql := "insert into cmn_template_tb values(?,?,?) "
	sql = ConvertSQL(sql, Getdbtype())
	preparestatment, _ := tr.Prepare(sql)
	_, err = preparestatment.Exec(template1.Templateid, template1.Templatename, template1.Formate)
	if err != nil {
		fmt.Println("insert into cmn_template_tb values(?,?,?) ")
		fmt.Println(err)
		err2 := tr.Rollback()
		if err2 != nil {
			err = err2
		}
	}

	err = AddMultiCMN_TEMPLATEITEM_TB(tr, rows, template1.Templateid, templateitems)
	if err == nil {
		err = tr.Commit()
	}
	//defer db.Close()
	return err
}
func AddMultiCMN_TEMPLATEITEM_TB(tr *sql.Tx, rows int, templateid string, templateitems []CMN_TEMPLATEITEM_TB) error {
	var deletesql string = ""
	var insertsql string = ""

	//fuck有一个陷阱，不同的数据库写法不同，坑爹。
	//MySQL               PostgreSQL            Oracle
	//WHERE col = ?       WHERE col = $1        WHERE col = :col
	//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		fmt.Println(err)
	}
	dbtype := iniconf.String("dbtype")
	switch dbtype {
	case "mysql":
		deletesql = "delete  from cmn_templateitem_tb where templateid=? "
		insertsql = "insert into cmn_templateitem_tb(templateid,colid,colname,coltype,required,length,accuracy,defaultvalue,pretype,sep) values(?,?,?,?,?,?,?,?,?,?)"
		break
	case "postgres":
		//deletesql = "delete  from cmn_flowuser_tb where flowid=$1"
		insertsql = "insert into cmn_flowuser_tb(flowid,varyid,varyname,varyvalue,pretype,sep) values($1,$2,$3,$4,$5,$6)"
		break
	case "sqlite3":
		//deletesql = "delete  from cmn_flowuser_tb where flowid=?"
		insertsql = "insert into cmn_flowuser_tb(flowid,varyid,varyname,varyvalue,pretype,sep) values(?,?,?,?,?,?)"
		break
	case "oracle":
		//deletesql = "delete  from cmn_flowuser_tb where flowid=:val1"
		insertsql = "insert into cmn_flowuser_tb(flowid,varyid,varyname,varyvalue,pretype,sep) values(:val1,:val2,:val3,:val4,:val5,:val6)"
		break
	}

	if rows > 0 {
		_, err = tr.Exec(deletesql, templateid)

		if err != nil {
			fmt.Println("delete cmn_templateitem_tb fail:==>")
			fmt.Println(err)
			err = tr.Rollback()
		}
	}

	for _, templateitem := range templateitems {

		_, err = tr.Exec(insertsql, templateitem.Templateid, templateitem.Colid, templateitem.Colname, templateitem.Coltype, templateitem.Required, templateitem.Length, templateitem.Accuracy, templateitem.Defaultvalue, templateitem.Pretype, templateitem.Sep)
		if err != nil {

			err = tr.Rollback()
		}
	}

	return err
}
func GetCMN_TEMPLATEITEM_TB(templateid string) (templateitems []CMN_TEMPLATEITEM_TB, err error) {
	var sql string
	templateitems = make([]CMN_TEMPLATEITEM_TB, 0)
	o := orm.NewOrm()
	if templateid != "" {
		sql = "select * from cmn_templateitem_tb where templateid=? order by colid"

	} else {
		return nil, nil
	}
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, templateid).QueryRows(&templateitems)

	return templateitems, err
}
func GetCMN_TEMPLATE_TB() (templates []CMN_TEMPLATE_TB, err error) {
	var sql string
	templates = make([]CMN_TEMPLATE_TB, 0)
	o := orm.NewOrm()

	sql = "select * from cmn_template_tb "

	_, err = o.Raw(sql).QueryRows(&templates)

	return templates, err
}
func GetCMN_TEMPLATE_TBbyid(templateid string) (template CMN_TEMPLATE_TB, err error) {
	var sql string

	o := orm.NewOrm()

	sql = "select * from cmn_template_tb where templateid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, templateid).QueryRow(&template)

	return template, err
}
func DeleteCMN_TEMPLATE_TB(templateid string) error {

	db, _ := orm.GetDB("default")
	tr, _ := db.Begin()

	deletesql := "delete  from cmn_template_tb where templateid=? "
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err := tr.Exec(deletesql, templateid)

	if err != nil {
		fmt.Println("delete cmn_template_tb fail:==>")
		fmt.Println(err)
		err = tr.Rollback()
	}
	deletesql = "delete  from cmn_templateitem_tb where templateid=? "
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err = tr.Exec(deletesql, templateid)

	if err != nil {
		fmt.Println("delete cmn_templateitem_tb fail:==>")
		fmt.Println(err)
		err = tr.Rollback()
	}
	if err == nil {
		err = tr.Commit()
	}
	//defer db.Close()
	return err
}
func DeleteCMN_EXPORTTEMPLATE_TB(u *CMN_EXPORTTEMPLATE_TB) error {

	db, _ := orm.GetDB("default")
	tr, _ := db.Begin()

	deletesql := "delete  from cmn_exporttemplate_tb where exporttemplateid=? "
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err := tr.Exec(deletesql, u.Exporttemplateid)

	if err != nil {
		fmt.Println("delete cmn_exporttemplate_tb fail:==>")
		fmt.Println(err)
		err = tr.Rollback()
	}

	if err == nil {
		err = tr.Commit()
	}
	//defer db.Close()
	return err
}
func AddCMN_EXPORTTEMPLATE_TB(u *CMN_EXPORTTEMPLATE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	}

	_, err = o.Insert(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func DeleteCMN_EXPORTTEMPLATE_TB2(u *CMN_EXPORTTEMPLATE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func GetCMN_EXPORTTEMPLATE_TB(Exporttemplateid string) (templates []CMN_EXPORTTEMPLATE_TB, err error) {
	var sql string
	templates = make([]CMN_EXPORTTEMPLATE_TB, 0)
	o := orm.NewOrm()
	sql = "select * from cmn_exporttemplate_tb "
	if Exporttemplateid != "" {
		sql = "select * from cmn_exporttemplate_tb where exporttemplateid='" + Exporttemplateid + "'"
	}

	_, err = o.Raw(sql).QueryRows(&templates)

	return templates, err
}

type CMN_IMPORTTEMPLATE_TB struct {
	Importtemplateid   string `orm:"pk;column(importtemplateid)"`
	Importtemplatename string `orm:"column(importtemplatename)"`
	Templateid         string `orm:"column(templateid)"`
	Importtable        string `orm:"column(importtable)"`
	Importtype         string `orm:"column(importtype)"`
	Importsql          string `orm:"column(importsql)"`
}

func (u *CMN_IMPORTTEMPLATE_TB) TableName() string {
	return "cmn_importtemplate_tb"
}

func GetCMN_IMPORTTEMPLATE_TB(Importtemplateid string) (templates []CMN_IMPORTTEMPLATE_TB, err error) {
	var sql string
	templates = make([]CMN_IMPORTTEMPLATE_TB, 0)
	o := orm.NewOrm()
	sql = "select * from cmn_importtemplate_tb "
	if Importtemplateid != "" {
		sql = "select * from cmn_importtemplate_tb where importtemplateid='" + Importtemplateid + "'"
	}

	_, err = o.Raw(sql).QueryRows(&templates)

	return templates, err
}
func DeleteCMN_IMPORTTEMPLATE_TB(u *CMN_IMPORTTEMPLATE_TB) error {

	db, _ := orm.GetDB("default")
	tr, _ := db.Begin()

	deletesql := "delete  from cmn_importtemplate_tb where importtemplateid=? "
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err := tr.Exec(deletesql, u.Importtemplateid)

	if err != nil {
		fmt.Println("delete cmn_importtemplate_tb fail:==>")
		fmt.Println(err)
		err = tr.Rollback()
	}

	if err == nil {
		err = tr.Commit()
	}
	//defer db.Close()
	return err
}
func AddCMN_IMPORTTEMPLATE_TB(u *CMN_IMPORTTEMPLATE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	}

	_, err = o.Insert(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func DeleteCMN_IMPORTTEMPLATE_TB2(u *CMN_IMPORTTEMPLATE_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		//fmt.Println(err)
		err2 := o.Rollback()
		if err2 != nil {
			err = err2
		}
	} else {
		err = o.Commit()
	}
	return err
}
func Getmetadata(tablename string) {
	//var cols interface{}

	db, _ := orm.GetDB("default")
	st := db.Stats()
	fmt.Println(st.OpenConnections)
	sql := "select * from " + tablename + " limit 0,1"
	fmt.Println(sql)
	rows, _ := db.Query(sql)

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	for _, col := range cols {
		fmt.Println(col)
	}

	var drv []driver.Value
	if rows.Next() {
		//drv = rows.GetLastcols()
		//fmt.Println(drv)
		for _, sv := range drv {
			//fmt.Println(sv)

			switch sv.(type) {
			case string:
				fmt.Println("string")
			case []byte:
				fmt.Println("string")
			case time.Time:
				fmt.Println("time")
			case bool:
				fmt.Println("bool")
			case float64:
				fmt.Println("float64")
			case int64:
				fmt.Println("int64")
			default:
				fmt.Println("string")

			}
			//			rv := reflect.ValueOf(sv)
			//			switch rv.Kind() {
			//			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			//				fmt.Println("Int")
			//			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			//				fmt.Println("Uint")
			//			case reflect.Float64:
			//				fmt.Println("Float64")
			//			case reflect.Float32:
			//				fmt.Println("Float32")
			//			case reflect.Bool:
			//				fmt.Println("Bool")
			//			}
		}
	}

}

func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	fmt.Println(rv.Kind())
	fmt.Println(rv.Type())
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

type TABLEINF struct {
	Tables string `json:"text"`
}
type COLINF struct {
	Id           string `json:"id"`
	Name         string `json:"text"`
	Type         string `json:"coltype"`
	Length       string `json:"length"`
	Isnull       string `json:"isnull"`
	Isprimary    string `json:"isprimary"`
	Isautoinc    string `json:"isautoinc"`
	Defaultvalue string `json:"defaultvalue"`
	Comment      string `json:"comment"`
}

func GetTABLEINF() (tableinfs []TABLEINF, err error) {
	var sql string
	tableinfs = make([]TABLEINF, 0)
	o := orm.NewOrm()
	sql = "show tables "

	_, err = o.Raw(sql).QueryRows(&tableinfs)

	db, _ := orm.GetDB("default")
	rows, err := db.Query(sql)

	cols, err := rows.Columns()
	fmt.Println(cols)
	tableinfss := make([]string, 0)
	err = rows.Scan(&tableinfss)
	fmt.Println(tableinfss)
	//values := rows.GetLastcols()
	//fmt.Println(values)
	return tableinfs, err
}
func Convert2time(input string) time.Time {
	var cvtvalue time.Time
	var err error
	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST
	if input == "" {
		input = "9999-01-01"
	}
	if len(input) == 10 {

		cvtvalue, err = time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Println(err)
			cvtvalue = time.Now()
			cvtvalue.Format("2006-01-02")
		}

	}
	if len(input) >= 19 {
		input = input[:19]
		cvtvalue, err = time.Parse("2006-01-02 15:04:05", input)
		if err != nil {
			fmt.Println(err)
			cvtvalue = time.Now()
			cvtvalue.Format("2006-01-02 15:04:05")
		}

	}

	return cvtvalue
}
func GetYYYYMMDDstring() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	return time.Now().Format("2006-01-02")

}
func GetHHmmssstring() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	return time.Now().Format("03:04:05")

}
func GetYMDtime() time.Time {
	var cvtvalue time.Time

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue = time.Now()
	cvtvalue.Format("2006-01-02")

	return cvtvalue
}
func GetYYYY() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("2006")

}
func GetMM() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("01")

}
func GetDD() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("02")

}
func GetYYYYMM() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("200601")

}

func GetYYYYMMDD() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("2006-01-02")

}

//返回20181118175858格式的数据，即2018年11月18日17点58分58秒
func GetYYYYMMDDHHMMSS() string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	cvtvalue := time.Now()
	return cvtvalue.Format("20060102030405")

}
func Convert2YYYYMMDD(times time.Time) string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	return times.Format("2006-01-02")

}
func Convert2YYYYMM(times time.Time) string {

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	return times.Format("2006-01")

}
func Convert2int64(input string) int64 {
	var cvtvalue int64
	var err error
	if input == "" {
		input = "-9999"
	}
	cvtvalue, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		cvtvalue = -9999
	}

	return cvtvalue
}
func Convert2float64(input string) float64 {
	var cvtvalue float64
	var err error

	cvtvalue, err = strconv.ParseFloat(input, 64)
	if err != nil {
		cvtvalue = 0.0
	}

	return cvtvalue
}
func Convert2bool(input string) bool {
	var cvtvalue bool
	var err error

	cvtvalue, err = strconv.ParseBool(input)
	if err != nil {
		cvtvalue = false
	}

	return cvtvalue
}
func Outputconvertleft(datatype string) string {
	switch datatype {
	case "time.Time":
		return "models.Convert2time("
	case "int64":
		return "models.Convert2int64("
	case "float64":
		return "models.Convert2float64("
	case "bool":
		return "models.Convert2bool("
	default:
		return ""
	}
}
func Outputconvertright(datatype string) string {
	switch datatype {
	case "string":
		return ""
	default:
		return ")"
	}
}
func Getexportfileinfomap(sql string) (exportfileinfmap []orm.Params, err error) {
	var expfmp []orm.Params
	expfmp = make([]orm.Params, 0)
	o := orm.NewOrm()
	_, err = o.Raw(sql).Values(&expfmp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return expfmp, nil
}

//case "mysql":
//deletesql = "delete  from cmn_flowaction_tb where flowid=? and taskid=?"
//insertsql = "insert into cmn_flowaction_tb(flowid,taskid,actionid,nexttaskid,backtotaskid,taskstatus,dispatcher) values(?,?,?,?,?,?,?)"

//case "postgres":
//deletesql = "delete  from cmn_flowaction_tb where flowid=$1 and taskid=$2"
//insertsql = "insert into cmn_flowaction_tb(flowid,taskid,actionid,nexttaskid,backtotaskid,taskstatus,dispatcher) values($1,$2,$3,$4,$5,$6,$7)"

//case "sqlite3":
//deletesql = "delete  from cmn_flowaction_tb where flowid=? and taskid=?"
//insertsql = "insert into cmn_flowaction_tb(flowid,taskid,actionid,nexttaskid,backtotaskid,taskstatus,dispatcher) values(?,?,?,?,?,?,?)"

//case "oracle":
//deletesql = "delete  from cmn_flowaction_tb where flowid=:val1 and taskid=:val2"
//insertsql = "insert into cmn_flowaction_tb(flowid,taskid,actionid,nexttaskid,backtotaskid,taskstatus,dispatcher) values(:val1,:val3,:val3,:val4,:val5,:val6,:val7)"

func ConvertSQL(sql string, databasetype string) string {
	//mysql:DATE_FORMAT(calltime,'%Y-%m-%d')
	//sqlite3:strftime('%Y-%m-%d',calltime)
	var cvtsql string
	symbol := "?"
	if databasetype == "oracle" {

		symbol = ":val"

		idx := strings.Index(sql, "?")
		if idx == -1 {
			return sql
		}
		r := strings.Split(sql, "?")
		length := len(r)
		if length == 1 {
			cvtsql = cvtsql + r[0]
			cvtsql = cvtsql + symbol + strconv.Itoa(1)
			return cvtsql
		}
		lastidx := strings.LastIndex(sql, "?")

		for i := 0; i < length; i++ {
			if i > 0 {
				cvtsql = cvtsql + symbol + strconv.Itoa(i)
			}
			cvtsql = cvtsql + r[i]
			if lastidx == len(sql)-1 && i == length-1 {
				cvtsql = cvtsql + symbol + strconv.Itoa(i)
			}
		}
	} else {
		if databasetype == "sqlite3" {
			reg := regexp.MustCompile(`DATE_FORMAT\(calltime,'\%Y-\%m-\%d'\)`)
			cvtsql = reg.ReplaceAllString(sql, "strftime('%Y-%m-%d',calltime)")
			reg = regexp.MustCompile(`DATE_FORMAT\(attdate,'\%Y-\%m-\%d'\)`)
			cvtsql = reg.ReplaceAllString(cvtsql, "strftime('%Y-%m-%d',attdate)")
			reg = regexp.MustCompile(`DATE_FORMAT\(attdate,'\%Y-\%m'\)`)
			cvtsql = reg.ReplaceAllString(cvtsql, "strftime('%Y-%m',attdate)")
			reg = regexp.MustCompile(`DATE_FORMAT\(calltime,'\%Y-\%m'\)`)
			cvtsql = reg.ReplaceAllString(cvtsql, "strftime('%Y-%m',calltime)")
			reg = regexp.MustCompile(`DATE_FORMAT\(flowstarttime,'\%Y-\%m-\%d'\)`)
			cvtsql = reg.ReplaceAllString(cvtsql, "strftime('%Y-%m-%d',flowstarttime)")
			reg = regexp.MustCompile(`DATE_FORMAT\(flowfinishtime,'\%Y-\%m-\%d'\)`)
			cvtsql = reg.ReplaceAllString(cvtsql, "strftime('%Y-%m-%d',flowfinishtime)")
		} else {
			cvtsql = sql
		}
	}

	return cvtsql
}
func Getdbtype() string {
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	dbtype := iniconf.String("dbtype")
	return dbtype
}
func SQLBRACKET2SPACE(sql string, databasetype string) string {
	var cvtsql string

	if databasetype == "postgres" || databasetype == "mysql" || databasetype == "sqlserver" {
		return sql
	}
	reg := regexp.MustCompile(`\(|\)`)
	cvtsql = reg.ReplaceAllString(sql, " ")
	return cvtsql
}

//文件转换为字符串
func Readfile2string(filePath string, charset string) (s string, err1 error) {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	switch charset {
	case "GBK":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	case "TGBK":
		decodeBytes, _ := traditionalchinese.Big5.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	default:
		s = string(b)

	}

	return s, nil
}
func GetProjectpath() (admin COMPONENT, err error) {

	o := orm.NewOrm()
	dbtype := Getdbtype()
	sql := "select * from dev_component_tb where componentlevel='0'"
	sql = ConvertSQL(sql, dbtype)
	err = o.Raw(sql).QueryRow(&admin)

	return admin, err
}
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
func Unescapedjs(x string) interface{} {
	return template.JS(x)
}
func UnescapedJSStr(x string) interface{} {
	return template.JSStr(x)
}
func Mod(s, m int) int {
	return s % m
}
func OutputFN(s string) string {
	return "function edit" + s + "(){"
}
func Toupper(s string) string {
	return strings.ToUpper(s)
}
func Calculate(source int, inc int) int {
	return source + inc
}
func Tofirstupper(s string) string {
	arrs := strings.Split(s, "")
	fmt.Println(arrs)
	ups := strings.ToUpper(arrs[0])
	fmt.Println(ups)
	lasts := strings.Join(arrs[1:], "")
	fmt.Println(lasts)
	return ups + lasts
}
func Tolower(s string) string {
	return strings.ToLower(s)
}
func Replace(s string, oldstr string, newstr string) string {
	return strings.Replace(s, oldstr, newstr, -1)
}

//获得两个日期的间隔天数
func Diffdays(startdate time.Time, enddate time.Time) int64 {
	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST
	// startyear := startdate.Format("2006")
	// startmonth := startdate.Format("01")
	// startday := startdate.Format("02")

	// endyear := enddate.Format("2006")
	// endmonth := enddate.Format("01")
	// endday := enddate.Format("02")
	//fmt.Println(enddate.Unix())
	fmt.Println(enddate.UnixNano())
	fmt.Println(startdate.UnixNano())
	fmt.Println(enddate.UnixNano() - startdate.UnixNano())
	return (enddate.UnixNano() - startdate.UnixNano()) / (24 * 60 * 60 * 1000 * 1000 * 1000)
}

//获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}

}

//获取相差时间
func GetMinuteDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 60
		return hour
	} else {
		return hour
	}

}
func Getfieldtype(value interface{}) (fieldtype string) {

	if value == nil {
		fieldtype = "varchar"
	} else {
		val := reflect.ValueOf(value)

		reflectfieldtype := val.Kind()
		switch reflectfieldtype {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldtype = "int"
			break
		case reflect.Float32, reflect.Float64:
			fieldtype = "float"
			break
		case reflect.Bool:
			fieldtype = "boolean"
			break
		case reflect.String:
			fieldtype = "varchar"
			break
		case reflect.Array:
			fieldtype = "Array"
			break
		case reflect.Slice:
			fieldtype = "Array"
			break
		case reflect.Map:
			fieldtype = "map"
			break
		default:
			fieldtype = "varchar"
			break
		}
	}
	//fmt.Println("Getfieldtype==>" + fieldtype)
	return
}

//a,b,c ==>'a','b','c'
func Convertstring2instring(commastring string) string {
	result := ""
	arr := strings.Split(commastring, ",")
	for idx, ar := range arr {
		if ar == "" {
			continue
		}
		result = result + "'" + ar + "'"
		if idx < len(arr)-1 {
			result = result + ","
		}

	}
	return result
}

//生成excel文件
func WriteExcelfile(excelfilename string, colnames []string, datamaparr []map[string]interface{}) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	filepath := GetCurrentDirectory() + "/static/files/"
	filename := filepath + excelfilename + ".xlsx"
	_, err = os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
	}
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	row = sheet.AddRow()
	for _, colname := range colnames {
		cell = row.AddCell()
		cell.Value = colname
	}
	for _, data := range datamaparr {
		row = sheet.AddRow()
		for _, colname := range colnames {
			cell = row.AddCell()
			Setcellvalue(cell, data[colname])
		}

	}

	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

//生成excel文件,同时带一行数据
func WriteExcelfileanddata(excelfilename string, colnames []string, datamaparr []map[string]interface{}, rowcolnames []string, rowdatamaparr []orm.Params) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	filepath := GetCurrentDirectory() + "/static/files/"
	filename := filepath + excelfilename + ".xlsx"
	_, err = os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}
	}
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	row = sheet.AddRow()
	for _, colname := range colnames {
		cell = row.AddCell()
		cell.Value = colname
	}
	for _, data := range datamaparr {

		row = sheet.AddRow()
		for _, colname := range colnames {
			cell = row.AddCell()
			Setcellvalue(cell, data[colname])
		}

	}
	sheet, err = file.AddSheet("Sheet2")
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	row = sheet.AddRow()
	for _, colname := range rowcolnames {
		cell = row.AddCell()
		cell.Value = colname
	}
	for _, data := range rowdatamaparr {
		row = sheet.AddRow()
		for _, colname := range rowcolnames {
			cell = row.AddCell()
			//fmt.Println(colname)
			Setcellvalue(cell, data[colname])
		}

	}

	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil

}

//获得当前执行程序所在的路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//struct转换为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//获得struct的所有属性
func GetStructFieldnames(structobj interface{}) []string {
	structfilednames := make([]string, 0)
	t := reflect.TypeOf(structobj)
	for i := 0; i < t.NumField(); i++ {
		structfilednames = append(structfilednames, t.Field(i).Name)
	}
	return structfilednames
}
func Convertinterface2value(arg interface{}) interface{} {
	if arg == nil {
		return ""
	}
	val := reflect.ValueOf(arg)
	kind := val.Kind()
	switch kind {
	case reflect.String:
		v := val.String()
		arg = v
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		arg = val.Int()
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		arg = val.Uint()
		break
	case reflect.Float32:
		arg, _ = orm.StrTo(orm.ToStr(arg)).Float64()
		break
	case reflect.Float64:
		arg = val.Float()
		break
	case reflect.Bool:
		arg = val.Bool()
		break
	case reflect.Array:
		var content []byte
		var err error
		err = json.Unmarshal(content, arg)
		if err != nil {
			beego.Error(err)
		}
		arg = StringsToJSON(string(content))
		break
	case reflect.Slice:
		var content []byte
		var err error
		err = json.Unmarshal(content, arg)
		if err != nil {
			beego.Error(err)
		}
		arg = StringsToJSON(string(content))
		break
	default:
		v := val.String()
		arg = v
		break
	}
	return arg
}
func ConvertInterface2valueByfieldtype(fieldtype string, arg interface{}) interface{} {
	//Getlog().Debug("ConvertInterface2valueByfieldtype()==>" + fieldtype + "==>" + orm.ToStr(arg))
	if arg == nil {
		switch fieldtype {
		case "date", "datetime", "time", "timestamp", "year":
			return nil
		case "int", "smallint", "int64", "bigint", "long", "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
			return nil
		default:
			return ""
		}
	}
	val := reflect.ValueOf(arg)
	kind := val.Kind()
	switch fieldtype {
	case "varchar", "char", "text", "longtext", "tinytext":
		//ObjectIdHex("
		Getlog().Debug("ConvertInterface2valueByfieldtype==>" + orm.ToStr(arg))
		Getlog().Debug("kind==>" + orm.ToStr(kind))
		if kind == reflect.Slice {
			textfld := ""
			argarr := arg.([]interface{})
			for idx, arg1 := range argarr {
				textfld = textfld + strings.TrimSpace(orm.ToStr(arg1))
				if idx < len(argarr)-1 {
					textfld = textfld + ","
				}
			}
			return textfld

		} else {
			textfld := orm.ToStr(arg)
			if strings.Contains(textfld, "ObjectIdHex") {
				textfld = strings.Replace(textfld, "ObjectIdHex(\"", "", -1)
				textfld = strings.Replace(textfld, "\")", "", -1)
			}
			return textfld
		}

	case "int", "smallint":
		arg, _ = orm.StrTo(orm.ToStr(arg)).Int()
		return arg
	case "int64", "bigint", "long":
		arg, _ = orm.StrTo(orm.ToStr(arg)).Int64()
		return arg
	case "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
		arg, _ = orm.StrTo(orm.ToStr(arg)).Float64()
		return arg
	case "tinyint":
		arg, _ = orm.StrTo(orm.ToStr(arg)).Bool()
		return arg
	case "date", "datetime", "time", "timestamp", "year":
		//fmt.Println("orm.ToStr(arg)==>" + orm.ToStr(arg))
		//fmt.Println("orm.ToStr(val)==>" + orm.ToStr(val))
		v := orm.ToStr(val)
		//return arg.(time.Time)
		//return Convert2time(orm.ToStr(arg))
		v = strings.Replace(v, "/", "-", -1)
		v = strings.Replace(v, "T", " ", -1)
		v = strings.Replace(v, "Z", "", -1)
		//fmt.Println(v)
		//fmt.Println(len(v))
		var t time.Time
		var err error
		if len(v) >= 19 {
			s := v[:19]
			t, err = time.ParseInLocation(formatDateTime, s, time.Local)
		} else if len(v) >= 10 {
			s := v
			if len(v) > 10 {
				s = v[:10]
			}
			t, err = time.ParseInLocation(formatDate, s, time.Local)
		} else {
			s := v
			if len(s) > 8 {
				s = v[:8]
			}
			t, err = time.ParseInLocation(formatTime, s, time.Local)
		}
		if err == nil {
			if fieldtype == "date" || fieldtype == "year" {
				v = t.In(time.Local).Format(formatDate)
			} else if fieldtype == "datetime" {
				v = t.In(time.Local).Format(formatDateTime)
			} else {
				v = t.In(time.Local).Format(formatTime)
			}
			return v
		} else {
			fmt.Println(err)
			Getlog().Error(err.Error())
			return ""
		}
	default:
		return val.String()
	}
	return ""
}

//write excel file cell value
func Setcellvalue(cell *xlsx.Cell, data interface{}) {
	//fmt.Println(data)
	if data == nil {
		cell.Value = ""
		return
	}
	val := reflect.ValueOf(data)
	kind := val.Kind()
	//fmt.Println(kind)
	//fmt.Println(val)
	switch kind {
	case reflect.String:
		//v := val.String()
		//cell.Value = v
		cell.SetValue(val)

		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := val.Int()
		cell.SetInt64(v)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := val.Uint()
		cell.SetValue(v)

		break
	case reflect.Float32:
		v, _ := orm.StrTo(orm.ToStr(val)).Float64()
		cell.SetFloat(v)
		break
	case reflect.Float64:
		v := val.Float()
		cell.SetFloat(v)
		break
	case reflect.Bool:
		v := val.Bool()
		cell.SetBool(v)
		break
	case reflect.Array:
		// var content []byte
		// var err error
		// err = json.Unmarshal(content, data)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// cell.Value = StringsToJSON(string(content))
		cell.SetValue(data)
		break
	case reflect.Slice:
		// var content []byte
		// var err error
		// err = json.Unmarshal(content, data)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// cell.Value = StringsToJSON(string(content))
		cell.SetValue(data)
		break
	default:
		cell.SetValue(data)
		break
	}
}

//转换json字节为string
func StringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}
//获得指定长度的随机数字符串
//size :指定长度
//kind:数字/大小写字母/
//KC_RAND_KIND_NUM   = 0 // 纯数字
//KC_RAND_KIND_LOWER = 1 // 小写字母
//KC_RAND_KIND_UPPER = 2 // 大写字母
//KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
func GetRandomStr(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
func CrypPWDtoSHA256(password string) (str, salt string) {
	salt = GetRandomStr(20, 3)
	//fmt.Println("random string:" + salt)
	sha := sha256.New()
	sha.Write([]byte(salt + password))

	str = hex.EncodeToString(sha.Sum(nil))
	//fmt.Println("sha256 password:" + str)
	return str, salt
}

//9ec9750e709431dad22365cabc5c625482e574c74adaebba7dd02f1129e4ce1d
//YzcmCZNvbXocrsz9dm8e
func VerifySHA256PWD(password string, salt string) string {
	sha := sha256.New()

	sha.Write([]byte(salt + password))
	str := hex.EncodeToString(sha.Sum(nil))
	fmt.Println("verify sha256 password:" + str)

	return str
}

//转换汉字为拼音
//张三===>[zhang san]
func ConvertChinese2Pinyin(chinese string, tone string) []string {
	pinyinarr := make([]string, 0)
	a := pinyin.NewArgs()
	switch tone {
	case "tone": //包含声调zhōng guó
		a.Style = pinyin.Tone
		twoarr := pinyin.Pinyin(chinese, a)
		pinyinarr = twoarr[0]
	case "tonenumber": //zhong1 guo2
		a.Style = pinyin.Tone3
		twoarr := pinyin.Pinyin(chinese, a)
		pinyinarr = twoarr[0]
	case "mutitone":
		a.Heteronym = true
		a.Style = pinyin.Tone3
		twoarr := pinyin.Pinyin(chinese, a)
		pinyinarr = twoarr[0]
	default: //zhong guo
		pinyinarr = pinyin.LazyPinyin(chinese, a)

	}
	fmt.Println(pinyinarr)
	return pinyinarr
}
//张三==>三
//张先生==>先生
//欧阳奋强==>奋强
func GetFirstnamecn(usernamecn string) string {
	firstnamecn := ""
	nameRune := []rune(usernamecn)
	runelen := len(nameRune)
	switch runelen {
	case 2:
		firstnamecn = string(nameRune[1])
		break
	case 3:
		firstnamecn = string(nameRune[1:3])
		break
	case 4:
		firstnamecn = string(nameRune[2:4])
		break
	case 5:
		firstnamecn = string(nameRune[2:5])
		break
	default:
		firstnamecn = string(nameRune[1])
		break
	}
	return firstnamecn

}
//张三==>三
//张先生==>先生
//欧阳奋强==>奋强
func GetLastnamecn(usernamecn string) string {
	lastnamecn := ""
	nameRune := []rune(usernamecn)
	runelen := len(nameRune)
	switch runelen {
	case 2:
		lastnamecn = string(nameRune[0])
		break
	case 3:
		lastnamecn = string(nameRune[0])
		break
	case 4:
		lastnamecn = string(nameRune[0:1])
		break
	case 5:
		lastnamecn = string(nameRune[0:1])
		break
	default:
		lastnamecn = string(nameRune[0])
		break
	}
	return lastnamecn

}
func IstotalDigital(digitalstr string) bool {
	pattern := "^\\d+$" 
	//Tregisteorg start
	isdigital, err := regexp.MatchString(pattern, digitalstr)
	if err != nil {
		return false
	}
	return isdigital
}
