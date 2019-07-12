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
type DATASOURCETABLEFIELD struct {
	Id           int64  `orm:"pk;auto"`                      //主键，自动增长
	Datasource   string `orm:"column(datasource)"`           //数据源
	Tablename    string `orm:"column(tablename)"`            //表名
	Fieldname    string `orm:"column(fieldname)"`            //字段名
	Fieldtype    string `orm:"column(fieldtype)"`            // 字段类型
	Isauto       string `orm:"column(isauto);default(0)"`    //是否自增
	Fieldlength  int    `orm:"column(fieldlength)"`          //字段长度
	Isprimary    string `orm:"column(isprimary);default(0)"` //是否主键
	Isnull       string `orm:"column(isnull);default(0)"`    //是否为空
	Defaultvalue string `orm:"column(defaultvalue)"`         //缺省值
	Isparent     string `orm:"column(isparent);default(0)"`  //是否为父，适用于mongodb的嵌套
	Comment      string `orm:"column(comment);"`             //备注
}

func (u *DATASOURCETABLEFIELD) TableName() string {
	return "skl_datasourcetablefield_tb"
}

// 多字段唯一键
func (u *DATASOURCETABLEFIELD) TableUnique() [][]string {
	return [][]string{
		[]string{"Datasource", "Tablename", "Fieldname"},
	}
}

func AddDATASOURCETABLEFIELD(u *DATASOURCETABLEFIELD) error {
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

func AddMultiDATASOURCETABLEFIELD(u []DATASOURCETABLEFIELD) error {
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: u[0].Datasource})
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	truncatesql := "delete from skl_datasourcetablefield_tb where datasource=? and tablename=?"
	sql := "select count(1) as ncount from skl_datasourcetablefield_tb where datasource=? and tablename=? and fieldname=?"
	updatesql := "update skl_datasourcetablefield_tb set datasource=?,tablename=?,fieldname=?,fieldtype=?,fieldlength=?,isprimary=?,isnull=?,defaultvalue=?,isauto=?,isparent=? where datasource=? and tablename=? and fieldname=?"

	insertsql := "insert into skl_datasourcetablefield_tb(datasource,tablename,fieldname,fieldtype,fieldlength,isprimary,isnull,defaultvalue,isauto,isparent) values(?,?,?,?,?,?,?,?,?,?)"

	insertsql = ConvertSQL(insertsql, Getdbtype())
	if ds.Sourcetargettype == "source" {
		for _, u1 := range u {
			if u1.Fieldtype == "longtext" {
				u1.Fieldlength = 0
			}
			if u1.Isauto == "" {
				u1.Isauto = "0"
			}
			if u1.Isnull == "" {
				u1.Isnull = "0"
			}
			if u1.Isparent == "" {
				u1.Isparent = "0"
			}
			if u1.Isprimary == "" {
				u1.Isprimary = "0"
			}

			ncount := 0
			err = o.Raw(sql, u1.Datasource, u1.Tablename, u1.Fieldname).QueryRow(&ncount)
			if err != nil {
				return err
			}
			if ncount > 0 {
				_, err = o.Raw(updatesql, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Fieldtype, u1.Fieldlength, u1.Isprimary, u1.Isnull, u1.Defaultvalue, u1.Isauto, u1.Isparent, u1.Datasource, u1.Tablename, u1.Fieldname).Exec()
				if err != nil {
					fmt.Println(err)
					Getlog().Error(updatesql + "==>" + err.Error())
					o.Rollback()
					return err
				}
			} else {
				_, err = o.Raw(insertsql, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Fieldtype, u1.Fieldlength, u1.Isprimary, u1.Isnull, u1.Defaultvalue, u1.Isauto, u1.Isparent).Exec()
				if err != nil {
					fmt.Println(err)
					Getlog().Error(insertsql + "==>" + err.Error())
					o.Rollback()
					return err
				}
			}

		}
	} else {
		_, err = o.Raw(truncatesql, u[0].Datasource, u[0].Tablename).Exec()
		if err != nil {
			Getlog().Error(truncatesql + "==>" + err.Error())
			o.Rollback()
			return err
		}
		for _, u1 := range u {
			if u1.Fieldtype == "longtext" {
				u1.Fieldlength = 0
			}
			if u1.Isauto == "" {
				u1.Isauto = "0"
			}
			if u1.Isnull == "" {
				u1.Isnull = "0"
			}
			if u1.Isparent == "" {
				u1.Isparent = "0"
			}
			if u1.Isprimary == "" {
				u1.Isprimary = "0"
			}

			_, err = o.Raw(insertsql, u1.Datasource, u1.Tablename, u1.Fieldname, u1.Fieldtype, u1.Fieldlength, u1.Isprimary, u1.Isnull, u1.Defaultvalue, u1.Isauto, u1.Isparent).Exec()
			if err != nil {
				fmt.Println(err)
				Getlog().Error(insertsql + "==>" + err.Error())
				o.Rollback()
				return err
			}

		}

	}
	err = o.Commit()
	return err
}

func GetAllDATASOURCETABLEFIELD() (admins []DATASOURCETABLEFIELD, err error) {
	admins = make([]DATASOURCETABLEFIELD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefield_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllDATASOURCETABLEFIELDbydatasourcetable(u DATASOURCETABLE) (admins []DATASOURCETABLEFIELD, err error) {
	admins = make([]DATASOURCETABLEFIELD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefield_tb where datasource=? and tablename=?"

	_, err = o.Raw(sql, u.Datasource, u.Tablename).QueryRows(&admins)

	return admins, err
}
func GetDATASOURCETABLEFIELDBYID(e DATASOURCETABLEFIELD) (admins []DATASOURCETABLEFIELD, err error) {

	admins = make([]DATASOURCETABLEFIELD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefield_tb where id=?"

	_, err = o.Raw(sql, e.Id).QueryRows(&admins)

	return admins, err
}

//获得数据条数
func Getdatasourcetablefieldcount() (page PAGE, err error) {
	o := orm.NewOrm()
	sql := "SELECT count(1) as total  from skl_datasourcetablefield_tb a"
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)
	return page, err
}

//获得分页数据
func Getdatasourcetablefieldbypageindex(l PAGE) (admins []DATASOURCETABLEFIELD, err error) {
	dbtype := Getdbtype()
	admins = make([]DATASOURCETABLEFIELD, 0)
	o := orm.NewOrm()

	sql := "select * from skl_datasourcetablefield_tb where 1=1 "

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

func GetAllDATASOURCETABLEFIELDoptions(u *DATASOURCETABLEFIELD) (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select fieldname as value,fieldname as label from skl_datasourcetablefield_tb where datasource=? and tablename=?"

	_, err = o.Raw(sql, u.Datasource, u.Tablename).QueryRows(&admins)

	return admins, err
}

func GetDATASOURCETABLEFIELD(u *DATASOURCETABLEFIELD) (admins []DATASOURCETABLEFIELD, err error) {
	admins = make([]DATASOURCETABLEFIELD, 0)
	o := orm.NewOrm()
	sql := "select * from skl_datasourcetablefield_tb where 1=1 "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func DeleteDATASOURCETABLEFIELD(u *DATASOURCETABLEFIELD) error {

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

func UpdateDATASOURCETABLEFIELD(u *DATASOURCETABLEFIELD) error {

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

//convert []DATASOURCETABLEFIELD==>map["datasource"_"tablename"][]DATASOURCETABLEFIELD
func ConvertField2map(admins []DATASOURCETABLEFIELD) (fieldmap map[string][]DATASOURCETABLEFIELD, err error) {
	fieldmap = make(map[string][]DATASOURCETABLEFIELD, 0)
	for _, admin := range admins {
		fieldmapkey := admin.Datasource + "&" + admin.Tablename
		_, ok := fieldmap[fieldmapkey]
		if !ok {
			dtfarr := make([]DATASOURCETABLEFIELD, 0)
			dtfarr = append(dtfarr, admin)
			fieldmap[fieldmapkey] = dtfarr

		} else {
			fieldmap[fieldmapkey] = append(fieldmap[fieldmapkey], admin)
		}
	}

	return fieldmap, err
}

//convert []DATASOURCETABLEFIELD==>map["datasource"_"tablename"][]DATASOURCETABLEFIELD
func ConvertFieldchild2map(admins []DATASOURCETABLEFIELDCHILD) (fieldmap map[string][]DATASOURCETABLEFIELDCHILD, err error) {
	fieldmap = make(map[string][]DATASOURCETABLEFIELDCHILD, 0)
	for _, admin := range admins {
		fieldmapkey := admin.Datasource + "&" + admin.Tablename + "&" + admin.Fieldname
		_, ok := fieldmap[fieldmapkey]
		if !ok {
			dtfarr := make([]DATASOURCETABLEFIELDCHILD, 0)
			dtfarr = append(dtfarr, admin)
			fieldmap[fieldmapkey] = dtfarr

		} else {
			fieldmap[fieldmapkey] = append(fieldmap[fieldmapkey], admin)
		}
	}

	return fieldmap, err
}

//获得表字段数组 ==>['a','b','c'...]
func GetDATASOURCETABLEFIELDARR(datasource, tablename string) (admins []string, err error) {
	admins = make([]string, 0)
	o := orm.NewOrm()

	sql := "select fieldname from skl_datasourcetablefield_tb where datasource=? and tablename=?"

	_, err = o.Raw(sql, datasource, tablename).QueryRows(&admins)

	return admins, err
}

//获得表字段数组 ==>['a','b','c'...]
func GetDATASOURCETABLEFIELDCHILDARR(datasource, tablename string, fieldname string) (admins []string, err error) {
	admins = make([]string, 0)
	o := orm.NewOrm()

	sql := "select fieldname from skl_datasourcetablefieldchild_tb where datasource=? and tablename=? and fieldname=?"

	_, err = o.Raw(sql, datasource, tablename, fieldname).QueryRows(&admins)

	return admins, err
}
func GetmetaDATASOURCETABLEFIELDbydatasource(datasource string) (err error) {

	dsarr, err := GetDATASOURCETABLEBYdatasource(datasource)
	if err != nil {
		Getlog().Error("GetsourceDATASOURCETABLE()==>" + err.Error())
		return err
	}

	datacount := len(dsarr)
	//500万以上时，启动1000个协程，每个协程处理5000条以上的数据
	var ndataOrGoroutineCount int   //每个协程处理的数据条数或需要启动的协程数
	isThousandGoroutineMode := true //true时，1000协程模式；false时非1000协程模式

	dividDataCount := 1000
	maxGoroutine := 100
	singleGoroutineDataCount := 10
	if datacount >= dividDataCount {
		ndataOrGoroutineCount = datacount / maxGoroutine

	} else { //dividDataCount以下时，启动n个协程，每个协程处理singleGoroutineDataCount条数据
		ndataOrGoroutineCount = datacount / singleGoroutineDataCount
		isThousandGoroutineMode = false
	}
	var dividedsourcedatamap []DATASOURCETABLE
	if isThousandGoroutineMode {

		for i := 1; i <= maxGoroutine; i++ {
			startcount := ndataOrGoroutineCount * (i)
			endcount := ndataOrGoroutineCount * (i + 1)
			if i == maxGoroutine {
				dividedsourcedatamap = dsarr[startcount:]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			} else {
				dividedsourcedatamap = dsarr[startcount:endcount]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			}

		}
	} else {

		for i := 0; i <= ndataOrGoroutineCount; i++ {
			startcount := singleGoroutineDataCount * (i)
			endcount := singleGoroutineDataCount * (i + 1)
			if i == ndataOrGoroutineCount {
				dividedsourcedatamap = dsarr[startcount:]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			} else {
				dividedsourcedatamap = dsarr[startcount:endcount]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			}

		}
	}
	//mongodb支持同时处理100个，受连接限制
	//dschan := make(chan DATASOURCE, 100)
	//本地mysql支持同时处理100个，受连接限制
	//adminschan := make(chan []DATASOURCETABLEFIELD, 100)

	//支持同时处理100个数据源表

	return err
}
func GetmetaDATASOURCETABLEFIELD() (err error) {

	dsarr, err := GetAllDATASOURCETABLE()
	if err != nil {
		Getlog().Error("GetAllDATASOURCETABLE()==>" + err.Error())
		return err
	}
	datacount := len(dsarr)
	//500万以上时，启动1000个协程，每个协程处理5000条以上的数据
	var ndataOrGoroutineCount int   //每个协程处理的数据条数或需要启动的协程数
	isThousandGoroutineMode := true //true时，1000协程模式；false时非1000协程模式

	dividDataCount := 1000
	maxGoroutine := 100
	singleGoroutineDataCount := 10
	if datacount >= dividDataCount {
		ndataOrGoroutineCount = datacount / maxGoroutine

	} else { //dividDataCount以下时，启动n个协程，每个协程处理singleGoroutineDataCount条数据
		ndataOrGoroutineCount = datacount / singleGoroutineDataCount
		isThousandGoroutineMode = false
	}
	var dividedsourcedatamap []DATASOURCETABLE
	if isThousandGoroutineMode {

		for i := 1; i <= maxGoroutine; i++ {
			startcount := ndataOrGoroutineCount * (i)
			endcount := ndataOrGoroutineCount * (i + 1)
			if i == maxGoroutine {
				dividedsourcedatamap = dsarr[startcount:]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			} else {
				dividedsourcedatamap = dsarr[startcount:endcount]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			}

		}
	} else {

		for i := 0; i <= ndataOrGoroutineCount; i++ {
			startcount := singleGoroutineDataCount * (i)
			endcount := singleGoroutineDataCount * (i + 1)
			if i == ndataOrGoroutineCount {
				dividedsourcedatamap = dsarr[startcount:]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			} else {
				dividedsourcedatamap = dsarr[startcount:endcount]
				go dividDatasourcetablefieldToTarget(dividedsourcedatamap)
			}

		}
	}
	//mongodb支持同时处理100个，受连接限制
	//dschan := make(chan DATASOURCE, 100)
	//本地mysql支持同时处理100个，受连接限制
	//adminschan := make(chan []DATASOURCETABLEFIELD, 100)

	//支持同时处理100个数据源表

	return err
}
func GetmetaDATASOURCETABLEFIELDbydstable(dstable DATASOURCETABLE) (err error) {
	admins := make([]DATASOURCETABLEFIELD, 0)
	ds2, err := GetDATASOURCETABLEBYID(dstable)
	if err != nil {
		Getlog().Error("GetDATASOURCETABLEBYID()==>" + err.Error())
		return err
	}
	ds3, err := GetDATASOURCEBYID(DATASOURCE{Datasource: dstable.Datasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID()==>" + err.Error())
		return err
	}

	optionsarr, err := GetAllDATASOURCETABLEFIELDoptions(&DATASOURCETABLEFIELD{Datasource: ds2.Datasource, Tablename: ds2.Tablename})
	if err != nil {
		fmt.Println(err)
		Getlog().Error("GetAllDATASOURCETABLEFIELDoptions()==>" + err.Error())

		return err
	}
	//fmt.Println(len(optionsarr))
	//目标表或未同步时才进行同步
	if ds3.Sourcetargettype == "target" || len(optionsarr) < 1 {
		ds, _ := GetDATASOURCEBYID(DATASOURCE{Datasource: ds2.Datasource})
		//fmt.Println(ds)
		switch ds.Dbtype {
		case "mongodb":
			sourcemongoconn, err := GetMongoConn(ds)

			if err == nil {
				fieldinfomap, err := sourcemongoconn.Getcollectionfieldinfo(ds.Schema, ds2.Tablename)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("sourcemongoconn.Getcollectionfieldinfo()==>" + err.Error())

					return err
				}
				for key, value := range fieldinfomap {
					admins = append(admins, DATASOURCETABLEFIELD{Datasource: ds2.Datasource, Tablename: ds2.Tablename, Fieldname: key, Fieldtype: value.Fieldtype, Fieldlength: value.Fieldlength, Isprimary: value.Isprimary, Isauto: value.Isauto, Isnull: value.Isnull, Isparent: value.Isparent, Defaultvalue: value.Defaultvalue})
				}

				err = AddMultiDATASOURCETABLEFIELD(admins)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())
					return err

				}

			} else {

				//defer sourcemongoconn.Session.Close()
				fmt.Println(err)
				Getlog().Error("GetMongoConn()==>" + err.Error())

			}
			break
		case "mysql":

			mysqladmins, err := Getmysqltablefields(ds, ds2.Tablename)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("Getmysqltablefields()==>" + err.Error())
				return err

			}
			if len(mysqladmins) > 0 {

				err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())
					return err

				}
			}
			break
		case "sqlserver":

			mysqladmins, err := Getmssqltablefields(ds, ds2.Tablename)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("Getmssqltablefields()==>" + err.Error())
				return err

			}
			if len(mysqladmins) > 0 {

				err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())
					return err

				}
			}
			break
		case "oracle":

			mysqladmins, err := Getoracletablefields(ds, ds2.Tablename)
			if err != nil {
				fmt.Println(err)
				Getlog().Error("Getoracletablefields()==>" + err.Error())
				return err

			}
			if len(mysqladmins) > 0 {

				err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())
					return err

				}
			}
			break
		}

	}

	return nil
}
func dividDatasourcetablefieldToTarget(dsarr []DATASOURCETABLE) {

	for _, ds2 := range dsarr {

		optionsarr, err := GetAllDATASOURCETABLEFIELDoptions(&DATASOURCETABLEFIELD{Datasource: ds2.Datasource, Tablename: ds2.Tablename})
		if err != nil {
			fmt.Println(err)
			Getlog().Error("GetAllDATASOURCETABLEFIELDoptions()==>" + err.Error())
			return
		}
		//fmt.Println(len(optionsarr))
		if len(optionsarr) < 1 {
			ds, _ := GetDATASOURCEBYID(DATASOURCE{Datasource: ds2.Datasource})
			//fmt.Println(ds)
			switch ds.Dbtype {
			case "mongodb":
				sourcemongoconn, err := GetMongoConn(ds)

				if err == nil {
					fieldinfomap, err := sourcemongoconn.Getcollectionfieldinfo(ds.Schema, ds2.Tablename)
					if err != nil {
						fmt.Println(err)
						Getlog().Error("sourcemongoconn.Getcollectionfieldinfo()==>" + err.Error())

					}
					admins := make([]DATASOURCETABLEFIELD, 0)
					for key, value := range fieldinfomap {
						admins = append(admins, DATASOURCETABLEFIELD{Datasource: ds2.Datasource, Tablename: ds2.Tablename, Fieldname: key, Fieldtype: value.Fieldtype, Fieldlength: value.Fieldlength, Isprimary: value.Isprimary, Isauto: value.Isauto, Isnull: value.Isnull, Isparent: value.Isparent, Defaultvalue: value.Defaultvalue})
					}
					if len(admins) > 0 {
						err = AddMultiDATASOURCETABLEFIELD(admins)
						if err != nil {
							fmt.Println(err)
							Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())

						}
					}

				} else {

					//defer sourcemongoconn.Session.Close()
					fmt.Println(err)
					Getlog().Error("GetMongoConn()==>" + err.Error())

				}
				break
			case "mysql":

				mysqladmins, err := Getmysqltablefields(ds, ds2.Tablename)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("Getmysqltablefields()==>" + err.Error())

				}
				if len(mysqladmins) > 0 {

					err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
					if err != nil {
						fmt.Println(err)
						Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())

					}
				}
				break
			case "sqlserver":

				mysqladmins, err := Getmssqltablefields(ds, ds2.Tablename)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("Getmssqltablefields()==>" + err.Error())

				}
				if len(mysqladmins) > 0 {

					err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
					if err != nil {
						fmt.Println(err)
						Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())

					}
				}
				break
			case "oracle":

				mysqladmins, err := Getoracletablefields(ds, ds2.Tablename)
				if err != nil {
					fmt.Println(err)
					Getlog().Error("Getoracletablefields()==>" + err.Error())

				}
				if len(mysqladmins) > 0 {

					err = AddMultiDATASOURCETABLEFIELD(mysqladmins)
					if err != nil {
						fmt.Println(err)
						Getlog().Error("AddMultiDATASOURCETABLEFIELD()==>" + err.Error())

					}
				}
				break
			}

		} //if
	} //for
}
