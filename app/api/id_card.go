package api

import (
	"gf-L/app/model"
	"gf-L/app/service"
	"gf-L/library"
	"github.com/gogf/gf/net/ghttp"
)

// 身份证相关接口
var IdCard = idCardApi{}

type idCardApi struct {
}

// GetFullNameBy6Code godoc
// @Summary 通过6位编码获取代表地区全称
// @Description  传入330302返回浙江省温州市鹿城区
// @Param code formData string true "编码"
// @Accept  json
// @Produce  json
// @Success 200 object} library.JsonRes
// @Failure 400,404 {object} library.JsonRes
// @Failure 500 {object} library.JsonRes
// @Router /idCard/get-full-name-by6-code [post]
func (idCardApi) GetFullNameBy6Code(r *ghttp.Request) {
	var idCardApiReq *model.IdCardApiReq
	if err := r.Parse(&idCardApiReq); err != nil {
		library.JsonExit(r, 1, err.Error())
	}
	if code, err := service.IdCard.GetNameByCode(r.Context(), idCardApiReq.Code); err != nil {
		library.JsonExit(r, 500, err.Error())
	} else {
		library.JsonExit(r, 0, "", code)
	}

}
