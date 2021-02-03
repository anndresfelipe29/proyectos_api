
CREATE SCHEMA IF NOT EXISTS proyectos;

SET search_path TO pg_catalog,public,proyectos;

CREATE TABLE IF NOT EXISTS proyectos.proyecto (
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion varchar(800) NOT NULL,
	fecha_inicio date,
	fecha_finalizacion date,
	id_estado integer,
	CONSTRAINT proyecto_pk PRIMARY KEY (id)

);

CREATE TABLE IF NOT EXISTS proyectos.estado (
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion varchar(800),
	CONSTRAINT estado_pk PRIMARY KEY (id)

);

CREATE TABLE IF NOT EXISTS proyectos.tarea (
	id serial NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion varchar(800),
	fecha_ejecucion smallint,
	id_proyecto integer,
	id_estado integer,
	CONSTRAINT tarea_pk PRIMARY KEY (id)

);

CREATE TABLE IF NOT EXISTS proyectos.usuario (
	id serial NOT NULL,
	usuario varchar(50) NOT NULL,
	contrasena varchar(50) NOT NULL,
	id_rol integer,
	activo bool NOT NULL,
	CONSTRAINT usuario_pk PRIMARY KEY (id),
	CONSTRAINT usuario_uk UNIQUE (usuario)

);

CREATE TABLE IF NOT EXISTS proyectos.rol (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(200),
	CONSTRAINT rol_pk PRIMARY KEY (id)

);

ALTER TABLE proyectos.tarea ADD CONSTRAINT proyecto_fk FOREIGN KEY (id_proyecto)
REFERENCES proyectos.proyecto (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE proyectos.proyecto ADD CONSTRAINT estado_fk FOREIGN KEY (id_estado)
REFERENCES proyectos.estado (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE proyectos.tarea ADD CONSTRAINT estado_fk FOREIGN KEY (id_estado)
REFERENCES proyectos.estado (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE proyectos.usuario ADD CONSTRAINT rol_fk FOREIGN KEY (id_rol)
REFERENCES proyectos.rol (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;


