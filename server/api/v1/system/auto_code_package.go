package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type AutoCodePackageApi struct{}

// Create
// @Tags      AutoCodePackage
// @Summary   创建package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/createPackage [post]
func (a *AutoCodePackageApi) Create(c *gin.Context) {
	var info request.SysAutoCodePackageCreate
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if strings.Contains(info.PackageName, "\\") || strings.Contains(info.PackageName, "/") || strings.Contains(info.PackageName, "..") {
		response.FailWithMessage("包名不合法", c)
		return
	} // PackageName可能导致路径穿越的问题 / 和 \ 都要防止
	err := autoCodePackageService.Create(c.Request.Context(), &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdatePackageDetail
// @Tags      AutoCode
// @Summary   更新package展示名字/描述
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "更新package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "更新package成功"
// @Router    /autoCode/updatePackageDetail [post]
func (a *AutoCodePackageApi) UpdatePackageDetail(c *gin.Context) {
	var info request.SysAutoCodePackageCreate
	var save model.SysAutoCodePackage
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.AutoPackageVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := global.GVA_DB.Where("id = ?", info.ID).First(&save).Error; err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 更新展示名字/描述
	save.Label = info.Label
	save.Desc = info.Desc
	if err := global.GVA_DB.Save(&save).Error; err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Delete
// @Tags      AutoCode
// @Summary   删除package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建package"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "删除package成功"
// @Router    /autoCode/delPackage [post]
func (a *AutoCodePackageApi) Delete(c *gin.Context) {
	var info common.GetById
	_ = c.ShouldBindJSON(&info)
	err := autoCodePackageService.Delete(c.Request.Context(), info)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetPackageById
// @Tags      AutoCode
// @Summary   根据ID获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  body      int                                                true  "package ID"
// @Success   200  {object}  response.Response{data=system.SysAutoCode,msg=string}  "根据ID获取package成功"
// @Router    /autoCode/getPackageByID [post]
func (a *AutoCodePackageApi) GetPackageById(c *gin.Context) {
	var info common.GetById
	_ = c.ShouldBindJSON(&info)
	data, err := autoCodePackageService.GetByID(c.Request.Context(), info)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"pkg": data}, "获取成功", c)

}

// All
// @Tags      AutoCodePackage
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getPackage [post]
func (a *AutoCodePackageApi) All(c *gin.Context) {
	data, err := autoCodePackageService.All(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"pkgs": data}, "获取成功", c)
}

// Templates
// @Tags      AutoCodePackage
// @Summary   获取package
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "创建package成功"
// @Router    /autoCode/getTemplates [get]
func (a *AutoCodePackageApi) Templates(c *gin.Context) {
	data, err := autoCodePackageService.Templates(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
