package utils

import (
	"strconv"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func InitTranslator() (*validator.Validate, ut.Translator) {
	// creating new validator instance
	validate := validator.New()

	en := en.New()

	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	return validate, trans

}

func TranslateError(validationErrs validator.ValidationErrors, trans ut.Translator) map[string]string {
	err_map := make(map[string]string)

	if len(validationErrs) != 0 {
		for _, e := range validationErrs {
			err_map[strings.ToLower(e.Field())] = e.Translate(trans)
		}
	}

	return err_map
}

// changing string to uint
func StringToUint(s string) uint32 {
	i, _ := strconv.Atoi(s)
	return uint32(i)
}
