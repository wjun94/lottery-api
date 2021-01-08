package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/translations/zh"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// Trans 语法校验
func Trans(data interface{}) *string {
	//验证
	str := ""
	zh_ch := zh.New()
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			str += err.Translate(trans) + "\n"
		}
		return &str
	}
	return nil
}
