package utils

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/translations/zh"
	"gopkg.in/go-playground/validator.v9"
)

// Trans 语法校验
func (_ Utils) Trans(data interface{}) *string {
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
		for k, e := range err.(validator.ValidationErrors) {
			tag := "，"
			if k == len(err.(validator.ValidationErrors))-1 {
				tag = "。"
			}
			str += e.Translate(trans) + tag
		}
		return &str
	}
	return nil
}
