// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/proyectos_api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/proyecto",
			beego.NSInclude(
				&controllers.ProyectoController{},
			),
		),

		beego.NSNamespace("/estado",
			beego.NSInclude(
				&controllers.EstadoController{},
			),
		),

		beego.NSNamespace("/tarea",
			beego.NSInclude(
				&controllers.TareaController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
