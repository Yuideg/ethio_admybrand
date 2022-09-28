package auth

import (
	"bytes"
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/Yideg/admybrand_challenge/internal/constant"
	CustomError "github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/Yideg/admybrand_challenge/internal/module"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

type authHandler struct {
	authUseCase module.AuthUseCase
	ut          model.Utils
}

var secret = os.Getenv("JWT_SECRET")

func NewAuthHandler(authUseCase module.AuthUseCase, ut model.Utils) server.AuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
		ut:          ut,
	}
}

func (a authHandler) UserLogin(c *gin.Context) {
	ctx := c.Request.Context()
	user_credential := model.LoginCredential{}
	err := c.Bind(&user_credential)
	if err != nil {
		errData := CustomError.ErrorService(CustomError.ErrorUnableToBindJsonToStruct)
		fmt.Println("errordata ", errData)
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		return
	}
	errData := constant.VerifyInput(user_credential, a.ut.GoValidator, a.ut.Translator)
	fmt.Println("verification =>", errData)
	if errData != nil {
		constant.ResponseJson(c, errData, http.StatusBadRequest)
		return
	}
	user := model.User{
		UserName: user_credential.UserName,
		Password: user_credential.Password,
	}
	access_token, err := a.authUseCase.UserLogin(ctx, user)
	fmt.Println("login error ", err)
	if err != nil {
		errdata := CustomError.ErrorService(err)
		fmt.Println("err data ", errdata)

		constant.ResponseJson(c, errdata, http.StatusBadRequest)
		c.Abort()
		return
	}
	token := access_token.AccessToken
	var bearer string = "Bearer " + token
	c.Header("Content-Type", "application/json")
	c.Header("token", token)
	c.Request.Header.Set("Authorization", bearer)
	constant.ResponseJson(c, access_token, http.StatusOK)
	return
}

// Authorizer is a middleware for authorization
func (n *authHandler) Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	log.Println("<<< authorizer >>>")
	return func(c *gin.Context) {
		role := []string{"anonymous"}
		token := c.Request.Header.Get("Authorization")
		ctx := c.Request.Context()
		bear_token_list := strings.Split(token, " ")
		if len(bear_token_list) == 2 && bear_token_list[0] == "Bearer" && len(strings.Split(bear_token_list[1], ".")) == 3 {
			token = strings.TrimPrefix(token, "Bearer ")
		}
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		verified := verify(token)
		log.Println("Verified", verified)
		if verified {
			tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			fmt.Println("error parse", err)
			if tkn != nil {
				if claim, ok := tkn.Claims.(*model.Claim); ok {
					id := claim.UserID
					user, err := n.authUseCase.GetUserByID(ctx, model.User{ID: id})
					if err != nil {
						errData := CustomError.ErrorService(CustomError.ErrorAccessDenied)
						constant.ResponseJson(c, errData, http.StatusBadRequest)
						c.Abort()
						return
					}

					c.Set("X-USER-ID", user.ID)
					c.Set("X-USER-USERNAME", user.ID)

					c.Request.Header.Set("X-USER-TIME-STAMP", fmt.Sprintf("%v", user.CreatedAt))
					fmt.Println("user id ", user.ID)
					fmt.Println("user username", user.UserName)
					fmt.Println("user role name", user.RoleName)
					verify_role := constant.ValidRole(user.RoleName)

					if verify_role {
						c.Set("X-USER-ROLE", user.RoleName)
						c.Request.Header.Set("X-USER-ROLE", user.RoleName)

						role = append(role, user.RoleName)
					}
					c.Set("X-USER-ROLES", role)
				} else {
					e := CustomError.ErrorService(CustomError.ErrUnableToParse)
					constant.ResponseJson(c, e, http.StatusBadRequest)
					c.Abort()
					return
				}
			}
		}

		log.Printf(" logging line --125 auth handler %v %v %v", role, c.Request.URL.Path, c.Request.Method)
		var res = false
		for _, v := range role {
			res, _ = e.EnforceSafe(v, c.Request.URL.Path, c.Request.Method)
			if res {
				break
			}
		}
		log.Println("res", res)
		if res {
			log.Println(" line 135 in res, before handler", res)
			//TODO  Save activity log
			//Middleware filtered request can get resources
			c.Next()
			log.Println("in res, after handler", res)
		} else {
			log.Println("no res", res)
			log.Println("No permission")
			e := CustomError.ErrorService(CustomError.ErrorAccessDenied)
			e.ErrorDescription = "Sorry,You don't have enough permission for this action."
			constant.ResponseJson(c, e, http.StatusForbidden)
			c.Abort()
			fmt.Println("ended middleware", c.IsAborted())
			return
		}
	}
}

var claims = &model.Claim{}

func verify(token string) bool {
	segments := strings.Split(token, ".")
	if len(segments) < 3 {
		return false
	}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	if !tkn.Valid {
		return false
	}
	fmt.Println("e ", err)
	if _, ok := tkn.Claims.(*model.Claim); !ok {
		fmt.Println("line 182")

		return false
	}
	return true
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
