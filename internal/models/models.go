package models

import (
	"encoding/json"
	"time"
)

// Banner represents a banner entity
type Banner struct {
	ID        int64     `json:"banner_id"`
	TagIDs    []int64   `json:"tag_ids"`
	FeatureID int64     `json:"feature_id"`
	Content   JSONData  `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JSONData is a helper type for marshaling/unmarshaling JSON data
type JSONData map[string]interface{}

// MarshalBinary implements the encoding.BinaryMarshaler interface
func (j JSONData) MarshalBinary() ([]byte, error) {
	return json.Marshal(j)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface
func (j *JSONData) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &j)
}

// ToMap converts a Banner struct to a map[string]interface{}
func (b *Banner) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["tag_ids"] = b.TagIDs
	data["feature_id"] = b.FeatureID
	data["content"] = b.Content
	data["is_active"] = b.IsActive
	return data
}

// BannerTag представляет связь между баннером и тегом
type BannerTag struct {
	BannerID int64 `gorm:"primaryKey"`
	TagID    int64 `gorm:"primaryKey"`
}

// TableName определяет имя таблицы для BannerTag
func (*BannerTag) TableName() string {
	return "banner_tags"
}
