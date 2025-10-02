package repository

import (
	"errors"
	"razor-blade/internal/model"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	// 内存存储（当数据库不可用时使用）
	memoryRazors      []model.Razor
	memoryBlades      []model.Blade
	memoryUsageRecords []model.UsageRecord
	nextRazorID       uint
	nextBladeID       uint
	nextUsageRecordID uint
	mu                sync.RWMutex
}

func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{
		db:                db,
		memoryRazors:      make([]model.Razor, 0),
		memoryBlades:      make([]model.Blade, 0),
		memoryUsageRecords: make([]model.UsageRecord, 0),
		nextRazorID:       1,
		nextBladeID:       1,
		nextUsageRecordID: 1,
	}

	// 如果数据库不可用，初始化一些演示数据
	if db == nil {
		r.initDemoData()
	}

	return r
}

// 初始化演示数据
func (r *Repository) initDemoData() {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	// 添加示例剃须刀
	r.memoryRazors = []model.Razor{
		{
			ID:           1,
			Brand:        "Gillette",
			Model:        "Fusion 5",
			PurchaseDate: &now,
			Price:        func() *float64 { p := 89.9; return &p }(),
			Notes:        "经典五刀头剃须刀",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			ID:           2,
			Brand:        "Philips",
			Model:        "OneBlade Pro",
			PurchaseDate: &now,
			Price:        func() *float64 { p := 299.0; return &p }(),
			Notes:        "电动剃须刀，干湿两用",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	// 添加示例刀片
	r.memoryBlades = []model.Blade{
		{
			ID:                1,
			Brand:             "Gillette",
			Model:             "Fusion 5 替换刀头",
			CompatibleRazors:  "[1]", // JSON格式存储兼容的剃须刀ID
			UnitPrice:         func() *float64 { p := 15.9; return &p }(),
			TotalQuantity:     10,    // 总共10个刀头
			RemainingQuantity: 8,     // 剩余8个刀头
			Notes:             "原装替换刀头",
			CreatedAt:         now,
			UpdatedAt:         now,
		},
		{
			ID:                2,
			Brand:             "Philips",
			Model:             "OneBlade 替换刀头",
			CompatibleRazors:  "[2]", // JSON格式存储兼容的剃须刀ID
			UnitPrice:         func() *float64 { p := 25.0; return &p }(),
			TotalQuantity:     5,     // 总共5个刀头
			RemainingQuantity: 4,     // 剩余4个刀头
			Notes:             "OneBlade专用刀头",
			CreatedAt:         now,
			UpdatedAt:         now,
		},
	}

	// 添加示例使用记录
	yesterday := now.Add(-24 * time.Hour)
	r.memoryUsageRecords = []model.UsageRecord{
		{
			ID:               1,
			RazorID:          1,
			BladeID:          1,
			UsageTime:        yesterday,
			BladeUsageCount:  5,
			Rating:           func() *int { r := 4; return &r }(),
			ExperienceText:   "剃得很干净，使用感受不错",
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}

	r.nextRazorID = 3
	r.nextBladeID = 3
	r.nextUsageRecordID = 2
}

// 自动迁移数据库表
func (r *Repository) AutoMigrate() error {
	if r.db == nil {
		return nil // 没有数据库连接时跳过迁移
	}
	return r.db.AutoMigrate(
		&model.Razor{},
		&model.Blade{},
		&model.UsageRecord{},
	)
}

// Razor相关方法
func (r *Repository) CreateRazor(razor *model.Razor) error {
	if r.db == nil {
		// 使用内存存储
		r.mu.Lock()
		defer r.mu.Unlock()

		razor.ID = r.nextRazorID
		r.nextRazorID++
		razor.CreatedAt = time.Now()
		razor.UpdatedAt = time.Now()

		r.memoryRazors = append(r.memoryRazors, *razor)
		return nil
	}
	return r.db.Create(razor).Error
}

func (r *Repository) GetRazorByID(id uint) (*model.Razor, error) {
	if r.db == nil {
		// 从内存存储中查找
		r.mu.RLock()
		defer r.mu.RUnlock()

		for _, razor := range r.memoryRazors {
			if razor.ID == id {
				return &razor, nil
			}
		}
		return nil, errors.New("razor not found")
	}
	var razor model.Razor
	err := r.db.First(&razor, id).Error
	if err != nil {
		return nil, err
	}
	return &razor, nil
}

func (r *Repository) GetRazors(offset, limit int) ([]model.Razor, int64, error) {
	if r.db == nil {
		// 从内存存储中返回数据
		r.mu.RLock()
		defer r.mu.RUnlock()

		total := int64(len(r.memoryRazors))
		if offset >= len(r.memoryRazors) {
			return []model.Razor{}, total, nil
		}

		end := offset + limit
		if end > len(r.memoryRazors) {
			end = len(r.memoryRazors)
		}

		razors := make([]model.Razor, end-offset)
		copy(razors, r.memoryRazors[offset:end])
		return razors, total, nil
	}
	var razors []model.Razor
	var total int64

	if err := r.db.Model(&model.Razor{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&razors).Error
	return razors, total, err
}

func (r *Repository) UpdateRazor(razor *model.Razor) error {
	if r.db == nil {
		return errors.New("database not available")
	}
	return r.db.Save(razor).Error
}

func (r *Repository) DeleteRazor(id uint) error {
	if r.db == nil {
		return errors.New("database not available")
	}
	return r.db.Delete(&model.Razor{}, id).Error
}

// Blade相关方法
func (r *Repository) CreateBlade(blade *model.Blade) error {
	if r.db == nil {
		// 使用内存存储
		r.mu.Lock()
		defer r.mu.Unlock()

		blade.ID = r.nextBladeID
		r.nextBladeID++
		blade.CreatedAt = time.Now()
		blade.UpdatedAt = time.Now()

		r.memoryBlades = append(r.memoryBlades, *blade)
		return nil
	}
	return r.db.Create(blade).Error
}

func (r *Repository) GetBladeByID(id uint) (*model.Blade, error) {
	if r.db == nil {
		// 从内存存储中查找
		r.mu.RLock()
		defer r.mu.RUnlock()

		for _, blade := range r.memoryBlades {
			if blade.ID == id {
				return &blade, nil
			}
		}
		return nil, errors.New("blade not found")
	}
	var blade model.Blade
	err := r.db.First(&blade, id).Error
	if err != nil {
		return nil, err
	}
	return &blade, nil
}

func (r *Repository) GetBlades(offset, limit int) ([]model.Blade, int64, error) {
	if r.db == nil {
		// 从内存存储中返回数据
		r.mu.RLock()
		defer r.mu.RUnlock()

		total := int64(len(r.memoryBlades))
		if offset >= len(r.memoryBlades) {
			return []model.Blade{}, total, nil
		}

		end := offset + limit
		if end > len(r.memoryBlades) {
			end = len(r.memoryBlades)
		}

		blades := make([]model.Blade, end-offset)
		copy(blades, r.memoryBlades[offset:end])
		return blades, total, nil
	}
	var blades []model.Blade
	var total int64

	if err := r.db.Model(&model.Blade{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&blades).Error
	return blades, total, err
}

func (r *Repository) UpdateBlade(blade *model.Blade) error {
	if r.db == nil {
		// 使用内存存储更新
		r.mu.Lock()
		defer r.mu.Unlock()

		for i := range r.memoryBlades {
			if r.memoryBlades[i].ID == blade.ID {
				blade.UpdatedAt = time.Now()
				r.memoryBlades[i] = *blade
				return nil
			}
		}
		return errors.New("blade not found")
	}
	return r.db.Save(blade).Error
}

func (r *Repository) DeleteBlade(id uint) error {
	if r.db == nil {
		return errors.New("database not available")
	}
	return r.db.Delete(&model.Blade{}, id).Error
}

// UsageRecord相关方法
func (r *Repository) CreateUsageRecord(record *model.UsageRecord) error {
	if r.db == nil {
		// 使用内存存储
		r.mu.Lock()
		defer r.mu.Unlock()

		record.ID = r.nextUsageRecordID
		r.nextUsageRecordID++
		record.CreatedAt = time.Now()
		record.UpdatedAt = time.Now()

		r.memoryUsageRecords = append(r.memoryUsageRecords, *record)
		return nil
	}
	return r.db.Create(record).Error
}

func (r *Repository) GetUsageRecordByID(id uint) (*model.UsageRecord, error) {
	if r.db == nil {
		// 从内存存储中查找
		r.mu.RLock()
		defer r.mu.RUnlock()

		for _, record := range r.memoryUsageRecords {
			if record.ID == id {
				// 创建副本并填充关联对象
				result := record

				// 查找关联的剃须刀
				for _, razor := range r.memoryRazors {
					if razor.ID == record.RazorID {
						result.Razor = razor
						break
					}
				}

				// 查找关联的刀片
				for _, blade := range r.memoryBlades {
					if blade.ID == record.BladeID {
						result.Blade = blade
						break
					}
				}

				return &result, nil
			}
		}
		return nil, errors.New("usage record not found")
	}
	var record model.UsageRecord
	err := r.db.Preload("Razor").Preload("Blade").First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *Repository) GetUsageRecords(offset, limit int) ([]model.UsageRecord, int64, error) {
	if r.db == nil {
		// 从内存存储中返回数据
		r.mu.RLock()
		defer r.mu.RUnlock()

		total := int64(len(r.memoryUsageRecords))
		if offset >= len(r.memoryUsageRecords) {
			return []model.UsageRecord{}, total, nil
		}

		end := offset + limit
		if end > len(r.memoryUsageRecords) {
			end = len(r.memoryUsageRecords)
		}

		records := make([]model.UsageRecord, end-offset)
		copy(records, r.memoryUsageRecords[offset:end])

		// 填充关联的Razor和Blade对象
		for i := range records {
			// 查找关联的剃须刀
			for _, razor := range r.memoryRazors {
				if razor.ID == records[i].RazorID {
					records[i].Razor = razor
					break
				}
			}
			// 查找关联的刀片
			for _, blade := range r.memoryBlades {
				if blade.ID == records[i].BladeID {
					records[i].Blade = blade
					break
				}
			}
		}

		return records, total, nil
	}
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
	if r.db == nil {
		return errors.New("database not available")
	}
	return r.db.Save(record).Error
}

func (r *Repository) DeleteUsageRecord(id uint) error {
	if r.db == nil {
		return errors.New("database not available")
	}
	return r.db.Delete(&model.UsageRecord{}, id).Error
}

// 统计相关方法
func (r *Repository) GetUsageStatistics() (map[string]interface{}, error) {
	if r.db == nil {
		// 从内存存储中计算统计数据
		r.mu.RLock()
		defer r.mu.RUnlock()

		stats := make(map[string]interface{})

		// 总使用次数
		stats["total_usage"] = int64(len(r.memoryUsageRecords))

		// 剃须刀数量
		stats["razor_count"] = int64(len(r.memoryRazors))

		// 刀片数量
		stats["blade_count"] = int64(len(r.memoryBlades))

		// 计算平均评分
		var totalRating float64
		var ratingCount int
		for _, record := range r.memoryUsageRecords {
			if record.Rating != nil {
				totalRating += float64(*record.Rating)
				ratingCount++
			}
		}

		var avgRating float64
		if ratingCount > 0 {
			avgRating = totalRating / float64(ratingCount)
		}
		stats["average_rating"] = avgRating

		return stats, nil
	}

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
	if r.db == nil {
		// 从内存存储中返回最近的记录
		r.mu.RLock()
		defer r.mu.RUnlock()

		if len(r.memoryUsageRecords) == 0 {
			return []model.UsageRecord{}, nil
		}

		// 取最后几条记录（最新的）
		start := 0
		if len(r.memoryUsageRecords) > limit {
			start = len(r.memoryUsageRecords) - limit
		}

		records := make([]model.UsageRecord, len(r.memoryUsageRecords)-start)
		copy(records, r.memoryUsageRecords[start:])

		// 反转顺序，使最新的记录在前
		for i := len(records)/2 - 1; i >= 0; i-- {
			opp := len(records) - 1 - i
			records[i], records[opp] = records[opp], records[i]
		}

		// 填充关联的Razor和Blade对象
		for i := range records {
			// 查找关联的剃须刀
			for _, razor := range r.memoryRazors {
				if razor.ID == records[i].RazorID {
					records[i].Razor = razor
					break
				}
			}
			// 查找关联的刀片
			for _, blade := range r.memoryBlades {
				if blade.ID == records[i].BladeID {
					records[i].Blade = blade
					break
				}
			}
		}

		return records, nil
	}
	var records []model.UsageRecord
	err := r.db.Preload("Razor").Preload("Blade").
		Order("usage_time DESC").
		Limit(limit).
		Find(&records).Error
	return records, err
}