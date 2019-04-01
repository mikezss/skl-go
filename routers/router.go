// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/mikezss/skl-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.LOGINController{})
	beego.AutoRouter(&controllers.MASTERController{})
	beego.AutoRouter(&controllers.FLOWController{})
	beego.AutoRouter(&controllers.COMMONController{})
	beego.AutoRouter(&controllers.ENUMController{})
	beego.AutoRouter(&controllers.COMPANYController{})
    beego.AutoRouter(&controllers.ENUMSEARCHController{})
}
