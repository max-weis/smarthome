package device

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db *gorm.DB
	}

	deviceEntity struct {
		ID     string
		Name   string
		Type   string
		Status string
	}
)

func (deviceEntity) TableName() string {
	return "devices"
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) GetDevices() ([]deviceEntity, error) {
	var devices []deviceEntity
	if err := r.db.Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *Repository) GetDevice(id string) (deviceEntity, error) {
	var device deviceEntity
	if err := r.db.First(&device, "id = ?", id).Error; err != nil {
		return deviceEntity{}, err
	}

	return device, nil
}

type configurationEntity struct {
	ID       string
	DeviceID string
	Name     string
	Active   bool
	Data     datatypes.JSON
}

func (configurationEntity) TableName() string {
	return "configurations"
}

func (r *Repository) ListConfigurations(id string) ([]configurationEntity, error) {
	var configurations []configurationEntity
	if err := r.db.Find(&configurations, "device_id = ?", id).Error; err != nil {
		return nil, err
	}

	return configurations, nil
}

func (r *Repository) CreateConfiguration(deviceId string, name string, data json.RawMessage) error {
	return r.db.Create(&configurationEntity{
		ID:       uuid.New().String(),
		DeviceID: deviceId,
		Name:     name,
		Data:     datatypes.JSON(data),
	}).Error
}
