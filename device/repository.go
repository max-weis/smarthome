package device

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Repository struct {
		db *gorm.DB
	}

	DeviceEntity struct {
		ID     string
		Name   string
		Type   string
		Status string
	}
)

func (DeviceEntity) TableName() string {
	return "devices"
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) GetDevices() ([]DeviceEntity, error) {
	var devices []DeviceEntity
	if err := r.db.Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *Repository) GetDevice(id string) (DeviceEntity, error) {
	var device DeviceEntity
	if err := r.db.First(&device, "id = ?", id).Error; err != nil {
		return DeviceEntity{}, err
	}

	return device, nil
}

func (r *Repository) CreateDevice(device DeviceEntity) (DeviceEntity, error) {
	if err := r.db.Create(device).Error; err != nil {
		return DeviceEntity{}, err
	}

	device, err := r.GetDevice(device.ID)
	if err != nil {
		return DeviceEntity{}, err
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

func (r *Repository) CreateConfiguration(deviceId string, name string, data json.RawMessage) (configurationEntity, error) {
	config := &configurationEntity{
		ID:       uuid.New().String(),
		DeviceID: deviceId,
		Name:     name,
		Active:   false,
		Data:     datatypes.JSON(data),
	}

	if err := r.db.Create(config).Error; err != nil {
		return configurationEntity{}, err
	}

	configuration, err := r.GetConfiguration(config.ID)
	if err != nil {
		return configurationEntity{}, err
	}

	return configuration, nil
}

func (r *Repository) GetConfiguration(id string) (configurationEntity, error) {
	var configuration configurationEntity
	if err := r.db.First(&configuration, "id = ?", id).Error; err != nil {
		return configurationEntity{}, err
	}

	return configuration, nil
}

func (r *Repository) UpdateConfiguration(config configurationEntity) (configurationEntity, error) {
	onConflict := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"Name", "Active", "Data"}),
	}

	if err := r.db.Clauses(onConflict).Create(&config).Error; err != nil {
		return configurationEntity{}, err
	}

	return config, nil
}

func (r *Repository) SetAllInactive(deviceId string) error {
	return r.db.Model(&configurationEntity{}).
		Where("device_id = ?", deviceId).
		Update("active", false).Error
}

func (r *Repository) ToggleConfigurationStatus(configId string) (bool, error) {
    var config configurationEntity
    
    // First update the status
    if err := r.db.Model(&configurationEntity{}).
        Where("id = ?", configId).
        Update("active", gorm.Expr("NOT active")).Error; err != nil {
        return false, err
    }
    
    // Then fetch the updated record to get the new status
    if err := r.db.Where("id = ?", configId).First(&config).Error; err != nil {
        return false, err
    }
    
    return config.Active, nil
}