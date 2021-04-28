// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gf-L/app/model/internal"
)

// User is the golang structure for table user.
type User internal.User

// Api用户登录
type UserApiLoginReq struct {
	Passport string `v:"required#请输入账号"` // 账号
	Password string `v:"required#请输入密码"` // 密码(明文)
}
