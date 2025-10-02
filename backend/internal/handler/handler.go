package handler

import (
	"net/http"
	"strconv"
	"time"

	"razor-blade/internal/model"
	"razor-blade/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	logger  *logrus.Logger
}

func NewHandler(service *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

// 通用响应方法
func (h *Handler) successResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func (h *Handler) errorResponse(c *gin.Context, statusCode int, err string) {
	h.logger.Error(err)
	c.JSON(statusCode, model.APIResponse{
		Success: false,
		Error:   err,
		Message: "操作失败",
	})
}

// 解析URL参数中的ID
func (h *Handler) parseIDParam(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// 剃须刀相关处理器
func (h *Handler) CreateRazor(c *gin.Context) {
	var req model.CreateRazorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	razor, err := h.service.CreateRazor(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, razor, "剃须刀创建成功")
}

func (h *Handler) GetRazor(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	razor, err := h.service.GetRazorByID(id)
	if err != nil {
		h.errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	h.successResponse(c, razor, "获取剃须刀成功")
}

func (h *Handler) GetRazors(c *gin.Context) {
	var req model.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	result, err := h.service.GetRazors(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, result, "获取剃须刀列表成功")
}

func (h *Handler) UpdateRazor(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	var req model.UpdateRazorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	razor, err := h.service.UpdateRazor(id, &req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, razor, "剃须刀更新成功")
}

func (h *Handler) DeleteRazor(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	if err := h.service.DeleteRazor(id); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, nil, "剃须刀删除成功")
}

// 刀片相关处理器
func (h *Handler) CreateBlade(c *gin.Context) {
	var req model.CreateBladeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	blade, err := h.service.CreateBlade(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, blade, "刀片创建成功")
}

func (h *Handler) GetBlade(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	blade, err := h.service.GetBladeByID(id)
	if err != nil {
		h.errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	h.successResponse(c, blade, "获取刀片成功")
}

func (h *Handler) GetBlades(c *gin.Context) {
	var req model.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	result, err := h.service.GetBlades(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, result, "获取刀片列表成功")
}

func (h *Handler) UpdateBlade(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	var req model.UpdateBladeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	blade, err := h.service.UpdateBlade(id, &req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, blade, "刀片更新成功")
}

func (h *Handler) DeleteBlade(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	if err := h.service.DeleteBlade(id); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, nil, "刀片删除成功")
}

// 使用记录相关处理器
func (h *Handler) CreateUsageRecord(c *gin.Context) {
	var req model.CreateUsageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	record, err := h.service.CreateUsageRecord(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, record, "使用记录创建成功")
}

func (h *Handler) GetUsageRecord(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	record, err := h.service.GetUsageRecordByID(id)
	if err != nil {
		h.errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	h.successResponse(c, record, "获取使用记录成功")
}

func (h *Handler) GetUsageRecords(c *gin.Context) {
	var req model.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	result, err := h.service.GetUsageRecords(&req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, result, "获取使用记录列表成功")
}

func (h *Handler) UpdateUsageRecord(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	var req model.UpdateUsageRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	record, err := h.service.UpdateUsageRecord(id, &req)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, record, "使用记录更新成功")
}

func (h *Handler) DeleteUsageRecord(c *gin.Context) {
	id, err := h.parseIDParam(c)
	if err != nil {
		h.errorResponse(c, http.StatusBadRequest, "无效的ID参数")
		return
	}

	if err := h.service.DeleteUsageRecord(id); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, nil, "使用记录删除成功")
}

// 统计相关处理器
func (h *Handler) GetDashboard(c *gin.Context) {
	data, err := h.service.GetDashboardData()
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, data, "获取仪表板数据成功")
}

func (h *Handler) GetStatistics(c *gin.Context) {
	stats, err := h.service.GetStatistics()
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.successResponse(c, stats, "获取统计数据成功")
}

// 健康检查
func (h *Handler) HealthCheck(c *gin.Context) {
	h.successResponse(c, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}, "服务正常运行")
}