package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/anndresfelipe29/proyectos_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// UsuarioController operations for Usuario
type UsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("put", c.CambiarContrasena)
	c.Mapping("put", c.DeshabilitarUsuario)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Usuario
// @Param   Authorization   header  string  false "Token"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {int} models.Usuario
// @Failure 403 body is empty
// @router / [post]
func (c *UsuarioController) Post() {

	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 1)
	if err != nil {
		logs.Error(err)
		c.Abort("401")
	}
	if role {
		var v models.Usuario
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if _, err := models.AddUsuario(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}

	}

	c.ServeJSON()
}

/*
// GetOne ...
// @Title Get One
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsuarioById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
*/
// GetAll ...
// @Title Get All
// @Description Puede ver a los usuarios y su información, no requiere token, es para facilitar el uso de los usuarios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuario
// @Failure 403
// @router / [get]
func (c *UsuarioController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsuario(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Usuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is not int
// @router /:id [put]
/*func (c *UsuarioController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Usuario{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUsuarioById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}*/

// DeshabilitarUsuario ...
// @Title DeshabilitarUsuario
// @Description Deshabilitar Usuario permite habilitar o inhabilitar un usuario
// @Param   Authorization   header  string  false "Token"
// @Param	username		path 	string	true		"El nombre de usuario a cambiar"
// @Param	activo		path 	boolean	true		"true para habiliar y false para deshabilitar"
// @Success 200 Ok
// @Failure 403 :id is not int
// @router /:username/:activo [put]
func (c *UsuarioController) DeshabilitarUsuario() {
	fmt.Println("Entro")
	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 1)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.Abort("401")
	}
	if role {

		username := c.Ctx.Input.Param(":username")
		//id, err1 := strconv.Atoi(idStr)
		activo, err2 := strconv.ParseBool(c.Ctx.Input.Param(":activo"))
		if err2 != nil {
			logs.Error(err)
			c.Data["json"] = err.Error()
			c.Abort("400")
		}
		/*v := models.Usuario{
			Id:     id,
			Activo: activo,
		}*/
		if v, err := models.GetUsuarioByUsuario(username); err != nil {
			c.Data["json"] = err.Error()
			c.Abort("404")
		} else {
			v.Activo = activo
			if err := models.UpdateUsuarioById(v); err == nil {
				c.Data["json"] = "OK"
			} else {
				c.Data["json"] = err.Error()
			}
		}

	}

	c.ServeJSON()
}

// CambiarContrasena ...
// @Title CambiarContrasena
// @Description Permite cambiar la contraseña de un usuario un usuario
// @Param   Authorization   header  string  false "Token"
// @Param	password	path 	string	true		"El nombre de usuario a cambiar"
// @Param	passwordNew	path 	string	true		"El nombre de usuario a cambiar"
// @Success 200 Ok
// @Failure 403 :id is not int
// @router /:password/:passwordNew [put]
func (c *UsuarioController) CambiarContrasena() {
	fmt.Println("Entro")
	role, username, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 2)
	if err != nil {
		logs.Error("Token: ", err)
		c.Data["json"] = err.Error()
		c.Abort("401")
	}
	if role {

		contrasena := c.Ctx.Input.Param(":password")
		contrasenaNew := c.Ctx.Input.Param(":passwordNew")
		if v, err := models.GetUsuarioByUsuario(username); err != nil {
			c.Data["json"] = err.Error()
			c.Abort("404")
		} else {
			fmt.Println(v.Contrasena, contrasena)
			if v.Contrasena == contrasena {
				v.Contrasena = contrasenaNew
				if err := models.UpdateUsuarioById(v); err == nil {
					c.Data["json"] = "OK"
				} else {
					logs.Error("Update: ", err)
					c.Data["json"] = err.Error()
				}
			} else {
				logs.Error("Password: fail ")
				c.Abort("401")
			}

		}

	}

	c.ServeJSON()
}

/*
// Delete ...
// @Title Delete
// @Description delete the Usuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuarioController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsuario(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}*/
