package handlers

import (
	"net/http"
	"strconv"

	"github.com/Esaak/banner-service/internal/banner/usecase"
	"github.com/Esaak/banner-service/internal/models"
	"github.com/Esaak/banner-service/pkg/auth"
	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	useCase usecase.BannerUseCase
	auth    auth.AuthService
}

// NewBannerHandler creates a new instance of BannerHandler
func NewBannerHandler(useCase usecase.BannerUseCase, auth auth.AuthService) *BannerHandler {
	return &BannerHandler{useCase: useCase, auth: auth}
}

// HandleGetUserBanner handles the "/user_banner" GET endpoint
func (h *BannerHandler) HandleGetUserBanner(c *gin.Context) {
	tagID, err := strconv.ParseInt(c.Query("tag_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag_id"})
		return
	}

	featureID, err := strconv.ParseInt(c.Query("feature_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feature_id"})
		return
	}

	useLastRevision, _ := strconv.ParseBool(c.Query("use_last_revision"))

	token := c.GetHeader("token")
	if _, err := h.auth.AuthenticateUser(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	banner, err := h.useCase.GetBanner(c.Request.Context(), tagID, featureID, useLastRevision)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if banner == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}

	c.JSON(http.StatusOK, banner)
}

// HandleGetBanners handles the "/banner" GET endpoint
func (h *BannerHandler) HandleGetBanners(c *gin.Context) {
	featureID, _ := strconv.ParseInt(c.Query("feature_id"), 10, 64)
	tagID, _ := strconv.ParseInt(c.Query("tag_id"), 10, 64)
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	token := c.GetHeader("token")
	if _, err := h.auth.AuthenticateAdmin(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var featureIDPtr, tagIDPtr *int64
	if featureID != 0 {
		featureIDPtr = &featureID
	}
	if tagID != 0 {
		tagIDPtr = &tagID
	}

	banners, err := h.useCase.GetBanners(c.Request.Context(), featureIDPtr, tagIDPtr, &limit, &offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banners)
}

// HandleCreateBanner handles the "/banner" POST endpoint
func (h *BannerHandler) HandleCreateBanner(c *gin.Context) {
	var banner models.Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("token")
	if _, err := h.auth.AuthenticateAdmin(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	bannerID, err := h.useCase.CreateBanner(c.Request.Context(), &banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"banner_id": bannerID})
}

// HandleUpdateBanner handles the "/banner/{id}" PATCH endpoint
func (h *BannerHandler) HandleUpdateBanner(c *gin.Context) {
	bannerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid banner ID"})
		return
	}

	var updates models.Banner
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("token")
	if _, err := h.auth.AuthenticateAdmin(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.useCase.UpdateBanner(c.Request.Context(), bannerID, updates.ToMap()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banner updated successfully"})
}

// HandleDeleteBanner handles the "/banner/{id}" DELETE endpoint
func (h *BannerHandler) HandleDeleteBanner(c *gin.Context) {
	bannerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid banner ID"})
		return
	}

	token := c.GetHeader("token")
	if _, err := h.auth.AuthenticateAdmin(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.useCase.DeleteBanner(c.Request.Context(), bannerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
