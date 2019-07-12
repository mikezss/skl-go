package models

import (
	"fmt"

	"database/sql"
	_ "log"
	_ "net/url"

	_ "github.com/astaxie/beego/orm"
	_ "github.com/denisenkom/go-mssqldb"
)

func newMssqlConn(ds DATASOURCE) (db *sql.DB, err error) {

	// query := url.Values{}
	// //query.Add("app name", "MyAppName")

	// u := &url.URL{
	// 	Scheme:   "sqlserver",
	// 	User:     url.UserPassword(ds.Username, ds.Password),
	// 	Host:     fmt.Sprintf("%s:%d", ds.Ip, ds.Port),
	// 	Path:     ds.Schema, // if connecting to an instance instead of a port
	// 	RawQuery: query.Encode(),
	// }
	// var conn *mssql.Connector
	// conn, err = mssql.NewConnector(u.String())
	// if err != nil {
	// 	return nil, err
	// }
	// db = sql.OpenDB(conn)
	//db, err := sql.Open("sqlserver", u.String())
	//db.SetMaxOpenConns(200)
	//db.SetMaxIdleConns(50)
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", ds.Ip, ds.Port, ds.Schema, ds.Username, ds.Password)
	db, err = sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	//`sqlserver://username:password@host/instance?param1=value&param2=value`

	return db, nil
}
func Getmssqltablenames(ds DATASOURCE) (tablenames []DATASOURCETABLE, err error) {
	tablenames = make([]DATASOURCETABLE, 0)
	sql := "SELECT '" + ds.Datasource + "' as datasource,name as tablename FROM sysobjects where xtype='U'"
	db, err := newMssqlConn(ds)
	if err != nil {
		Getlog().Error("Getalltablenames()==>" + err.Error())
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
func Getmssqltablefields(ds DATASOURCE, tablename string) (fieldsmap []DATASOURCETABLEFIELD, err error) {
	fieldsmap = make([]DATASOURCETABLEFIELD, 0)
	sql := "SELECT '" + ds.Datasource + "' as datasource,a.name as tablename,b.name as fieldname "
	sql = sql + ",c.name as fieldtype,b.isnullable as isnull "
	sql = sql + ",(case d.keyno when '1' then '1' else '0' end  ) as isprimary,b.length as fieldlength  "
	sql = sql + " from sysobjects a "
	sql = sql + " inner join syscolumns b on a.id=b.id  "
	sql = sql + " inner join SYSTYPES c on b.xtype=c.xtype  "
	sql = sql + " left join sysindexkeys d on d.id=a.id and b.colid=d.colid  "
	sql = sql + " where a.name =?"
	db, err := newMssqlConn(ds)
	if err != nil {
		Getlog().Error("newMssqlConn()==>" + err.Error())
		return nil, err
	}
	rows, err := db.Query(sql, tablename)
	if err != nil {
		Getlog().Error("db.Query(sql)==>" + err.Error())
		return nil, err
	}
	for rows.Next() {
		dst := DATASOURCETABLEFIELD{}
		err = rows.Scan(&dst.Datasource, &dst.Tablename, &dst.Fieldname, &dst.Fieldtype, &dst.Isnull, &dst.Isprimary, &dst.Fieldlength)
		if err != nil {
			Getlog().Error("rows.Scan(&dst.Datasource, &dst.Tablename)==>" + err.Error())
			return nil, err
		}
		fieldsmap = append(fieldsmap, dst)
	}

	return fieldsmap, nil
}
func Getmssqltabledatacount(ds DATASOURCE, tablename string) (ncount int, err error) {
	countsql := "select count(1) as ncount from " + tablename
	db, err := newMssqlConn(ds)
	if err != nil {
		Getlog().Error("newMssqlConn()==>" + err.Error())
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
