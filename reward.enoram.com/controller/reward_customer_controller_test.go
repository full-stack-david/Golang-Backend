package controller

import (
	"encoding/json"
	"log"
	"mauappa-go/repository/firestore"
	"mauappa-go/service"
	"mauappa-go/util"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type RewardCustomerSuite struct {
	suite.Suite
	RewardOrderKeys []string
}

func TestRewardCustomerSuite(t *testing.T) {
	// This is what actually runs our suite
	util.InitTestConfig()
	suite.Run(t, new(RewardCustomerSuite))
}

func createRewardCustomerController() (*RewardCustomerController, error) {
	rewardCustomerRepo, err := firestore.NewRewardCustomerFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	rewardCustomerService := service.NewRewardCustomerService(rewardCustomerRepo)
	rewardCustomerController := NewRewardCustomerController(rewardCustomerService)

	return rewardCustomerController, nil
}

func checkRewardCustomerResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (suite *RewardCustomerSuite) TestGetRewardCustomerList() {

	dropshipId := "dropshipdll_test_buah"

	rewardCustomerController, err := createRewardCustomerController()
	if err != nil {
		suite.T().Errorf("Unable to create controller %v", err)
	}
	log.Printf("## Using dropshipId %s", dropshipId)
	req, _ := http.NewRequest("GET", "/_api/reward/{dropshipId}/customers", nil)

	req = mux.SetURLVars(req, map[string]string{
		"dropshipId": dropshipId,
	})

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(rewardCustomerController.GetRewardCustomerList)
	handler.ServeHTTP(response, req)

	checkRewardCustomerResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		suite.T().Errorf("Expected the 'status' key of the response to be set to 'true'. Got '%s'", m["status"])
	}
}

func (suite *RewardCustomerSuite) TestGetRewardCustomer() {

	dropshipId := "dropshipdll_test_buah"
	customerId := "142"

	rewardCustomerController, err := createRewardCustomerController()
	if err != nil {
		suite.T().Errorf("Unable to create controller %v", err)
	}
	log.Printf("## Using dropshipId %s, customerId %v", dropshipId, customerId)
	req, _ := http.NewRequest("GET", "/_api/reward/{dropshipId}/customers/{customerId}", nil)

	req = mux.SetURLVars(req, map[string]string{
		"dropshipId": dropshipId,
		"customerId": customerId,
	})

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(rewardCustomerController.GetRewardCustomer)
	handler.ServeHTTP(response, req)

	checkRewardCustomerResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		suite.T().Errorf("Expected the 'status' key of the response to be set to 'true'. Got '%s'", m["status"])
	}
}
