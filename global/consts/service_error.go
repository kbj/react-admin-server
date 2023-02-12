// Package consts 自定义业务相关的错误
// @author: kbj
// @date: 2023/2/2
package consts

import (
	"react-admin-server/entity/vo"
)

type ServiceError vo.Response

func (e ServiceError) Error() string {
	return e.Msg
}

func NewServiceError(msg string) error {
	return ServiceError{
		Code: StatusFailure,
		Msg:  msg,
	}
}
