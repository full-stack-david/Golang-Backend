package controller

import (
	"encoding/json"
	"log"
	"mauappa-go/repository/firestore"
	"mauappa-go/service"
	"mauappa-go/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RewardSuite struct {
	suite.Suite
	RewardKeys []string
}

func TestRewardSuite(t *testing.T) {
	// This is what actually runs our suite
	util.InitTestConfig()
	suite.Run(t, new(RewardSuite))
}

func createRewardController() (*RewardController, error) {

	rewardOrderRepo, err := firestore.NewRewardOrderFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	invoiceRepo, err := firestore.NewInvoiceFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firebase storage repository", err)
	}

	storeProfileRepo, err := firestore.NewStoreProfileFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firebase storage repository", err)
	}

	customerRepo, err := firestore.NewCustomerFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firebase storage repository", err)
	}

	rewardOrderService := service.NewRewardOrderService(rewardOrderRepo, invoiceRepo, customerRepo, storeProfileRepo)

	rewardCustomerRepo, err := firestore.NewRewardCustomerFirestoreRepository()
	if err != nil {
		log.Panicln("## Error while creating firestore repository", err)
	}

	rewardCustomerService := service.NewRewardCustomerService(rewardCustomerRepo)

	rewardController := NewRewardController(rewardOrderService, rewardCustomerService)

	return rewardController, nil
}

func (suite *RewardSuite) TestCalculateReward() {

	storeId := "dropship_test_buah"
	dropshipId := "dropshipdll_test_buah"
	orderId := "190807016-o"

	rewardController, err := createRewardController()
	if err != nil {
		suite.T().Errorf("Unable to create controller %v", err)
	}
	log.Printf("## Using dropshipId %s", dropshipId)
	req, _ := http.NewRequest("POST", "/_api/reward/calc", nil)

	//req = mux.SetURLVars(req, map[string]string{
	//	"dropshipId": dropshipId,
	//})

	q := req.URL.Query()
	q.Add("dropshipId", dropshipId)
	q.Add("storeId", storeId)
	q.Add("orderId", orderId)
	req.URL.RawQuery = q.Encode()

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(rewardController.CalculateReward)
	handler.ServeHTTP(response, req)

	checkRewardResponseCode(suite.T(), http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["status"] != true {
		suite.T().Errorf("Expected the 'status' key of the response to be set to 'true'. Got '%s'", m["status"])
	}
}

func checkRewardResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
