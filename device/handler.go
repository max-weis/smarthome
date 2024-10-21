//go:generate oapi-codegen --config=oapi-config.yaml api.yaml
package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/max-weis/smarthome/internal"
)

type Handler struct {
	e          *echo.Echo
	repository Repository
	producer   *Producer
}

func NewHandler(e *echo.Echo, repository Repository, producer *Producer) ServerInterface {
	handler := &Handler{
		e:          e,
		repository: repository,
	}

	RegisterHandlers(e, handler)

	return handler
}

func (h *Handler) GetDevices(ctx echo.Context) error {
	entities, err := h.repository.GetDevices()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapDevices(entities))
}

func (h *Handler) GetDevice(ctx echo.Context, id string) error {
	entity, err := h.repository.GetDevice(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapDevice(entity))
}

func mapDevices(devices []DeviceEntity) []Device {
	mappedDevices := make([]Device, len(devices))
	for i, device := range devices {
		mappedDevices[i] = mapDevice(device)
	}

	return mappedDevices
}

func mapDevice(device DeviceEntity) Device {
	return Device{
		Id:     device.ID,
		Name:   device.Name,
		Type:   device.Type,
		Status: device.Status,
	}
}

func (h *Handler) GetDeviceConfigurations(ctx echo.Context, id string) error {
	if _, err := h.findDevice(ctx, id); err != nil {
		return err
	}

	configurations, err := h.repository.ListConfigurations(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapConfigurations(configurations))
}

func mapConfigurations(configurations []configurationEntity) []ConfigurationListItem {
	mappedConfigurations := make([]ConfigurationListItem, len(configurations))
	for i, configuration := range configurations {
		mappedConfigurations[i] = ConfigurationListItem{
			Id:     configuration.ID,
			Name:   configuration.Name,
			Active: configuration.Active,
		}
	}

	return mappedConfigurations
}

func (h *Handler) CreateConfiguration(ctx echo.Context, id string) error {
	if _, err := h.findDevice(ctx, id); err != nil {
		return err
	}

	var configuration Configuration
	if err := ctx.Bind(&configuration); err != nil {
		return err
	}

	data, err := json.Marshal(configuration.Data)
	if err != nil {
		return err
	}

	newConfiguration, err := h.repository.CreateConfiguration(id, configuration.Name, data)
	if err != nil {
		return err
	}

	config, err := mapConfiguration(newConfiguration)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, config)
}

func (h *Handler) GetConfiguration(ctx echo.Context, id string, configurationId string) error {
	if _, err := h.findDevice(ctx, id); err != nil {
		return err
	}

	config, err := h.repository.GetConfiguration(configurationId)
	if err != nil {
		return err
	}

	configuration, err := mapConfiguration(config)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, configuration)
}

func mapConfiguration(configuration configurationEntity) (Configuration, error) {
	var data map[string]any
	if err := json.Unmarshal(configuration.Data, &data); err != nil {
		return Configuration{}, err
	}

	return Configuration{
		Id:     configuration.ID,
		Name:   configuration.Name,
		Active: configuration.Active,
		Data:   &data,
	}, nil
}

func (h *Handler) UpdateConfiguration(ctx echo.Context, id string, configurationId string) error {
	if _, err := h.findDevice(ctx, id); err != nil {
		return err
	}

	var newConfiguration Configuration
	if err := ctx.Bind(&newConfiguration); err != nil {
		return err
	}

	newDeviceEntity, err := mapToEntity(id, configurationId, newConfiguration)
	if err != nil {
		return err
	}

	if newDeviceEntity.Active {
		// check if device is active and set all configurations inactive,
		// but only if the new configuration is active. So the device state can be
		// updated to only one configuration at a time.
		if err := h.repository.SetAllInactive(id); err != nil {
			return err
		}
	}

	config, err := h.repository.UpdateConfiguration(newDeviceEntity)
	if err != nil {
		return err
	}

	cfg, err := mapConfiguration(config)
	if err != nil {
		return err
	}

	if config.Active {
		if err := h.producer.PublishConfiguration(id, config.ID, *cfg.Data); err != nil {
			return err
		}
	}

	return ctx.JSON(http.StatusOK, cfg)
}

func mapToEntity(deviceId, configId string, config Configuration) (configurationEntity, error) {
	data, err := json.Marshal(config.Data)
	if err != nil {
		return configurationEntity{}, err
	}

	return configurationEntity{
		ID:       configId,
		DeviceID: deviceId,
		Name:     config.Name,
		Active:   config.Active,
		Data:     data,
	}, nil
}

func (h *Handler) findDevice(ctx echo.Context, id string) (*DeviceEntity, error) {
	device, err := h.repository.GetDevice(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %w", internal.ErrNotFound, err)
		}

		return nil, fmt.Errorf("%w: %w", internal.ErrInternalServerError, err)
	}

	return &device, nil
}
