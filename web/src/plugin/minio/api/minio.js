import service from '@/utils/request'
// @Tags Minio
// @Summary 创建minio存储
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Minio true "创建minio存储"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MI/createMinio [post]
export const createMinio = (data) => {
  return service({
    url: '/MI/createMinio',
    method: 'post',
    data,
  })
}

// @Tags Minio
// @Summary 删除minio存储
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Minio true "删除minio存储"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MI/deleteMinio [delete]
export const deleteMinio = (params) => {
  return service({
    url: '/MI/deleteMinio',
    method: 'delete',
    params,
  })
}

// @Tags Minio
// @Summary 批量删除minio存储
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除minio存储"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MI/deleteMinio [delete]
export const deleteMinioByIds = (params) => {
  return service({
    url: '/MI/deleteMinioByIds',
    method: 'delete',
    params,
  })
}

// @Tags Minio
// @Summary 更新minio存储
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Minio true "更新minio存储"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MI/updateMinio [put]
export const updateMinio = (data) => {
  return service({
    url: '/MI/updateMinio',
    method: 'put',
    data,
  })
}

// @Tags Minio
// @Summary 用id查询minio存储
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Minio true "用id查询minio存储"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MI/findMinio [get]
export const findMinio = (params) => {
  return service({
    url: '/MI/findMinio',
    method: 'get',
    params,
  })
}

// @Tags Minio
// @Summary 分页获取minio存储列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取minio存储列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MI/getMinioList [get]
export const getMinioList = (params) => {
  return service({
    url: '/MI/getMinioList',
    method: 'get',
    params,
  })
}
// @Tags Minio
// @Summary 不需要鉴权的minio存储接口
// @accept application/json
// @Produce application/json
// @Param data query request.MinioSearch true "分页获取minio存储列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MI/getMinioPublic [get]
export const getMinioPublic = () => {
  return service({
    url: '/MI/getMinioPublic',
    method: 'get',
  })
}
