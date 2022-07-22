package service

import (
	"log"
	"mauappa-go/repository/firestore"
	"mauappa-go/util"
	"strconv"
	"testing"

	"mauappa-go/model"

	"github.com/stretchr/testify/suite"
)

const dropshipId string = "dropshipdll_test_buah"

type RewardCustomerServiceSuite struct {
	suite.Suite
	RewardCustomerKeys []string
}

func TestRewardCustomerServiceSuite(t *testing.T) {
	// This is what actually runs our suite
	util.InitTestConfig()
	suite.Run(t, new(RewardCustomerServiceSuite))
}

func createRewardCustomerService() RewardCustomerService {
	rewardCustomerRepository, err := firestore.NewRewardCustomerFirestoreRepository()
	if err != nil {
		return nil
	}
	rewardCustomerService := NewRewardCustomerService(rewardCustomerRepository)
	return rewardCustomerService
}

func (suite *RewardCustomerServiceSuite) TestDeleteRewardCustomerContent() {

	orderId := "190809002-o"

	rewardCustomerService := createRewardCustomerService()
	log.Printf("## Using dropshipId %s", dropshipId)

	success, err := rewardCustomerService.DeleteRewardCustomerContent(dropshipId, orderId)

	if err != nil {
		suite.T().Errorf("Error while deleting customer rewards contents %v", err)
	}

	if success == false {
		suite.T().Errorf("Expected the DeleteRewardCustomerContent in response. Got %v", true)
	}
}

func (suite *RewardCustomerServiceSuite) TestCalculateRewardCustomer() {

	orderId := "190809002-o"

	rewardCustomerService := createRewardCustomerService()
	log.Printf("## Using dropshipId %s", dropshipId)

	success, err := rewardCustomerService.CalculateRewardCustomer(dropshipId, orderId)

	if err != nil {
		suite.T().Errorf("Error while calculating customer rewards %v", err)
	}

	if success == false {
		suite.T().Errorf("Expected the CalculateRewardCustomer in response. Got %v", true)
	}
}

func (suite *RewardCustomerServiceSuite) TestCreateRewardCustomer() string {

	mockRewardCustomerModel := mockRewardCustomerModel()
	rewardCustomerService := createRewardCustomerService()

	resp, err := rewardCustomerService.CreateRewardCustomer(dropshipId, mockRewardCustomerModel)

	if err != nil {
		suite.T().Errorf("Error while creating project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected  response to be set to 'true'. Got '%v'", resp)
	}

	suite.RewardCustomerKeys = append(suite.RewardCustomerKeys, mockRewardCustomerModel.Id)
	return mockRewardCustomerModel.Id
}

func (suite *RewardCustomerServiceSuite) TestUpdateRewardCustomer() {

	customerId := suite.TestCreateRewardCustomer()

	mockRewardCustomerModel := mockRewardCustomerModel()
	mockRewardCustomerModel.Id = customerId

	rewardCustomerService := createRewardCustomerService()

	resp, err := rewardCustomerService.UpdateRewardCustomer(dropshipId, mockRewardCustomerModel)

	if err != nil {
		suite.T().Errorf("Error while updating project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected  response to be set to 'true'. Got '%v'", resp)
	}
}

func (suite *RewardCustomerServiceSuite) TestGetRewardCustomer() {

	customerId := suite.TestCreateRewardCustomer()

	rewardCustomerService := createRewardCustomerService()

	resp, err := rewardCustomerService.GetRewardCustomer(dropshipId, customerId)

	if err != nil {
		suite.T().Errorf("Error while getting project %v", err)
	}
	if resp == nil {
		suite.T().Errorf("Expected  response to be project model. Got '%v'", resp)
	}
}

func (suite *RewardCustomerServiceSuite) TestGetRewardCustomerList() {

	suite.TestCreateRewardCustomer()

	rewardCustomerService := createRewardCustomerService()

	resp, err := rewardCustomerService.GetRewardCustomerList(dropshipId)

	if err != nil {
		suite.T().Errorf("Error while getting project list %v", err)
	}
	if resp == nil || len(*resp) <= 0 {
		suite.T().Errorf("Expected list of project models. Got '%v'", resp)
	}
}

func (suite *RewardCustomerServiceSuite) TestDeleteRewardCustomer() {
	customerId := suite.TestCreateRewardCustomer()
	deleteRewardCustomer(*suite, customerId)
}

func deleteRewardCustomer(suite RewardCustomerServiceSuite, customerId string) {
	rewardCustomerService := createRewardCustomerService()

	resp, err := rewardCustomerService.DeleteRewardCustomer(dropshipId, customerId)

	if err != nil {
		suite.T().Errorf("Error while deleting project %v", err)
	}
	if resp != true {
		suite.T().Errorf("Expected response to be set to 'true'. Got '%v'", resp)
	}

}

func mockRewardCustomerModel() model.RewardCustomer {
	rand := strconv.Itoa(util.GenerateRandom())
	points := float32(util.GenerateRandomInRange(100, 1000))
	sPoints := float32(util.GenerateRandomInRange(100, 1000))
	mockRewardCustomerModel := &model.RewardCustomer{
		Id:                "RewardCustomer-" + rand,
		Name:              "Ruslan" + rand,
		Email:             rand + "@" + rand + ".com",
		ReferalCode:       "USE1" + rand,
		TotalRewardPoints: points,
		TotalSpentPoints:  sPoints,
		BalancePoints:     points - sPoints,
	}
	return *mockRewardCustomerModel
}

func (suite *RewardCustomerServiceSuite) TearDownSuite() {
	log.Println("## In Suite Tear Down ##")
	for _, elem := range suite.RewardCustomerKeys {
		deleteRewardCustomer(*suite, elem)
	}
}
