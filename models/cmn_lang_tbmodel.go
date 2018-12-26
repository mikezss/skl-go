package models

import (
	_ "encoding/json"
	_ "errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	_ "time"

	"github.com/astaxie/beego/orm"
	_ "golang.org/x/text/encoding/simplifiedchinese"
	_ "golang.org/x/text/encoding/traditionalchinese"
	_ "golang.org/x/text/transform"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`

type CMN_LANG_TB struct {
	Langid   string `orm:"pk;column(langid)"`
	Chinese  string `orm:"column(chinese);null"`
	Tchinese string `orm:"column(tchinese);null"`
	English  string `orm:"column(english);null"`
	Japanese string `orm:"column(jpanese);null"`
}
type LANG struct {
	Langid   string
	Langname string
}
type LANGPAGEINDEX struct {
	Langid    string
	Langname  string
	Pageindex int
	Pagesize  int
}

func (u *CMN_LANG_TB) TableName() string {
	return "cmn_lang_tb"
}

func AddCMN_LANG_TB(u *CMN_LANG_TB) error {
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

func AddMultiCMN_LANG_TB(u []CMN_LANG_TB) error {
	o := orm.NewOrm()
	err := o.Begin()

	//_, err = o.InsertMulti(len(u), &u)
	sql := "select count(1) as ncount from cmn_lang_tb where langid=?"
	updatesql := "update cmn_lang_tb set chinese=?,tchinese=?,english=?,jpanese=? where langid=?"
	insertsql := "insert into cmn_lang_tb(langid,chinese,tchinese,english,jpanese) values(?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		ncount := 0
		err = o.Raw(sql, u1.Langid).QueryRow(&ncount)
		if ncount > 0 {
			_, err = o.Raw(updatesql, u1.Chinese, u1.Tchinese, u1.English, u1.Japanese, u1.Langid).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		} else {
			_, err = o.Raw(insertsql, u1.Langid, u1.Chinese, u1.Tchinese, u1.English, u1.Japanese).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}

	}
	err = o.Commit()
	return err
}

func GetAllCMN_LANG_TB() (admins []CMN_LANG_TB, err error) {
	admins = make([]CMN_LANG_TB, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_lang_tb order by langid "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetAllCMN_LANG_TBbypageindex(l LANGPAGEINDEX) (admins []CMN_LANG_TB, err error) {
	dbtype := Getdbtype()
	admins = make([]CMN_LANG_TB, 0)
	o := orm.NewOrm()

	sql := "select * from cmn_lang_tb where 1=1 "
	if l.Langid != "" {
		sql = sql + " and langid like '%" + l.Langid + "%'"
	}
	if l.Langname != "" {
		sql = sql + " and (chinese like '%" + l.Langname + "%'"
		sql = sql + " or tchinese like '%" + l.Langname + "%'"
		sql = sql + " or english like '%" + l.Langname + "%'"
		sql = sql + " or jpanese like '%" + l.Langname + "%')"
	}
	var limitstr string = " limit "
	if dbtype == "postgres" {
		limitstr = limitstr + strconv.Itoa(l.Pagesize) + " offset " + strconv.Itoa((l.Pageindex-1)*l.Pagesize)

	} else if dbtype == "mysql" {
		limitstr = limitstr + strconv.Itoa((l.Pageindex-1)*l.Pagesize) + "," + strconv.Itoa(l.Pagesize)

	} else {
		limitstr = limitstr + strconv.Itoa((l.Pageindex-1)*l.Pagesize) + "," + strconv.Itoa(l.Pagesize)

	}
	sql = sql + limitstr
	_, err = o.Raw(ConvertSQL(sql, dbtype)).QueryRows(&admins)

	return admins, err
}
func GetAllCMN_LANG_TBcount(l LANG) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "select count(1) as total from cmn_lang_tb where 1=1 "
	if l.Langid != "" {
		sql = sql + " and langid like '%" + l.Langid + "%'"
	}
	if l.Langname != "" {
		sql = sql + " and (chinese like '%" + l.Langname + "%'"
		sql = sql + " or tchinese like '%" + l.Langname + "%'"
		sql = sql + " or english like '%" + l.Langname + "%'"
		sql = sql + " or jpanese like '%" + l.Langname + "%')"
	}

	err = o.Raw(sql).QueryRow(&page)

	return page, err
}

func GetCMN_LANG_TB(u *CMN_LANG_TB) (admins []CMN_LANG_TB, err error) {
	admins = make([]CMN_LANG_TB, 0)
	o := orm.NewOrm()
	sql := "select langid,langname from cmn_lang_tb where 1=1 "

	if u.Langid != "" {
		sql = sql + " and langid like '%" + u.Langid + "%'"
	}

	if u.Chinese != "" {
		sql = sql + " and chinese like '%" + u.Chinese + "%'"
	}
	if u.Tchinese != "" {
		sql = sql + " and tchinese like '%" + u.Tchinese + "%'"
	}
	if u.English != "" {
		sql = sql + " and english like '%" + u.English + "%'"
	}
	if u.Japanese != "" {
		sql = sql + " and jpanese like '%" + u.Japanese + "%'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func DeleteCMN_LANG_TB(u *CMN_LANG_TB) error {

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

func UpdateCMN_LANG_TB(u *CMN_LANG_TB) error {

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
func Updatelangjson(u2 []CMN_LANG_TB) error {
	project, _ := GetProjectpath()
	ngpath := project.Ngdirectory
	enjsonpath := ngpath + "/src/assets/i18n/English.json"
	zhjsonpath := ngpath + "/src/assets/i18n/Chinese.json"
	jpjsonpath := ngpath + "/src/assets/i18n/Japanese.json"
	tzhjsonpath := ngpath + "/src/assets/i18n/Tchinese.json"
	f, err := os.OpenFile(enjsonpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	filecontent := ""
	filecontent = filecontent + "{"
	u, err2 := GetAllCMN_LANG_TB()
	if err2 != nil {
		fmt.Println("WriteString")
		fmt.Println(err2)
		return err2
	}
	for idx, clt := range u {
		if idx < len(u)-1 {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + clt.English + "\",\n"
		} else {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + clt.English + "\"\n"
		}

	}
	filecontent = filecontent + "}\r\n"
	_, err = f.WriteString(filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}

	f1, err1 := os.OpenFile(zhjsonpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f1.Close()
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	filecontent = ""
	filecontent = filecontent + "{"

	for idx, clt := range u {
		if idx < len(u)-1 {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Chinese) + "\",\n"
		} else {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Chinese) + "\"\n"
		}
	}
	filecontent = filecontent + "}\r\n"
	_, err1 = f1.WriteString(filecontent)
	if err1 != nil {
		fmt.Println("WriteString")
		fmt.Println(err1)
		return err1
	}
	f2, err2 := os.OpenFile(jpjsonpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f2.Close()
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	filecontent = ""
	filecontent = filecontent + "{"

	for idx, clt := range u {
		if idx < len(u)-1 {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Japanese) + "\",\n"
		} else {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Japanese) + "\"\n"
		}
	}
	filecontent = filecontent + "}\r\n"
	_, err2 = f2.WriteString(filecontent)
	if err2 != nil {
		fmt.Println("WriteString")
		fmt.Println(err2)
		return err2
	}
	f3, err3 := os.OpenFile(tzhjsonpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f3.Close()
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	filecontent = ""
	filecontent = filecontent + "{"

	for idx, clt := range u {
		if idx < len(u)-1 {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Tchinese) + "\",\n"
		} else {
			filecontent = filecontent + "\"" + clt.Langid + "\":\"" + ConvertGBK2unicode(clt.Tchinese) + "\"\n"
		}
	}
	filecontent = filecontent + "}\r\n"
	_, err3 = f3.WriteString(filecontent)
	if err3 != nil {
		fmt.Println("WriteString")
		fmt.Println(err3)
		return err3
	}
	return nil
}
func Loadlangjson() error {

	project, _ := GetProjectpath()
	ngpath := project.Ngdirectory
	enjsonpath := ngpath + "/src/assets/i18n/English.json"
	zhjsonpath := ngpath + "/src/assets/i18n/Chinese.json"
	jpjsonpath := ngpath + "/src/assets/i18n/Japanese.json"
	tzhjsonpath := ngpath + "/src/assets/i18n/Tchinese.json"
	data, err := Readfile2string(enjsonpath, "utf8")
	if err != nil {
		return err
	}
	reg := regexp.MustCompile(`\{|\}|\"|\"`)
	data = reg.ReplaceAllString(data, "")
	dataarr := strings.Split(data, ",")
	clbarr := make([]CMN_LANG_TB, 0)
	for _, data1 := range dataarr {
		fmt.Print(data1)
		data1arr := strings.Split(data1, ":")
		clb := CMN_LANG_TB{Langid: strings.TrimSpace(data1arr[0]), English: strings.TrimSpace(data1arr[1])}
		clbarr = append(clbarr, clb)
	}
	err = InsertorUpdatelang(clbarr, "english")

	data, err = Readfile2string(jpjsonpath, "utf8")
	if err != nil {
		return err
	}
	reg = regexp.MustCompile(`\{|\}|\"|\"`)
	data = reg.ReplaceAllString(data, "")
	dataarr = strings.Split(data, ",")
	clbarr = make([]CMN_LANG_TB, 0)
	for _, data1 := range dataarr {
		data1arr := strings.Split(data1, ":")
		context := Convertunicode2GBK(data1arr[1])
		fmt.Print("context------->" + context)
		clb := CMN_LANG_TB{Langid: strings.TrimSpace(data1arr[0]), Japanese: context}
		clbarr = append(clbarr, clb)
	}
	err = InsertorUpdatelang(clbarr, "jpanese")

	data, err = Readfile2string(tzhjsonpath, "utf8")
	if err != nil {
		return err
	}
	reg = regexp.MustCompile(`\{|\}|\"|\"`)
	data = reg.ReplaceAllString(data, "")
	dataarr = strings.Split(data, ",")
	clbarr = make([]CMN_LANG_TB, 0)
	for _, data1 := range dataarr {
		data1arr := strings.Split(data1, ":")
		context := Convertunicode2GBK(data1arr[1])
		fmt.Print("context------->" + context)
		clb := CMN_LANG_TB{Langid: strings.TrimSpace(data1arr[0]), Tchinese: context}
		clbarr = append(clbarr, clb)
	}
	err = InsertorUpdatelang(clbarr, "tchinese")

	data, err = Readfile2string(zhjsonpath, "utf8")
	if err != nil {
		return err
	}

	reg = regexp.MustCompile(`\{|\}|\"|\"`)
	data = reg.ReplaceAllString(data, "")
	dataarr = strings.Split(data, ",")
	clbarr = make([]CMN_LANG_TB, 0)
	for _, data1 := range dataarr {
		data1arr := strings.Split(data1, ":")
		context := Convertunicode2GBK(data1arr[1])
		fmt.Print("context------->" + context)
		clb := CMN_LANG_TB{Langid: strings.TrimSpace(data1arr[0]), Chinese: context}
		clbarr = append(clbarr, clb)
	}
	err = InsertorUpdatelang(clbarr, "chinese")

	return err

}
func InsertorUpdatelang(cltarr []CMN_LANG_TB, langtype string) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	sql := "select count(1) as ncount from cmn_lang_tb where langid=?"
	sql = ConvertSQL(sql, Getdbtype())
	updatesql := "update cmn_lang_tb set " + langtype + "=? where langid=?"
	insertsql := "insert into cmn_lang_tb (langid," + langtype + ") values(?,?)"
	for _, clt := range cltarr {
		ncount := 0
		err = o.Raw(sql, clt.Langid).QueryRow(&ncount)
		if ncount > 0 {
			switch langtype {
			case "chinese":
				_, err = o.Raw(updatesql, clt.Chinese, clt.Langid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "english":
				_, err = o.Raw(updatesql, clt.English, clt.Langid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "jpanese":
				_, err = o.Raw(updatesql, clt.Japanese, clt.Langid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "tchinese":
				_, err = o.Raw(updatesql, clt.Tchinese, clt.Langid).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}
		} else {
			switch langtype {
			case "chinese":
				_, err = o.Raw(insertsql, clt.Langid, clt.Chinese).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "english":
				_, err = o.Raw(insertsql, clt.Langid, clt.English).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "jpanese":
				_, err = o.Raw(insertsql, clt.Langid, clt.Japanese).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			case "tchinese":
				_, err = o.Raw(insertsql, clt.Langid, clt.Tchinese).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}
		}
	}
	err = o.Commit()
	return err
}
func Convertunicode2GBK(unicodestr string) string {
	sUnicodev := strings.Split(strings.TrimSpace(unicodestr), "\\u")
	var context string
	for _, v := range sUnicodev {
		v1 := ""
		v2 := ""
		if len(v) < 1 {
			continue
		}
		if len(v) > 4 {
			v1 = v[0:4]
			v2 = v[4:]
		} else {
			v1 = v
			v2 = ""
		}

		temp, err := strconv.ParseInt(v1, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp) + v2
	}
	return context
}
func ConvertGBK2unicode(gbkstr string) string {
	textQuoted := strconv.QuoteToASCII(gbkstr)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}
