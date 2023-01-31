package api

import "github.com/eddwinpaz/device-svc/domain"

var MockResponseLog = &domain.DeviceResponse{
	DeviceId:           "5luNMc",
	AverageTemperature: 1,
	MostRecentLogDate:  "2019-01-01T00:00:00Z",
	Logs: []*domain.Log{{
		DeviceId:    "5luNMc",
		LogDate:     "2019-01-01T00:00:00Z",
		Temperature: -5},
	},
}
