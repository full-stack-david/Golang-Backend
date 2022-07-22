package service

import (
	"log"
	"mauappa-go/repository/firestore"
	"mauappa-go/util"
	"strconv"
	"testing"
	"time"

	"mauappa-go/model"

	"github.com/stretchr/testify/suite"
)

type RewardOrderServiceSuite struct {
	suite.Suite
	RewardOrderKeys []string
}

func TestRewardOrderServiceSuite(t *testing.T) {
	// This is what actually runs our suite
	util.InitTestConfig()
	suite.Run(t, new(RewardOrderServiceSuite))
}

func createRewardOrderService() RewardOrderService {
	rewardOrderRepository, err := firestore.NewRewardOrderFirestoreRepository()
	invoiceRepository, err := firestore.NewInvoiceFirestoreRepository()
	customerRepository, err := firestore.NewCustomerFirestoreRepository()
	storeProfileRepository, err := firestore.NewStoreProfileFirestoreRepository()
	if err != nil {
		return nil
	}
	rewardOrderService := NewRewardOrderService(rewardOrderRepository, invoiceRepository, customerRepository, storeProfileRepository)
	return rewardOrderService
}

func (suite *RewardOrderServiceSuite) TestCalculateRewardOrder() {

	storeId := "dropship_test_buah"
	orderId := "190801004-o"

	rewardOrderService := createRewardOrderService()
	log.Printf("## Using dropshipId %s", dropshipId)

	success, err := rewardOrderService.CalculateRewardOrder(storeId, dropshipId, orderId)

	if err != nil {
		suite.T().Errorf("Error while calculating reward points per order %v", err)
	}

	if success == false {
		suite.T().Errorf("Expected the CalculateRewardOrder in response. Got %v", true)
	}
}

func (suite *RewardOrderServiceSuite) TestCreateRewardOrder() string {

	mockRewardOrderModel := mockRewardOrderModel()
	rewardOrderService := createRewardOrderService()

	resp, err := rewardOrderService.CreateRewardOrder(dropshipId, mockRewardOrderModel)

	if err != nil {
		suite.T().Errorf("Error while creating project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected  response to be set to 'true'. Got '%v'", resp)
	}

	suite.RewardOrderKeys = append(suite.RewardOrderKeys, mockRewardOrderModel.Id)
	return mockRewardOrderModel.Id
}

func (suite *RewardOrderServiceSuite) TestUpdateRewardOrder() {

	orderId := suite.TestCreateRewardOrder()

	mockRewardOrderModel := mockRewardOrderModel()
	mockRewardOrderModel.Id = orderId

	rewardOrderService := createRewardOrderService()

	resp, err := rewardOrderService.UpdateRewardOrder(dropshipId, mockRewardOrderModel)

	if err != nil {
		suite.T().Errorf("Error while updating project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected  response to be set to 'true'. Got '%v'", resp)
	}
}

func (suite *RewardOrderServiceSuite) TestGetRewardOrder() {

	orderId := suite.TestCreateRewardOrder()

	rewardOrderService := createRewardOrderService()

	resp, err := rewardOrderService.GetRewardOrder(dropshipId, orderId)

	if err != nil {
		suite.T().Errorf("Error while getting project %v", err)
	}
	if resp == nil {
		suite.T().Errorf("Expected  response to be project model. Got '%v'", resp)
	}
}

func (suite *RewardOrderServiceSuite) TestGetRewardOrderList() {

	suite.TestCreateRewardOrder()

	rewardOrderService := createRewardOrderService()

	resp, err := rewardOrderService.GetRewardOrderList(dropshipId)

	if err != nil {
		suite.T().Errorf("Error while getting project list %v", err)
	}
	if resp == nil || len(*resp) <= 0 {
		suite.T().Errorf("Expected list of project models. Got '%v'", resp)
	}
}

func (suite *RewardOrderServiceSuite) TestDeleteRewardOrder() {
	orderId := suite.TestCreateRewardOrder()
	deleteRewardOrder(*suite, orderId)
}

func deleteRewardOrder(suite RewardOrderServiceSuite, orderId string) {
	rewardOrderService := createRewardOrderService()

	resp, err := rewardOrderService.DeleteRewardOrder(dropshipId, orderId)

	if err != nil {
		suite.T().Errorf("Error while deleting project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected response to be set to 'true'. Got '%v'", resp)
	}

}

func mockRewardOrderModel() model.RewardOrder {
	rand := strconv.Itoa(util.GenerateRandom())
	total := float32(util.GenerateRandomInRange(1000, 10000))
	mockRewardOrderModel := &model.RewardOrder{
		Id:           "RewardOrder-" + rand,
		InvoiceNo:    rand + rand + "-i",
		InvoiceDate:  time.Now(),
		OrderNo:      rand + rand + "-o",
		OrderDate:    time.Now(),
		CustomerId:   rand,
		CustomerName: "Ruslan" + rand,
		CreatedBy:    rand + "@" + rand + ".com",
		Total:        float32(total),
		User1:        "USE1" + rand,
		Level1:       float32(total * 0.015),
		User2:        "USE2" + rand,
		Level2:       float32(total * 0.01),
		User3:        "USE3" + rand,
		Level3:       float32(total * 0.005),
	}
	return *mockRewardOrderModel
}

func (suite *RewardOrderServiceSuite) TearDownSuite() {
	log.Println("## In Suite Tear Down ##")
	for _, elem := range suite.RewardOrderKeys {
		deleteRewardOrder(*suite, elem)
	}
}
