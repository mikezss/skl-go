package models

import (
	_ "errors"
	"fmt"
	"log"
	_ "strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`

type CMN_USER_TB struct {
	Userid                  string    `orm:"pk;column(userid)"`
	Username                string    `orm:"column(username);null;"`
	Adminid                 string    `orm:"column(adminid);null;"`
	Orgid                   string    `orm:"column(orgid);null;"`
	Password                string    `orm:"column(password)"`
	Userlevel               string    `orm:"column(userlevel);null;"`
	Expireddate             time.Time `orm:"column(expireddate);null;"`
	Logintime               time.Time `orm:"column(logintime);null;"`
	Loginip                 string    `orm:"column(loginip);null;"`
	Lasttime                time.Time `orm:"column(lasttime);null;"`
	Lastip                  string    `orm:"column(lastip);null;"`
	Skin                    string    `orm:"column(skin);null;"`
	Langcode                string    `orm:"column(langcode);null;"`
	Sex                     string    `orm:"column(sex);null;"`
	Birthday                time.Time `orm:"column(birthday);null;"`
	Idcard                  string    `orm:"column(idcard);null;"`
	School                  string    `orm:"column(school);null;"`
	Graduation              string    `orm:"column(graduation);null;"`
	Degree                  string    `orm:"column(degree);null;"`
	Major                   string    `orm:"column(major);null;"`
	Country                 string    `orm:"column(country);null;"`
	Province                string    `orm:"column(province);null;"`
	City                    string    `orm:"column(city);null;"`
	Address                 string    `orm:"column(address);null;"`
	Postcode                string    `orm:"column(postcode);null;"`
	Phone                   string    `orm:"column(phone);null;"`
	Fax                     string    `orm:"column(fax);null;"`
	Mobile                  string    `orm:"column(mobile);null;"`
	Email                   string    `orm:"column(email);null;"`
	Remark                  string    `orm:"column(remark);null;"`
	Creator                 string    `orm:"column(creator);null;"`
	Createtime              time.Time `orm:"column(createtime);null;"`
	Modifier                string    `orm:"column(modifier);null;"`
	Modifytime              time.Time `orm:"column(modifytime);null;"`
	Usertype                string    `orm:"column(usertype);null;"`
	Postid                  string    `orm:"column(postid);null;"`
	Isleader                bool      `orm:"column(isleader);null;default(false)"`
	Expired                 string    `orm:"column(expired);null;default(0)"`
	Ipconfig                string    `orm:"column(ipconfig);null;"`
	English_name            string    `orm:"column(english_name);null;"`
	Nationality             string    `orm:"column(nationality);null;"`
	Employeeid              string    `orm:"column(employeeid);null;"`
	Entrydate               time.Time `orm:"column(entrydate);null;"`
	Residence_addres        string    `orm:"column(residence_addres)"`
	Residence_type          string    `orm:"column(residence_type);null;"`
	Marital_status          string    `orm:"column(marital_status);null;"`
	Native_place            string    `orm:"column(native_place);null;"`
	Work_date               time.Time `orm:"column(work_date);null;"`
	Contact_way             string    `orm:"column(contact_way);null;"`
	Contact_person          string    `orm:"column(contact_person);null;"`
	Professional_title      string    `orm:"column(professional_title);null;"`
	Computer_level          string    `orm:"column(computer_level);null;"`
	Computer_cert           string    `orm:"column(computer_cert);null;"`
	English_level           string    `orm:"column(english_level);null;"`
	English_cert            string    `orm:"column(english_cert);null;"`
	Japanese_level          string    `orm:"column(japanese_level);null;"`
	Japanese_cert           string    `orm:"column(japanese_cert);null;"`
	Speciality              string    `orm:"column(speciality);null;"`
	Speciality_cert         string    `orm:"column(speciality_cert);null;"`
	Hobby_sport             string    `orm:"column(hobby_sport);null;"`
	Hobby_art               string    `orm:"column(hobby_art);null;"`
	Hobby_other             string    `orm:"column(hobby_other);null;"`
	Key_user                string    `orm:"column(key_user);null;"`
	Work_card               string    `orm:"column(work_card);null;"`
	Guard_card              string    `orm:"column(guard_card);null;"`
	Computer                string    `orm:"column(computer);null;"`
	Ext                     string    `orm:"column(ext);null;"`
	Msn                     string    `orm:"column(msn);null;"`
	Carborrow_qualification string    `orm:"column(carborrow_qualification);null;"`
	Rank                    string    `orm:"column(rank);null;"`
}

type PROFILE struct {
	Userid      string
	Username    string
	Companycode string
	Loginip     string
	Logintime   time.Time
}

type CMN_USERROLE_TB struct {
	Userid      string
	Roleid      string
	Hole        string
	Expireddate time.Time
}
type CMN_USERMODUAL_TB struct {
	Userid   string
	Modualid string
}

func (u *CMN_USER_TB) TableName() string {
	return "cmn_user_tb"
}

func AddCMN_USER_TB(u CMN_USER_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Delete(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.Insert(&u)

	insertsql := "insert into cmn_user_tb(userid,username,isleader,userlevel,orgid,expireddate,loginip,postid,email,remark,password) values(?,?,?,?,?,?,?,?,?,?,?)"
	//_, err = o.InsertMulti(1, &u)
	insertsql = ConvertSQL(insertsql, Getdbtype())
	_, err = o.Raw(insertsql, u.Userid, u.Username, u.Isleader, u.Userlevel, u.Orgid, u.Expireddate, u.Loginip, u.Postid, u.Email, u.Remark, "666666").Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}

func AddMultiCMN_USER_TB(u []CMN_USER_TB) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.InsertMulti(len(u), u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func GetAllCMN_USER_TB() (admins []CMN_USER_TB, err error) {
	admins = make([]CMN_USER_TB, 0)
	o := orm.NewOrm()

	sql := "select userid,username,adminid,orgid,password,userlevel,expireddate,logintime,loginip,lasttime,lastip,skin,langcode,sex,birthday,idcard,school,graduation,degree,major,country,province,city,address,postcode,phone,fax,mobile,email,remark,creator,createtime,modifier,modifytime,usertype,postid,isleader,expired,ipconfig,english_name,nationality,employeeid,entrydate,residence_addres,residence_type,marital_status,native_place,work_date,contact_way,contact_person,professional_title,computer_level,computer_cert,english_level,english_cert,japanese_level,japanese_cert,speciality,speciality_cert,hobby_sport,hobby_art,hobby_other,key_user,work_card,guard_card,computer,ext,msn from cmn_user_tb "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetCMN_USER_TB(u CMN_USER_TB) (admins []CMN_USER_TB, err error) {
	admins = make([]CMN_USER_TB, 0)
	o := orm.NewOrm()
	sql := "select * from cmn_user_tb where 1=1 "

	if u.Userid != "" {
		sql = sql + " and userid='" + u.Userid + "'"
	}

	if u.Username != "" {
		sql = sql + " and username like '%" + u.Username + "%'"
	}
	if u.Orgid != "" && u.Orgid != "0" && u.Orgid != "root" {
		sql = sql + " and orgid='" + u.Orgid + "'"
	}
	if u.Userlevel != "" {
		sql = sql + " and userlevel='" + u.Userlevel + "'"
	}
	if u.Expired == "0" || u.Expired == "1" {
		sql = sql + " and expired='" + u.Expired + "'"
	}

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func Getuseroptionsbyorgid(u CMN_ORG_TB) (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()
	sql := "select userid as value,username as label from cmn_user_tb where 1=1 and orgid=?"

	_, err = o.Raw(sql, u.Orgid).QueryRows(&admins)

	return admins, err
}
func GetALLCMN_USER_TB() (admins []CMN_USER_TB, err error) {
	admins = make([]CMN_USER_TB, 0)
	o := orm.NewOrm()
	sql := "select userid,username from cmn_user_tb where expired='0' "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetALLCMN_USER_TBoptions() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()
	sql := "select userid as value,username as label from cmn_user_tb where expired='0' "

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func DeleteCMN_USER_TB(u CMN_USER_TB) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(&u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}

func UpdateCMN_USER_TB(u CMN_USER_TB) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Update(&u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	}
	//	err = FirstHR_LEAVE_YEAR(o, u)
	//	if err != nil {
	//		fmt.Println(err)
	//		o.Rollback()
	//	}
	err = o.Commit()

	return err

}
func Updatepassword(companycode, userid, password string) error {

	o := orm.NewOrm()

	err := o.Begin()
	sql := "update cmn_user_tb set password=? where companycode=? and userid=?"
	sql = ConvertSQL(sql, Getdbtype())
	rs := o.Raw(sql, password, companycode, userid)
	_, err = rs.Exec()

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}
func Checklogin(companycode, userid, password string) bool {

	user := make([]CMN_USER_TB, 0)
	o := orm.NewOrm()

	sql := "select userid,password from cmn_user_tb  where userid=? and password=?"
	//rs := o.Raw(sql, companycode, userid, password)
	//sql = ConvertSQL(sql, Getdbtype())
	rows, _ := o.Raw(sql, userid, password).QueryRows(&user)

	return rows > 0

}

func GetCMN_USERROLE_TB(u CMN_USER_TB) (admins []CMN_USERROLE_TB, err error) {
	admins = make([]CMN_USERROLE_TB, 0)
	o := orm.NewOrm()
	sql := "select userid,roleid from cmn_userrole_tb where userid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid).QueryRows(&admins)

	return admins, err
}
func GetALLCMN_USERROLE_TB(u CMN_USER_TB) (admins []CMN_USERROLE_TB, err error) {
	admins = make([]CMN_USERROLE_TB, 0)
	o := orm.NewOrm()
	sql := "select userid,roleid from cmn_userrole_tb where userid=?  union SELECT  b.userid,a.roleid  FROM  cmn_grouprole_tb a inner join  cmn_usergroup_tb b on a.groupid=b.groupid where b.userid=? union SELECT b.userid,a.roleid FROM  cmn_orgrole_tb a inner join  cmn_user_tb b on a.orgid=b.orgid where b.userid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid, u.Userid, u.Userid).QueryRows(&admins)

	return admins, err
}
func GetALLCMN_USERROLEMODUAL_TB(u CMN_USER_TB) (admins2 []CMN_MODUAL_TB, err error) {
	admins := make([]CMN_MODUAL_TB, 0)
	admins2 = make([]CMN_MODUAL_TB, 0)
	o := orm.NewOrm()
	sql := " select distinct a.modualid,parentid,modualname,url,remark from cmn_modual_tb a "
	sql = sql + " inner join cmn_roleprivilege_tb b on a.modualid=b.modualid inner join "
	sql = sql + " (select  userid,roleid from cmn_userrole_tb where userid=? "
	sql = sql + " union SELECT   b.userid,a.roleid  FROM  cmn_grouprole_tb a inner join  cmn_usergroup_tb b on a.groupid=b.groupid where b.userid=? "
	sql = sql + " union SELECT  b.userid,a.roleid FROM  cmn_orgrole_tb a inner join  cmn_user_tb b on a.orgid=b.orgid where b.userid=?) c on b.roleid=c.roleid "
	sql = sql + " union all select a.modualid,parentid,modualname,url,remark from cmn_modual_tb a where parentid='root' "
	sql = sql + " union select a.modualid,parentid,modualname,url,remark from cmn_modual_tb a where url is null or url=' ' or url=''"
	sql = sql + " union select distinct a.modualid,parentid,modualname,url,remark from cmn_modual_tb a "
	sql = sql + " inner join cmn_roleprivilege_tb b on a.modualid=b.modualid where b.roleid='root'"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid, u.Userid, u.Userid).QueryRows(&admins)
	for _, admin := range admins {
		if admin.Url != "" {
			admins2 = append(admins2, admin)
		} else {
			if isparentmodual(admin.Modualid, admins) {
				admins2 = append(admins2, admin)
			}
		}
	}

	return admins2, err
}
func GetNAVIGATORMODUALBYUSER(u CMN_USER_TB) (admins2 []CMN_MODUAL_TB, err error) {
	admins := make([]CMN_MODUAL_TB, 0)
	admins2 = make([]CMN_MODUAL_TB, 0)
	o := orm.NewOrm()
	sql := "select distinct d.modualid,d.parentid,d.modualname,d.url,d.remark from cmn_modual_tb a "
	sql = sql + " inner join cmn_roleprivilege_tb b on a.modualid=b.modualid "
	sql = sql + " inner join (select  userid,roleid from cmn_userrole_tb where userid=? "
	sql = sql + " union SELECT   b.userid,a.roleid  FROM  cmn_grouprole_tb a inner join  cmn_usergroup_tb b on a.groupid=b.groupid where b.userid=? "
	sql = sql + " union SELECT  b.userid,a.roleid FROM  cmn_orgrole_tb a inner join  cmn_user_tb b on a.orgid=b.orgid where b.userid=? "
	sql = sql + " union select '',roleid from  cmn_roleprivilege_tb  where roleid='root')c on b.roleid=c.roleid "
	sql = sql + " inner join cmn_modual_tb d on d.modualid=a.parentid "
	sql = sql + " where a.url!='' and a.url is not null order by d.displayno"

	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid, u.Userid, u.Userid).QueryRows(&admins)
	for _, admin := range admins {
		admins2 = append(admins2, admin)
	}

	return admins2, err
}
func GetMENUMODUALBYPARENT(u CMN_USERMODUAL_TB) (admins2 []CMN_MODUAL_TB, err error) {
	admins := make([]CMN_MODUAL_TB, 0)
	admins2 = make([]CMN_MODUAL_TB, 0)
	o := orm.NewOrm()
	sql := "select distinct a.modualid,a.parentid,a.modualname,a.url,a.remark from cmn_modual_tb a "
	sql = sql + " inner join cmn_roleprivilege_tb b on a.modualid=b.modualid "
	sql = sql + " inner join (select  userid,roleid from cmn_userrole_tb where userid=? "
	sql = sql + " union SELECT   b.userid,a.roleid  FROM  cmn_grouprole_tb a inner join  cmn_usergroup_tb b on a.groupid=b.groupid where b.userid=? "
	sql = sql + " union SELECT  b.userid,a.roleid FROM  cmn_orgrole_tb a inner join  cmn_user_tb b on a.orgid=b.orgid where b.userid=? "
	sql = sql + " union select '',roleid from  cmn_roleprivilege_tb  where roleid='root')c on b.roleid=c.roleid "
	sql = sql + " where a.url!='' and a.url is not null and a.parentid=? order by a.displayno"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid, u.Userid, u.Userid, u.Modualid).QueryRows(&admins)
	for _, admin := range admins {
		if admin.Url != "" {
			admins2 = append(admins2, admin)
		} else {
			if isparentmodual(admin.Modualid, admins) {
				admins2 = append(admins2, admin)
			}
		}
	}

	return admins2, err
}
func isparentmodual(modualid string, admins []CMN_MODUAL_TB) bool {
	for _, admin := range admins {
		if admin.Parentid == modualid {
			return true
			break
		}
	}
	return false
}
func DeleteCMN_USERROLE_TB(u CMN_USER_TB) (err error) {

	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete from cmn_userrole_tb where userid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Userid).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	err = o.Commit()

	return err
}
func AddMultiCMN_USERROLE_TB(userid string, u []CMN_USERROLE_TB) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	sql := "delete  from cmn_userrole_tb where userid=?"
	dbtype := Getdbtype()
	sql = ConvertSQL(sql, dbtype)
	_, err = o.Raw(sql, userid).Exec()

	if err != nil {
		o.Rollback()
		return err
	}
	sql = "insert into cmn_userrole_tb(userid,roleid) values(?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, userrole := range u {
		_, err = o.Raw(sql, userid, userrole.Roleid).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	return err
}
func GetLeaders() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()
	dbtype := Getdbtype()
	sql := "select userid as value,username as label from cmn_user_tb where isleader=1 "
	if dbtype == "postgres" {
		sql = "select userid as value,username as label from cmn_user_tb where isleader=true "
	}
	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
func GetUsersbyorgid(cot CMN_ORG_TB) (admins []CMN_USER_TB, err error) {
	admins = make([]CMN_USER_TB, 0)
	o := orm.NewOrm()
	sql := "select userid,username,expired from cmn_user_tb where orgid=? "
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, cot.Orgid).QueryRows(&admins)

	return admins, err
}
func PASSWORDCHANGE(u CMN_USER_TB) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	sql := "update cmn_user_tb set password=? where userid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Password, u.Userid).Exec()

	if err != nil {
		o.Rollback()
		return err
	}

	err = o.Commit()

	return err
}
func PASSWORDRESET(u CMN_USER_TB) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	sql := "update cmn_user_tb set password=? where userid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Password, u.Userid).Exec()

	if err != nil {
		o.Rollback()
		return err
	}

	err = o.Commit()

	return err
}
func Getexcelfileinfo(filepath string) (fileinfomaparr []map[string]interface{}, err error) {
	fileinfomaparr = make([]map[string]interface{}, 0)

	colnames := make([]string, 0)
	var file *xlsx.File
	var sheet *xlsx.Sheet
	//var row *xlsx.Row
	//var cell *xlsx.Cell
	file, err = xlsx.OpenFile(filepath)

	if err != nil {
		return nil, err
	}
	sheet = file.Sheets[0]
	for index, row := range sheet.Rows {
		fileinfomap := make(map[string]interface{})
		if row == nil {
			continue
		}

		if index == 1 {
			continue
		}
		for index2, cell := range row.Cells {

			str, err := cell.FormattedValue()
			if err != nil {
				return nil, err
			}

			if str == "" {
				break
			}
			if index == 0 {
				colnames = append(colnames, str)
				continue
			}
			colname := colnames[index2]
			fileinfomap[colname] = str

		}
		fileinfomaparr = append(fileinfomaparr, fileinfomap)
	}
	return fileinfomaparr, nil
}
func Uploadusers(filepath CMN_FILEINFO_TB) (err1 error) {
	fileinfomaparr, err := Getexcelfileinfo(filepath.Filepath)

	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(fileinfomaparr)
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	sql := "select count(1) as ncount from cmn_user_tb where userid=?"
	updatesql := "update cmn_user_tb set username=?,isleader=?,userlevel=?,orgid=?,postid=? where userid=?"
	insertsql := "insert into cmn_user_tb(userid,username,isleader,userlevel,orgid,postid) values(?,?,?,?,?,?)"
	sql = ConvertSQL(sql, Getdbtype())
	updatesql = ConvertSQL(updatesql, Getdbtype())
	insertsql = ConvertSQL(insertsql, Getdbtype())
	var userid, username, isleader, userlevel, orgid, postid string
	var isleaderbool bool = false
	for _, fileinfomap := range fileinfomaparr {
		userid = ""
		username = ""
		isleader = ""
		userlevel = ""
		orgid = ""
		postid = ""
		isleaderbool = false
		var ncount int = 0
		if fileinfomap["Userid"] == nil {
			continue
		}
		userid = fileinfomap["Userid"].(string)

		err = o.Raw(sql, userid).QueryRow(&ncount)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		username = fileinfomap["Username"].(string)
		isleader = fileinfomap["Isleader"].(string)

		if isleader == "是" {
			isleaderbool = true
		}
		userlevel = fileinfomap["Userlevel"].(string)
		switch userlevel {
		case "一般":
			userlevel = "0"
		case "管理员":
			userlevel = "1"
		case "超级用户":
			userlevel = "2"
		case "开发者":
			userlevel = "3"
		case "PM":
			userlevel = "4"
		default:
			userlevel = "0"
		}
		orgid = fileinfomap["Orgid"].(string)
		postid = fileinfomap["Postid"].(string)
		switch postid {
		case "总经理":
			postid = "1"
		case "副总经理":
			postid = "2"
		case "部长":
			postid = "3"
		case "副部长":
			postid = "4"
		case "科长":
			postid = "5"
		case "副课长":
			postid = "6"
		case "主查":
			postid = "7"
		case "一般员工":
			postid = "8"
		case "特殊":
			postid = "9"
		case "总管":
			postid = "10"
		case "本部长":
			postid = "11"
		case "上海总经理":
			postid = "12"
		case "上海副总经理":
			postid = "13"
		case "广州总经理":
			postid = "14"
		case "广州副总经理":
			postid = "15"
		case "室长":
			postid = "16"
		case "副室长":
			postid = "17"
		case "代理课长":
			postid = "18"
		case "代理部长":
			postid = "19"
		default:
			postid = "7"
		}
		if ncount > 0 {
			_, err = o.Raw(updatesql, username, isleaderbool, userlevel, orgid, postid, userid).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		} else {
			_, err = o.Raw(insertsql, userid, username, isleaderbool, userlevel, orgid, postid).Exec()
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
