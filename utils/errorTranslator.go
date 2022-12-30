package util

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func TranslateError(s interface{}) (errs []string) {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	validatorErrors := err.(validator.ValidationErrors)
	fmt.Println(validatorErrors)
	for _, e := range validatorErrors {
		translatedErrors := e.Translate(trans)
		errs = append(errs, translatedErrors)
	}

	return errs
}
