package users

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Secret to sign JSON Web Tokens 
var jwt_secret = []byte("OH MY GAH")

var once sync.Once
var users map[string]string

func Users() map[string]string {
	once.Do(func() {
		users = make(map[string]string)
	})

	return users
}

func LoginFunc(ctx *gin.Context) {
	user := ctx.PostForm("username")
	fmt.Println("username=", user)
	pass := ctx.PostForm("password")
	fmt.Println("password=", pass)

	if passV, ok := users[user]; ok && pass == passV {
		ctx.Header("HX-Redirect", "/users")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": user,
			"exp": time.Now().Unix() + 3600,
		})

		token_str, err := token.SignedString(jwt_secret)
		if err != nil { 
			panic("Could not sign token for user: " + err.Error())
		}

		ctx.SetCookie("JWT", token_str, 3600, "/", ctx.Request.Host, false, true)

		ctx.HTML(http.StatusOK, "greatSuccess", gin.H{
			"db":      users,
			"thisUsr": user,
		})
		
	} else {
		ctx.HTML(http.StatusOK, "he cannot afford", gin.H{})
	}
}


func AddUserFunc(ctx *gin.Context) {
	fmt.Println("Adding new user!")
	user := ctx.PostForm("username")
	pass := ctx.PostForm("password")
	fmt.Printf("\tUser: '%s' \n\tPass: '%s'\n", user, pass)

	_, ok := users[user]
	if ok {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	
	users[user] = pass

	ctx.HTML(http.StatusOK, "userListElement", user)
}

func ListUsersFunc(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user-list", users)
}

func DeleteUserFunc(ctx *gin.Context) {
	user := ctx.Param("user")
	if _, ok := users[user]; !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	
	delete(users, user)
	//time.Sleep(1 * time.Second)
	ctx.AbortWithStatus(http.StatusOK)
}


func LoadUsersHandlers(r *gin.Engine) {
	r.POST("/login", LoginFunc)

	// TODO: Auth middleware.
	// Aqui agregariamos un middleware de lo del JWT si lo tuvieramos
	authed := r.Group("/", AuthMiddleware())
	
	authed.GET("/users", func(ctx *gin.Context) {
		user := ctx.MustGet("AuthUser")
		if user == "admin" {
			ctx.HTML(http.StatusOK, "users.html", Users())
		} else {
			ctx.Redirect(http.StatusPermanentRedirect, "/my_properties")
		}
	})

	authed.POST("/users/new", AddUserFunc)

	authed.DELETE("/users/:user", DeleteUserFunc)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tkn, err := ctx.Cookie("JWT")
		if err != nil {
			fmt.Println("Could not get cookie: " + err.Error())
			//ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}

		token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return jwt_secret, nil
		})

		if err != nil {
			fmt.Println("Could not parse JWT: " + err.Error())
			//ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			usr := claims["user"].(string)
			fmt.Println("Setting (", usr, ")", claims["exp"])
			ctx.Set("AuthUser", usr)
		} else {
			//ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}

		ctx.Next()
	}
}

