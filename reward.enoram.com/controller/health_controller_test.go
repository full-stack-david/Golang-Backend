package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HealthSuite struct {
	suite.Suite
}

func TestHealthSuite(t *testing.T) {
	suite.Run(t, new(HealthSuite))
}

func (suite *HealthSuite) TestHealthCheck() {

	healthController := NewHealthController()
	req, _ := http.NewRequest("GET", "/_api/reward/health", nil)

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(healthController.HealthCheck)
	handler.ServeHTTP(response, req)

	checkResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		suite.T().Errorf("Expected the 'status' id of the response to be set to 'true'. Got '%s'", m["status"])
	}

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
