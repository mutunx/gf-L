package service

import (
	"context"
	"gf-L/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 头像服务
var Avatar = avatarService{}

type avatarService struct {
}

// 上传头像到目录中 使用id作为名称
func (s avatarService) Upload(ctx context.Context, files ghttp.UploadFiles, id string) (err error) {
	sourcePath := "public/"
	filePath := "image/avatar/"
	_, err = files.Save(sourcePath + filePath + id)
	if err != nil {
		return
	}
	_, err = dao.User.UserDao.Update(g.Map{dao.User.Columns.Avatar: filePath + id}, dao.User.Columns.Id, id)
	return
}
