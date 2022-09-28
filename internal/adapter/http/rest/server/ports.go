package server

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	RolesHandler(c *gin.Context)
	DeleteRoleHandler(c *gin.Context)
	StoreRoleHandler(c *gin.Context)
}

type UserHandler interface {
	GetUserByIDHandler(c *gin.Context)
	UsersHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	StoreUserHandler(c *gin.Context)
	AssignRoleToUser(c *gin.Context)
}
type AuthHandler interface {
	UserLogin(c *gin.Context)
	Authorizer(e *casbin.Enforcer) gin.HandlerFunc
}
