//go:generate oapi-codegen --config=oapi-config.yaml api.yaml
package device

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	e          *echo.Echo
	repository Repository
}

func NewHandler(e *echo.Echo, repository Repository) ServerInterface {
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

func (h *Handler) GetDeviceId(ctx echo.Context, id string) error {
	entity, err := h.repository.GetDevice(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapDevice(entity))
}

func mapDevices(devices []deviceEntity) []Device {
	mappedDevices := make([]Device, len(devices))
	for i, device := range devices {
		mappedDevices[i] = mapDevice(device)
	}

	return mappedDevices
}

func mapDevice(device deviceEntity) Device {
	return Device{
		Id:     device.ID,
		Name:   device.Name,
		Type:   device.Type,
		Status: device.Status,
	}
}

func (h *Handler) GetDeviceIdConfigurations(ctx echo.Context, id string) error {
	configurations, err := h.repository.ListConfigurations(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, mapConfigurations(configurations))
}

func mapConfigurations(configurations []configurationEntity) []Configuration {
	mappedConfigurations := make([]Configuration, len(configurations))
	for i, configuration := range configurations {
		mappedConfigurations[i] = Configuration{
			Id:     &configuration.ID,
			Name:   configuration.Name,
			Active: &configuration.Active,
		}
	}

	return mappedConfigurations
}

func (h *Handler) PostDeviceIdConfigurations(ctx echo.Context, id string) error {
	var configuration Configuration
	if err := ctx.Bind(&configuration); err != nil {
		return err
	}

	if err := h.repository.CreateConfiguration(id, configuration.Name, configuration.Data); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, Configuration{
		Id:     configuration.Id,
		Name:   configuration.Name,
		Active: configuration.Active,
		Data:   configuration.Data,
	})
}
