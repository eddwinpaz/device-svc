package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/eddwinpaz/device-svc/domain"
)

type RepositoryInterface interface {
	Ingest(logs []domain.Log) error
	GetDeviceByID(deviceID string) ([]*domain.Log, error)
}

type PostgresRepository struct {
	database *sql.DB
}

func NewPostgresClient() (*sql.DB, error) {
	psqlconn := "host=localhost port=5432 user=pguser password=pgpassword dbname=code_challenge sslmode=disable"

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgresRepository() (RepositoryInterface, error) {
	postgresClient, err := NewPostgresClient()
	if err != nil {
		panic(err.Error())
	}
	return &PostgresRepository{
		database: postgresClient,
	}, nil
}

func (r *PostgresRepository) Ingest(logs []domain.Log) error {

	tx, err := r.database.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, log := range logs {
		_, err := tx.Exec("INSERT INTO logs (event_date, device_id, temp_farenheit) VALUES ($1, $2, $3)", log.LogDate, log.DeviceId, log.Temperature)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *PostgresRepository) GetDeviceByID(deviceID string) ([]*domain.Log, error) {

	logs := []*domain.Log{}

	sql := fmt.Sprintf(`SELECT event_date , device_id , temp_farenheit FROM logs WHERE device_id = '%s' ORDER BY event_date ASC`, deviceID)

	rows, err := r.database.Query(sql)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var eventDate string
		var deviceID string
		var tempFarenheit int

		if err := rows.Scan(&eventDate, &deviceID, &tempFarenheit); err != nil {
			return nil, err
		}

		newLog := domain.Log{
			LogDate:     eventDate,
			DeviceId:    deviceID,
			Temperature: tempFarenheit,
		}

		newLog.AddFlag()
		newLog.ConvertTemperatureToFarenheit()

		logs = append(logs, &newLog)
	}

	return logs, nil
}
