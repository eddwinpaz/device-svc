package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/eddwinpaz/device-svc/domain"
)

type MockRepositoryInterface struct {
	mock.Mock
}

func (m *MockRepositoryInterface) GetDeviceByID(deviceId string) ([]*domain.Log, error) {

	logs := []*domain.Log{}

	args := m.Called(deviceId)

	if rf, ok := args.Get(0).(func(string) []*domain.Log); ok {
		logs = rf(deviceId)
	} else {
		if args.Get(0) != nil {
			logs = append(logs, &domain.Log{
				DeviceId:    "5luNMc",
				LogDate:     "2022-03-28 15:56:02.568 -0300",
				Temperature: 10,
				Alert:       false,
			})
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(deviceId)
	} else {
		r1 = args.Error(1)
	}

	return logs, r1
}

func (m *MockRepositoryInterface) Ingest(logs []domain.Log) error {
	return nil
}
