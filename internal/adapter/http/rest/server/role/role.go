package user

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/Yideg/admybrand_challenge/internal/constant"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	roleUseCase module.RoleUseCase
	ut          model.Utils
}

//NewRoleHandler Creates new object of RoleHandler
func NewRoleHandler(roleUseCase module.RoleUseCase, ut model.Utils) server.RoleHandler {
	return &roleHandler{
		roleUseCase: roleUseCase,
		ut:          ut,
	}
}

func (r roleHandler) RolesHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	method := c.Request.Method
	name := strings.TrimSpace(c.Request.URL.Query().Get("name"))
	if method == http.MethodGet && name != "" {
		successData, err := r.roleUseCase.GetRoleByName(ctx, model.Role{Name: name})
		if err != nil {
			errData := CustomError.ErrorService(err)
			constant.ResponseJson(c, errData, http.StatusBadRequest)
			c.Abort()
			return
		}
		constant.ResponseJson(c, successData, http.StatusOK)
		return
	} else {

		successData, err := r.roleUseCase.Roles(ctx)
		if err != nil {
			errData := CustomError.ErrorService(err)
			constant.ResponseJson(c, errData, http.StatusNotFound)
			c.Abort()
			return
		}
		constant.ResponseJson(c, successData, http.StatusOK)
		return
	}
}

func (r roleHandler) DeleteRoleHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	name := strings.TrimSpace(c.Request.URL.Query().Get("name"))
	errData := r.roleUseCase.DeleteRole(ctx, model.Role{Name: name})
	fmt.Println("error ", errData)
	if errData != nil {
		e := CustomError.ErrorService(errData)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		c.Abort()
		return
	}
	dd := &model.DeleteData{
		Result: "Role with id " + name + " deleted successfully!",
	}
	constant.ResponseJson(c, dd, http.StatusOK)
	return
}

func (r roleHandler) StoreRoleHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	role := model.Role{}
	err := c.BindJSON(&role)
	if err != nil {
		errData := CustomError.ErrorService(CustomError.ErrorUnableToBindJsonToStruct)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}
	errData := constant.VerifyInput(role, r.ut.GoValidator, r.ut.Translator)
	fmt.Println("verification =>", errData)
	if errData != nil {
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}

	successData, err := r.roleUseCase.StoreRole(ctx, role)
	if err != nil {
		errdata := CustomError.ErrorService(err)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
	return
}
