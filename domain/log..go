package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	DeviceId    string `json:"deviceId,omitempty"`
	LogDate     string `json:"logDate"`
	Temperature int    `json:"temperature"`
	Humidity    string `json:"humidity,omitempty"`
	Alert       bool   `json:"alert,omitempty"`
}

type LogRequest struct {
	Data []string `json:"data,omitempty"`
}

type DeviceResponse struct {
	DeviceId           string `json:"deviceId"`
	AverageTemperature int    `json:"averageTemperature"`
	MostRecentLogDate  string `json:"mostRecentLogDate"`
	TotalAlerts        int    `json:"totalAlerts,omitempty"`
	Logs               []*Log `json:"logs"`
}

func (l *Log) ConvertTemperatureToFarenheit() {
	l.Temperature = (2 * l.Temperature) + 30
}

func (l *Log) AddFlag() {
	if l.Temperature > 32 {
		l.Alert = true
	}
}

func (p *LogRequest) Parse() []Log {

	newLogs := []Log{}

	for _, log := range p.Data {

		splitLog := strings.Split(log, "|")

		temperature, err := strconv.Atoi(splitLog[2])

		if err != nil {
			fmt.Println(err)
		}

		newLog := Log{
			DeviceId:    splitLog[0],
			LogDate:     splitLog[1],
			Temperature: temperature,
			Humidity:    splitLog[2],
			Alert:       false,
		}

		newLogs = append(newLogs, newLog)
	}

	return newLogs
}

func averageTemperature(logs []*Log) int {
	total := 0
	for _, log := range logs {
		total += log.Temperature
	}
	return total / len(logs)
}

func totalAlerts(logs []*Log) int {
	total := 0
	for _, log := range logs {
		if log.Temperature > 32 {
			total += 1
		}
	}
	return total
}

func LogResponse(logs []*Log, deviceId string) *DeviceResponse {

	if len(logs) == 0 {
		return &DeviceResponse{}
	}

	lastElement := logs[len(logs)-1]

	return &DeviceResponse{
		DeviceId:           deviceId,
		AverageTemperature: averageTemperature(logs),
		MostRecentLogDate:  lastElement.LogDate,
		TotalAlerts:        totalAlerts(logs),
		Logs:               logs,
	}
}
