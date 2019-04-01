package main

import (
	_ "github.com/mikezss/skl-go/routers"

	_ "os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"fmt"
	"log"
	"github.com/mikezss/skl-go/controllers"
	"github.com/mikezss/skl-go/models"
	"strings"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.AddFuncMap("Calculate", models.Calculate)
	beego.AddFuncMap("Unescaped", models.Unescaped)
	beego.AddFuncMap("Unescapedjs", models.Unescapedjs)
	beego.AddFuncMap("UnescapedJSStr", models.UnescapedJSStr)
	beego.AddFuncMap("Tofirstupper", models.Tofirstupper)
	beego.AddFuncMap("Tolower", models.Tolower)
	beego.AddFuncMap("Replace", models.Replace)
	beego.AddFuncMap("Toupper", models.Toupper)
	beego.AddFuncMap("Mod", models.Mod)
	beego.AddFuncMap("OutputFN", models.OutputFN)

	beego.Run()
}
func init() {
	dbtype := registerDB()
	createTable(dbtype)

	orm.Debug = true

	//静态文件处理
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/file", "file")
	//beego.SetViewsPath("/views")
	/*CORP,增加一个跨域的header：*/
	//context.NewOutput().Header("Access-Control-Allow-Origin", "*")
	//	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//		AllowOrigins:     []string{"http://*"},
	//		AllowMethods:     []string{"PUT", "PATCH"},
	//		AllowHeaders:     []string{"Origin"},
	//		ExposeHeaders:    []string{"Content-Length"},
	//		AllowCredentials: true,
	//	}))
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//settingLocales()
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

}
func registerDB() string {
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		log.Fatal(err)
	}
	dbtype := iniconf.String("dbtype")
	fmt.Println("dbtype:-->" + dbtype)
	datasourcename := ""
	ds := make([]string, 0)
	switch dbtype {
	case "sqlite3":
		//oa.db

		datasourcename = iniconf.String(dbtype + "::dbname")

		orm.RegisterDriver(dbtype, orm.DRSqlite)

	case "mysql":

		ds = append(ds, iniconf.String(dbtype+"::username")+":")
		ds = append(ds, iniconf.String(dbtype+"::password")+"@tcp(")
		ds = append(ds, iniconf.String(dbtype+"::hostname")+":")
		ds = append(ds, iniconf.String(dbtype+"::port")+")/")
		ds = append(ds, iniconf.String(dbtype+"::dbname")+"?charset=utf8")

		datasourcename = strings.Join(ds, "")
		fmt.Println(datasourcename)
		//datasourcename:-->root:root@tcp(localhost:3306)/skl-ticket?charset=utf8
		//tcp:localhost:3306*mydb/root/rootroot
		orm.RegisterDriver("mysql", orm.DRMySQL)

	case "postgres":
		//tcp:localhost:5432*mydb/postgres/postgres

		ds = append(ds, "user=")
		ds = append(ds, iniconf.String(dbtype+"::username")+" ")
		ds = append(ds, "password=")
		ds = append(ds, iniconf.String(dbtype+"::password")+" ")
		ds = append(ds, "dbname=")
		ds = append(ds, iniconf.String(dbtype+"::dbname")+" ")
		ds = append(ds, "host=")
		ds = append(ds, iniconf.String(dbtype+"::hostname")+" ")
		ds = append(ds, "port=")
		ds = append(ds, iniconf.String(dbtype+"::port")+" ")
		ds = append(ds, "sslmode=disable")

		datasourcename = strings.Join(ds, "")
		fmt.Println(datasourcename)
		orm.RegisterDriver(dbtype, orm.DRPostgres)

	}
	fmt.Println("datasourcename:-->" + datasourcename)
	orm.RegisterDataBase("default", dbtype, datasourcename)
	//orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxOpen)
	return "default"
}

//自动建表
func createTable(dbtype string) {
	name := dbtype                             //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		log.Fatal(err)
	}
	dbtype2 := iniconf.String("dbtype")
	fmt.Println("dbtype:-->" + dbtype2)

	inittable := iniconf.String("inittable")
	fmt.Println("inittable:-->" + inittable)

	if inittable == "true" {
		initTable(dbtype2)
	}

}

//
func initTable(dbtype string) {
	cctl := &controllers.COMMONController{}
	filepath := cctl.GetCurrentDirectory() + "/" + dbtype + ".sql"
	s, err := cctl.Readfile2string(filepath, "GBK")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	sqlarr := strings.Split(s, ";")
	o := orm.NewOrm()
	//err = o.Begin()

	for _, sql := range sqlarr {
		fmt.Println(sql)
		_, err = o.Raw(sql).Exec()

		if err != nil {
			fmt.Println(err)
			//o.Rollback()
			continue
		}
	}
	//o.Commit()
}
