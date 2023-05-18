package tool

import (
	"fmt"
	"github.com/gookit/goutil/arrutil"
	"path"
	"react-admin-server/global/consts"
)

// FileNameCheck 文件名称格式合法性检查
func FileNameCheck(fileName string) error {
	if len(fileName) > consts.MaxFileNameSize {
		return consts.NewServiceError("文件长度过长")
	}
	ext := path.Ext(fileName)
	if ext == "" {
		return consts.NewServiceError("请上传正确格式的文件")
	}
	if !arrutil.Contains(consts.FileAllowedExtension, ext) {
		return consts.NewServiceError(fmt.Sprintf("暂时不支持上传%s格式", ext))
	}
	return nil
}

// FileMimeTypeCheck 文件MimeType合法性检查
func FileMimeTypeCheck(dir string) error {
	return nil
}
