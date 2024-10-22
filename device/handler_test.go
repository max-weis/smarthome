package device

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

	_ = NewHandler(suite.echo, suite.repository, nil)
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
