package postgres

import (
	"context"
	"errors"
	"github.com/Esaak/banner-service/internal/models"
	"gorm.io/gorm"
)

// BannerRepository is an interface for banner repository
type BannerRepository interface {
	GetBanner(ctx context.Context, tagID, featureID int64, useLastRevision bool) (*models.Banner, error)
	GetBanners(ctx context.Context, featureID, tagID *int64, limit, offset *int) ([]models.Banner, error)
	CreateBanner(ctx context.Context, banner *models.Banner) (int64, error)
	UpdateBanner(ctx context.Context, bannerID int64, updates map[string]interface{}) error
	DeleteBanner(ctx context.Context, bannerID int64) error
}

type bannerRepository struct {
	db *gorm.DB
}

// NewBannerRepository creates a new instance of BannerRepository
func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{db: db}
}

// GetBanner retrieves a banner for a given tag and feature
func (r *bannerRepository) GetBanner(ctx context.Context, tagID, featureID int64, useLastRevision bool) (*models.Banner, error) {
	var banner models.Banner
	query := r.db.Preload("Tags").Where("feature_id = ?", featureID)

	if !useLastRevision {
		query = query.Where("updated_at <= NOW() - INTERVAL '5 MINUTES'")
	}

	if err := query.Joins("JOIN banner_tags ON banners.id = banner_tags.banner_id").
		Where("banner_tags.tag_id = ?", tagID).
		First(&banner).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &banner, nil
}

// GetBanners retrieves banners with filtering by feature and/or tag
func (r *bannerRepository) GetBanners(ctx context.Context, featureID, tagID *int64, limit, offset *int) ([]models.Banner, error) {
	var banners []models.Banner
	query := r.db.Preload("Tags")

	if featureID != nil {
		query = query.Where("feature_id = ?", *featureID)
	}

	if tagID != nil {
		query = query.Joins("JOIN banner_tags ON banners.id = banner_tags.banner_id").
			Where("banner_tags.tag_id = ?", *tagID)
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	if err := query.Find(&banners).Error; err != nil {
		return nil, err
	}

	return banners, nil
}

// CreateBanner creates a new banner
func (r *bannerRepository) CreateBanner(ctx context.Context, banner *models.Banner) (int64, error) {
	result := r.db.Create(banner)
	if result.Error != nil {
		return 0, result.Error
	}

	for _, tagID := range banner.TagIDs {
		if err := r.db.Create(&models.BannerTag{
			BannerID: banner.ID,
			TagID:    tagID,
		}).Error; err != nil {
			return 0, err
		}
	}

	return banner.ID, nil
}

// UpdateBanner updates an existing banner
func (r *bannerRepository) UpdateBanner(ctx context.Context, bannerID int64, updates map[string]interface{}) error {
	result := r.db.Model(&models.Banner{}).Where("id = ?", bannerID).Updates(updates)
	return result.Error
}

// DeleteBanner deletes a banner by its ID
func (r *bannerRepository) DeleteBanner(ctx context.Context, bannerID int64) error {
	return r.db.Where("id = ?", bannerID).Delete(&models.Banner{}).Error
}
