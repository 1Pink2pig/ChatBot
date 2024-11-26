package server

import (
	"ChatBot/api"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx gin.Context, req *api.RegisterRequest) error
	Login(ctx gin.Context, req *api.LoginRequest) (string, error)
}
