package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/eddwinpaz/device-svc/repository"
)

func TestHttpGetDeviceById(t *testing.T) {

	path := fmt.Sprintf("/device/%s", "5luNMc")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCase := new(repository.MockRepositoryInterface)
	mockCase.On("GetDeviceByID", "5luNMc").Return(MockResponseLog, nil)

	svc, _ := repository.NewPostgresRepository()
	handler := NewHandler(svc)

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/device/{id}", handler.Get).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "{\"deviceId\":\"5luNMc\",\"averageTemperature\":37,\"mostRecentLogDate\":\"2022-03-28T18:56:28.679Z\",\"totalAlerts\":11,\"logs\":[{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:17.357Z\",\"temperature\":10},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:19.369Z\",\"temperature\":30},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:23.387Z\",\"temperature\":28},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:29.406Z\",\"temperature\":58},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:35.433Z\",\"temperature\":12},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:43.469Z\",\"temperature\":64},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:44.475Z\",\"temperature\":16},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:48.493Z\",\"temperature\":44},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:49.496Z\",\"temperature\":14},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:50.506Z\",\"temperature\":44},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:54.521Z\",\"temperature\":60},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:57.534Z\",\"temperature\":58},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:58.539Z\",\"temperature\":30},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:55:59.559Z\",\"temperature\":36},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:02.568Z\",\"temperature\":66},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:12.615Z\",\"temperature\":50},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:23.656Z\",\"temperature\":40},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:24.661Z\",\"temperature\":56},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:27.678Z\",\"temperature\":18},{\"deviceId\":\"5luNMc\",\"logDate\":\"2022-03-28T18:56:28.679Z\",\"temperature\":18}]}\n"
	assert.Equal(t, rr.Body.String(), expected)
}
