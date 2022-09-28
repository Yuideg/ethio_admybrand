package initiator

import (
	"github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

func GetUtils(dbUrl string) (model.Utils, *errors.ErrorModel) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		SkipDefaultTransaction: true, //30% performance increases
		Logger:                 newLogger,
	})
	if err != nil {
		log.Printf("Error when Opening database connection: %v", err)
		os.Exit(1)
	}
	trans, validate, err := GetValidation()
	if err != nil {
		log.Fatal("*errors.ErrorModel ", err)
	}
	duration, _ := strconv.Atoi(os.Getenv("TIMEOUT"))
	timeoutContext := time.Duration(duration) * time.Second
	return model.Utils{
		Timeout:     timeoutContext,
		Translator:  trans,
		GoValidator: validate,
		Conn:        conn,
	}, nil
}
