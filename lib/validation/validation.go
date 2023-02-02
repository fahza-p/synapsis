package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fahza-p/synapsis/lib/response"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	/* Init Translator */
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")

	/* Init Validate */
	validate = validator.New()

	validate.RegisterTagNameFunc(getJsonTagName)

	/* Register Translator */
	en_translations.RegisterDefaultTranslations(validate, trans)

	/* Register Custom Rule */

	/* Register Custom Rule Message */
}

/* Run Valiedation for Custom Error Message */
func RunValidate(i interface{}, r *response.Build) error {
	var errMap = make(map[string]string)
	var sb strings.Builder
	err := validate.Struct(i)
	sep := " | "

	if err != nil {
		for i, v := range err.(validator.ValidationErrors) {
			if len(err.(validator.ValidationErrors)) == i+1 {
				sep = ""
			}
			sb.WriteString(v.Translate(trans) + sep)
			errMap[v.Field()] = v.Translate(trans)
		}

		r.Err = errMap
		r.Msg = sb.String()
	}

	return err
}

/* Add Message Translation */
func addTranslation(tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		tag := fe.Tag()
		t, err := ut.T(tag, fe.Field(), fmt.Sprintf("%v", fe.Value()))

		if err != nil {
			return fe.(error).Error()
		}

		return t
	}

	validate.RegisterTranslation(tag, trans, registerFn, transFn)
}

/* Register Tag Name */
func getJsonTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
