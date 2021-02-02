package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquema_20210202_154037 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquema_20210202_154037{}
	m.Created = "20210202_154037"

	migration.Register("CrearEsquema_20210202_154037", m)
}

// Run the migrations
func (m *CrearEsquema_20210202_154037) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210202_154037_crear_esquema_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearEsquema_20210202_154037) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
