package controllers

import (
	"strconv"
	"time"

	"github.com/anndresfelipe29/proyectos_api/models"
	"github.com/anndresfelipe29/proyectos_api/token"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// EstadoController operations for Estado
type TokenController struct {
	beego.Controller
}

// URLMapping ...
func (c *TokenController) URLMapping() {
	c.Mapping("GetToken", c.GetToken)
}

// @Title getToken
// @Description Get token from user session
// @Param	username		query 	string	true		"The username for get token"
// @Param	password		query 	string	true		"The password for get token"
// @Success 200 {string} Obtain Token
// @router / [post]
func (c *TokenController) GetToken() {
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")

	tokenString := ""

	//fmt.Println(models.GetUsuarioByUsuario(username))
	if usuario, err := models.GetUsuarioByUsuario(username); err != nil {
		logs.Error(err)
		c.Abort("401")
	} else {
		if username == usuario.Usuario && password == usuario.Contrasena && usuario.Activo {
			et := token.EasyToken{
				Username: username,
				Expires:  time.Now().Unix() + 3600, //Segundos
				Role:     strconv.Itoa(usuario.IdRol.Id),
			}
			tokenString, _ = et.GetToken()

		}

		c.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	}

	c.ServeJSON()
	return
}

// generar llave rsa
// http://security.stackexchange.com/questions/40566/create-rsa-public-private-key-pair-having-numbers-generated-already
// openssl genrsa -out rsakey.pem 2048
// openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub
