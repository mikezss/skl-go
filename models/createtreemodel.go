package models

import (
	_ "errors"
	"fmt"
	_ "strconv"
	_ "time"

	_ "github.com/astaxie/beego/orm"
	_ "github.com/beego/i18n"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//var json string = ""

type Node struct {
	Key        string
	Parentid   string
	Title      string
	Url        string
	Checked    string
	Childrens  []string
	Orgtype    string
	Orglevel   string
	Rolelevel  string
	Grouplevel string
}
type Json struct {
	Treejson string
}

func CreateModualTreeJson() string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllModual()
	orgs2 := orgs
	//fmt.Println(orgs)

	for _, org := range orgs {
		//fmt.Println("parentid:" + org.Parentid)
		//childorgs, _ := Getchildrenorgids(org.Orgid)
		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Modualid {
				childorgs = append(childorgs, org21.Modualid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "false", Childrens: childorgs}
			m[org.Modualid] = node
		} else {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "false"}
			m[org.Modualid] = node
		}

	}
	fmt.Println(m)

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateModualTreeJsonForuser(user CMN_USER_TB) string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetALLCMN_USERROLEMODUAL_TB(user)
	orgs2 := orgs
	//fmt.Println(orgs)

	for _, org := range orgs {
		//fmt.Println("parentid:" + org.Parentid)
		//childorgs, _ := Getchildrenorgids(org.Orgid)
		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Modualid {
				childorgs = append(childorgs, org21.Modualid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "true", Childrens: childorgs}
			m[org.Modualid] = node
		} else {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "true"}
			m[org.Modualid] = node
		}

	}
	fmt.Println(m)

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateOrgTreeJson() string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllOrg()
	orgs2 := orgs
	//fmt.Println(orgs)

	for _, org := range orgs {
		//fmt.Println("parentid:" + org.Parentid)
		//childorgs, _ := Getchildrenorgids(org.Orgid)
		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Orgid {
				childorgs = append(childorgs, org21.Orgid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Orgid, Parentid: org.Parentid, Title: org.Orgname, Checked: "false", Childrens: childorgs, Orgtype: org.Orgtype, Orglevel: org.Orglevel}
			m[org.Orgid] = node
		} else {
			node := Node{Key: org.Orgid, Parentid: org.Parentid, Title: org.Orgname, Checked: "false", Orgtype: org.Orgtype, Orglevel: org.Orglevel}
			m[org.Orgid] = node
		}

	}
	fmt.Println(m)

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateRoleTreeJson() string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllRole()
	orgs2 := orgs
	//fmt.Println(orgs)

	for _, org := range orgs {
		//fmt.Println("parentid:" + org.Parentid)
		//childorgs, _ := Getchildrenorgids(org.Orgid)
		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Roleid {
				childorgs = append(childorgs, org21.Roleid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false", Childrens: childorgs, Rolelevel: org.Rolelevel}
			m[org.Roleid] = node
		} else {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false", Rolelevel: org.Rolelevel}
			m[org.Roleid] = node
		}

	}
	fmt.Println(m)

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateRolePrivilegeTreeJson(roleprivilege []CMN_ROLEPRIVILEGE_TB) string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllModual()
	orgs2 := orgs

	for _, org := range orgs {

		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Modualid {
				childorgs = append(childorgs, org21.Modualid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "false", Childrens: childorgs}
			m[org.Modualid] = node
		} else {
			node := Node{Key: org.Modualid, Parentid: org.Parentid, Title: org.Modualname, Url: org.Url, Checked: "false"}
			m[org.Modualid] = node
		}

	}
	fmt.Println(m)
	for _, grprole := range roleprivilege {
		//map里面的元素不可寻址,不能用m[grprole.Roleid].Checked = "true"这种方法赋值
		//m[grprole.Roleid].Checked = "true"
		p := m[grprole.Modualid]
		p.Checked = "true"
		m[grprole.Modualid] = p
	}

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateOrgRoleTreeJson(orgrole []CMN_ORGROLE_TB) string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllRole()
	orgs2 := orgs

	for _, org := range orgs {

		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Roleid {
				childorgs = append(childorgs, org21.Roleid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false", Childrens: childorgs}
			m[org.Roleid] = node
		} else {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false"}
			m[org.Roleid] = node
		}

	}
	fmt.Println(m)
	for _, grprole := range orgrole {
		//map里面的元素不可寻址,不能用m[grprole.Roleid].Checked = "true"这种方法赋值
		//m[grprole.Roleid].Checked = "true"
		p := m[grprole.Roleid]
		p.Checked = "true"
		m[grprole.Roleid] = p
	}

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateGroupTreeJson() string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllCMN_GROUP_TB()
	orgs2 := orgs
	//fmt.Println(orgs)

	for _, org := range orgs {
		//fmt.Println("parentid:" + org.Parentid)
		//childorgs, _ := Getchildrenorgids(org.Orgid)
		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Groupid {
				childorgs = append(childorgs, org21.Groupid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Groupid, Parentid: org.Parentid, Title: org.Groupname, Checked: "false", Childrens: childorgs, Grouplevel: org.Grouplevel}
			m[org.Groupid] = node
		} else {
			node := Node{Key: org.Groupid, Parentid: org.Parentid, Title: org.Groupname, Checked: "false", Grouplevel: org.Grouplevel}
			m[org.Groupid] = node
		}

	}
	fmt.Println(m)

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateGroupRoleTreeJson(grouprole []CMN_GROUPROLE_TB) string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllRole()
	orgs2 := orgs

	for _, org := range orgs {

		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Roleid {
				childorgs = append(childorgs, org21.Roleid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false", Childrens: childorgs}
			m[org.Roleid] = node
		} else {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false"}
			m[org.Roleid] = node
		}

	}
	fmt.Println(m)
	for _, grprole := range grouprole {
		//map里面的元素不可寻址,不能用m[grprole.Roleid].Checked = "true"这种方法赋值
		//m[grprole.Roleid].Checked = "true"
		p := m[grprole.Roleid]
		p.Checked = "true"
		m[grprole.Roleid] = p
	}

	return "[" + Addnode(&json, m, "root") + "]"

}
func CreateUserRoleTreeJson(userrole []CMN_USERROLE_TB) string {
	json := Json{}
	m := make(map[string]Node)

	orgs, _ := GetAllRole()
	orgs2 := orgs

	for _, org := range orgs {

		childorgs := make([]string, 0)
		for _, org21 := range orgs2 {
			if org21.Parentid == org.Roleid {
				childorgs = append(childorgs, org21.Roleid)
				//处理过的移除，减少下次循环反而更慢，估计和slice收缩有关
				//orgs2 = Remove2(orgs2, idx, idx)
			}

		}
		fmt.Println(childorgs)
		if childorgs != nil && len(childorgs) > 0 {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false", Childrens: childorgs}
			m[org.Roleid] = node
		} else {
			node := Node{Key: org.Roleid, Parentid: org.Parentid, Title: org.Rolename, Checked: "false"}
			m[org.Roleid] = node
		}

	}
	fmt.Println(m)
	for _, grprole := range userrole {
		//map里面的元素不可寻址,不能用m[grprole.Roleid].Checked = "true"这种方法赋值
		//m[grprole.Roleid].Checked = "true"
		p := m[grprole.Roleid]
		p.Checked = "true"
		m[grprole.Roleid] = p
	}

	return "[" + Addnode(&json, m, "root") + "]"

}

func Addnode(json *Json, m2 map[string]Node, id string) string {

	n := m2[id]
	json.Treejson = json.Treejson + "{"
	json.Treejson = json.Treejson + "\"key\":\"" + n.Key + "\","
	json.Treejson = json.Treejson + "\"title\":\"" + n.Title + "\","
	json.Treejson = json.Treejson + "\"isChecked\":" + n.Checked + ","
	json.Treejson = json.Treejson + "\"checked\":" + n.Checked
	if n.Url != "" {
		json.Treejson = json.Treejson + ",\"url\":\"" + n.Url + "\"\r\n"

	}
	if n.Orgtype != "" {
		json.Treejson = json.Treejson + ",\"orgtype\":\"" + n.Orgtype + "\"\r\n"

	}
	if n.Orglevel != "" {
		json.Treejson = json.Treejson + ",\"orglevel\":\"" + n.Orglevel + "\"\r\n"

	}
	if n.Rolelevel != "" {
		json.Treejson = json.Treejson + ",\"rolelevel\":\"" + n.Rolelevel + "\"\r\n"

	}

	if n.Childrens != nil {
		if len(n.Childrens) > 0 {
			json.Treejson = json.Treejson + ",\"children\":[\r\n"

			for i := 0; i < len(n.Childrens)+1; i++ {
				if i == len(n.Childrens) {

					json.Treejson = json.Treejson + "]}"

					break
				} else {
					if i > 0 && i <= len(n.Childrens)-1 {

						json.Treejson = json.Treejson + ","

					}
					Addnode(json, m2, n.Childrens[i])

				}
			}
		}
	} else {
		json.Treejson = json.Treejson + "}\r\n"

	}
	return json.Treejson

}
