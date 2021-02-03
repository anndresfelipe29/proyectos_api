package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CargarTablasParametricas_20210202_204936 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CargarTablasParametricas_20210202_204936{}
	m.Created = "20210202_204936"

	migration.Register("CargarTablasParametricas_20210202_204936", m)
}

// Run the migrations
func (m *CargarTablasParametricas_20210202_204936) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO proyectos.estado(nombre, descripcion) VALUES ('Pendiente', 'Estado de una tarea sin culminar')")
	m.SQL("INSERT INTO proyectos.estado(nombre, descripcion) VALUES ('Realizada', 'Estado de una tarea culminada')")
	m.SQL("INSERT INTO proyectos.estado(nombre, descripcion) VALUES ('En proceso', 'Estado de un proceso sin culminar')")
	m.SQL("INSERT INTO proyectos.estado(nombre, descripcion) VALUES ('Finalizada', 'Estado de un proceso culminado')")

	m.SQL("INSERT INTO proyectos.rol(nombre, descripcion) VALUES ('Administrador', 'Rol de usuario administrativo')")
	m.SQL("INSERT INTO proyectos.rol(nombre, descripcion) VALUES ('Operador', 'Rol de usuario operativo')")
}

// Reverse the migrations
func (m *CargarTablasParametricas_20210202_204936) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
