package domain

import (
	role_handler "github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server/role"
	role_persist "github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence/role"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	glue "github.com/Yideg/admybrand_challenge/internal/glue/routing"
	role_module "github.com/Yideg/admybrand_challenge/internal/module/role"
	"github.com/gin-gonic/gin"
)

func RoleInit(utils model.Utils, router *gin.RouterGroup) {
	rolePersistence := role_persist.RoleInit(utils.Conn)
	roleUsecase := role_module.Initialize(rolePersistence, utils)
	roleHandler := role_handler.NewRoleHandler(roleUsecase, utils)
	glue.RoleRoutes(router, roleHandler)
}
