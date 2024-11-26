package main

import (
	"ChatBot/internal/controller"
	"ChatBot/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var db db.MyDB
	var uc controller.UserController
	db.InitDB()

	r := gin.Default()
	r.Use(cors.Default())

	// r.Use()
	r.LoadHTMLGlob("data/src/*.html")
	r.Static("assets/", "data/src/assets/")
	// r.Static("user/assets/", "data/src/assets/")
	controllers.LoadIndex(r)
	ctrl.Uc.LoadUser(r)
	ctrl.Ac.LoadArticle(r)
	ctrl.Mc.LoadMap(r)

	r.Run(":9998")
}
