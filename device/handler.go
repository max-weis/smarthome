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

func mapDevices(devices []deviceEntity) []Device {
	mappedDevices := make([]Device, len(devices))
	for i, device := range devices {
		mappedDevices[i] = Device{
			Id:     device.ID,
			Name:   device.Name,
			Type:   device.Type,
			Status: device.Status,
		}
	}
	return mappedDevices
}
