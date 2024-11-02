package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/max-weis/smarthome/internal"
)

type DeviceTestSuite struct {
	suite.Suite
	echo             *echo.Echo
	responseRecorder *httptest.ResponseRecorder

	db         *gorm.DB
	repository Repository
}

func (suite *DeviceTestSuite) SetupTest() {
	suite.echo = internal.NewEchoServer()
	db, err := internal.NewDatabase()
	suite.Assert().NoError(err)

	suite.db = db
	suite.repository = NewRepository(db)

	producer := &mockProducer{}
	producer.On("PublishConfiguration", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	_ = NewHandler(suite.echo, suite.repository, producer)
}

func (suite *DeviceTestSuite) TestGetDevices() {
	suite.newRequest("GET", "/devices", nil)
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var devices []Device
	suite.mapFromResp(&devices)

	suite.Require().Equal([]Device{
		{
			Id:     "1",
			Name:   "Living Room Light",
			Type:   "light",
			Status: "on",
		},
		{
			Id:     "2",
			Name:   "Bedroom Light",
			Type:   "light",
			Status: "off",
		},
		{
			Id:     "3",
			Name:   "Kitchen Light",
			Type:   "light",
			Status: "idle",
		},
		{
			Id:     "4",
			Name:   "Bedroom Thermostat",
			Type:   "thermostat",
			Status: "idle",
		},
		{
			Id:     "5",
			Name:   "Kitchen Thermostat",
			Type:   "thermostat",
			Status: "idle",
		},
	}, devices)
}

func (suite *DeviceTestSuite) TestGetDevice() {
	suite.newRequest("GET", "/device/1", nil)
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var device Device
	suite.mapFromResp(&device)

	suite.Require().Equal(Device{
		Id:     "1",
		Name:   "Living Room Light",
		Type:   "light",
		Status: "on",
	}, device)
}

func (suite *DeviceTestSuite) TestGetDeviceConfigurations() {
	suite.newRequest("GET", "/device/1/configurations", nil)
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var configurations []ConfigurationListItem
	suite.mapFromResp(&configurations)

	suite.Require().Equal([]ConfigurationListItem{
		{
			Id:     "1",
			Name:   "Daymode",
			Active: true,
		},
		{
			Id:   "2",
			Name: "Nightmode",
		},
	}, configurations)
}

func (suite *DeviceTestSuite) TestCreateConfiguration() {
	configuration := Configuration{
		Name: "Test Configuration",
		Data: &map[string]any{"test": true},
	}
	suite.newRequest("POST", "/device/1/configurations", mapToReq(configuration))
	suite.Require().Equal(http.StatusCreated, suite.responseRecorder.Code)

	var newConfiguration Configuration
	suite.mapFromResp(&newConfiguration)

	suite.Require().NotEmpty(newConfiguration.Id)
	suite.Require().Equal(configuration.Name, newConfiguration.Name)
	suite.Require().Equal(configuration.Data, newConfiguration.Data)
}

func (suite *DeviceTestSuite) TestGetConfiguration() {
	suite.newRequest("GET", "/device/1/configuration/1", nil)
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var configuration Configuration
	suite.mapFromResp(&configuration)

	suite.Require().Equal("1", configuration.Id)
	suite.Require().Equal("Daymode", configuration.Name)
	suite.Require().True(configuration.Active)
	suite.Require().Equal(map[string]any{"brightness": float64(100)}, *configuration.Data)
}

func (suite *DeviceTestSuite) TestUpdateConfiguration() {
	configuration := Configuration{
		Name: "Test Configuration",
		Data: &map[string]any{"test": true},
	}
	suite.newRequest("PUT", "/device/1/configuration/1", mapToReq(configuration))
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var updatedConfiguration Configuration
	suite.mapFromResp(&updatedConfiguration)

	suite.Require().Equal(configuration.Name, updatedConfiguration.Name)
	suite.Require().Equal(configuration.Data, updatedConfiguration.Data)
}

func (suite *DeviceTestSuite) TestToggleConfigurationStatus() {
	device, err := suite.repository.CreateDevice(DeviceEntity{
		ID:     uuid.NewString(),
		Name:   "Test Device",
		Type:   "light",
		Status: "on",
	})
	suite.Require().NoError(err)

	configuration, err := suite.repository.CreateConfiguration(device.ID, "Test Config", []byte("{}"))
	suite.Require().NoError(err)

	// check that the active flag is set to false. This will change after the toggle
	suite.Require().False(configuration.Active)

	route := fmt.Sprintf("/device/%s/configuration/%s/status", device.ID, configuration.ID)
	suite.newRequest("POST", route, nil)
	suite.Require().Equal(http.StatusNoContent, suite.responseRecorder.Code)

	updatedConfig, err := suite.repository.GetConfiguration(configuration.ID)
	suite.Require().NoError(err)

	suite.Require().True(updatedConfig.Active)
}

func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceTestSuite))
}

func (suite *DeviceTestSuite) newRequest(method string, route string, requestBody []byte) {
	request := httptest.NewRequest(method, route, bytes.NewReader(requestBody))

	request.Header.Set("content-type", "application/json; charset=utf-8")
	suite.responseRecorder = httptest.NewRecorder()
	suite.echo.ServeHTTP(suite.responseRecorder, request)
}

func mapToReq(v any) []byte {
	jsonObject, err := json.Marshal(v)
	if err != nil {
		return []byte("")
	}

	return append(jsonObject, 10)
}

func (suite *DeviceTestSuite) mapFromResp(v any) {
	if err := json.Unmarshal(suite.responseRecorder.Body.Bytes(), v); err != nil {
		suite.T().Log(err)
		suite.T().Fatal()
	}
}

type mockProducer struct {
	mock.Mock
}

func (m *mockProducer) PublishConfiguration(deviceId string, configId string, configData map[string]any) error {
	args := m.Called(deviceId, configId, configData)
	return args.Error(0)
}
