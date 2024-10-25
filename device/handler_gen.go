// Package device provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package device

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Configuration defines model for Configuration.
type Configuration struct {
	// Active Whether the configuration is active or not
	Active bool `json:"active"`

	// Data Configuration data
	Data *map[string]interface{} `json:"data,omitempty"`

	// Id Unique identifier for the configuration
	Id string `json:"id"`

	// Name Human-readable name of the configuration
	Name string `json:"name"`
}

// ConfigurationListItem defines model for ConfigurationListItem.
type ConfigurationListItem struct {
	// Active Whether the configuration is active or not
	Active bool `json:"active"`

	// Id Unique identifier for the configuration
	Id string `json:"id"`

	// Name Human-readable name of the configuration
	Name string `json:"name"`
}

// Device defines model for Device.
type Device struct {
	// Id Unique identifier for the device
	Id string `json:"id"`

	// Name Human-readable name of the device
	Name string `json:"name"`

	// Status Current status of the device (e.g., on, off, idle)
	Status string `json:"status"`

	// Type Type of the smart device (e.g., light, thermostat)
	Type string `json:"type"`
}

// UpdateConfigurationJSONRequestBody defines body for UpdateConfiguration for application/json ContentType.
type UpdateConfigurationJSONRequestBody = Configuration

// CreateConfigurationJSONRequestBody defines body for CreateConfiguration for application/json ContentType.
type CreateConfigurationJSONRequestBody = Configuration

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get device
	// (GET /device/{id})
	GetDevice(ctx echo.Context, id string) error
	// Get device configuration
	// (GET /device/{id}/configuration/{configurationId})
	GetConfiguration(ctx echo.Context, id string, configurationId string) error
	// Update device configuration
	// (PUT /device/{id}/configuration/{configurationId})
	UpdateConfiguration(ctx echo.Context, id string, configurationId string) error
	// Toggle device configuration status
	// (POST /device/{id}/configuration/{configurationId}/status)
	ToggleConfigurationStatus(ctx echo.Context, id string, configurationId string) error
	// Get device configurations
	// (GET /device/{id}/configurations)
	GetDeviceConfigurations(ctx echo.Context, id string) error
	// Create device configuration
	// (POST /device/{id}/configurations)
	CreateConfiguration(ctx echo.Context, id string) error
	// List all devices
	// (GET /devices)
	GetDevices(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDevice converts echo context to params.
func (w *ServerInterfaceWrapper) GetDevice(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDevice(ctx, id)
	return err
}

// GetConfiguration converts echo context to params.
func (w *ServerInterfaceWrapper) GetConfiguration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "configurationId" -------------
	var configurationId string

	err = runtime.BindStyledParameterWithOptions("simple", "configurationId", ctx.Param("configurationId"), &configurationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter configurationId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetConfiguration(ctx, id, configurationId)
	return err
}

// UpdateConfiguration converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateConfiguration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "configurationId" -------------
	var configurationId string

	err = runtime.BindStyledParameterWithOptions("simple", "configurationId", ctx.Param("configurationId"), &configurationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter configurationId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateConfiguration(ctx, id, configurationId)
	return err
}

// ToggleConfigurationStatus converts echo context to params.
func (w *ServerInterfaceWrapper) ToggleConfigurationStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "configurationId" -------------
	var configurationId string

	err = runtime.BindStyledParameterWithOptions("simple", "configurationId", ctx.Param("configurationId"), &configurationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter configurationId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ToggleConfigurationStatus(ctx, id, configurationId)
	return err
}

// GetDeviceConfigurations converts echo context to params.
func (w *ServerInterfaceWrapper) GetDeviceConfigurations(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDeviceConfigurations(ctx, id)
	return err
}

// CreateConfiguration converts echo context to params.
func (w *ServerInterfaceWrapper) CreateConfiguration(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateConfiguration(ctx, id)
	return err
}

// GetDevices converts echo context to params.
func (w *ServerInterfaceWrapper) GetDevices(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDevices(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/device/:id", wrapper.GetDevice)
	router.GET(baseURL+"/device/:id/configuration/:configurationId", wrapper.GetConfiguration)
	router.PUT(baseURL+"/device/:id/configuration/:configurationId", wrapper.UpdateConfiguration)
	router.POST(baseURL+"/device/:id/configuration/:configurationId/status", wrapper.ToggleConfigurationStatus)
	router.GET(baseURL+"/device/:id/configurations", wrapper.GetDeviceConfigurations)
	router.POST(baseURL+"/device/:id/configurations", wrapper.CreateConfiguration)
	router.GET(baseURL+"/devices", wrapper.GetDevices)

}
