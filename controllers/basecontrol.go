package controllers

import (
	_ "encoding/json"
	"strconv"
	"strings"

	"fmt"
	"github.com/mikezss/skl-go/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/i18n"
)

//（2）建立一个全局session mananger对象
var GlobalSessions *session.Manager

// Operations about Users
type BASEController struct {
	beego.Controller
	i18n.Locale
	User    models.CMN_USER_TB
	IsLogin bool
}

func init() {
	fmt.Println("BASEController.init()==>")
	var err error
	GlobalSessions, err = session.NewManager("memory", &session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600})
	go GlobalSessions.GC()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GlobalSessions:")
	fmt.Println(GlobalSessions)
	fmt.Println("BASEController.init()<==")
}

// Prepare implemented Prepare method for baseRouter.
func (this *BASEController) Prepare() {
	fmt.Println("BASEController.Prepare==>")
	rquri := this.Ctx.Request.RequestURI
	fmt.Println("rquest uri==>" + rquri)
	if this.CruSession == nil {
		this.StartSession()
	} else {
		PROFILE := this.CruSession.Get("PROFILE")
		fmt.Println(PROFILE)
		if PROFILE == nil && !strings.Contains(rquri, "/login/login") {
			status := "expired"
			this.Data["json"] = map[string]string{"status": status}
			this.ServeJSON()
			return

		}
	}

	fmt.Println("BASEController.Prepare<==")
}
func (this *BASEController) getUserid() string {
	var p models.PROFILE
	pr := this.CruSession.Get("PROFILE")
	if pr != nil {
		p = pr.(models.PROFILE)
		return p.Userid
	}

	return ""
}
func stringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}
