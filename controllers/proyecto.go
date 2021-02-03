package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/anndresfelipe29/proyectos_api/models"
	"github.com/anndresfelipe29/proyectos_api/token"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ProyectoController operations for Proyecto
type ProyectoController struct {
	beego.Controller
}

var et = token.EasyToken{}

// URLMapping ...
func (c *ProyectoController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("CompletarProyecto", c.CompletarProyecto)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Proyecto
// @Param   Authorization   header  string  false "Token"
// @Param	body		body 	models.Proyecto	true		"body for Proyecto content"
// @Success 201 {int} models.Proyecto
// @Failure 403 body is empty
// @router / [post]
func (c *ProyectoController) Post() {
	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 2)
	if err != nil {
		logs.Error(err)
		c.Abort("401")
	}
	if role {

		var v models.Proyecto
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if _, err := models.AddProyecto(&v); err == nil {
				fmt.Println("entra aca")
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else {
				logs.Error(err)
				c.Data["json"] = err.Error()
			}
		} else {
			logs.Error(err)
			c.Data["json"] = err.Error()
		}
	}

	//fmt.Println(et.ValidateRoleToken(tokenString, 1))

	c.ServeJSON()
}

/*// GetOne ...
// @Title Get One
// @Description get Proyecto by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Proyecto
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProyectoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProyectoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}*/

// GetAll ...
// @Title Get All
// @Description get Proyecto
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Proyecto
// @Failure 403
// @router / [get]
func (c *ProyectoController) GetAll() {
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

	l, err := models.GetAllProyecto(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Proyecto
// @Param   Authorization   header  string  false "Token"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Proyecto	true		"body for Proyecto content"
// @Success 200 {object} models.Proyecto
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProyectoController) Put() {

	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 2)
	if err != nil {
		logs.Error(err)
		c.Abort("401")
	}
	if role {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		v := models.Proyecto{Id: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateProyectoById(&v); err == nil {
				c.Data["json"] = "OK"
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}

	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Proyecto
// @Param   Authorization   header  string  false "Token"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProyectoController) Delete() {

	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 2)
	if err != nil {
		logs.Error(err)
		c.Abort("401")
	}
	if role {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		if err := models.DeleteProyecto(id); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Proyecto
// @Param   Authorization   header  string  false "Token"
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} models.Proyecto
// @Failure 403 :id is not int
// @router /completar/:id [put]
func (c *ProyectoController) CompletarProyecto() {

	role, _, err := et.ValidateRoleToken(c.Ctx.Input.Header("Authorization"), 2)
	if err != nil {
		logs.Error(err)
		c.Abort("401")
	}
	if role {
		idStr := c.Ctx.Input.Param(":id")
		//id, _ := strconv.Atoi(idStr)
		//v := models.Proyecto{Id: id}
		//fmt.Println(models.GetTareaByProyecto(id))
		v := make(map[string]string)
		x := []string{}
		v["id_proyecto"] = idStr
		tareas1, err := models.GetAllTarea(v, x, x, x, -1, -1)
		if err != nil {
			logs.Error(err)
			c.Abort("401")
		}
		v["id_estado"] = "2"
		tareas2, err := models.GetAllTarea(v, x, x, x, -1, -1)
		if err != nil {
			logs.Error(err)
			c.Abort("401")
		}
		fmt.Println(len(tareas1), len(tareas2))
		if len(tareas1) == len(tareas2) {
			id, _ := strconv.Atoi(idStr)
			/*v := models.Proyecto{
				Id:       id,
				IdEstado: &models.Estado{Id: 4},
			}*/
			if v, err := models.GetProyectoById(id); err == nil {
				v.IdEstado = &models.Estado{Id: 4}
				fmt.Println("entro")
				fmt.Println(v)
				if err := models.UpdateProyectoById(v); err == nil {
					c.Data["json"] = "OK"
				} else {
					c.Data["json"] = err.Error()
					c.Abort("401")
				}
			} else {
				logs.Error(v, err)
				c.Abort("401")
			}

		}

	}

	c.ServeJSON()
}

/*
{

	"IdEstado": {

	  "Id": 2,

	},


  }
  eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTIzMjg5MTksImlzcyI6ImZlbGlwZSIsInN1YiI6IjEifQ.jEyJ136J-UtEP2PwmCDYSn60MAxmo9cFmPbjcACNjKzVWUUoQ6uQaNHrU9FUp7h_SI8lVHg0xADa55FgjuUPh9jSpJkIRQTRkTnr_w5ppxwOUsu8cc8ZG4mUCtfw4JDhYk4weoyRwp7XjgJLZO9GyLJbAkYaRHhKtBnY_nyQnjH40yNfV9GqXQjI3IDCvbaL6CMWD5xxbBGpMVkK29q3nMV196KwjDJTgnsx_mWv-kWoEQixz28kWeYx9PRUM_-XE_PV690XxScldyAbdPILAty_Ab8GTp22He35_2fcQT5geqbazDLstJZ8FZzc8dG-YKijs2gXMF16kgC8ONL3-w
*/
