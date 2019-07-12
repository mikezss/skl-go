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
	_ "github.com/mattn/go-oci8"
)

func GetOracleConn2(ds DATASOURCE) (db *sql.DB, err error) {
	connString := fmt.Sprintf("%s/%s@%s:%d/%s", ds.Username, ds.Password, ds.Ip, ds.Port, ds.Schema)
	Getlog().Debug("connString==>" + orm.ToStr(connString))
	db, err = sql.Open("goracle", connString)
	if err != nil {
		Getlog().Error("sql.Open('goracle', connString)==>" + err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		Getlog().Error("db.Ping==>" + err.Error())
		return nil, err
	}
	Getlog().Debug("GetOracleConn2 db==>ok" + orm.ToStr(db))
	return db, nil
}

func GetOracleConn(ds DATASOURCE) (db *sql.DB, err error) {
	db, err = orm.GetDB(ds.Datasource)
	if err != nil {
		Getlog().Error("orm.GetOracleConn()==>" + err.Error())

	}
	Getlog().Debug("orm.GetOracleConn()==>ok")
	return db, nil
}
func NewOracleConn(ds DATASOURCE) error {

	sb := make([]string, 0)
	sb = append(sb, ds.Username+"/")
	sb = append(sb, ds.Password+"@")
	sb = append(sb, ds.Ip+":")
	sb = append(sb, strconv.Itoa(ds.Port)+"/")
	sb = append(sb, ds.Schema)

	datasourcename := strings.Join(sb, "")
	fmt.Println(datasourcename)

	err := orm.RegisterDriver("oci8", orm.DROracle)
	if err != nil {
		Getlog().Error("orm.RegisterDriver()==>" + err.Error())
	}

	err = orm.RegisterDataBase(ds.Datasource, "oci8", datasourcename)
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

func Getoracletablenames(ds DATASOURCE) (tablenames []DATASOURCETABLE, err error) {
	tablenames = make([]DATASOURCETABLE, 0)
	sql := "SELECT '" + ds.Datasource + "' as datasource,table_name as tablename FROM USER_TABLES where  not REGEXP_LIKE (table_name, '.*_COPY.*' ) and  not REGEXP_LIKE(table_name, '.*[0-9]{8}.*') "
	db, err := GetOracleConn(ds)

	if err != nil {
		Getlog().Error("GetOracleConn()==>" + err.Error())
		return nil, err
	}
	Getlog().Debug("oracle db==>" + orm.ToStr(db))
	o, err := orm.NewOrmWithDB("oci8", ds.Datasource, db)
	if err != nil {
		Getlog().Error("orm.NewOrmWithDB()==>" + err.Error())
		return nil, err
	}
	Getlog().Debug("oracle orm==>" + orm.ToStr(o))
	_, err = o.Raw(sql).QueryRows(&tablenames)

	if err != nil {
		Getlog().Error("o.Raw(sql).QueryRows()==>" + err.Error())
		return nil, err
	}

	return tablenames, nil
}
func Getoracletablenames2(ds DATASOURCE) (tablenames []DATASOURCETABLE, err error) {
	tablenames = make([]DATASOURCETABLE, 0)
	sql := "SELECT '" + ds.Datasource + "' as datasource,table_name as tablename FROM USER_TABLES where  not REGEXP_LIKE (table_name, '.*_COPY.*' ) and  not REGEXP_LIKE(table_name, '.*[0-9]{8}.*') "
	db, err := GetOracleConn(ds)

	if err != nil {
		Getlog().Error("GetOracleConn()==>" + err.Error())
		return nil, err
	}
	rows, err := db.Query(sql)
	if err != nil {
		Getlog().Error("db.Query(sql)==>" + err.Error())
		return nil, err
	}
	for rows.Next() {
		dst := DATASOURCETABLE{}
		err = rows.Scan(&dst.Datasource, &dst.Tablename)
		if err != nil {
			Getlog().Error("rows.Scan(&dst.Datasource, &dst.Tablename)==>" + err.Error())
			return nil, err
		}
		tablenames = append(tablenames, dst)
	}

	return tablenames, nil
}

func Getoracletablefields(ds DATASOURCE, tablename string) (fieldsmap []DATASOURCETABLEFIELD, err error) {
	//SELECT column_name,data_type,data_length,nullable FROM ALL_TAB_COLUMNS  WHERE TABLE_NAME='CMN_TEST_TB'
	sql := "SELECT '" + ds.Datasource + "' as datasource,a.TABLE_NAME as tablename,a.COLUMN_NAME as fieldname "
	sql = sql + ",DATA_TYPE as fieldtype,case NULLABLE when 'Y' then '1' else '0' end as isnull   "
	sql = sql + ", DATA_LENGTH as fieldlength, (case c.UNIQUENESS when 'UNIQUE' then '1' else '0' end ) as isprimary "
	sql = sql + " FROM ALL_TAB_COLUMNS a "
	sql = sql + " left join USER_IND_COLUMNS b on a.TABLE_NAME=b.TABLE_NAME and a.COLUMN_NAME=b.COLUMN_NAME "
	sql = sql + " left join USER_INDEXES c on b.TABLE_NAME=c.TABLE_NAME and b.INDEX_NAME=c.INDEX_NAME "
	sql = sql + " WHERE  a.TABLE_NAME =:s_tablename"

	db, err := GetOracleConn(ds)
	if err != nil {
		Getlog().Error("GetOracleConn()==>" + err.Error())
		return nil, err
	}
	// o, err := orm.NewOrmWithDB("oci8", ds.Datasource, db)
	// if err != nil {
	// 	Getlog().Error("orm.NewOrmWithDB()==>" + err.Error())
	// 	return nil, err
	// }

	fieldsmap = make([]DATASOURCETABLEFIELD, 0)
	// _, err = o.Raw(sql, strings.ToUpper(tablename)).QueryRows(&fieldsmap)

	// if err != nil {
	// 	Getlog().Error("o.Raw(sql).QueryRows()==>" + err.Error())
	// 	return nil, err
	// }
	s_tablename := strings.ToUpper(tablename)
	rows, err := db.Query(sql, s_tablename)
	if err != nil {
		Getlog().Error("db.Query==>" + err.Error())
		return nil, err
	}
	for rows.Next() {
		dst := DATASOURCETABLEFIELD{}
		err = rows.Scan(&dst.Datasource, &dst.Tablename, &dst.Fieldname, &dst.Fieldtype, &dst.Isnull, &dst.Fieldlength, &dst.Isprimary)
		if err != nil {
			Getlog().Error("rows.Scan(&dst.Datasource, &dst.Tablename,...)==>" + err.Error())
			return nil, err
		}
		dst.Fieldtype = strings.ToLower(dst.Fieldtype)
		fieldsmap = append(fieldsmap, dst)
	}
	Getlog().Debug("fieldsmap==>" + orm.ToStr(fieldsmap))
	return fieldsmap, nil
}
func Getoracletabledatacount(ds DATASOURCE, tablename string) (ncount int, err error) {
	countsql := "select count(1) as ncount from " + tablename
	db, err := GetOracleConn(ds)
	if err != nil {
		Getlog().Error("GetOracleConn()==>" + err.Error())
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
