package validator

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
)

var Validate *validator.Validate
var Trans ut.Translator

func InitValidator()  {
	translator := ut.New(zh.New())
	Trans ,_=translator.GetTranslator("zh")
	// 校验器
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	// 注册翻译器到校验器
	err := zhTranslations.RegisterDefaultTranslations(validate, Trans)
	if err != nil {
		log.Panicln(err)
	}
}

// ValidateStruct 验证结构体
func ValidateStruct(s interface{}) map[string]string{
	errs:=make(map[string]string)
	err := Validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Tag()]=e.Translate(Trans)
		}
	}
	return  errs
}