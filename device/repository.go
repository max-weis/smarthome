package device

import "gorm.io/gorm"

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
