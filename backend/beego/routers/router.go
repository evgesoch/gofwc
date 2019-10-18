// @APIVersion 1.0.0
// @Title Beego API project
// @Description Beego project
// @Contact eugenesocho@hotmail.com
// @License MIT
// @LicenseUrl https://github.com/evgesoch/gofwc/blob/master/LICENSE
package routers

import (
	"github.com/astaxie/beego"
	"github.com/evgesoch/gofwc/backend/beego/controllers"
)

func init() {
	ns := beego.NewNamespace("/posts", beego.NSInclude(
		&controllers.PostController{}),
	)
	beego.AddNamespace(ns)
}
