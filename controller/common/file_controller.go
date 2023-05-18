package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/goutil/fsutil"
	"go.uber.org/zap"
	"os"
	"path"
	"react-admin-server/entity/vo"
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
}

// UploadFile 上传文件
func (*FileController) UploadFile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		g.Logger.Error("读取上传文件失败", zap.Error(err))
		return consts.NewServiceError("上传失败")
	} else if err = tool.FileNameCheck(file.Filename); err != nil {
		return err
	}
	fileName := file.Filename

	// 保存文件
	now := time.Now()
	newFileName := fmt.Sprintf("%s_%s%s", strings.TrimSuffix(fileName, path.Ext(fileName)), strconv.FormatInt(now.UnixNano(), 10), path.Ext(fileName))
	relativePath := fmt.Sprintf("/%d/%s", now.Year(), now.Format("01"))
	targetPath := g.Env.System.UploadPath + relativePath + "/" + newFileName
	if !fsutil.PathExists(g.Env.System.UploadPath + relativePath) {
		_ = os.MkdirAll(g.Env.System.UploadPath+relativePath, os.ModePerm)
	}
	if err = ctx.SaveFile(file, targetPath); err != nil {
		g.Logger.Error("保存上传附件失败", zap.String("path", targetPath), zap.Error(err))
		return consts.NewServiceError("上传失败")
	}

	resp := vo.FileInfo{
		OriginFileName: fileName,
		NewFileName:    newFileName,
		Path:           relativePath + "/" + newFileName,
	}

	return r.Ok(ctx, r.Data(&resp), r.Msg("ok"))
}
