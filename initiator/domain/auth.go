package domain

import (
	auth_handler "github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server/auth"
	auth_persist "github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence/user"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	glue "github.com/Yideg/admybrand_challenge/internal/glue/routing"
	auth_module "github.com/Yideg/admybrand_challenge/internal/module/auth"
	"github.com/gin-gonic/gin"
)

func AuthInit(utils model.Utils, router *gin.RouterGroup) {
	authPersistence := auth_persist.UserInit(utils.Conn)
	authUsecase := auth_module.Initialize(authPersistence, utils)
	authHandler := auth_handler.NewAuthHandler(authUsecase, utils)
	glue.AuthRoutes(router, authHandler)
}
