package initiator

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/initiator/domain"
	auth_handler "github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server/auth"
	auth_persist "github.com/Yideg/admybrand_challenge/internal/adapter/storage/persistence/user"
	"github.com/Yideg/admybrand_challenge/internal/constant"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	auth_module "github.com/Yideg/admybrand_challenge/internal/module/auth"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func Initialize() {
	if os.Getenv("ENV") == "" {
		err := godotenv.Load("./../../.env")
		fmt.Println("err ", err, "os host ", os.Getenv("DB_USER"))
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	DATABASE_URL, err := constant.DbConnectionString()
	if err != nil {
		log.Fatal("database connection failed!")
	}
	common, er := GetUtils(DATABASE_URL)
	if err != nil {
		log.Fatal(er)
	}
	err = common.Conn.Migrator().AutoMigrate(
		new(model.Role),
		new(model.User),
	)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.Use(corsMW())

	authPersistence := auth_persist.UserInit(common.Conn)
	authUsecase := auth_module.Initialize(authPersistence, common)
	authHandler := auth_handler.NewAuthHandler(authUsecase, common)
	var e *casbin.Enforcer
	if os.Getenv("ENV") == "PROD" {
		e = casbin.NewEnforcer("config/rbac_model.conf", "config/policy.csv")
	} else {
		e = casbin.NewEnforcer("./../../config/rbac_model.conf", "./../../config/policy.csv")

	}

	router.Use(authHandler.Authorizer(e))
	v1 := router.Group("/api/v1")
	// initialize domains
	domain.AuthInit(common, v1)
	domain.UserInit(common, v1)
	domain.RoleInit(common, v1)
	port := os.Getenv("PORT")

	if port == "" {
		port = "1190"
	}
	router.Run(":" + port)

	logrus.WithFields(logrus.Fields{
		"host": os.Getenv("SERVER_HOST"),
		"port": ":" + os.Getenv("PORT"),
	}).Info("Starts Serving on HTTP")

}
