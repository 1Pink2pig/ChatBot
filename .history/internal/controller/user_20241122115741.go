package controller

import (
	"ChatBot/api"
	"ChatBot/internal/model"
	"ChatBot/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 负责API的路由，定向到指定服务方法

var nowUser *model.User

func GetNowUser() *model.User {
	return nowUser
}

type UserController struct {
	userService service.UserService
}

func (uc *UserController) LoadUser(e *gin.Engine) {
	e.GET("/login.html", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	e.GET("/register.html", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{})
	})
	e.GET("/user_info.html", func(c *gin.Context) {
		c.HTML(200, "user_info.html", gin.H{})
	})

	e.GET("/api/user/:id/info", uc.userInfo)
	e.POST("/api/login", uc.login)
	e.POST("/api/register", uc.register)
	e.PUT("/api/user/:id", uc.modify)
}

func (uc *UserController) userInfo(c *gin.Context) {
	var reqUser model.User
	id := c.Param("id")
	if err := uc.DB.Table("users").Where("id = ?", id).Find(&reqUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid user id " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user": reqUser,
	})
}

func (uc *UserController) modify(c *gin.Context) {
	var reqUser model.User
	c.BindJSON(&reqUser)
	id := c.Param("id")
	reqUser.Id = id
	if err := uc.DB.Table("users").Where("id = ?", id).Update("name", reqUser.Name).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid area name " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "modify successfully"})
}

func (uc *UserController) login(c *gin.Context) {
	uc.DB.Table("users")
	username := c.Query("username")
	passwd := c.Query("passwd")

	fmt.Println("guest ", username, passwd)

	var user model.User
	uc.DB.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		c.JSON(500, gin.H{"error": "username already exist"})
		return
	}

	if flag := strings.Compare(passwd, user.Passwd); flag != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	nowUser = &user

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Register(c *gin.Context) {
	req := new(api.RegisterRequest)
	uc.DB.Table("users")
	var reqUser model.User
	if err := c.BindJSON(&reqUser); err != nil {
		return
	}

	if err := uc.DB.Create(&reqUser).Error; err != nil {
		c.JSON(500, "cannot create user "+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "register successfully",
	})

	var cnt int64
	if err := uc.DB.Model(&model.User{}).Count(&cnt).Error; err != nil {
		panic("failed to count users")
	}
	reqUser.Id = strconv.Itoa(int(cnt))
	uc.DB.Model(&reqUser).Update("id", reqUser.Id)

	nowUser = &reqUser
}