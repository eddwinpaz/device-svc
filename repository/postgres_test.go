package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/eddwinpaz/device-svc/repository"
)

func Test_repository_GetDeviceByID_valid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close() //nolint

	rows := sqlmock.NewRows([]string{"event_id", "event_date", "device_id", "temp_farenheit"}).
		AddRow("1", "2022-03-28 15:56:02.568 -0300", "5luNMc", 20)

	mock.ExpectQuery("1").WillReturnRows(rows)
	h, _ := repository.NewPostgresRepository()

	device_id := "5luNMc"

	log, err := h.GetDeviceByID(device_id)
	assert.NoError(t, err)
	assert.NotNil(t, log)
}
