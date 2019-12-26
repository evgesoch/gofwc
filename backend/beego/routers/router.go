// @APIVersion 1.0.0
// @Title Beego API project
// @Description Beego project
// @Contact eugenesocho@hotmail.com
// @License MIT
// @LicenseUrl https://github.com/evgesoch/gofwc/blob/master/LICENSE
package routers

import (
	"github.com/astaxie/beego"
	beegoController "github.com/evgesoch/gofwc/backend/beego/controllers"
)

func init() {
	// Api
	ns := beego.NewNamespace("/posts", beego.NSInclude(
		&beegoController.PostController{}),
	)
	beego.AddNamespace(ns)

	// Frontend
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/frontend", "../../frontend")
	beego.SetStaticPath("/speak4env", "../../frontend/index.html")
}
