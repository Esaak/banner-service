package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Esaak/banner-service/internal/banner/repository/postgres/repository_mocks"
	"github.com/Esaak/banner-service/internal/models"
	"github.com/Esaak/banner-service/pkg/auth/auth_mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGetUserBanner(t *testing.T) {
	// Arrange
	tagID := int64(1)
	featureID := int64(2)
	useLastRevision := false

	expectedBanner := &models.Banner{
		ID:        1,
		TagIDs:    []int64{1},
		FeatureID: 2,
		Content:   models.JSONData{"title": "Test Banner", "text": "This is a test banner."},
		IsActive:  true,
	}

	mockRepo := new(repository_mocks.BannerRepositoryMock)
	mockRepo.On("GetBanner", tagID, featureID, useLastRevision).Return(expectedBanner, nil)

	mockAuth := new(auth_mocks.AuthServiceMock)
	mockAuth.On("AuthenticateUser", "valid-token").Return(int64(1), nil)

	router := gin.Default()
	handler := NewBannerHandler(mockRepo, mockAuth)
	router.GET("/user_banner", handler.HandleGetUserBanner)

	// Act
	req, err := http.NewRequest("GET", "/user_banner?tag_id=1&feature_id=2", nil)
	require.NoError(t, err)
	req.Header.Set("token", "valid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)

	var actualBanner models.Banner
	err = json.Unmarshal(w.Body.Bytes(), &actualBanner)
	require.NoError(t, err)

	assert.Equal(t, expectedBanner, &actualBanner)

	mockRepo.AssertExpectations(t)
	mockAuth.AssertExpectations(t)
}
