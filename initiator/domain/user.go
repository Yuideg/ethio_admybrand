package domain

import (
	user_handler "github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server/user"
	user_persist "github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence/user"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	glue "github.com/Yideg/admybrand_challenge/internal/glue/routing"
	user_module "github.com/Yideg/admybrand_challenge/internal/module/user"
	"github.com/gin-gonic/gin"
)

func UserInit(utils model.Utils, router *gin.RouterGroup) {
	userPersistence := user_persist.UserInit(utils.Conn)
	userUsecase := user_module.Initialize(userPersistence, utils)
	userHandler := user_handler.NewUserHandler(userUsecase, utils)
	glue.UserRoutes(router, userHandler)
}
