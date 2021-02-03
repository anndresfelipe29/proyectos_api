# Proyectos API
Proyecto backend realizado en Golang con el framework Beego que usa Postgresql como persistencia.
El proyecto usa tokens JWT para hacer seguro el acceso a cada endpoint que lo requiere, los metodos get no requieren token.
La idea de la api es poder llevar el registro de proyectos y tareas por medio de dos roles, administrador y operario, donde operario tiene algunas limitaciones de acceso.
## Modelo de datos
![Esquema de datos](https://i.imgur.com/bUKKF0H.png)
## Ejecución en repositorio local

Requerimientos:
- Docker

Instalación:
```shell
#1. Obtener el repositorio con git clone
git clone https://github.com/anndresfelipe29/proyectos_api.git

#2. Moverse a la carpeta del repositorio
cd proyectos_api

# 3. Ejecutar docker-compose
docker-compose up
```
## Uso
Puede usar este aplicativo accediento a:
- Endpoint base http://localhost:8092  y agregar las rutas de acceso a cada controlador.
- Url de la aplicación swagger que facilita el uso del microservicio http://localhost:8092/swagger-ui.html#

## Funcionalidades

- /token [Post] Genera un token, requiere un usuario y contraseña validos.
- /usuario/ [Post] Crea un nuevo usuario, solo lo puede hacer el administrador que tenga un token valido.
- /usuario/{username}/{activo} Habilita o deshabilita un usuario, solo lo puede hacer un administrador.
- /usuario/{password}/{passwordNew} Permite cambiar una contraseña de usuario, lo puede hacer cualquier usuario, requiere token y contraseña actual.
- /proyecto/ [Post] Crea un nuevo proyecto, lo puede hacer un rol operio o superior, requiere un token valido.
- /proyecto/{id} [Put] Edita un proyecto existente,lo puede hacer un rol operio o superior, requiere un token valido.
- /proyecto/{id} [Delete] Elimina un proyecto y sus tareas asociadas, lo puede hacer un rol operio o superior, requiere un token valido.(Falta ajustar la consulta sql)
- /proyecto/completar/{id} [Put] Cambia el estado del proyecto a finalizado si tiene todas las tareas en estado realizado.
- /tarea/ [Post] Crea una nueva tarea y la asocia a un proyecto, lo puede hacer un rol operio o superior, requiere un token valido.
- /tarea/{id} [Put] Edita una tarea existente,lo puede hacer un rol operio o superior, requiere un token valido.

