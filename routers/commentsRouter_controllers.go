package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:ProyectoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TareaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	/*beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
	beego.ControllerComments{
		Method:           "GetOne",
		Router:           "/:id",
		AllowHTTPMethods: []string{"get"},
		MethodParams:     param.Make(),
		Filters:          nil,
		Params:           nil})*/
	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "CambiarContrasena",
			Router:           "/:password/:passwordNew ",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "DeshabilitarUsuario",
			Router:           "/:username/:activo",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TokenController"] = append(beego.GlobalControllerRouter["github.com/anndresfelipe29/proyectos_api/controllers:TokenController"],
		beego.ControllerComments{
			Method:           "GetToken",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
