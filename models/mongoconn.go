package models

import (
	_ "errors"

	"fmt"
	"log"
	_ "reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const OplogNS = "oplog.rs"

var con *MongoConn

type MongoConn struct {
	Session *mgo.Session
	URL     string
}

type Result struct {
	Key   string
	Value string
	Type  string
}

type COLLECTIONINFO struct {
	Dbname         string
	Collectionname string
	Ncount         int
	Elements       []bson.DocElem
}

func GetMongoConn(ds DATASOURCE) (*MongoConn, error) {
	var err error
	var c1 *MongoConn
	if con == nil {
		c1, err = newMongoConn(ds)
		if err == nil {
			con = c1
		}

	}
	return con, err
}

func newMongoConn(ds DATASOURCE) (*MongoConn, error) {
	mode := "strong"
	url := ""

	hostname := ds.Ip
	port := ds.Port
	username := ds.Username
	password := ds.Password
	if username != "" {
		url = "mongodb://" + username + ":" + password + "@" + hostname + ":" + strconv.Itoa(port)
	} else {
		url = "mongodb://" + hostname + ":" + strconv.Itoa(port)
	}
	session, err := mgo.Dial(url)
	if err != nil {
		Getlog().Error("Connect to %s failed=>" + url + err.Error())
		return nil, err
	}
	// maximum pooled connections. the overall established sockets
	// should be lower than this value(will block otherwise)
	session.SetPoolLimit(256)
	session.SetSocketTimeout(10 * time.Minute)

	if err := session.Ping(); err != nil {
		Getlog().Error("session.Ping=>" + url + err.Error())
		return nil, err
	}

	// Switch the session to a eventually behavior. In that case session
	// may read for any secondary node. default mode is mgo.Strong
	switch mode {
	case "primary":
		session.SetMode(mgo.Primary, true)
		break
	case "secondary":
		session.SetMode(mgo.SecondaryPreferred, true)
		break
	default:
		session.SetMode(mgo.Strong, true)
		break

	}

	log.Printf("New session to %s successfully", url)
	return &MongoConn{Session: session, URL: url}, nil
}

func (conn *MongoConn) Close() {
	log.Printf("Close session with %s", conn.URL)
	conn.Session.Close()
}

func (conn *MongoConn) IsGood() bool {
	if err := conn.Session.Ping(); err != nil {
		return false
	}

	return true
}
func (conn *MongoConn) HasOplogNs() bool {
	if ns, err := conn.Session.DB("local").CollectionNames(); err == nil {
		for _, table := range ns {
			if table == OplogNS {
				return true
			}
		}
	}
	return false
}

func (conn *MongoConn) Getdatabasenames() (databases []string, err error) {
	if databases, err = conn.Session.DatabaseNames(); err != nil {
		Getlog().Error("Couldn't get databases from remote server==>" + err.Error())
		return nil, err
	}
	return databases, nil
}

//获得指定schema所有的collection名，3.0以上才能获得
func (conn *MongoConn) Getcollectionnames(databasename string) (collectionnames []string, err error) {
	//collectionnames = make([]map[string]interface{}, 0)
	//query := conn.Session.DB(databasename).C("system.namespaces").Find(nil)
	//err = query.All(&collectionnames)
	collectionnames, err = conn.Session.DB(databasename).CollectionNames()
	if err != nil {
		Getlog().Error("Couldn't get collections from database==>" + databasename + err.Error())
		return nil, err
	}

	return collectionnames, nil
}

//
func (conn *MongoConn) Getcollectionnames2(databasename string) (collectionnames []map[string]interface{}, err error) {
	collectionnames = make([]map[string]interface{}, 0)
	query := conn.Session.DB(databasename).C("tap.system.namespaces").Find(nil)

	err = query.All(&collectionnames)
	//collectionnames, err = conn.Session.DB(databasename).CollectionNames()
	if err != nil {
		Getlog().Error("Couldn't get collections from database==>" + databasename + err.Error())
		return nil, err
	}

	return collectionnames, nil
}

//获得指定schema指定collection的数据条数
func (conn *MongoConn) Getcollectionrows(databasename string, collectionname string) (cltinf COLLECTIONINFO, err error) {

	var ncount int

	ncount, err = conn.Session.DB(databasename).C(collectionname).Count()

	if err != nil {
		Getlog().Error("Couldn't get collections rows from database,collection==>" + databasename + collectionname + err.Error())
		ncount = -1
	}
	cltinf = COLLECTIONINFO{Dbname: databasename, Collectionname: collectionname, Ncount: ncount}

	return cltinf, err

}

//获得指定schema指定collection的所有数据
//返回数据为[]map[string]interface{}
func (conn *MongoConn) Getcollectiondatas(databasename string, collectionname string) (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 0)
	query := conn.Session.DB(databasename).C(collectionname).Find(nil)

	err = query.All(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil

}

//获得collection的所有key，不包含子key
func (conn *MongoConn) Getcollectionkeys(databasename string, collectionname string) (keys []string, err error) {
	keys = make([]string, 0)
	datas, err := conn.Getcollectiondatas(databasename, collectionname)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for idx, datamap := range datas {
		if idx > 100000 {
			break
		}
		for key, _ := range datamap {
			// fmt.Println(reflect.TypeOf(value))
			if !Isinarry(keys, key) {
				keys = append(keys, key)
			}

		}

	}

	return keys, nil

}

//通过mapreduce获得collection的所有key，不包含子key
// mr = db.runCommand({
//   "mapreduce" : "my_collection",
//   "map" : function() {
//     for (var key in this) { emit(key, null); }
//   },
//   "reduce" : function(key, stuff) { return null; },
//   "out": "my_collection" + "_keys"
// })

//     job := &mgo.MapReduce{
//             Map:      "function() { emit(this.n, 1) }",
//             Reduce:   "function(key, values) { return Array.sum(values) }",
//     }
//     _, err := collection.Find(nil).MapReduce(job, &result)
//     if err != nil {
//         return err
//     }
func (conn *MongoConn) Getcollectionkeysbymapreduce(databasename string, collectionname string) (keys []string, err error) {
	keys = make([]string, 0)

	job := &mgo.MapReduce{
		Map:    "function() { for (var key in this) { emit(key, null);} }",
		Reduce: "function(key, value) { return   Array.sum(key); }",
		Out:    "_keys",
	}
	//result := make([]map[string]interface{}, 0)
	var rsult2 interface{}
	_, err = conn.Session.DB(databasename).C(collectionname).Find(nil).MapReduce(job, &rsult2)
	if err != nil {
		return nil, err
	}
	fmt.Println(rsult2)
	//db[mr.result].distinct("_id")
	// for key, value := range result {
	// 	fmt.Println(key)
	// 	fmt.Println(value)
	// }
	return keys, nil

}

//获得collection的字段信息，包括
//字段名称、字段类型、字段长度、是否主键、是否允许为空、是否为父、是否自动增长
func (conn *MongoConn) Getcollectionfieldinfo(databasename string, collectionname string) (fieldinfomap map[string]DATASOURCETABLEFIELD, err error) {
	mongodbkeys := getmongodbkeys()
	fieldinfomap = make(map[string]DATASOURCETABLEFIELD, 0)
	datas, err := conn.Getcollectiondatas(databasename, collectionname)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Getlog().Debug("conn.Getcollectiondatas==>" + orm.ToStr(datas))
	for i := 0; i < len(datas); i++ {

		datamap := datas[i]

		for key, value := range datamap {
			_, ok := fieldinfomap[key]
			if !ok {
				isprimary := "0"
				isauto := "0"
				isnull := "1"
				isparent := "0"
				fieldlength := 255
				fieldtype := ""
				fieldtype = Getfieldtype(value)
				if fieldtype == "Array" {
					isparent = "1"
				}
				if strings.Contains(mongodbkeys, key) {
					isprimary = "1"
					isauto = "1"
					isnull = "0"
				}

				dtf := DATASOURCETABLEFIELD{Datasource: databasename, Tablename: collectionname, Fieldname: key, Fieldtype: fieldtype, Isprimary: isprimary, Isnull: isnull, Isparent: isparent, Isauto: isauto, Fieldlength: fieldlength}
				fieldinfomap[key] = dtf
			}

		}

	}

	return fieldinfomap, nil

}

//获得collection的字段信息，包括
//字段名称、字段类型、字段长度、是否主键、是否允许为空、是否为父、是否自动增长
func (conn *MongoConn) Getcollectionfieldinfochild(databasename string, collectionname string, fieldname string, maptype string) (fieldinfomap map[string]DATASOURCETABLEFIELDCHILD, err error) {
	mongodbkeys := getmongodbkeys()
	fieldinfomap = make(map[string]DATASOURCETABLEFIELDCHILD, 0)
	datas, err := conn.Getcollectiondatasbyfieldnames(databasename, collectionname, fieldname)
	//fmt.Println("len(datas)==>")
	//fmt.Println(len(datas))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for i := 0; i < len(datas); i++ {

		datamap := datas[i]

		for key, value := range datamap {
			//fmt.Println(key)
			//fmt.Println(value)
			if key != fieldname {
				continue
			}
			if maptype == "map" {
				valuemap := value.(map[string]interface{})
				for key2, value2 := range valuemap {
					_, ok := fieldinfomap[key2]
					if !ok {
						isprimary := "0"
						isauto := "0"
						isnull := "1"
						isparent := "0"
						fieldlength := 255
						fieldtype := ""
						fieldtype = Getfieldtype(value2)
						if fieldtype == "Array" {
							isparent = "1"
						}
						if strings.Contains(mongodbkeys, key2) {
							isprimary = "1"
							isauto = "0"
							isnull = "0"
						}

						dtf := DATASOURCETABLEFIELDCHILD{Datasource: databasename, Tablename: collectionname, Fieldname: fieldname, Childfieldname: key2, Fieldtype: fieldtype, Isprimary: isprimary, Isnull: isnull, Isparent: isparent, Isauto: isauto, Fieldlength: fieldlength}
						fieldinfomap[key2] = dtf
					}
				}
			} else if maptype == "maparray" {
				valuemap := value.([]interface{})
				// for _, valuemap2 := range valuemap {
				// 	valuemap3 := valuemap2.(map[string]interface{})
				// 	for key3, value3 := range valuemap3 {
				// 		_, ok := fieldinfomap[key3]
				// 		if !ok {
				// 			isprimary := "0"
				// 			isauto := "0"
				// 			isnull := "1"
				// 			isparent := "0"
				// 			fieldlength := 255
				// 			fieldtype := ""
				// 			fieldtype = Getfieldtype(value3)
				// 			if fieldtype == "Array" {
				// 				isparent = "1"
				// 			}
				// 			if key3 == "_id" || key3 == "idNo" || key3 == "key" {
				// 				isprimary = "1"
				// 				isauto = "0"
				// 				isnull = "0"
				// 			}

				// 			dtf := DATASOURCETABLEFIELDCHILD{Datasource: databasename, Tablename: collectionname, Fieldname: fieldname, Childfieldname: key3, Fieldtype: fieldtype, Isprimary: isprimary, Isnull: isnull, Isparent: isparent, Isauto: isauto, Fieldlength: fieldlength}
				// 			fieldinfomap[key3] = dtf
				// 		}
				// 	}
				// }
				loopfieldnames(databasename, collectionname, fieldname, &fieldinfomap, valuemap)
			}

		}

	}

	return fieldinfomap, nil

}

//获得mongodb的版本号 OK
func (conn *MongoConn) GetDBVersion() (string, error) {
	var result bson.M
	err := conn.Session.Run(bson.D{{"buildInfo", 1}}, &result)
	if err != nil {
		return "", err
	}

	if version, ok := result["version"]; ok {
		if s, ok := version.(string); ok {
			return s, nil
		}
		return "", fmt.Errorf("version type assertion error[%v]", version)
	}
	return "", fmt.Errorf("version not found")
}

//获得collection的索引
func (conn *MongoConn) Getcollectionindexs(databasename string, collectionname string) (index []mgo.Index, err error) {

	index, err = conn.Session.DB(databasename).C(collectionname).Indexes()

	if err != nil {
		Getlog().Error("Couldn't get collections indexs from database,collection" + databasename + collectionname + err.Error())
		return nil, err
	}

	return index, nil

}

//判断字符串数组中是否包含指定字符串
func Isinarry(target []string, source string) bool {
	for _, t := range target {
		if t == source {
			return true
		}
	}
	return false
}

//获得指定key的数据，如fieldnames为'a','b','c',获得a b c列数据
func (conn *MongoConn) Getcollectiondatasbyfieldnames(databasename string, collectionname string, fieldnames string, fdt ...FromDatasourceTableinfo2) (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 0)

	// only获得 name field数据

	projection := bson.D{}
	fieldnamesarr := strings.Split(fieldnames, ",")
	for _, fieldname := range fieldnamesarr {
		fieldname2 := strings.Split(fieldname, ".")
		if len(fieldname2) > 1 {
			if fieldname2[1] == "" || fieldname2[1] == "no" {
				projection = append(projection, bson.DocElem{Name: fieldname2[0], Value: 1})
			} else {
				projection = append(projection, bson.DocElem{Name: fieldname, Value: 1})
			}
		} else {
			projection = append(projection, bson.DocElem{Name: fieldname, Value: 1})
		}

	}
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {

		return nil, err
	}
	limitDataCountFortest := iniconf.DefaultInt("parameters::limitDataCountFortest", 0)
	var query *mgo.Query
	filter := ""
	if fdt != nil {
		if len(fdt) > 0 {
			filter, err = Getfromdatasourcefilter(fdt[0])

			if err != nil {
				Getlog().Error("Getfromdatasourcefilter(fdt[0])==>" + err.Error())
				return nil, err
			}

		}

	}

	if filter != "" {
		filterc := bson.D{}
		filterarr := strings.Split(filter, ",")
		for _, filtea := range filterarr {
			filterc = append(filterc, bson.DocElem{Name: "key", Value: filtea})
		}
		if limitDataCountFortest == 0 {
			query = conn.Session.DB(databasename).C(collectionname).Find(filterc).Select(projection)
		} else {
			query = conn.Session.DB(databasename).C(collectionname).Find(filterc).Select(projection).Limit(limitDataCountFortest)
		}
	} else {
		if limitDataCountFortest == 0 {
			query = conn.Session.DB(databasename).C(collectionname).Find(nil).Select(projection)
		} else {
			query = conn.Session.DB(databasename).C(collectionname).Find(nil).Select(projection).Limit(limitDataCountFortest)
		}
	}

	err = query.All(&result)

	if err != nil {
		Getlog().Error("query.All(&result)==>" + err.Error())
		return nil, err
	}
	//fmt.Println(result)
	return result, nil
}

//获得指定key的数据，如fieldnames为'a','b','c',获得a b c列数据
func (conn *MongoConn) Getcollectionchilddatasbyfieldnames(databasename string, collectionname string, parentfieldname, childfieldnames string) (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 0)

	projection := bson.D{}
	childfieldnamesarr := strings.Split(childfieldnames, ",")
	for _, childfieldname := range childfieldnamesarr {
		projection = append(projection, bson.DocElem{Name: parentfieldname + "." + childfieldname, Value: 1})
	}
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {

		return nil, err
	}
	limitDataCountFortest := iniconf.DefaultInt("parameters::limitDataCountFortest", 0)
	var query *mgo.Query

	if limitDataCountFortest == 0 {
		query = conn.Session.DB(databasename).C(collectionname).Find(nil).Select(projection)

	} else {
		query = conn.Session.DB(databasename).C(collectionname).Find(nil).Select(projection).Limit(limitDataCountFortest)

	}

	err = query.All(&result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

//获得指定key的一行数据，如fieldnames为'a','b','c',获得a b c列数据
func (conn *MongoConn) Getcollectiondatasbyfieldnamesonerow(databasename string, collectionname string, fieldnames string) (result []orm.Params, err error) {
	result = make([]orm.Params, 0)
	bm := make(bson.M, 0)
	fieldnamearr := strings.Split(fieldnames, ",")
	for _, fieldname := range fieldnamearr {
		bm[fieldname] = 1
	}
	// only获得 name field数据
	//err := collection.Find(nil).Select(bson.M{"name": 1}).One(&result)
	query := conn.Session.DB(databasename).C(collectionname).Find(nil).Select(bm).Limit(1)

	err = query.All(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil

}

//获得一行数据
func (conn *MongoConn) Getcollectiondatasrow(databasename string, collectionname string) (result []orm.Params, err error) {
	result = make([]orm.Params, 0)
	query := conn.Session.DB(databasename).C(collectionname).Find(nil).Limit(1)

	err = query.All(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil

}
func loopfieldnames(databasename, collectionname, fieldname string, fieldinfomap *map[string]DATASOURCETABLEFIELDCHILD, value interface{}) {
	mongodbkeys := getmongodbkeys()
	valuemap := value.([]interface{})
	for _, valuemap2 := range valuemap {
		valuemap3 := valuemap2.(map[string]interface{})
		for key3, value3 := range valuemap3 {
			_, ok := (*fieldinfomap)[key3]
			if !ok {
				isprimary := "0"
				isauto := "0"
				isnull := "1"
				isparent := "0"
				fieldlength := 255
				fieldtype := ""
				fieldtype = Getfieldtype(value3)
				if fieldtype == "Array" {
					isparent = "1"
				}
				if strings.Contains(mongodbkeys, key3) {
					isprimary = "1"
					isauto = "0"
					isnull = "0"
				}

				dtf := DATASOURCETABLEFIELDCHILD{Datasource: databasename, Tablename: collectionname, Fieldname: fieldname, Childfieldname: key3, Fieldtype: fieldtype, Isprimary: isprimary, Isnull: isnull, Isparent: isparent, Isauto: isauto, Fieldlength: fieldlength}
				(*fieldinfomap)[key3] = dtf
				if fieldtype == "Array" && key3 == "fieldname" {
					fmt.Println("loop Array==>" + key3)
					//fmt.Println(value3)
					loopfieldnames(databasename, collectionname, fieldname, fieldinfomap, value3)
				}
			}
		}
	}
}
func getmongodbkeys() string {
	iniconf, err := config.NewConfig("ini", "conf/myconf.ini")
	if err != nil {
		Getlog().Error(err.Error())
		return ""
	}
	return iniconf.String("parameters::mongodbkeys")

}

//filterkey:Country;fieldnames:  subCatalogs.displayEN,subCatalogs.displayZH
//filterkey:Country;fieldnames:  subCatalogs.displayEN,subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.displayZH
//filterkey:Country;fieldnames:  subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayZH
//通过key获得数据
func (conn *MongoConn) Getcollectiondataforregion(databasename, collectionname, filterkey string, fieldnames string) (result []map[string]interface{}, err error) {

	projection := bson.D{}
	fieldnamesarr := strings.Split(fieldnames, ",")
	for _, fieldname := range fieldnamesarr {

		projection = append(projection, bson.DocElem{Name: fieldname, Value: 1})

	}
	var query *mgo.Query

	filterc := bson.D{}

	filterc = append(filterc, bson.DocElem{Name: "key", Value: filterkey})

	query = conn.Session.DB(databasename).C(collectionname).Find(filterc).Select(projection)

	err = query.All(&result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(result)
	return result, nil
}

//filterkey:Country;fieldnames:  subCatalogs.displayEN,subCatalogs.displayZH
//filterkey:Country;fieldnames:  subCatalogs.displayEN,subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.displayZH
//filterkey:Country;fieldnames:  subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayEN,subCatalogs.subCatalogs.subCatalogs.displayZH
//获得父子表数据
func (conn *MongoConn) Getcollectiondataforparentchild(databasename, collectionname, fieldnames string, filterkey ...interface{}) (result []map[string]interface{}, err error) {

	projection := bson.D{}
	fieldnamesarr := strings.Split(fieldnames, ",")
	for _, fieldname := range fieldnamesarr {

		projection = append(projection, bson.DocElem{Name: fieldname, Value: 1})

	}
	var query *mgo.Query
	if filterkey != nil {

		if len(filterkey) > 0 {

			filterc := bson.D{}

			filterc = append(filterc, bson.DocElem{Name: "key", Value: filterkey[0].(string)})

			query = conn.Session.DB(databasename).C(collectionname).Find(filterc).Select(projection)
		}
	} else {
		query = conn.Session.DB(databasename).C(collectionname).Find(nil).Select(projection)
	}

	err = query.All(&result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(result)
	return result, nil
}
