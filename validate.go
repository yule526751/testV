package validator

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"strings"
)

var Trans ut.Translator

// InitValidator 初始化中文验证器
func InitValidator() {
	// 中文翻译器
	t := ut.New(zh.New())
	Trans, _ = t.GetTranslator("zh")

	// 校验器
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})

	// 注册翻译器到校验器
	err := zhTranslations.RegisterDefaultTranslations(v, Trans)
	if err != nil {
		log.Panicln(err)
	}
}

// FormatErr 格式化错误
func FormatErr(err error) map[string]string {
	rs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		key := strings.SplitN(e.StructNamespace(), ".", 2)[1]
		rs[key] = e.Translate(Trans)
	}
	return rs
}
