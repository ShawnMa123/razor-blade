package service

import (
	"errors"
	"math"
	"razor-blade/internal/model"
	"razor-blade/internal/repository"

	"gorm.io/gorm"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Razor服务方法
func (s *Service) CreateRazor(req *model.CreateRazorRequest) (*model.Razor, error) {
	razor := &model.Razor{
		Brand:        req.Brand,
		Model:        req.Model,
		PurchaseDate: req.PurchaseDate,
		Price:        req.Price,
		Notes:        req.Notes,
	}

	if err := s.repo.CreateRazor(razor); err != nil {
		return nil, err
	}

	return razor, nil
}

func (s *Service) GetRazorByID(id uint) (*model.Razor, error) {
	return s.repo.GetRazorByID(id)
}

func (s *Service) GetRazors(req *model.PaginationRequest) (*model.PaginationResponse, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	razors, total, err := s.repo.GetRazors(offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.PaginationResponse{
		Items:      razors,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *Service) UpdateRazor(id uint, req *model.UpdateRazorRequest) (*model.Razor, error) {
	razor, err := s.repo.GetRazorByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("剃须刀不存在")
		}
		return nil, err
	}

	if req.Brand != "" {
		razor.Brand = req.Brand
	}
	if req.Model != "" {
		razor.Model = req.Model
	}
	if req.PurchaseDate != nil {
		razor.PurchaseDate = req.PurchaseDate
	}
	if req.Price != nil {
		razor.Price = req.Price
	}
	razor.Notes = req.Notes

	if err := s.repo.UpdateRazor(razor); err != nil {
		return nil, err
	}

	return razor, nil
}

func (s *Service) DeleteRazor(id uint) error {
	// 检查是否存在使用记录
	_, err := s.repo.GetRazorByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("剃须刀不存在")
		}
		return err
	}

	return s.repo.DeleteRazor(id)
}

// Blade服务方法
func (s *Service) CreateBlade(req *model.CreateBladeRequest) (*model.Blade, error) {
	blade := &model.Blade{
		Brand:             req.Brand,
		Model:             req.Model,
		CompatibleRazors:  req.CompatibleRazors,
		PurchaseDate:      req.PurchaseDate,
		UnitPrice:         req.UnitPrice,
		TotalQuantity:     req.TotalQuantity,
		RemainingQuantity: req.RemainingQuantity,
		Notes:             req.Notes,
	}

	if err := s.repo.CreateBlade(blade); err != nil {
		return nil, err
	}

	return blade, nil
}

func (s *Service) GetBladeByID(id uint) (*model.Blade, error) {
	return s.repo.GetBladeByID(id)
}

func (s *Service) GetBlades(req *model.PaginationRequest) (*model.PaginationResponse, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	blades, total, err := s.repo.GetBlades(offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.PaginationResponse{
		Items:      blades,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *Service) UpdateBlade(id uint, req *model.UpdateBladeRequest) (*model.Blade, error) {
	blade, err := s.repo.GetBladeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("刀片不存在")
		}
		return nil, err
	}

	if req.Brand != "" {
		blade.Brand = req.Brand
	}
	if req.Model != "" {
		blade.Model = req.Model
	}
	blade.CompatibleRazors = req.CompatibleRazors
	if req.PurchaseDate != nil {
		blade.PurchaseDate = req.PurchaseDate
	}
	if req.UnitPrice != nil {
		blade.UnitPrice = req.UnitPrice
	}
	blade.TotalQuantity = req.TotalQuantity
	blade.RemainingQuantity = req.RemainingQuantity
	blade.Notes = req.Notes

	if err := s.repo.UpdateBlade(blade); err != nil {
		return nil, err
	}

	return blade, nil
}

func (s *Service) DeleteBlade(id uint) error {
	_, err := s.repo.GetBladeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("刀片不存在")
		}
		return err
	}

	return s.repo.DeleteBlade(id)
}

// UsageRecord服务方法
func (s *Service) CreateUsageRecord(req *model.CreateUsageRecordRequest) (*model.UsageRecord, error) {
	// 验证剃须刀和刀片是否存在
	_, err := s.repo.GetRazorByID(req.RazorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("剃须刀不存在")
		}
		return nil, err
	}

	_, err = s.repo.GetBladeByID(req.BladeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("刀片不存在")
		}
		return nil, err
	}

	record := &model.UsageRecord{
		UsageTime:       req.UsageTime,
		RazorID:         req.RazorID,
		BladeID:         req.BladeID,
		BladeUsageCount: req.BladeUsageCount,
		Rating:          req.Rating,
		ExperienceText:  req.ExperienceText,
		NeedBladeChange: req.NeedBladeChange,
	}

	if record.BladeUsageCount == 0 {
		record.BladeUsageCount = 1
	}

	if err := s.repo.CreateUsageRecord(record); err != nil {
		return nil, err
	}

	return s.repo.GetUsageRecordByID(record.ID)
}

func (s *Service) GetUsageRecordByID(id uint) (*model.UsageRecord, error) {
	return s.repo.GetUsageRecordByID(id)
}

func (s *Service) GetUsageRecords(req *model.PaginationRequest) (*model.PaginationResponse, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	records, total, err := s.repo.GetUsageRecords(offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.PaginationResponse{
		Items:      records,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (s *Service) UpdateUsageRecord(id uint, req *model.UpdateUsageRecordRequest) (*model.UsageRecord, error) {
	record, err := s.repo.GetUsageRecordByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("使用记录不存在")
		}
		return nil, err
	}

	record.UsageTime = req.UsageTime
	record.RazorID = req.RazorID
	record.BladeID = req.BladeID
	record.BladeUsageCount = req.BladeUsageCount
	record.Rating = req.Rating
	record.ExperienceText = req.ExperienceText
	record.NeedBladeChange = req.NeedBladeChange

	if err := s.repo.UpdateUsageRecord(record); err != nil {
		return nil, err
	}

	return s.repo.GetUsageRecordByID(record.ID)
}

func (s *Service) DeleteUsageRecord(id uint) error {
	_, err := s.repo.GetUsageRecordByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("使用记录不存在")
		}
		return err
	}

	return s.repo.DeleteUsageRecord(id)
}

// 统计服务方法
func (s *Service) GetDashboardData() (map[string]interface{}, error) {
	stats, err := s.repo.GetUsageStatistics()
	if err != nil {
		return nil, err
	}

	recentRecords, err := s.repo.GetRecentUsageRecords(5)
	if err != nil {
		return nil, err
	}

	dashboardData := map[string]interface{}{
		"statistics":     stats,
		"recent_records": recentRecords,
	}

	return dashboardData, nil
}

func (s *Service) GetStatistics() (map[string]interface{}, error) {
	return s.repo.GetUsageStatistics()
}