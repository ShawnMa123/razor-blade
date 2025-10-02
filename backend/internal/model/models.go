package model

import (
	"time"
)

// Razor 剃须刀模型
type Razor struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Brand        string     `json:"brand" gorm:"not null"`
	Model        string     `json:"model" gorm:"not null"`
	PurchaseDate *time.Time `json:"purchase_date"`
	Price        *float64   `json:"price"`
	Notes        string     `json:"notes"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// 关联关系
	UsageRecords []UsageRecord `json:"usage_records,omitempty" gorm:"foreignKey:RazorID"`
}

// Blade 刀片模型
type Blade struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	Brand             string     `json:"brand" gorm:"not null"`
	Model             string     `json:"model" gorm:"not null"`
	CompatibleRazors  string     `json:"compatible_razors"` // JSON格式存储兼容的剃须刀ID列表
	PurchaseDate      *time.Time `json:"purchase_date"`
	UnitPrice         *float64   `json:"unit_price"`
	TotalQuantity     int        `json:"total_quantity" gorm:"default:0"`
	RemainingQuantity int        `json:"remaining_quantity" gorm:"default:0"`
	Notes             string     `json:"notes"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`

	// 关联关系
	UsageRecords []UsageRecord `json:"usage_records,omitempty" gorm:"foreignKey:BladeID"`
}

// UsageRecord 使用记录模型
type UsageRecord struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UsageTime       time.Time `json:"usage_time" gorm:"not null"`
	RazorID         uint      `json:"razor_id" gorm:"not null"`
	BladeID         uint      `json:"blade_id" gorm:"not null"`
	BladeUsageCount int       `json:"blade_usage_count" gorm:"default:1"`
	Rating          *int      `json:"rating"` // 1-5评分
	ExperienceText  string    `json:"experience_text"`
	NeedBladeChange bool      `json:"need_blade_change" gorm:"default:false"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// 关联关系
	Razor Razor `json:"razor" gorm:"foreignKey:RazorID"`
	Blade Blade `json:"blade" gorm:"foreignKey:BladeID"`
}

// CreateRazorRequest 创建剃须刀请求
type CreateRazorRequest struct {
	Brand        string     `json:"brand" binding:"required"`
	Model        string     `json:"model" binding:"required"`
	PurchaseDate *time.Time `json:"purchase_date"`
	Price        *float64   `json:"price"`
	Notes        string     `json:"notes"`
}

// UpdateRazorRequest 更新剃须刀请求
type UpdateRazorRequest struct {
	Brand        string     `json:"brand"`
	Model        string     `json:"model"`
	PurchaseDate *time.Time `json:"purchase_date"`
	Price        *float64   `json:"price"`
	Notes        string     `json:"notes"`
}

// CreateBladeRequest 创建刀片请求
type CreateBladeRequest struct {
	Brand             string     `json:"brand" binding:"required"`
	Model             string     `json:"model" binding:"required"`
	CompatibleRazors  string     `json:"compatible_razors"`
	PurchaseDate      *time.Time `json:"purchase_date"`
	UnitPrice         *float64   `json:"unit_price"`
	TotalQuantity     int        `json:"total_quantity"`
	RemainingQuantity int        `json:"remaining_quantity"`
	Notes             string     `json:"notes"`
}

// UpdateBladeRequest 更新刀片请求
type UpdateBladeRequest struct {
	Brand             string     `json:"brand"`
	Model             string     `json:"model"`
	CompatibleRazors  string     `json:"compatible_razors"`
	PurchaseDate      *time.Time `json:"purchase_date"`
	UnitPrice         *float64   `json:"unit_price"`
	TotalQuantity     int        `json:"total_quantity"`
	RemainingQuantity int        `json:"remaining_quantity"`
	Notes             string     `json:"notes"`
}

// CreateUsageRecordRequest 创建使用记录请求
type CreateUsageRecordRequest struct {
	UsageTime       time.Time `json:"usage_time" binding:"required"`
	RazorID         uint      `json:"razor_id" binding:"required"`
	BladeID         uint      `json:"blade_id" binding:"required"`
	BladeUsageCount int       `json:"blade_usage_count"`
	Rating          *int      `json:"rating"`
	ExperienceText  string    `json:"experience_text"`
	NeedBladeChange bool      `json:"need_blade_change"`
}

// UpdateUsageRecordRequest 更新使用记录请求
type UpdateUsageRecordRequest struct {
	UsageTime       time.Time `json:"usage_time"`
	RazorID         uint      `json:"razor_id"`
	BladeID         uint      `json:"blade_id"`
	BladeUsageCount int       `json:"blade_usage_count"`
	Rating          *int      `json:"rating"`
	ExperienceText  string    `json:"experience_text"`
	NeedBladeChange bool      `json:"need_blade_change"`
}

// APIResponse API响应格式
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
}

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
}