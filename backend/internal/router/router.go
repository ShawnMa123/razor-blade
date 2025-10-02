package router

import (
	"razor-blade/internal/handler"
	"razor-blade/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter(h *handler.Handler, logger *logrus.Logger) *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.ErrorHandlerMiddleware(logger))
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", h.HealthCheck)

	// API路由组
	api := r.Group("/api/v1")
	{
		// 剃须刀路由
		razors := api.Group("/razors")
		{
			razors.POST("", h.CreateRazor)
			razors.GET("", h.GetRazors)
			razors.GET("/:id", h.GetRazor)
			razors.PUT("/:id", h.UpdateRazor)
			razors.DELETE("/:id", h.DeleteRazor)
		}

		// 刀片路由
		blades := api.Group("/blades")
		{
			blades.POST("", h.CreateBlade)
			blades.GET("", h.GetBlades)
			blades.GET("/:id", h.GetBlade)
			blades.PUT("/:id", h.UpdateBlade)
			blades.DELETE("/:id", h.DeleteBlade)
		}

		// 使用记录路由
		usageRecords := api.Group("/usage-records")
		{
			usageRecords.POST("", h.CreateUsageRecord)
			usageRecords.GET("", h.GetUsageRecords)
			usageRecords.GET("/:id", h.GetUsageRecord)
			usageRecords.PUT("/:id", h.UpdateUsageRecord)
			usageRecords.DELETE("/:id", h.DeleteUsageRecord)
		}

		// 统计和仪表板路由
		api.GET("/dashboard", h.GetDashboard)
		api.GET("/statistics", h.GetStatistics)
	}

	return r
}