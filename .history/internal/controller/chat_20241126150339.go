package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func (cc *ChatController) LoadChat(e *gin.Engine) {
	e.POST("/:id/newchat", cc.NewChat)
}

func (cc *ChatController) NewChat(ctx *gin.Context) {

}

func (cc *ChatController) EndChat(e *gin.Engine) {
	
}

func (cc *ChatController) Chat(e *gin.Engine) {
	
}