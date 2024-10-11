package device_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/max-weis/smarthome/device"
	"github.com/max-weis/smarthome/internal"
)

type HandlerTestSuite struct {
	suite.Suite
	e          *echo.Echo
	db         *gorm.DB
	repository device.Repository
	handler    device.ServerInterface
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupTest() {
	e := echo.New()
	db, err := internal.NewDatabase()
	suite.Assert().NoError(err)

	repository := device.NewRepository(db)
	h := device.NewHandler(e, repository)

	suite.e = e
	suite.db = db
	suite.repository = repository
	suite.handler = h
}

func (suite *HandlerTestSuite) TestUpdateConfiguration() {
	dev, err := suite.repository.CreateDevice(device.DeviceEntity{
		Name:   "Test Device",
		Type:   "Test Type",
		Status: "Test Status",
	})
	suite.Assert().NoError(err)

	config, err := suite.repository.CreateConfiguration(dev.ID, "Test Configuration", []byte(`{"test": "data"}`))
	suite.Assert().NoError(err)

	config.Name = "Updated Configuration"

	updateConfigJSON, err := json.Marshal(config)
	suite.Assert().NoError(err)

	req := httptest.NewRequest(http.MethodPost, updateConfigPath(dev.ID, config.ID), bytes.NewReader(updateConfigJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := suite.e.NewContext(req, rec)

	err = suite.handler.UpdateConfiguration(c, dev.ID, config.ID)
	fmt.Println(err)

	if suite.Assert().NoError(err) {
		suite.Assert().Equal(http.StatusCreated, rec.Code)
		suite.Assert().Equal("", rec.Body.String())
	}
}

func updateConfigPath(deviceId, configId string) string {
	return fmt.Sprintf("/device/%s/configuration/%s", deviceId, configId)
}
