package main

import (
	//"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"evaluator/properties"
	"evaluator/users"

	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/Geocoding.js", "./assets/Geocoding.js")
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

	err := godotenv.Load()
	if err != nil {
		panic("Could not load .env!: " + err.Error())
	}

	api_keys := gin.H{ "googleapikey" :os.Getenv("google_maps")}
	
	// Se puede crear un grupo que no extiende el relative path,
	// significa que agregar middlewares es trivial dentro de una
	// ruta. (Lo digo para que lo de autenticacion este en su propio modulo)
	authed := r.Group("/")
	//authed.Use(func(ctx *gin.Context) {fmt.Println("Heloo")})
	authed.GET("/new_home", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "new_home.html", api_keys)
	})

	users.LoadUsersHandlers(r)
	properties.LoadPropertiesHandlers(r)
	

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
