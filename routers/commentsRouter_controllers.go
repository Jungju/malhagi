package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:GameController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:PlayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"] = append(beego.GlobalControllerRouter["github.com/jungju/malhagi/controllers:SentenceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
