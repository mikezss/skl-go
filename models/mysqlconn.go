package models

import (
	"fmt"

	"database/sql"
	_ "log"
	"strconv"
	"strings"
	_ "sync"
	"time"

	"github.com/astaxie/beego/config"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConn(ds DATASOURCE) (db *sql.DB, err error) {
	db, err = orm.GetDB(ds.Datasource)
	if err != nil {
		Getlog().Error("orm.GetDB()==>" + err.Error())

	}
	Getlog().Debug("orm.GetDB()==>ok")
	return db, nil
}
func NewMysqlConn(ds DATASOURCE) error {

	sb := make([]string, 0)
	sb = append(sb, ds.Username+":")
	sb = append(sb, ds.Password+"@tcp(")
	sb = append(sb, ds.Ip+":")
	sb = append(sb, strconv.Itoa(ds.Port)+")/")
	sb = append(sb, ds.Schema+"?charset=utf8&loc=Asia%2FShanghai&parseTime=true")

	datasourcename := strings.Join(sb, "")
	fmt.Println(datasourcename)
	//datasourcename:-->root:root@tcp(localhost:3306)/skl-ticket?charset=utf8
	//tcp:localhost:3306*mydb/root/rootroot
	// err := orm.RegisterDriver("mysql", orm.DRMySQL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := orm.RegisterDataBase(ds.Datasource, "mysql", datasourcename)
	if err != nil {

		return err
	}
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		return err
	}
	maxOpenConns := iniconf.DefaultInt("parameters::maxOpenConns", 150)
	maxIdleConns := iniconf.DefaultInt("parameters::maxIdleConns", 100)
	orm.SetMaxOpenConns(ds.Datasource, maxOpenConns)
	orm.SetMaxIdleConns(ds.Datasource, maxIdleConns)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {

		return err
	}
	err = orm.SetDataBaseTZ(ds.Datasource, location)
	if err != nil {

		return err
	}
	//orm.DefaultTimeLoc = location
	return nil

}
func Getmysqltablenames(ds DATASOURCE) (tablenames []DATASOURCETABLE, err error) {
	tablenames = make([]DATASOURCETABLE, 0)
	sql := "SELECT '" + ds.Datasource + "' as datasource,table_name as tablename FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema = DATABASE() and table_name not REGEXP '.*_copy.*' and table_name not REGEXP '.*[0-9]{8}.*' "
	db, err := GetMysqlConn(ds)
	if err != nil {
		Getlog().Error("Getalltablenames()==>" + err.Error())
		return nil, err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, db)
	if err != nil {
		Getlog().Error("orm.NewOrmWithDB()==>" + err.Error())
		return nil, err
	}
	_, err = o.Raw(sql).QueryRows(&tablenames)

	if err != nil {
		Getlog().Error("o.Raw(sql).QueryRows()==>" + err.Error())
		return nil, err
	}

	return tablenames, nil
}
func Getmysqltablefields(ds DATASOURCE, tablename string) (fieldsmap []DATASOURCETABLEFIELD, err error) {
	sql := "SELECT '" + ds.Datasource + "' as datasource,TABLE_NAME as tablename,COLUMN_NAME as fieldname "
	sql = sql + ",DATA_TYPE as fieldtype,case IS_NULLABLE when 'YES' then '1' else '0' end as Isnull ,COLUMN_COMMENT as comment,COLUMN_DEFAULT as defaultvalue "
	sql = sql + ",case COLUMN_KEY when 'PRI' then '1' else '0' end as isprimary,CHARACTER_MAXIMUM_LENGTH as fieldlength  "
	sql = sql + " FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = ?"
	db, err := GetMysqlConn(ds)
	if err != nil {
		Getlog().Error("GetMysqlConn()==>" + err.Error())
		return nil, err
	}
	o, err := orm.NewOrmWithDB("mysql", ds.Datasource, db)
	if err != nil {
		Getlog().Error("orm.NewOrmWithDB()==>" + err.Error())
		return nil, err
	}

	fieldsmap = make([]DATASOURCETABLEFIELD, 0)
	_, err = o.Raw(sql, tablename).QueryRows(&fieldsmap)

	if err != nil {
		Getlog().Error("o.Raw(sql).QueryRows()==>" + err.Error())
		return nil, err
	}

	return fieldsmap, nil
}
func Getmysqltabledatacount(ds DATASOURCE, tablename string) (ncount int, err error) {
	countsql := "select count(1) as ncount from " + tablename
	db, err := GetMysqlConn(ds)
	if err != nil {
		Getlog().Error("GetMysqlConn()==>" + err.Error())
		return 0, err
	}
	rows, err := db.Query(countsql)
	if err != nil {
		Getlog().Error("db.Query(sql)==>" + err.Error())
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&ncount)
		if err != nil {
			Getlog().Error("rows.Scan()==>" + err.Error())
			return 0, err
		}
	}
	return ncount, nil

}
