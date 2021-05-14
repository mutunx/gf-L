package service

import (
	"context"
	"gf-L/app/dao"
	"gf-L/app/model"
	"github.com/gogf/gf/errors/gerror"
)

// 身份证相关服务
var IdCard = idCardService{}

type idCardService struct {
}

// 根据code长度获取不同的值 2位取省市 4位取市县 6取全称
func (s idCardService) GetNameByCode(ctx context.Context, code string) (*model.ProvinceCityCode, error) {
	if len(code) < 6 {
		return nil, gerror.New("错误编码数")
	}
	code = code[0:6]
	var codeField string
	var nameField string
	switch len(code) {
	case 2:
		codeField = dao.ProvinceCityCode.Columns.ProvinceCode
		nameField = dao.ProvinceCityCode.Columns.ProvinceName
	case 4:
		codeField = dao.ProvinceCityCode.Columns.CityCode
		nameField = dao.ProvinceCityCode.Columns.CityName
	case 6:
		codeField = dao.ProvinceCityCode.Columns.Code
		nameField = dao.ProvinceCityCode.Columns.Name
	}
	return s.getDbNameByCode(ctx, code, codeField, nameField)
}

func (s idCardService) getDbNameByCode(ctx context.Context, code, codeField, nameField string) (*model.ProvinceCityCode, error) {
	return dao.ProvinceCityCode.Where(codeField, code).One()
}
