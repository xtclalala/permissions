package utils

import (
	"errors"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// Validate 验证字段，结构体内的结构体属性也会被验证
func Validate(data any) error {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err := validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return errors.New(v.Translate(trans))
		}
	}
	return nil
}
