package service

import (
	"github.com/eddwinpaz/device-svc/domain"
)

type LogServiceInterface interface {
	Ingest(logs []domain.Log) error
	GetDeviceByID(deviceID string) ([]*domain.Log, error)
}
