// Package tool 校验参数工具类
// @author: kbj
// @date: 2023/2/3
package tool

import (
	"github.com/go-playground/validator/v10"
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
	"reflect"
	"strings"
)

// ValidateParams 校验参数
func ValidateParams(param any) error {
	// 只允许传入结构体指针
	if reflect.TypeOf(param).Kind() != reflect.Pointer {
		return consts.NewServiceError("类型有误，校验失败")
	}
	err := g.Validator.Struct(param)
	//如果数据效验不通过，则将所有err以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			sliceErrs = append(sliceErrs, e.Translate(*g.ValidatorTranslator))
		}
		return consts.NewServiceError(strings.Join(sliceErrs, ","))
	}
	return nil
}
