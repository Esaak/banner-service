package repository_mocks

import (
	"context"

	"github.com/Esaak/banner-service/internal/models"
	"github.com/stretchr/testify/mock"
)

type BannerRepositoryMock struct {
	mock.Mock
}

func (m *BannerRepositoryMock) GetBanner(ctx context.Context, tagID, featureID int64, useLastRevision bool) (*models.Banner, error) {
	args := m.Called(ctx, tagID, featureID, useLastRevision)
	return args.Get(0).(*models.Banner), args.Error(1)
}

func (m *BannerRepositoryMock) GetBanners(ctx context.Context, featureID, tagID *int64, limit, offset *int) ([]models.Banner, error) {
	args := m.Called(ctx, featureID, tagID, limit, offset)
	return args.Get(0).([]models.Banner), args.Error(1)
}

func (m *BannerRepositoryMock) CreateBanner(ctx context.Context, banner *models.Banner) (int64, error) {
	args := m.Called(ctx, banner)
	return args.Get(0).(int64), args.Error(1)
}

func (m *BannerRepositoryMock) UpdateBanner(ctx context.Context, bannerID int64, updates map[string]interface{}) error {
	args := m.Called(ctx, bannerID, updates)
	return args.Error(0)
}

func (m *BannerRepositoryMock) DeleteBanner(ctx context.Context, bannerID int64) error {
	args := m.Called(ctx, bannerID)
	return args.Error(0)
}
