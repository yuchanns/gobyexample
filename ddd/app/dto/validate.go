package dto

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

func validate(v interface{}) error {
	zhCh := zh.New()
	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	vd := validator.New()
	if err := zhtranslations.RegisterDefaultTranslations(vd, trans); err != nil {
		return errors.WithStack(err)
	}
	if err := vd.Struct(v); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}

	return nil
}
