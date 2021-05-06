package service

import (
	"context"
	"fmt"
	"gf-L/app/dao"
	"gf-L/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var User = userService{}

type userService struct {
}

// 登录操作 登录成功后更新session和上下文中的用户信息
func (u userService) Login(ctx context.Context, userEntity *model.UserApiLoginReq) error {
	user, error := u.getUserByPassportAndPassword(
		userEntity.Passport,
		userEntity.Password,
	)
	if error != nil {
		return error
	}
	if user == nil {
		return gerror.New("账号密码错误")
	}
	error = Session.SetUser(ctx, user)
	if error != nil {
		return error
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		IsAdmin:  false,
	})
	return nil
}

// 通过账号密码获取数据库中的用户
func (userService) getUserByPassportAndPassword(passport, password string) (*model.User, error) {
	return dao.User.Where(g.Map{
		dao.User.Columns.Passport: passport,
		dao.User.Columns.Password: password,
	}).One()
}

// 登出 删除session中的用户
func (userService) Logout(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// 注册
func (u userService) Register(ctx context.Context, user *model.User) (string, error) {
	return u.insertUser(user)
}

func (userService) insertUser(user *model.User) (string, error) {
	r, err := dao.User.Data(user).Insert()
	if err != nil {
		return "", err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		return "", err
	}
	if affected == 0 {
		err = gerror.New("无数据插入数据库")
	}
	id, err := r.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(id), err
}
