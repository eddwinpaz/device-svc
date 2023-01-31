package service

import (
	"github.com/eddwinpaz/device-svc/domain"
	"github.com/eddwinpaz/device-svc/repository"
)

type LogService struct {
	database repository.RepositoryInterface
}

func NewLogService(database repository.RepositoryInterface) LogServiceInterface {
	return &LogService{
		database: database,
	}
}

func (l *LogService) GetDeviceByID(deviceID string) ([]*domain.Log, error) {
	logs, err := l.database.GetDeviceByID(deviceID)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (l *LogService) Ingest(logs []domain.Log) error {
	err := l.database.Ingest(logs)
	if err != nil {
		return err
	}
	return nil
}
