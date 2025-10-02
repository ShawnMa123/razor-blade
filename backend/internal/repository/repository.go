package repository

import (
	"razor-blade/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// 自动迁移数据库表
func (r *Repository) AutoMigrate() error {
	return r.db.AutoMigrate(
		&model.Razor{},
		&model.Blade{},
		&model.UsageRecord{},
	)
}

// Razor相关方法
func (r *Repository) CreateRazor(razor *model.Razor) error {
	return r.db.Create(razor).Error
}

func (r *Repository) GetRazorByID(id uint) (*model.Razor, error) {
	var razor model.Razor
	err := r.db.First(&razor, id).Error
	if err != nil {
		return nil, err
	}
	return &razor, nil
}

func (r *Repository) GetRazors(offset, limit int) ([]model.Razor, int64, error) {
	var razors []model.Razor
	var total int64

	if err := r.db.Model(&model.Razor{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&razors).Error
	return razors, total, err
}

func (r *Repository) UpdateRazor(razor *model.Razor) error {
	return r.db.Save(razor).Error
}

func (r *Repository) DeleteRazor(id uint) error {
	return r.db.Delete(&model.Razor{}, id).Error
}

// Blade相关方法
func (r *Repository) CreateBlade(blade *model.Blade) error {
	return r.db.Create(blade).Error
}

func (r *Repository) GetBladeByID(id uint) (*model.Blade, error) {
	var blade model.Blade
	err := r.db.First(&blade, id).Error
	if err != nil {
		return nil, err
	}
	return &blade, nil
}

func (r *Repository) GetBlades(offset, limit int) ([]model.Blade, int64, error) {
	var blades []model.Blade
	var total int64

	if err := r.db.Model(&model.Blade{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&blades).Error
	return blades, total, err
}

func (r *Repository) UpdateBlade(blade *model.Blade) error {
	return r.db.Save(blade).Error
}

func (r *Repository) DeleteBlade(id uint) error {
	return r.db.Delete(&model.Blade{}, id).Error
}

// UsageRecord相关方法
func (r *Repository) CreateUsageRecord(record *model.UsageRecord) error {
	return r.db.Create(record).Error
}

func (r *Repository) GetUsageRecordByID(id uint) (*model.UsageRecord, error) {
	var record model.UsageRecord
	err := r.db.Preload("Razor").Preload("Blade").First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *Repository) GetUsageRecords(offset, limit int) ([]model.UsageRecord, int64, error) {
	var records []model.UsageRecord
	var total int64

	if err := r.db.Model(&model.UsageRecord{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Razor").Preload("Blade").
		Order("usage_time DESC").
		Offset(offset).Limit(limit).
		Find(&records).Error
	return records, total, err
}

func (r *Repository) UpdateUsageRecord(record *model.UsageRecord) error {
	return r.db.Save(record).Error
}

func (r *Repository) DeleteUsageRecord(id uint) error {
	return r.db.Delete(&model.UsageRecord{}, id).Error
}

// 统计相关方法
func (r *Repository) GetUsageStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 总使用次数
	var totalUsage int64
	if err := r.db.Model(&model.UsageRecord{}).Count(&totalUsage).Error; err != nil {
		return nil, err
	}
	stats["total_usage"] = totalUsage

	// 剃须刀数量
	var razorCount int64
	if err := r.db.Model(&model.Razor{}).Count(&razorCount).Error; err != nil {
		return nil, err
	}
	stats["razor_count"] = razorCount

	// 刀片数量
	var bladeCount int64
	if err := r.db.Model(&model.Blade{}).Count(&bladeCount).Error; err != nil {
		return nil, err
	}
	stats["blade_count"] = bladeCount

	// 平均评分
	var avgRating float64
	if err := r.db.Model(&model.UsageRecord{}).
		Where("rating IS NOT NULL").
		Select("AVG(rating)").
		Scan(&avgRating).Error; err != nil {
		return nil, err
	}
	stats["average_rating"] = avgRating

	return stats, nil
}

func (r *Repository) GetRecentUsageRecords(limit int) ([]model.UsageRecord, error) {
	var records []model.UsageRecord
	err := r.db.Preload("Razor").Preload("Blade").
		Order("usage_time DESC").
		Limit(limit).
		Find(&records).Error
	return records, err
}