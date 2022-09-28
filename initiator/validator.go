package initiator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func GetValidation() (ut.Translator, *validator.Validate, error) {
	v := validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, nil, errors.New("translator not found")
	}

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, nil, err
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
	_ = v.RegisterTranslation("address", trans, func(ut ut.Translator) error {
		return ut.Add("address", "{0} must be a valid address", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("address", fe.Field())
		return t
	})

	_ = v.RegisterValidation("address", func(fl validator.FieldLevel) bool {
		var isAlpha = regexp.MustCompile(`^[A-Za-z]+$`)
		return isAlpha.MatchString(fl.Field().String())
	})

	_ = v.RegisterTranslation("passwordcheck", trans, func(ut ut.Translator) error {
		return ut.Add("passwordcheck", "{0} is not strong enough", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwordcheck", fe.Field())
		return t
	})

	_ = v.RegisterValidation("passwordcheck", func(fl validator.FieldLevel) bool {
		fmt.Println("password rule failed :", len(fl.Field().String()) > 8)
		return (len(fl.Field().String()) > 8) && (len(fl.Field().String()) < 30)
	})
	_ = v.RegisterTranslation("ethiopianhone", trans, func(ut ut.Translator) error {
		return ut.Add("ethphone", "{0} is not a valid ethiopian phone number.", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ethphone", fe.Field())
		return t
	})

	_ = v.RegisterValidation("ethiopianhone", func(fl validator.FieldLevel) bool {
		var reg = regexp.MustCompile(`^((91)|(\+91)|0)?9\d{8}$`)
		return reg.MatchString(fl.Field().String())
	})

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return trans, v, nil
}
