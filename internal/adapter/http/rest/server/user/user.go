package user

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/Yideg/admybrand_challenge/internal/constant"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

//userHandler pors receiver
type userHandler struct {
	userUseCase module.UserUseCase
	ut          model.Utils
}

//NewUserHandler Creates new object of UserHandler
func NewUserHandler(userUseCase module.UserUseCase, ut model.Utils) server.UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
		ut:          ut,
	}
}

//GetUserByIDHandler fetches user who matches the given username
func (t userHandler) GetUserByIDHandler(c *gin.Context) {
	//TODO implement me
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		//conversion error
		e := CustomError.ErrorService(CustomError.ErrorUnableToConvert)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		c.Abort()
		return
	}
	ctx := c.Request.Context()
	successData, err := t.userUseCase.GetUserByID(ctx, model.User{ID: id})
	if err != nil {
		errData := CustomError.ErrorService(err)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
	return
}

//UsersHandler returns all registered account
func (t userHandler) UsersHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	successData, err := t.userUseCase.Users(ctx)
	if err != nil {
		errData := CustomError.ErrorService(err)
		constant.ResponseJson(c, errData, http.StatusNotFound)
		c.Abort()
		return
	}

	constant.ResponseJson(c, successData, http.StatusOK)
	return

}

//UpdateUserHandler edits user account
func (t userHandler) UpdateUserHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		//conversion error
		e := CustomError.ErrorService(CustomError.ErrorUnableToConvert)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		c.Abort()
		return
	}
	user := model.User{}
	err = c.Bind(&user)
	if err != nil {
		errData := CustomError.ErrorService(CustomError.ErrorUnableToBindJsonToStruct)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return

	}
	user.ID = id
	successData, err := t.userUseCase.UpdateUser(ctx, user)
	if err != nil {
		errdata := CustomError.ErrorService(err)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}

	constant.ResponseJson(c, successData, http.StatusOK)
	return

}

//DeleteUserHandler delete or remove an account
func (t userHandler) DeleteUserHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		//conversion error
		e := CustomError.ErrorService(CustomError.ErrorUnableToConvert)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		c.Abort()
		return
	}
	errData := t.userUseCase.DeleteUser(ctx, model.User{ID: id})
	fmt.Println("error ", errData)
	if errData != nil {
		e := CustomError.ErrorService(errData)
		constant.ResponseJson(c, e, http.StatusBadRequest)
		c.Abort()
		return
	}

	dd := &model.DeleteData{
		Result: "User with id " + param + " deleted successfuly!",
	}

	constant.ResponseJson(c, dd, http.StatusOK)
	return

}

//StoreUserHandler handles user request to create new account
func (t userHandler) StoreUserHandler(c *gin.Context) {
	//TODO implement me
	ctx := c.Request.Context()
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		errData := CustomError.ErrorService(CustomError.ErrorUnableToBindJsonToStruct)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}
	errData := constant.VerifyInput(user, t.ut.GoValidator, t.ut.Translator)
	fmt.Println("verification =>", errData)
	if errData != nil {
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}
	hashedPassword, err := constant.HashPassword(c, user.Password)
	if err != nil {
		errData := CustomError.ErrorService(CustomError.ErrInvalidField)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		c.Abort()
		return
	}
	user.Password = hashedPassword

	successData, err := t.userUseCase.StoreUser(ctx, user)
	if err != nil {
		errdata := CustomError.ErrorService(err)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}

	constant.ResponseJson(c, successData, http.StatusOK)
	return
}

//AssignRoleToUser  handles  an incoming admin user request to assign role to user
func (t userHandler) AssignRoleToUser(c *gin.Context) {
	//TODO implement me
	//gets quesry param from request url
	user_id := strings.TrimSpace(c.Request.URL.Query().Get("id"))
	id, err := uuid.FromString(user_id)
	if err != nil {
		errdata := CustomError.ErrorService(CustomError.ErrInvalidRequest)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}
	name := strings.TrimSpace(c.Request.URL.Query().Get("role"))
	if name == "" {
		errdata := CustomError.ErrorService(CustomError.ErrInvalidRequest)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}
	successData, err := t.userUseCase.AssignRoleToUser(c.Request.Context(), model.User{ID: id, RoleName: name})
	if err != nil {
		errdata := CustomError.ErrorService(err)
		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}
	constant.ResponseJson(c, successData, http.StatusOK)
	return
}
