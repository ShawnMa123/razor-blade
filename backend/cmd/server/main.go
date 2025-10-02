package main

import (
	"log"
	"razor-blade/internal/config"
	"razor-blade/internal/handler"
	"razor-blade/internal/repository"
	"razor-blade/internal/router"
	"razor-blade/internal/service"
	"razor-blade/pkg/database"
	"razor-blade/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	appLogger := logger.InitLogger(&cfg.Log)
	appLogger.Info("Starting Razor-Blade server...")

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库
	db, err := database.InitDB(cfg.Database.Path)
	if err != nil {
		appLogger.Warnf("Failed to initialize database: %v", err)
		appLogger.Info("Running without database (CGO disabled)")
		db = nil
	} else {
		appLogger.Info("Database connected successfully")
	}

	// 初始化各层
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, appLogger)

	// 自动迁移数据库 (仅在数据库可用时)
	if db != nil {
		if err := repo.AutoMigrate(); err != nil {
			appLogger.Fatalf("Failed to migrate database: %v", err)
		}
		appLogger.Info("Database migration completed")
	} else {
		appLogger.Info("Skipping database migration (no database connection)")
	}

	// 设置路由
	r := router.SetupRouter(h, appLogger)

	// 启动服务器
	appLogger.Infof("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		appLogger.Fatalf("Failed to start server: %v", err)
	}
}