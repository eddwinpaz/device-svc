package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogDomain(t *testing.T) {

	log := Log{
		LogDate:     "2020-01-01 00:00:00",
		DeviceId:    "0000",
		Temperature: 9,
	}
	assert.Equal(t, log.DeviceId, "0000")
	assert.Equal(t, log.Temperature, 9)
	// converted temperature
	log.ConvertTemperatureToFarenheit()
	assert.Equal(t, log.Temperature, 48)
	// converted temperature

}
