package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:postID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:postID`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/evgesoch/gofwc/backend/beego/controllers:PostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:postID`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
