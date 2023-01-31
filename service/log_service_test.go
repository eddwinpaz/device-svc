package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eddwinpaz/device-svc/domain"
	"github.com/eddwinpaz/device-svc/repository"
	"github.com/eddwinpaz/device-svc/service"
)

func TestServiceGetDeviceById(t *testing.T) {

	deviceId := "5luNMc"

	mockDevice := []domain.Log{{
		DeviceId:    "5luNMc",
		LogDate:     "2022-03-28 15:56:02.568 -0300",
		Temperature: 10,
		Alert:       false,
	}}

	mockCase := new(repository.MockRepositoryInterface)
	mockCase.On("GetDeviceByID", deviceId).Return(&mockDevice, nil)

	svc := service.NewLogService(mockCase)

	// Execute
	logs, err := svc.GetDeviceByID(deviceId)

	assert.NoError(t, err)
	assert.NotNil(t, logs)
}
