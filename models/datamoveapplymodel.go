package models

import (
	"database/sql"

	_ "errors"
	"fmt"
	_ "math"
	_ "reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
)

type DATAMOVEAPPLY struct {
	Ip                  string
	Sourcetargettype    string //源/目标区分，source/target
	Datasource          string
	Tablename           string
	Targettabletruncate string //目标表是否truncate
	Onlyinsert          string //只是插入。不进行有则更新判断
	Pageindex           int
	Pagesize            int
}

type DATAMOVEAPPLYITEM struct {
	Sourcetargettype string `json:"Sourcetargettype"` //源/目标区分，source/target
	Datasource       string `json:"Datasource"`
	Tablename        string `json:"Tablename"`
	Movestatus       string `json:"Movestatus"`
	Starttime        string `json:"Starttime"`
	Endtime          string `json:"Endtime"`
	Spendtime        string `json:"Spendtime"`
	Operator         string `json:"Operator"`
	Fromdatasource   string `json:"Fromdatasource"`
	Cancel           string `json:"Cancel"`
	Log              string `json:"Log"`
	Pageindex        int    `json:"Pageindex"`
	Pagesize         int    `json:"Pagesize"`
}

type FromDatasourceTableinfo struct {
	Id                     int64  `orm:"pk;auto"`
	Fromdatasource         string `orm:"column(fromdatasource)"`         //源数据源
	Fromtablename          string `orm:"column(fromtablename)"`          //源表
	Fromfieldnames         string `orm:"column(fromfieldnames)"`         //源表字段，用逗号分隔；例如 a,b,c
	Fromchildfieldnames    string `orm:"column(fromchildfieldnames)"`    //源子表字段，用逗号分隔；例如 a,b,c
	Fromfieldandchildnames string `orm:"column(fromfieldandchildnames)"` //父字段和子字段组合，用逗号分隔；例如 activity.userName,activity.email
	Datasource             string `orm:"column(datasource)"`             //目标数据源
	Tablename              string `orm:"column(tablename)"`              //目标表
	Targetfieldnames       string `orm:"column(targetfieldnames)"`       //目标表字段，用逗号分隔；例如 a,b,c
	TruncateTargetSql      string `orm:"column(truncateTargetSql)"`      //truncate目标表
	QuerySourceSql         string `orm:"column(querySourceSql)"`         //查询源表SQL，例如select a,b,c from table
	InsertTargetSql        string `orm:"column(insertTargetSql)"`        //插入目标表SQL，例如insert into table(a,b,c) values(?,?,?)
	UpdateTargetSql        string `orm:"column(updateTargetSql)"`        //更新目标表SQL，例如update table set a=?,b=?,c=? where primarykey1=? and primarykey2=?
	SelectcountTargetSql   string `orm:"column(selectcountTargetSql)"`   //判断目标表是否存在记录SQL，例如select count(1) as ncount from table where primarykey1=? and primarykey2=?
	Sourcetableprimarykeys string `orm:"column(sourcetableprimarykeys)"` //源表主键字段
	Targettableprimarykeys string `orm:"column(targettableprimarykeys)"` //目标表主键字段

}
type FromDatasourceTableinfo2 struct {
	Id                     int64
	Fromdatasource         string            //源数据源
	Fromtablename          string            //源表
	Fromfieldnames         string            //源表字段，用逗号分隔；例如 a,b,c
	Fromchildfieldnames    string            //源子表字段，用逗号分隔；例如 a,b,c
	Fromfieldandchildnames string            //父字段和子字段组合，用逗号分隔；例如 activity.userName,activity.email
	Datasource             string            //目标数据源
	Tablename              string            //目标表
	Targetfieldnames       string            //目标表字段，用逗号分隔；例如 a,b,c
	TruncateTargetSql      string            //truncate目标表
	QuerySourceSql         string            //查询源表SQL，例如select a,b,c from table
	InsertTargetSql        string            //插入目标表SQL，例如insert into table(a,b,c) values(?,?,?)
	UpdateTargetSql        string            //更新目标表SQL，例如update table set a=?,b=?,c=? where primarykey1=? and primarykey2=?
	SelectcountTargetSql   string            //判断目标表是否存在记录SQL，例如select count(1) as ncount from table where primarykey1=? and primarykey2=?
	Sourcetableprimarykeys string            //源表主键字段
	Targettableprimarykeys string            //目标表主键字段
	Targettablefieldtype   map[string]string //目标表字段类型
}

func (u *FromDatasourceTableinfo) TableName() string {
	return "skl_fromdatasourcetableinfo_tb"
}
func AddFromDatasourceTableinfo(u FromDatasourceTableinfo) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Insert(&u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

//获得数据条数
func Getdatamoveapplycount(u DATAMOVEAPPLY) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total  from skl_datamoveapplyitem_tb a  where 1=1 "
	if u.Sourcetargettype != "" {
		sql = sql + " and sourcetargettype='" + u.Sourcetargettype + "'"
	}
	if u.Datasource != "" {
		sql = sql + " and datasource='" + u.Datasource + "'"
	}
	if u.Tablename != "" {
		sql = sql + " and tablename='" + u.Tablename + "'"
	}
	if u.Targettabletruncate != "" {
		sql = sql + " and targettabletruncate='" + u.Targettabletruncate + "'"
	}

	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}

//获得分页数据
func Getdatamoveapplybypageindex(u DATAMOVEAPPLY) (admins []DATAMOVEAPPLYITEM, err error) {
	dbtype := Getdbtype()
	admins = make([]DATAMOVEAPPLYITEM, 0)
	o := orm.NewOrm()

	sql := "SELECT a.*,b.vvalue as amount,c.flowstatusname from skl_datamoveapplyitem_tb a "
	sql = sql + " inner join fi_flowstatus c on a.flowstatus=c.flowstatus "
	sql = sql + " left join fi_var b on a.flowinstid=b.fiid and b.vid='money' where 1=1 "
	if u.Sourcetargettype != "" {
		sql = sql + " and sourcetargettype='" + u.Sourcetargettype + "'"
	}
	if u.Datasource != "" {
		sql = sql + " and datasource='" + u.Datasource + "'"
	}
	if u.Tablename != "" {
		sql = sql + " and tablename='" + u.Tablename + "'"
	}
	if u.Targettabletruncate != "" {
		sql = sql + " and targettabletruncate='" + u.Targettabletruncate + "'"
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

func writelog(fdstiptr FromDatasourceTableinfo, errchan <-chan map[string]interface{}) {
	beegolog := Getlog()
	var logmap map[string]interface{}
	logmap = <-errchan
	value, ok := logmap["starttime"].(int64)
	if !ok {
		beegolog.Debug("It's not ok for type string")

	}
	beegolog.Debug(string(value))
	beegolog.Debug(fdstiptr.Fromdatasource)
	beegolog.Debug(fdstiptr.Fromtablename)
	beegolog.Debug(fdstiptr.Fromfieldnames)
	beegolog.Debug(fdstiptr.Sourcetableprimarykeys)
	beegolog.Debug(fdstiptr.Datasource)
	beegolog.Debug(fdstiptr.Tablename)
	beegolog.Debug(fdstiptr.Targetfieldnames)
	beegolog.Debug(fdstiptr.SelectcountTargetSql)
	beegolog.Debug(fdstiptr.Targettableprimarykeys)
	beegolog.Debug(fdstiptr.InsertTargetSql)
	beegolog.Debug(fdstiptr.UpdateTargetSql)
	err, ok := logmap["error"].(error)
	if !ok {
		beegolog.Debug("It's not ok for type string")

	}
	if err != nil {
		beegolog.Error(err.Error())
	} else {
		beegolog.Debug("同步OK")
	}
	value, ok = logmap["starttime"].(int64)
	if !ok {
		beegolog.Debug("It's not ok for type string")
	}
	beegolog.Debug(string(value))

}
func Datamoveapply(dm DATAMOVEAPPLY) error {
	tablenamearr := make([]string, 0)
	switch dm.Sourcetargettype {
	case "target":

		if strings.Index(dm.Tablename, "ALL") > -1 {
			tableoptions, err := GetAllDATASOURCETABLEoptions(&DATASOURCE{Datasource: dm.Datasource})
			if err != nil {
				panic(err)
				return err
			}
			for _, tableoption := range tableoptions {
				tablenamearr = append(tablenamearr, tableoption.Value)
			}

		} else {
			tablenamearr = strings.Split(dm.Tablename, ",")
		}

		for _, tablename := range tablenamearr {
			if tablename == "" {
				continue
			}
			newdm := DATAMOVEAPPLY{Sourcetargettype: dm.Sourcetargettype, Datasource: dm.Datasource, Tablename: tablename, Targettabletruncate: dm.Targettabletruncate}

			fromDatasourceTableinfo, err := Getdatamovetableinfobytarget(newdm)
			if err != nil {
				Getlog().Error(err.Error())
				return err
			}
			for _, fdsti := range fromDatasourceTableinfo {
				fdstiptr := FromDatasourceTableinfo{}
				fdstiptr2 := fdsti
				fdstiptr2.GetSQLByTableInfo()

				fdstiptr.Datasource = fdstiptr2.Datasource
				fdstiptr.Fromchildfieldnames = fdstiptr2.Fromchildfieldnames
				fdstiptr.Fromdatasource = fdstiptr2.Fromdatasource
				fdstiptr.Fromfieldandchildnames = fdstiptr2.Fromfieldandchildnames
				fdstiptr.Fromfieldnames = fdstiptr2.Fromfieldnames
				fdstiptr.Fromtablename = fdstiptr2.Fromtablename
				fdstiptr.InsertTargetSql = fdstiptr2.InsertTargetSql
				fdstiptr.QuerySourceSql = fdstiptr2.QuerySourceSql
				fdstiptr.SelectcountTargetSql = fdstiptr2.SelectcountTargetSql
				fdstiptr.Sourcetableprimarykeys = fdstiptr2.Sourcetableprimarykeys
				fdstiptr.Tablename = fdstiptr2.Tablename
				fdstiptr.Targetfieldnames = fdstiptr2.Targetfieldnames
				fdstiptr.Targettableprimarykeys = fdstiptr2.Targettableprimarykeys
				fdstiptr.TruncateTargetSql = fdstiptr2.TruncateTargetSql
				fdstiptr.UpdateTargetSql = fdstiptr2.UpdateTargetSql

				resultchan := make(chan map[string]interface{}, 1)
				go syndatafromsource2target(dm.Onlyinsert, dm.Targettabletruncate, fdstiptr2, resultchan, dm.Ip)

				go writelog(fdstiptr, resultchan)
				go AddFromDatasourceTableinfo(fdstiptr)
			}
		}

		break
	case "source":
		if strings.Index(dm.Tablename, "ALL") > -1 {
			tableoptions, err := GetAllDATASOURCETABLEoptions(&DATASOURCE{Datasource: dm.Datasource})
			if err != nil {
				panic(err)
				return err
			}
			for _, tableoption := range tableoptions {
				tablenamearr = append(tablenamearr, tableoption.Value)
			}

		} else {
			tablenamearr = strings.Split(dm.Tablename, ",")
		}

		for _, tablename := range tablenamearr {
			newdm := DATAMOVEAPPLY{Sourcetargettype: dm.Sourcetargettype, Datasource: dm.Datasource, Tablename: tablename, Targettabletruncate: dm.Targettabletruncate}

			fromDatasourceTableinfo, err := Getdatamovetableinfobysource(newdm)
			if err != nil {
				Getlog().Error(err.Error())
				return err
			}
			for _, fdsti := range fromDatasourceTableinfo {
				fdstiptr := FromDatasourceTableinfo{}
				fdstiptr2 := fdsti
				fdstiptr2.GetSQLByTableInfo()

				fdstiptr.Datasource = fdstiptr2.Datasource
				fdstiptr.Fromchildfieldnames = fdstiptr2.Fromchildfieldnames
				fdstiptr.Fromdatasource = fdstiptr2.Fromdatasource
				fdstiptr.Fromfieldandchildnames = fdstiptr2.Fromfieldandchildnames
				fdstiptr.Fromfieldnames = fdstiptr2.Fromfieldnames
				fdstiptr.Fromtablename = fdstiptr2.Fromtablename
				fdstiptr.InsertTargetSql = fdstiptr2.InsertTargetSql
				fdstiptr.QuerySourceSql = fdstiptr2.QuerySourceSql
				fdstiptr.SelectcountTargetSql = fdstiptr2.SelectcountTargetSql
				fdstiptr.Sourcetableprimarykeys = fdstiptr2.Sourcetableprimarykeys
				fdstiptr.Tablename = fdstiptr2.Tablename
				fdstiptr.Targetfieldnames = fdstiptr2.Targetfieldnames
				fdstiptr.Targettableprimarykeys = fdstiptr2.Targettableprimarykeys
				fdstiptr.TruncateTargetSql = fdstiptr2.TruncateTargetSql
				fdstiptr.UpdateTargetSql = fdstiptr2.UpdateTargetSql

				resultchan := make(chan map[string]interface{}, 1)
				go syndatafromsource2target(dm.Onlyinsert, dm.Targettabletruncate, fdstiptr2, resultchan, dm.Ip)
				go writelog(fdstiptr, resultchan)
				go AddFromDatasourceTableinfo(fdstiptr)
			}
		}
		break
	}
	return nil

}

//同步源数据到目标表中
func syndatafromsource2target(onlyinsert, truncatetargettable string, fdsti FromDatasourceTableinfo2, errorchan chan<- map[string]interface{}, ip string) {
	queryfromdatasourcesql := fdsti.QuerySourceSql
	//Getlog().Debug("queryfromdatasourcesql==>" + queryfromdatasourcesql)
	//Getlog().Debug("truncatetargettable==>" + truncatetargettable)
	//Getlog().Debug("fdsti==>" + orm.ToStr(fdsti))
	//resultmap := make(map[string]interface{}, 0)
	//resultmap["starttime"] = time.Now().Unix()
	//获得目标数据源信息
	datasource, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fdsti.Datasource})
	if err != nil {
		Getlog().Error(err.Error())
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}
	//获得源数据源信息
	fromdatasource, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fdsti.Fromdatasource})
	if err != nil {
		Getlog().Error(err.Error())
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}
	var sourcemongoconn *MongoConn //源mongodb连接
	//var targetmongoconn *MongoConn //目标mongodb连接
	var sourcedb *sql.DB //源(非mongodb)数据库类型连接
	var targetdb *sql.DB //目标(非mongodb)数据库类型连接
	var rows *sql.Rows
	//源表所有行数据，但只包含目标表需要的列数据
	sourcedatamap := make([]map[string]interface{}, 0)
	mysqlsourcedatamap := make([]orm.Params, 0)
	//根据源DB类型获得不同的源数据库连接
	switch fromdatasource.Dbtype {
	case "mongodb":
		sourcemongoconn, err = GetMongoConn(fromdatasource)
		if err != nil {
			Getlog().Error(err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}

		sourcedatamap, err = sourcemongoconn.Getcollectiondatasbyfieldnames(fromdatasource.Schema, fdsti.Fromtablename, fdsti.Fromfieldandchildnames)
		if err != nil {
			Getlog().Error(err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		break
	case "mysql":
		sourcedb, err = GetMysqlConn(fromdatasource)
		if err != nil {
			Getlog().Error(err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		o, err := orm.NewOrmWithDB("mysql", fromdatasource.Datasource, sourcedb)
		if err != nil {

			Getlog().Error("orm.NewOrmWithDB('mysql', fromdatasource.Datasource, sourcedb)==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		//Getlog().Debug("mysql fdsti.QuerySourceSql==>" + fdsti.QuerySourceSql)

		//sourcedatamap2 := make([]map[string]interface{}, 0)
		//_, err = o.Raw(queryfromdatasourcesql).QueryRows(&sourcedatamap2)
		_, err = o.Raw(queryfromdatasourcesql).Values(&mysqlsourcedatamap)
		if err != nil {
			Getlog().Error("o.Raw(fdsti.QuerySourceSql).QueryRows(&sourcedatamap)==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		//Getlog().Debug("mysql sourcedatamap==>" + orm.ToStr(admins))
		//sourcedatamap = sourcedatamap2
		sourcedatamap = ConvertParam2map(mysqlsourcedatamap)
		break
	case "sqlserver":
		sourcedb, err = newMssqlConn(fromdatasource)
		if err != nil {
			Getlog().Error("newMssqlConn(fromdatasource)==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		rows, err = sourcedb.Query(fdsti.QuerySourceSql)
		if err != nil {
			Getlog().Error("sourcedb.Query(fdsti.QuerySourceSql)==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		//sourcedatamap = make([]map[string]interface{}, len(rows.Columns()))
		totalcolumns, err := rows.Columns()
		if err != nil {
			Getlog().Error("rows.Columns()==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		//Getlog().Debug("totalcolumns==>" + orm.ToStr(totalcolumns))
		for rows.Next() {
			columns, err := rows.Columns()
			if err != nil {
				Getlog().Error("rows.Columns()==>" + err.Error())
				//resultmap["error"] = err
				//resultmap["endtime"] = time.Now().Unix()
				//errorchan <- resultmap
				return
			}
			refs := make([]interface{}, len(totalcolumns))
			for i, _ := range refs {
				//var ref sql.NullString
				refs[i] = &columns[i]

			}
			err = rows.Scan(refs...)
			if err != nil {
				Getlog().Error("rows.Scan(refs...)==>" + err.Error())
				//resultmap["error"] = err
				//resultmap["endtime"] = time.Now().Unix()
				//errorchan <- resultmap
				return
			}
			datamap := make(map[string]interface{}, 0)
			for idx, column := range totalcolumns {
				//Getlog().Debug("column==>" + column)
				//Getlog().Debug("columns[idx]==>" + orm.ToStr(columns[idx]))
				//Getlog().Debug("refs[idx]==>" + orm.ToStr(refs[idx]))
				datamap[column] = columns[idx]
			}
			sourcedatamap = append(sourcedatamap, datamap)

		}

		break
	case "oracle":
		sourcedb, err = GetOracleConn(fromdatasource)
		if err != nil {
			Getlog().Error("GetOracleConn(fromdatasource)==>" + err.Error())
			return
		}
		rows, err = sourcedb.Query(fdsti.QuerySourceSql)
		if err != nil {
			Getlog().Error("sourcedb.Query(fdsti.QuerySourceSql)==>" + err.Error())
			return
		}
		//sourcedatamap = make([]map[string]interface{}, len(rows.Columns()))
		totalcolumns, err := rows.Columns()
		if err != nil {
			Getlog().Error("rows.Columns()==>" + err.Error())
			return
		}
		//Getlog().Debug("totalcolumns==>" + orm.ToStr(totalcolumns))
		for rows.Next() {
			columns, err := rows.Columns()
			if err != nil {
				Getlog().Error("rows.Columns()==>" + err.Error())
				return
			}
			refs := make([]interface{}, len(totalcolumns))
			for i, _ := range refs {
				//var ref sql.NullString
				refs[i] = &columns[i]

			}
			err = rows.Scan(refs...)
			if err != nil {
				Getlog().Error("rows.Scan(refs...)==>" + err.Error())
				return
			}
			datamap := make(map[string]interface{}, 0)
			for idx, column := range totalcolumns {
				//Getlog().Debug("column==>" + column)
				//Getlog().Debug("columns[idx]==>" + orm.ToStr(columns[idx]))
				//Getlog().Debug("refs[idx]==>" + orm.ToStr(refs[idx]))
				datamap[column] = columns[idx]
			}
			sourcedatamap = append(sourcedatamap, datamap)

		}
		break
	}
	//根据目标DB类型获得不同的目标数据库连接
	//本项目目标DB类型为mysql
	switch datasource.Dbtype {
	// case "mongodb":
	// 	targetmongoconn, err = GetMongoConn(datasource)
	// 	if err != nil {
	// 		resultmap["error"] = err
	// 		resultmap["endtime"] = time.Now().Unix()
	// 		errorchan <- resultmap
	// 		return
	// 	}
	// 	break
	case "mysql":
		targetdb, err = GetMysqlConn(datasource)
		if err != nil {
			Getlog().Error("GetMysqlConn(datasource)==>" + err.Error())
			//resultmap["error"] = err
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
		break
		// case "sqlserver":
		// 	targetdb, err = newMssqlConn(datasource)
		// 	if err != nil {
		// 		resultmap["error"] = err
		// 		resultmap["endtime"] = time.Now().Unix()
		// 		errorchan <- resultmap
		// 		return
		// 	}
		// 	break
		// case "oracle":
		// 	break
	}
	//TO DO 根据目标DB类型动态使用连接句柄，本项目目标数据库类型为"mysql"
	o, err := orm.NewOrmWithDB("mysql", datasource.Datasource, targetdb)

	if err != nil {
		Getlog().Error("orm.NewOrmWithDB('mysql', datasource.Datasource, targetdb)==>" + err.Error())
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}
	err = o.Begin()
	if err != nil {
		Getlog().Error("o.Begin()==>" + err.Error())
		//resultmap["error"] = err
		o.Rollback()
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}
	//truncate目标表
	if truncatetargettable == "1" {
		_, err = o.Raw(fdsti.TruncateTargetSql).Exec()
		if err != nil {
			Getlog().Error("o.Raw(fdsti.TruncateTargetSql).Exec()==>" + err.Error())
			//resultmap["error"] = err
			o.Rollback()
			//resultmap["endtime"] = time.Now().Unix()
			//errorchan <- resultmap
			return
		}
	}
	err = o.Commit()
	if err != nil {
		Getlog().Error("o.Commit()==>" + err.Error())
		//resultmap["error"] = err
		o.Rollback()
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}

	//循环处理源表中的数据
	//根据主键（或唯一约束）对目标表进行有则更新，无则插入
	//目标表中的字段，除了主键（或唯一约束）外，其它字段必须设置为可以为空
	//需要根据目标表定义判断是否存在字段值为空的数据。如果有,需要记录在日志中。
	//数据条数
	//sourcedatamap = sourcedatamap[:1]
	datacount := len(sourcedatamap)
	//500万以上时，启动1000个协程，每个协程处理5000条以上的数据
	var ndataOrGoroutineCount int   //每个协程处理的数据条数或需要启动的协程数
	isThousandGoroutineMode := true //true时，1000协程模式；false时非1000协程模式
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		Getlog().Error("config.NewConfig==>" + err.Error())
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		return
	}
	dividDataCount := iniconf.DefaultInt("parameters::dividDataCount", 5000000)
	maxGoroutine := iniconf.DefaultInt("parameters::maxGoroutine", 1000)
	singleGoroutineDataCount := iniconf.DefaultInt("parameters::singleGoroutineDataCount", 20000)
	if datacount >= dividDataCount {
		ndataOrGoroutineCount = datacount / maxGoroutine

	} else { //dividDataCount以下时，启动n个协程，每个协程处理singleGoroutineDataCount条数据
		ndataOrGoroutineCount = datacount / singleGoroutineDataCount
		isThousandGoroutineMode = false
	}
	var dividedsourcedatamap []map[string]interface{}
	if isThousandGoroutineMode {

		for i := 1; i <= maxGoroutine; i++ {
			startcount := ndataOrGoroutineCount * (i)
			endcount := ndataOrGoroutineCount * (i + 1)
			if i == maxGoroutine {
				dividedsourcedatamap = sourcedatamap[startcount:]
				go dividsyndatatotarget(onlyinsert, datasource, fromdatasource, fdsti, dividedsourcedatamap, ip)
			} else {
				dividedsourcedatamap = sourcedatamap[startcount:endcount]
				go dividsyndatatotarget(onlyinsert, datasource, fromdatasource, fdsti, dividedsourcedatamap, ip)
			}

		}
	} else {

		for i := 0; i <= ndataOrGoroutineCount; i++ {
			startcount := singleGoroutineDataCount * (i)
			endcount := singleGoroutineDataCount * (i + 1)
			if i == ndataOrGoroutineCount {
				dividedsourcedatamap = sourcedatamap[startcount:]
				go dividsyndatatotarget(onlyinsert, datasource, fromdatasource, fdsti, dividedsourcedatamap, ip)
			} else {
				dividedsourcedatamap = sourcedatamap[startcount:endcount]
				go dividsyndatatotarget(onlyinsert, datasource, fromdatasource, fdsti, dividedsourcedatamap, ip)
			}

		}
	}

}

//根据要迁移的目标表获得表配置信息
//一张目标表可对应多张源表
func Getdatamovetableinfobytarget(dm DATAMOVEAPPLY) (admins []FromDatasourceTableinfo2, err error) {
	dbtype := Getdbtype()

	admins = make([]FromDatasourceTableinfo2, 0)
	o := orm.NewOrm()
	sql := "SELECT max(datasource) as datasource,max(tablename) as tablename,GROUP_CONCAT(targetfieldname) as targetfieldnames ,fromdatasource ,fromtablename,GROUP_CONCAT(fromfieldname) as fromfieldnames,GROUP_CONCAT(fromchildfieldname) as fromchildfieldnames,GROUP_CONCAT(fromfieldname,'.',fromchildfieldname) as fromfieldandchildnames   FROM skl_targetfromsourceitem_tb where datasource=? and tablename=? group by fromdatasource,fromtablename"
	_, err = o.Raw(ConvertSQL(sql, dbtype), dm.Datasource, dm.Tablename).QueryRows(&admins)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

//根据要迁移的源表获得表配置信息
//一张源表可对于多张目标表
func Getdatamovetableinfobysource(dm DATAMOVEAPPLY) (admins []FromDatasourceTableinfo2, err error) {
	dbtype := Getdbtype()

	admins = make([]FromDatasourceTableinfo2, 0)
	o := orm.NewOrm()
	sql := "SELECT datasource,tablename,GROUP_CONCAT(targetfieldname) as targetfieldnames ,max(fromdatasource) as fromdatasource ,max(fromtablename) as fromtablename,GROUP_CONCAT(fromfieldname) as fromfieldnames,GROUP_CONCAT(fromchildfieldname) as fromchildfieldnames,GROUP_CONCAT(fromfieldname,'.',fromchildfieldname) as fromfieldandchildnames   FROM skl_targetfromsourceitem_tb where fromdatasource=? and fromtablename=? group by datasource,tablename"
	_, err = o.Raw(ConvertSQL(sql, dbtype), dm.Datasource, dm.Tablename).QueryRows(&admins)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

//根据要迁移的源表获得表配置信息
//一张源表可对于多张目标表
func Getfromdatasourcefilter(fdt FromDatasourceTableinfo2) (filter string, err error) {
	dbtype := Getdbtype()

	o := orm.NewOrm()
	sql := "SELECT filter from skl_stablefilter_tb where datasource=? and tablename=? and fromdatasource=? and fromtablename=?"
	err = o.Raw(ConvertSQL(sql, dbtype), fdt.Datasource, fdt.Tablename, fdt.Fromdatasource, fdt.Fromtablename).QueryRow(&filter)
	if err != nil {
		return "", err
	}
	return filter, nil
}

//根据目标/源表对照信息获得SELECT/INSERT/UPDATE SQL

func (tblinf *FromDatasourceTableinfo2) GetSQLByTableInfo() {
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		fmt.Println(err)
	}
	limitDataCountFortest := iniconf.DefaultInt("parameters::limitDataCountFortest", 0)
	//truncate目标表SQL
	truncateTargetSql := ""
	//查询源表SQL
	querySourceSql := ""
	//插入目标表SQL
	insertTargetSql := ""
	//更新目标表SQL
	updateTargetSql := ""
	//判断目标表中是否存在该记录SQL
	selectcountTargetSql := ""
	//目标表where条件SQL，用于更新、查询条件。用主键拼接成。
	targettablewheresql := ""

	fromdatasource, err := GetDATASOURCEBYID(DATASOURCE{Datasource: tblinf.Fromdatasource})
	if err != nil {
		Getlog().Error(err.Error())
	}

	truncateTargetSql = "truncate table " + tblinf.Tablename
	if limitDataCountFortest == 0 {
		if fromdatasource.Dbtype == "sqlserver" {
			querySourceSql = "select " + tblinf.Fromfieldnames + " from " + fromdatasource.Schema + ".guest." + tblinf.Fromtablename
		} else {
			querySourceSql = "select " + tblinf.Fromfieldnames + " from " + tblinf.Fromtablename
		}

	} else {
		switch fromdatasource.Dbtype {
		case "sqlserver":
			querySourceSql = "select top " + strconv.Itoa(limitDataCountFortest) + " " + tblinf.Fromfieldnames + " from " + fromdatasource.Schema + ".guest." + tblinf.Fromtablename
			break
		case "oracle":
			querySourceSql = "select  " + tblinf.Fromfieldnames + " from " + tblinf.Fromtablename + " where rownum>=1 and rownum<=" + strconv.Itoa(limitDataCountFortest)
			break
		case "mysql":
			querySourceSql = "select " + tblinf.Fromfieldnames + " from " + tblinf.Fromtablename + " limit 1," + strconv.Itoa(limitDataCountFortest)
			break
		default:
			querySourceSql = "select " + tblinf.Fromfieldnames + " from " + tblinf.Fromtablename + " limit 1," + strconv.Itoa(limitDataCountFortest)
			break
		}
	}

	insertTargetSql = "insert into " + tblinf.Tablename + "(" + tblinf.Targetfieldnames + ") values(" + getmarkbyfieldnames(tblinf.Targetfieldnames) + ")"

	updateTargetSql = "update " + tblinf.Tablename + " set " + getupdatemarkbyfieldnames(tblinf.Targetfieldnames)

	targettableprimarykeys, _ := Gettableprimarykeys(tblinf.Datasource, tblinf.Tablename)

	targettablewheresql = getwheremarkbyfieldnames(targettableprimarykeys)

	updateTargetSql = updateTargetSql + targettablewheresql

	selectcountTargetSql = "select count(1) as ncount from " + tblinf.Tablename + targettablewheresql

	sourcetableprimarykeys, _ := Gettableprimarykeys(tblinf.Fromdatasource, tblinf.Fromtablename)

	targettablefieldtypemap, _ := Gettablefieldtypemap(tblinf.Datasource, tblinf.Tablename)

	tblinf.TruncateTargetSql = truncateTargetSql
	tblinf.QuerySourceSql = querySourceSql
	tblinf.InsertTargetSql = insertTargetSql
	tblinf.UpdateTargetSql = updateTargetSql
	tblinf.SelectcountTargetSql = selectcountTargetSql
	tblinf.Sourcetableprimarykeys = sourcetableprimarykeys
	tblinf.Targettableprimarykeys = targettableprimarykeys
	tblinf.Targettablefieldtype = targettablefieldtypemap

}

//根据字段获得?部分,例如insert table(a,b,c) values(?,?,?)
//a,b,c ==>?,?,?
func getmarkbyfieldnames(fieldnames string) string {
	markstring := ""
	fieldarr := strings.Split(fieldnames, ",")
	for idx, _ := range fieldarr {
		markstring = markstring + "?"
		if idx < len(fieldarr)-1 {
			markstring = markstring + ","
		}
	}
	return markstring
}

//根据字段获得?部分，例如update table set a=?,b=?,c=?
//a,b,c ==>a=?,b=?,c=?
func getupdatemarkbyfieldnames(fieldnames string) string {
	markstring := ""
	fieldarr := strings.Split(fieldnames, ",")
	for idx, field := range fieldarr {
		markstring = markstring + field + "=?"
		if idx < len(fieldarr)-1 {
			markstring = markstring + ","
		}
	}
	return markstring
}

//根据字段获得where条件
//[a,b,c] ==>where a=? and b=? and c=?
func getwheremarkbyfieldnames(fieldnames string) string {
	fieldnamesarr := strings.Split(fieldnames, ",")
	markstring := " where "

	for idx, field := range fieldnamesarr {
		if field == "" {
			continue
		}
		markstring = markstring + field + "=?"
		if idx < len(fieldnamesarr)-1 {
			markstring = markstring + " and "
		}
	}
	return markstring
}

//获得表的主键字段，根据数据源名、表名
func Gettableprimarykeys(datasource, tablename string) (primaryfields string, err error) {
	dbtype := Getdbtype()

	primaryfieldsarr := make([]string, 0)
	o := orm.NewOrm()
	sql := "SELECT fieldname FROM skl_datasourcetablefield_tb  where  datasource=? and  tablename=? and isprimary='1'"
	_, err = o.Raw(ConvertSQL(sql, dbtype), datasource, tablename).QueryRows(&primaryfieldsarr)
	if err != nil {
		return "", err
	}
	primaryfields = strings.Join(primaryfieldsarr, ",")
	return primaryfields, nil
}

//获得表的主键字段，根据数据源名、表名、字段名。适用于mongodb嵌套子表
func Getchildtableprimarykeys(datasource, tablename, fieldname string) (primaryfields string, err error) {
	dbtype := Getdbtype()

	primaryfieldsarr := make([]string, 0)
	o := orm.NewOrm()
	sql := "SELECT CONCAT_WS('.',fieldname,childfieldname) as fieldname from skl_datasourcetablefieldchild_tb  where  datasource=? and  tablename=? and fieldname=? and isprimary='1'"
	_, err = o.Raw(ConvertSQL(sql, dbtype), datasource, tablename, fieldname).QueryRows(&primaryfieldsarr)
	if err != nil {
		return "", err
	}
	primaryfields = strings.Join(primaryfieldsarr, ",")
	return primaryfields, nil
}

//获得表的字段类型map，根据数据源名、表名
func Gettablefieldtypemap(datasource, tablename string) (fieldtypemap map[string]string, err error) {
	dbtype := Getdbtype()
	fieldtypesd := make([]DATASOURCETABLEFIELD, 0)
	fieldtypemap = make(map[string]string, 0)
	o := orm.NewOrm()
	sql := "SELECT  fieldname,fieldtype from skl_datasourcetablefield_tb  where  datasource=? and  tablename=? "
	_, err = o.Raw(ConvertSQL(sql, dbtype), datasource, tablename).QueryRows(&fieldtypesd)
	if err != nil {
		return nil, err
	}
	for _, fieldtypesd1 := range fieldtypesd {
		fieldtypemap[fieldtypesd1.Fieldname] = fieldtypesd1.Fieldtype
	}
	return fieldtypemap, nil
}

//获得字段在字段数组中的索引，根据字段名.fieldname:a;fieldnamesarr:b,a,c===>1
func GetFieldIndexOfArray(fieldname string, fieldnames string) int {
	//Getlog().Debug("GetFieldIndexOfArray()==>" + fieldname + "==>" + fieldnames)
	fieldnamesarr := strings.Split(fieldnames, ",")
	for idx, fieldnamesa := range fieldnamesarr {
		if fieldnamesa == fieldname {
			return idx
		}
	}
	return 0
}
func ConvertParam2map(sourcemap []orm.Params) []map[string]interface{} {
	destmap := make([]map[string]interface{}, 0)
	for _, source := range sourcemap {
		dest := make(map[string]interface{}, 0)
		for key, value := range source {
			dest[key] = value
		}

		destmap = append(destmap, dest)
	}
	return destmap
}
func dividsyndatatotarget(onlyinsert string, datasource DATASOURCE, fromdatasource DATASOURCE, fdsti FromDatasourceTableinfo2, dividedsourcedatamap []map[string]interface{}, ip string) {
	dmai := DATAMOVEAPPLYITEM{}
	dmai.Sourcetargettype = datasource.Sourcetargettype
	dmai.Datasource = datasource.Datasource
	dmai.Tablename = fdsti.Tablename
	dmai.Movestatus = "unfinished"
	d1 := time.Now()
	dmai.Starttime = orm.ToStr(d1)
	dmai.Fromdatasource = fromdatasource.Datasource
	dmai.Log = ""

	targetdb, err := GetMysqlConn(datasource)
	if err != nil {
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		Getlog().Error("GetMysqlConn==>" + err.Error())
		dmai.Movestatus = "false"
		dmai.Log = err.Error()

		GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
		return
	}
	o, err := orm.NewOrmWithDB("mysql", datasource.Datasource, targetdb)

	if err != nil {
		//resultmap["error"] = err
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		Getlog().Error("NewOrmWithDB==>" + err.Error())
		dmai.Log = err.Error()
		dmai.Movestatus = "unfinished"
		GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
		return
	}
	err = o.Begin()
	if err != nil {
		//resultmap["error"] = err
		o.Rollback()
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		Getlog().Error("o.Begin==>" + err.Error())
		dmai.Log = err.Error()
		GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
		return
	}
	for _, sdm := range dividedsourcedatamap {

		selectcountparams := make([]interface{}, 0)
		var ncount int
		if onlyinsert == "0" {
			if fromdatasource.Dbtype != "mongodb" {
				//Getlog().Debug("fdsti.Sourcetableprimarykeys==>" + fdsti.Sourcetableprimarykeys)
				primarykeysarr := strings.Split(fdsti.Sourcetableprimarykeys, ",")
				//Getlog().Debug("fdsti.Targettablefieldtype==>" + orm.ToStr(fdsti.Targettablefieldtype))
				for _, pkfld := range primarykeysarr {
					paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[pkfld], sdm[pkfld])
					selectcountparams = append(selectcountparams, paramvalue)
				}

				err = o.Raw(fdsti.SelectcountTargetSql, selectcountparams...).QueryRow(&ncount)
				if err != nil {
					//resultmap["error"] = err
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("SelectcountTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					dmai.Movestatus = "unfinished"
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return
				}
			} else {
				fromfieldandchildnamesarr := strings.Split(fdsti.Fromfieldandchildnames, ",")
				//targetfieldnamesarr := strings.Split(fdsti.Targetfieldnames, ",")
				targetprimarykeysarr := strings.Split(fdsti.Targettableprimarykeys, ",")
				//Getlog().Debug("Fromfieldandchildnames==>" + fdsti.Fromfieldandchildnames)
				//Getlog().Debug("Targetfieldnames==>" + fdsti.Targetfieldnames)
				//Getlog().Debug("Targettableprimarykeys==>" + fdsti.Targettableprimarykeys)
				for _, targetprimarykeysar := range targetprimarykeysarr {
					//Getlog().Debug("targetprimarykeysar==>" + targetprimarykeysar)
					targetprimarykeyindex := GetFieldIndexOfArray(targetprimarykeysar, fdsti.Targetfieldnames)
					//Getlog().Debug("orm.ToStr(targetprimarykeyindex)==>" + orm.ToStr(targetprimarykeyindex))
					sourceprimarykeyfield := fromfieldandchildnamesarr[targetprimarykeyindex]
					//Getlog().Debug("sourceprimarykeyfield==>" + sourceprimarykeyfield)
					ffnarr := strings.Split(sourceprimarykeyfield, ".")
					if ffnarr[1] == "" || ffnarr[1] == "no" {
						paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetprimarykeysar], sdm[ffnarr[0]])
						selectcountparams = append(selectcountparams, paramvalue)
					} else { //comm_user.user_id<==user.activity.userName
						//Getlog().Debug("sdm[ffnarr[0]]==>" + orm.ToStr(sdm[ffnarr[0]]))
						childvaluemap := sdm[ffnarr[0]].(map[string]interface{})
						//Getlog().Debug("childvaluemap==>" + orm.ToStr(childvaluemap))
						paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetprimarykeysar], childvaluemap[ffnarr[1]])
						selectcountparams = append(selectcountparams, paramvalue)

					}
				}

				err = o.Raw(fdsti.SelectcountTargetSql, selectcountparams...).QueryRow(&ncount)
				if err != nil {
					//resultmap["error"] = err
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("o.Raw(fdsti.SelectcountTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return
				}
			}
		} else {
			ncount = -1
		}

		if fromdatasource.Dbtype != "mongodb" {

			ffnarr := strings.Split(fdsti.Fromfieldnames, ",")
			targetfieldnamesarr := strings.Split(fdsti.Targetfieldnames, ",")
			if ncount < 1 { //无则插入
				inserttargetparams := make([]interface{}, 0)
				for idx, ffn := range ffnarr {
					paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], sdm[ffn])
					inserttargetparams = append(inserttargetparams, paramvalue)
				}
				_, err := o.Raw(fdsti.InsertTargetSql, inserttargetparams...).Exec()
				if err != nil {
					//resultmap["error"] = err
					o.Rollback()
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("o.Raw(fdsti.InsertTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return

				}

			} else { //有则更新
				updatetargetparams := make([]interface{}, 0)
				for idx, ffn := range ffnarr {
					paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], sdm[ffn])
					updatetargetparams = append(updatetargetparams, paramvalue)
				}
				primarykeysarr := strings.Split(fdsti.Sourcetableprimarykeys, ",")
				for _, pkfld := range primarykeysarr {
					paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[pkfld], sdm[pkfld])
					updatetargetparams = append(updatetargetparams, paramvalue)
				}
				_, err := o.Raw(fdsti.UpdateTargetSql, updatetargetparams...).Exec()
				if err != nil {
					//resultmap["error"] = err
					o.Rollback()
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("o.Raw(fdsti.UpdateTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return
				}
			}
		} else { //源是mongodb需要特殊处理
			//为insert sql中 ?赋值
			fromfieldandchildnamesarr := strings.Split(fdsti.Fromfieldandchildnames, ",")
			targetfieldnamesarr := strings.Split(fdsti.Targetfieldnames, ",")
			if ncount < 1 { //无则插入
				inserttargetparams := make([]interface{}, 0)
				for idx, ffn := range fromfieldandchildnamesarr {
					//Getlog().Debug("ffn==>" + orm.ToStr(ffn))
					ffnarr := strings.Split(ffn, ".")
					//lenghffnarr := len(ffnarr)
					//Getlog().Debug("len(ffnarr)==>" + orm.ToStr(lenghffnarr))
					if ffnarr[1] == "" || ffnarr[1] == "no" {
						paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], sdm[ffnarr[0]])
						inserttargetparams = append(inserttargetparams, paramvalue)
					} else { //email<==activity.email
						//Getlog().Debug("sdm[ffnarr[lenghffnarr-2]]==>" + orm.ToStr(sdm[ffnarr[lenghffnarr-2]]))

						mongovalue := loopmongomap(sdm, ffnarr)
						if mongovalue != nil {
							//childvaluemap := sdm[ffnarr[lenghffnarr-2]].(map[string]interface{})
							paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], mongovalue)
							inserttargetparams = append(inserttargetparams, paramvalue)
						} else {
							switch fdsti.Targettablefieldtype[targetfieldnamesarr[idx]] {
							case "date", "datetime", "time", "timestamp", "year":
								inserttargetparams = append(inserttargetparams, nil)
							case "int", "smallint", "int64", "bigint", "long", "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
								inserttargetparams = append(inserttargetparams, nil)
							default:
								inserttargetparams = append(inserttargetparams, "")
							}

						}

						// if sdm[ffnarr[lenghffnarr-2]] != nil {

						// 	childvaluemap := sdm[ffnarr[lenghffnarr-2]].(map[string]interface{})
						// 	paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], childvaluemap[ffnarr[lenghffnarr-1]])
						// 	inserttargetparams = append(inserttargetparams, paramvalue)

						// } else {

						// 	switch fdsti.Targettablefieldtype[targetfieldnamesarr[idx]] {
						// 	case "date", "datetime", "time", "timestamp", "year":
						// 		inserttargetparams = append(inserttargetparams, nil)
						// 	case "int", "smallint", "int64", "bigint", "long", "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
						// 		inserttargetparams = append(inserttargetparams, nil)
						// 	default:
						// 		inserttargetparams = append(inserttargetparams, "")
						// 	}

						// }

					}

				}
				_, err := o.Raw(fdsti.InsertTargetSql, inserttargetparams...).Exec()
				if err != nil {
					//resultmap["error"] = err
					o.Rollback()
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("o.Raw(fdsti.InsertTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return

				}

			} else { //有则更新
				//为update sql中set 位置?赋值
				fromfieldandchildnamesarr := strings.Split(fdsti.Fromfieldandchildnames, ",")
				targetfieldnamesarr := strings.Split(fdsti.Targetfieldnames, ",")
				updatetargetparams := make([]interface{}, 0)
				for idx, ffn := range fromfieldandchildnamesarr {
					ffnarr := strings.Split(ffn, ".")
					if ffnarr[1] == "" || ffnarr[1] == "no" {
						paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], sdm[ffnarr[0]])
						updatetargetparams = append(updatetargetparams, paramvalue)
					} else { //email<==activity.email
						//childvaluemap := sdm[ffnarr[0]].(map[string]interface{})
						//paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], childvaluemap[ffnarr[1]])
						//updatetargetparams = append(updatetargetparams, paramvalue)
						mongovalue := loopmongomap(sdm, ffnarr)
						if mongovalue != nil {
							//childvaluemap := sdm[ffnarr[lenghffnarr-2]].(map[string]interface{})
							paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetfieldnamesarr[idx]], mongovalue)
							updatetargetparams = append(updatetargetparams, paramvalue)
						} else {
							switch fdsti.Targettablefieldtype[targetfieldnamesarr[idx]] {
							case "date", "datetime", "time", "timestamp", "year":
								updatetargetparams = append(updatetargetparams, nil)
							case "int", "smallint", "int64", "bigint", "long", "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
								updatetargetparams = append(updatetargetparams, nil)
							default:
								updatetargetparams = append(updatetargetparams, "")
							}

						}

					}
				}
				//为update sql中where 位置?赋值
				targetprimarykeysarr := strings.Split(fdsti.Targettableprimarykeys, ",")
				for _, targetprimarykeysar := range targetprimarykeysarr {
					targetprimarykeyindex := GetFieldIndexOfArray(targetprimarykeysar, fdsti.Targetfieldnames)
					sourceprimarykeyfield := fromfieldandchildnamesarr[targetprimarykeyindex]
					ffnarr := strings.Split(sourceprimarykeyfield, ".")
					if ffnarr[1] == "" || ffnarr[1] == "no" {
						paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetprimarykeysar], sdm[ffnarr[0]])
						updatetargetparams = append(updatetargetparams, paramvalue)
					} else { //comm_user.user_id<==user.activity.userName
						//childvaluemap := sdm[ffnarr[0]].(map[string]interface{})
						//paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetprimarykeysar], childvaluemap[ffnarr[1]])
						//updatetargetparams = append(updatetargetparams, paramvalue)
						mongovalue := loopmongomap(sdm, ffnarr)
						if mongovalue != nil {
							//childvaluemap := sdm[ffnarr[lenghffnarr-2]].(map[string]interface{})
							paramvalue := ConvertInterface2valueByfieldtype(fdsti.Targettablefieldtype[targetprimarykeysar], mongovalue)
							updatetargetparams = append(updatetargetparams, paramvalue)
						} else {
							switch fdsti.Targettablefieldtype[targetprimarykeysar] {
							case "date", "datetime", "time", "timestamp", "year":
								updatetargetparams = append(updatetargetparams, nil)
							case "int", "smallint", "int64", "bigint", "long", "real", "decimal", "double", "float", "money", "number", "smallmoney", "numeric":
								updatetargetparams = append(updatetargetparams, nil)
							default:
								updatetargetparams = append(updatetargetparams, "")
							}

						}

					}
				}

				_, err := o.Raw(fdsti.UpdateTargetSql, updatetargetparams...).Exec()
				if err != nil {
					//resultmap["error"] = err
					o.Rollback()
					//resultmap["endtime"] = time.Now().Unix()
					//errorchan <- resultmap
					Getlog().Error("o.Raw(fdsti.UpdateTargetSql==>" + err.Error())
					dmai.Log = err.Error()
					GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
					return
				}
			}
		}

	} //for
	err = o.Commit()
	if err != nil {
		//resultmap["error"] = err
		//o.Rollback()
		//resultmap["endtime"] = time.Now().Unix()
		//errorchan <- resultmap
		Getlog().Error("o.Commit()==>" + err.Error())
		dmai.Log = err.Error()
		GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}
	}
	Getlog().Info("o.Commit()==>ok")

	d2 := time.Now()
	dmai.Endtime = orm.ToStr(d2)
	//dmai.Spendtime = orm.ToStr(GetMinuteDiffer(dmai.Starttime, dmai.Endtime)) + " minutes"
	dmai.Log = "syn data is ok"
	dmai.Movestatus = "finished"
	//resultmap["error"] = nil
	//resultmap["endtime"] = time.Now().Unix()
	//errorchan <- resultmap
	GetWebsocketServer().Datamessage <- &SynDataMessage{Ip: ip, Message: dmai}

}

//arryflag 数组标识 0 非数组 1数组
func loopmongomap(sdm map[string]interface{}, fieldnamearr []string) interface{} {

	newsdam := sdm

	for idx, fieldname := range fieldnamearr {
		if idx < len(fieldnamearr)-1 {
			if newsdam[fieldname] != nil {
				childmap := newsdam[fieldname].(map[string]interface{})
				newsdam = childmap

			} else {
				return nil
			}

		} else {
			return newsdam[fieldname]
		}
	}

	return nil
}

//获得大区、国家。
//{"subCatalogs.displayEN":"","subCatalogs.displayZH":""}
//subCatalogs是数组
//displayEN是map
func loopmongomapregion(sdm map[string]interface{}, fieldnamearr []string) interface{} {

	newsdam := sdm
	var childmap []interface{}
	for idx, fieldname := range fieldnamearr {
		if idx < len(fieldnamearr)-1 {
			if newsdam[fieldname] != nil {
				childmap = newsdam[fieldname].([]interface{})

			} else {
				return nil
			}

		} else {
			returnvalue := make([]interface{}, 0)
			for _, childmapvalue := range childmap {
				childmap2 := childmapvalue.(map[string]interface{})
				returnvalue = append(returnvalue, childmap2[fieldname])
			}
			return returnvalue
		}
	}

	return nil
}

//获得省，父为大区或国家。
//{"subCatalogs.displayEN":"","subCatalogs.subCatalogs.displayEN":"","subCatalogs.subCatalogs.displayZH":""}
//subCatalogs是数组
//displayEN是map
func loopmongomapprovince(sdm map[string]interface{}, fieldnamearr []string) interface{} {

	newsdam := sdm
	var childmap []interface{}
	for idx, fieldname := range fieldnamearr {
		if idx < len(fieldnamearr)-1 {
			if newsdam[fieldname] != nil {
				childmap = newsdam[fieldname].([]interface{})

			} else {
				return nil
			}

		} else {
			returnvalue := make([]interface{}, 0)
			for _, childmapvalue := range childmap {
				childmap2 := childmapvalue.(map[string]interface{})
				returnvalue = append(returnvalue, childmap2[fieldname])
			}
			return returnvalue
		}
	}

	return nil
}

//从mongodb中同步大区和国家
func Synregiontable(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "Country", "subCatalogs.displayEN,subCatalogs.displayZH")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(displayZH,displayEN) values(?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		_, err = o.Raw(insertsql, orm.ToStr(regionmap["displayZH"]), orm.ToStr(regionmap["displayEN"])).Exec()
		if err != nil {
			o.Rollback()
			Getlog().Error(err.Error())
			return err
		}
	}
	err = o.Commit()

	return nil
}

//从mongodb中同步省
func Synprovincetable(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "Country", "subCatalogs.displayEN,subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.displayZH")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(parent,displayZH,displayEN) values(?,?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		parent := orm.ToStr(regionmap["displayEN"])
		provincearr := regionmap["subCatalogs"].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			_, err = o.Raw(insertsql, parent, orm.ToStr(provincemap["displayZH"]), orm.ToStr(provincemap["displayEN"])).Exec()
			if err != nil {
				o.Rollback()
				Getlog().Error(err.Error())
				return err
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步市
func Syncitytable(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "Country", "subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayZH,subCatalogs.subCatalogs.subCatalogs.cityLevel")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(parent,displayZH,displayEN,cityLevel) values(?,?,?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})

		provincearr := regionmap["subCatalogs"].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			parent := orm.ToStr(provincemap["displayEN"])
			cityarr := provincemap["subCatalogs"].([]interface{})
			for _, city := range cityarr {
				citymap := city.(map[string]interface{})
				_, err = o.Raw(insertsql, parent, orm.ToStr(citymap["displayZH"]), orm.ToStr(citymap["displayEN"]), orm.ToStr(citymap["cityLevel"])).Exec()
				if err != nil {
					o.Rollback()
					Getlog().Error(err.Error())
					return err
				}
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步大区和国家
func Synclassroom1table(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "RoomCategory", "subCatalogs.displayEN,subCatalogs.displayZH,subCatalogs.status,subCatalogs.allowBatchUpload")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(displayZH,displayEN,status,allowBatchUpload) values(?,?,?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		_, err = o.Raw(insertsql, orm.ToStr(regionmap["displayZH"]), orm.ToStr(regionmap["displayEN"]), orm.ToStr(regionmap["status"]), orm.ToStr(regionmap["allowBatchUpload"])).Exec()
		if err != nil {
			o.Rollback()
			Getlog().Error(err.Error())
			return err
		}
	}
	err = o.Commit()

	return nil
}

//从mongodb中同步省
func Synclassroom2table(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "RoomCategory", "subCatalogs.displayEN,subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.displayZH,subCatalogs.subCatalogs.status,subCatalogs.subCatalogs.allowBatchUpload")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(parent,displayZH,displayEN,status,allowBatchUpload) values(?,?,?,?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		parent := orm.ToStr(regionmap["displayEN"])
		provincearr := regionmap["subCatalogs"].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			_, err = o.Raw(insertsql, parent, orm.ToStr(provincemap["displayZH"]), orm.ToStr(provincemap["displayEN"]), orm.ToStr(provincemap["status"]), orm.ToStr(provincemap["allowBatchUpload"])).Exec()
			if err != nil {
				o.Rollback()
				Getlog().Error(err.Error())
				return err
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步市
func Synclassroom3table(fromdatasource, datasource, tablename string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, "catalog", "RoomCategory", "subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayZH,subCatalogs.subCatalogs.subCatalogs.status,subCatalogs.subCatalogs.subCatalogs.allowBatchUpload")
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	regionarr := data0["subCatalogs"].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	insertsql := "insert into " + tablename + "(parent,displayZH,displayEN,status,allowBatchUpload) values(?,?,?,?,?)"
	o.Begin()
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})

		provincearr := regionmap["subCatalogs"].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			parent := orm.ToStr(provincemap["displayEN"])
			cityarr := provincemap["subCatalogs"].([]interface{})
			for _, city := range cityarr {
				citymap := city.(map[string]interface{})
				_, err = o.Raw(insertsql, parent, orm.ToStr(citymap["displayZH"]), orm.ToStr(citymap["displayEN"]), orm.ToStr(citymap["status"]), orm.ToStr(citymap["allowBatchUpload"])).Exec()
				if err != nil {
					o.Rollback()
					Getlog().Error(err.Error())
					return err
				}
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步一级嵌套的数据，例如从catalog取key=TrainerCategory的一级数据
//fromdatasource：源数据源。例如tap_mongodb
//fromtable：源collectionname。例如catalog
//fromkey：源表过滤条件。TrainerCategory
//fromfieldname：源表查询字段。例如：subCatalogs.displayEN,subCatalogs.displayZH,subCatalogs.status
//datasource:目标数据源。例如target_mysql
//tablename：目标表。例如t_s_trainer
//fieldname：目标字段trainer_cn_name,trainer_en_name,trainer_status
func Syncatalog1table(fromdatasource, fromtable, fromfieldname, fromkey, datasource, tablename, fieldname string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, fromtable, fromkey, fromfieldname)
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	fromfieldnamearr := strings.Split(fromfieldname, ",")
	fromfieldnamearr2 := strings.Split(fromfieldnamearr[0], ".")
	regionarr := data0[fromfieldnamearr2[0]].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}

	fieldnamearr := strings.Split(fieldname, ",")
	insertsql1 := "("
	for idx, fieldname1 := range fieldnamearr {
		insertsql1 = insertsql1 + fieldname1
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ") values("
		}
	}
	for idx, _ := range fieldnamearr {
		insertsql1 = insertsql1 + "?"
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ")"
		}
	}

	insertsql := "insert into " + tablename + insertsql1
	truncatetablesql := "truncate table " + tablename
	o.Begin()
	_, err = o.Raw(truncatetablesql).Exec()
	if err != nil {
		o.Rollback()
		Getlog().Error(err.Error())
		return err
	}
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		insertparams := make([]interface{}, 0)
		for _, fromfieldname1 := range fromfieldnamearr {
			fromfieldnamearr2 := strings.Split(fromfieldname1, ".")
			insertparams = append(insertparams, orm.ToStr(regionmap[fromfieldnamearr2[1]]))
		}
		_, err = o.Raw(insertsql, insertparams...).Exec()
		if err != nil {
			o.Rollback()
			Getlog().Error(err.Error())
			return err
		}
	}
	err = o.Commit()

	return nil
}

//从mongodb中同步二级嵌套的数据，例如从catalog取key=TrainerCategory的二级数据
//fromdatasource：源数据源。例如tap_mongodb
//fromtable：源collectionname。例如catalog
//fromkey：源表过滤条件。TrainerCategory
//fromfieldname：源表查询字段。例如：subCatalogs.displayEN,subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.displayZH,subCatalogs.subCatalogs.status
//datasource:目标数据源。例如target_mysql
//tablename：目标表。例如t_s_trainer
//fieldname：目标字段parent,trainer_cn_name,trainer_en_name,trainer_status
func Syncatalog2table(fromdatasource, fromtable, fromfieldname, fromkey, datasource, tablename, fieldname string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, fromtable, fromkey, fromfieldname)
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	fromfieldnamearr := strings.Split(fromfieldname, ",")
	fromfieldnamearr2 := strings.Split(fromfieldnamearr[0], ".")
	regionarr := data0[fromfieldnamearr2[0]].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}
	fieldnamearr := strings.Split(fieldname, ",")
	insertsql1 := "("
	for idx, fieldname1 := range fieldnamearr {
		insertsql1 = insertsql1 + fieldname1
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ") values("
		}
	}
	for idx, _ := range fieldnamearr {
		insertsql1 = insertsql1 + "?"
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ")"
		}
	}

	insertsql := "insert into " + tablename + insertsql1
	fromfieldnamearr3 := strings.Split(fromfieldnamearr[1], ".")
	truncatetablesql := "truncate table " + tablename
	o.Begin()
	_, err = o.Raw(truncatetablesql).Exec()
	if err != nil {
		o.Rollback()
		Getlog().Error(err.Error())
		return err
	}
	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})
		parent := orm.ToStr(regionmap[fromfieldnamearr2[1]])
		provincearr := regionmap[fromfieldnamearr3[1]].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			insertparams := make([]interface{}, 0)
			insertparams = append(insertparams, parent)
			for idx, fromfieldname1 := range fromfieldnamearr {
				if idx != 0 {
					fromfieldnamearr2 := strings.Split(fromfieldname1, ".")
					insertparams = append(insertparams, orm.ToStr(provincemap[fromfieldnamearr2[len(fromfieldnamearr2)-1]]))
				}
			}
			_, err = o.Raw(insertsql, insertparams...).Exec()
			if err != nil {
				o.Rollback()
				Getlog().Error(err.Error())
				return err
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步三级嵌套的数据，例如从catalog取key=TrainerCategory的三级级数据
//fromdatasource：源数据源。例如tap_mongodb
//fromtable：源collectionname。例如catalog
//fromkey：源表过滤条件。TrainerCategory
//fromfieldname：源表查询字段。例如：subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayZH,subCatalogs.subCatalogs.subCatalogs.status
//datasource:目标数据源。例如target_mysql
//tablename：目标表。例如t_s_trainer
//fieldname：目标字段parent,trainer_cn_name,trainer_en_name,trainer_status
func Syncatalog3table(fromdatasource, fromtable, fromfieldname, fromkey, datasource, tablename, fieldname string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data, err := monconn.Getcollectiondataforregion(fromds.Schema, fromtable, fromkey, fromfieldname)
	if err != nil {
		Getlog().Error("Getcollectiondataforregion==>" + err.Error())
		return err
	}
	data0 := data[0]
	fromfieldnamearr := strings.Split(fromfieldname, ",")
	fromfieldnamearr2 := strings.Split(fromfieldnamearr[0], ".")
	regionarr := data0[fromfieldnamearr2[0]].([]interface{})
	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}

	fieldnamearr := strings.Split(fieldname, ",")
	insertsql1 := "("
	for idx, fieldname1 := range fieldnamearr {
		insertsql1 = insertsql1 + fieldname1
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ") values("
		}
	}
	for idx, _ := range fieldnamearr {
		insertsql1 = insertsql1 + "?"
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ")"
		}
	}

	insertsql := "insert into " + tablename + insertsql1
	fromfieldnamearr3 := strings.Split(fromfieldnamearr[1], ".")
	truncatetablesql := "truncate table " + tablename
	o.Begin()
	_, err = o.Raw(truncatetablesql).Exec()
	if err != nil {
		o.Rollback()
		Getlog().Error(err.Error())
		return err
	}

	for _, region := range regionarr {
		regionmap := region.(map[string]interface{})

		provincearr := regionmap[fromfieldnamearr2[1]].([]interface{})
		for _, province := range provincearr {
			provincemap := province.(map[string]interface{})
			parent := orm.ToStr(provincemap[fromfieldnamearr2[2]])

			cityarr := provincemap[fromfieldnamearr3[2]].([]interface{})
			for _, city := range cityarr {
				citymap := city.(map[string]interface{})
				insertparams := make([]interface{}, 0)
				insertparams = append(insertparams, parent)
				for idx, fromfieldname1 := range fromfieldnamearr {
					if idx != 0 {
						fromfieldnamearr2 := strings.Split(fromfieldname1, ".")
						insertparams = append(insertparams, orm.ToStr(citymap[fromfieldnamearr2[len(fromfieldnamearr2)-1]]))
					}
				}
				_, err = o.Raw(insertsql, insertparams...).Exec()
				if err != nil {
					o.Rollback()
					Getlog().Error(err.Error())
					return err
				}
			}
		}

	}
	err = o.Commit()

	return nil
}

//从mongodb中同步父子表的数据，例如从trainingRoom取培训教室的年度目标
//fromdatasource：源数据源。例如tap_mongodb
//fromtable：源collectionname。例如trainingRoom
//fromkey：源表过滤条件。可以没有过滤条件
//fromfieldname：源表查询字段。例如：roomNameEN,targetList.target,targetList.year
//datasource:目标数据源。例如target_mysql
//tablename：目标表。例如t_s_classroom_annual_target
//fieldname：目标字段classroom_id,report_annual_target,annual
func Synparentchildtable(fromdatasource, fromtable, fromfieldname, fromkey, datasource, tablename, fieldname string) (err error) {
	fromds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: fromdatasource})
	if err != nil {
		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	monconn, err := GetMongoConn(fromds)
	if err != nil {
		Getlog().Error("GetMongoConn==>" + err.Error())
		return err
	}
	data := make([]map[string]interface{}, 0)
	if fromkey == "" {
		data, err = monconn.Getcollectiondataforparentchild(fromds.Schema, fromtable, fromfieldname)
		if err != nil {
			Getlog().Error("Getcollectiondataforparentchild==>" + err.Error())
			return err
		}
	} else {
		fromkeyparam := make([]interface{}, 0)
		fromkeyparam = append(fromkeyparam, fromkey)
		data, err = monconn.Getcollectiondataforparentchild(fromds.Schema, fromtable, fromfieldname, fromkeyparam)
		if err != nil {
			Getlog().Error("Getcollectiondataforparentchild==>" + err.Error())
			return err
		}
	}

	ds, err := GetDATASOURCEBYID(DATASOURCE{Datasource: datasource})
	if err != nil {

		Getlog().Error("GetDATASOURCEBYID==>" + err.Error())
		return err
	}
	targetdb, err := GetMysqlConn(ds)
	if err != nil {

		Getlog().Error("GetMysqlConn==>" + err.Error())
		return err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, targetdb)

	if err != nil {

		Getlog().Error("NewOrmWithDB==>" + err.Error())
		return err
	}

	fieldnamearr := strings.Split(fieldname, ",")
	insertsql1 := "("
	for idx, fieldname1 := range fieldnamearr {
		insertsql1 = insertsql1 + fieldname1
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ") values("
		}
	}
	for idx, _ := range fieldnamearr {
		insertsql1 = insertsql1 + "?"
		if idx < len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ","
		}
		if idx == len(fieldnamearr)-1 {
			insertsql1 = insertsql1 + ")"
		}
	}

	insertsql := "insert into " + tablename + insertsql1
	truncatetablesql := "truncate table " + tablename
	o.Begin()
	_, err = o.Raw(truncatetablesql).Exec()
	if err != nil {
		o.Rollback()
		Getlog().Error(err.Error())
		return err
	}

	fromfieldnamearr := strings.Split(fromfieldname, ",")
	fromfieldnamearr2 := strings.Split(fromfieldnamearr[1], ".")
	childfromfieldnamearr := fromfieldnamearr[1:]

	for _, datamap := range data {
		if datamap[fromfieldnamearr2[0]] == nil {
			continue
		}
		childdata := datamap[fromfieldnamearr2[0]].([]interface{})
		parent := ""
		if fromfieldnamearr[0] == "_id" {
			parent = strings.Replace(orm.ToStr(datamap[fromfieldnamearr[0]]), "ObjectIdHex(\"", "", -1)
			parent = strings.Replace(parent, "\")", "", -1)
		} else {
			parent = orm.ToStr(datamap[fromfieldnamearr[0]])
		}

		for _, child := range childdata {
			insertparams := make([]interface{}, 0)

			insertparams = append(insertparams, parent)

			if child == nil {
				continue
			}
			childmap := child.(map[string]interface{})
			for _, childfromfieldname := range childfromfieldnamearr {
				childfromfieldnamearr1 := strings.Split(childfromfieldname, ".")
				insertparams = append(insertparams, orm.ToStr(childmap[childfromfieldnamearr1[1]]))
			}
			_, err = o.Raw(insertsql, insertparams...).Exec()
			if err != nil {
				o.Rollback()
				Getlog().Error(err.Error())
				return err
			}
		}
	}
	err = o.Commit()

	return nil
}
