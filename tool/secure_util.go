// Package tool 加解密相关工具类
// @author: kbj
// @date: 2023/2/4
package tool

import (
	"github.com/gookit/goutil/strutil"
)

// Md5Encode MD5加密
func Md5Encode(str string, count int) string {
	for i := 0; i < count; i++ {
		str = strutil.MD5(str)
	}
	return str
}
