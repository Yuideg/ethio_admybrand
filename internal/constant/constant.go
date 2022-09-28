package constant

import (
	"fmt"
	"github.com/Yideg/admybrand_challenge/internal/constant/errors"
	"github.com/Yideg/admybrand_challenge/internal/constant/model"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"math"
	"os"
	"regexp"
	"strings"
)

type SuccessData struct {
	Sucess bool
	Data   interface{}
}

func CompareHashAndPassword(new, old []byte) error {
	err := bcrypt.CompareHashAndPassword(new, old)
	return err
}
func HashPassword(c *gin.Context, pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

//StructValidator validates specific struct
func StructValidator(structName interface{}, validate *validator.Validate, trans ut.Translator) []string {
	var errorList []string

	errV := validate.Struct(structName)
	if errV != nil {
		errs := errV.(validator.ValidationErrors)
		for _, e := range errs {
			errorList = append(errorList, e.Translate(trans))
		}
		return errorList
	}
	return nil
}

//ResponseJson creates new json object
func ResponseJson(c *gin.Context, responseData interface{}, statusCode int) {
	c.JSON(statusCode, responseData)
	return
}

// wrap field validator with error code
func VerifyInput(structName interface{}, validate *validator.Validate, trans ut.Translator) *errors.ErrorModel {
	errs := StructValidator(structName, validate, trans)
	if errs == nil {
		return nil
	}
	return &errors.ErrorModel{
		ErrorCode:        errors.ErrCodes[errors.ErrInvalidField],
		ErrorDescription: errors.Descriptions[errors.ErrInvalidField],
		ErrorMessage:     errors.ErrOneOrMoreFieldsInvalid.Error(),
		ErrorDetail:      errs,
	}
}

//DbConnectionString connction string finder from the .env file
func DbConnectionString() (string, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbsslmode := os.Getenv("DB_SSL_MODE")

	addr := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=%v", user, password, host, dbname, dbsslmode)
	fmt.Println("addr", addr)
	return addr, nil
}
func ValidRole(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) != 0 && s != "anonymous" {
		return true
	}
	return false

}

// IsNumber checks if the string is a number
func IsNumber(str string) bool {
	number_regex := regexp.MustCompile(`\d+`)
	if str == "" || str == "0" {
		return false
	}

	return number_regex.MatchString(str)
}

// Contains checks if a role is present in a slice of roles
func Contains(roles []string, role string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}

// DistanceBetweenTwoPoints find the distance between two points which is represented by latitude and longitude values
func DistanceBetweenTwoPoints(arival model.Coordinate, destination model.Coordinate) float64 {
	radlat1 := float64(math.Pi * arival.Latitude / 180)
	radlat2 := float64(math.Pi * destination.Latitude / 180)

	theta := float64(arival.Longitude - destination.Longitude)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344

	return dist
}
