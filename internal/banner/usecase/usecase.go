package usecase

import (
	"context"

	"github.com/Esaak/banner-service/internal/banner/repository/postgres"
	"github.com/Esaak/banner-service/internal/models"
)

// BannerUseCase представляет логику работы с баннерами
type BannerUseCase interface {
	GetBanner(ctx context.Context, tagID, featureID int64, useLastRevision bool) (*models.Banner, error)
	GetBanners(ctx context.Context, featureID, tagID *int64, limit, offset *int) ([]models.Banner, error)
	CreateBanner(ctx context.Context, banner *models.Banner) (int64, error)
	UpdateBanner(ctx context.Context, bannerID int64, updates map[string]interface{}) error
	DeleteBanner(ctx context.Context, bannerID int64) error
}

type bannerUseCase struct {
	repo postgres.BannerRepository
}

// NewBannerUseCase создает новый экземпляр BannerUseCase
func NewBannerUseCase(repo postgres.BannerRepository) BannerUseCase {
	return &bannerUseCase{repo: repo}
}

// GetBanner получает баннер для указанного тега и функции
func (uc *bannerUseCase) GetBanner(ctx context.Context, tagID, featureID int64, useLastRevision bool) (*models.Banner, error) {
	return uc.repo.GetBanner(ctx, tagID, featureID, useLastRevision)
}

// GetBanners получает списки баннеров с фильтрацией по функции и/или тегу
func (uc *bannerUseCase) GetBanners(ctx context.Context, featureID, tagID *int64, limit, offset *int) ([]models.Banner, error) {
	return uc.repo.GetBanners(ctx, featureID, tagID, limit, offset)
}

// CreateBanner создает новый баннер
func (uc *bannerUseCase) CreateBanner(ctx context.Context, banner *models.Banner) (int64, error) {
	return uc.repo.CreateBanner(ctx, banner)
}

// UpdateBanner обновляет существующий баннер
func (uc *bannerUseCase) UpdateBanner(ctx context.Context, bannerID int64, updates map[string]interface{}) error {
	return uc.repo.UpdateBanner(ctx, bannerID, updates)
}

// DeleteBanner удаляет баннер по его ID
func (uc *bannerUseCase) DeleteBanner(ctx context.Context, bannerID int64) error {
	return uc.repo.DeleteBanner(ctx, bannerID)
}
