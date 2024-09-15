package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"evaluator/users"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.Static("/assets", "./assets")

	users.Users()["admin"] = "secret"
	users.Users()["user1"] = "omg"
	users.Users()["user2"] = "wow"
	users.Users()["user3"] = "lol"


	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/login")
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	
	// Se puede crear un grupo que no extiende el relative path,
	// significa que agregar middlewares es trivial dentro de una
	// ruta. (Lo digo para que lo de autenticacion este en su propio modulo)
	authed := r.Group("/")
	authed.Use(func(ctx *gin.Context) {fmt.Println("Heloo")})
	authed.GET("/lol")

	users.LoadUsersHandlers(r)
	

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
