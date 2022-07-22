package service

import (
	"fmt"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"
)

type rewardCustomerServiceImpl struct {
	rewardCustomerRepository repository.RewardCustomerRepository
}

func NewRewardCustomerService(
	rewardCustomerRepo repository.RewardCustomerRepository,
) RewardCustomerService {
	return &rewardCustomerServiceImpl{rewardCustomerRepo}
}

func (x rewardCustomerServiceImpl) CreateRewardCustomer(dropshipId string, rewardCustomerModel model.RewardCustomer) (bool, error) {
	log.Println("##  Calling repository to create reward customer ##")

	if rewardCustomerModel.Id == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and reward customer id can not be empty")
	}

	return x.rewardCustomerRepository.CreateRewardCustomer(dropshipId, rewardCustomerModel)
}

func (x rewardCustomerServiceImpl) UpdateRewardCustomer(dropshipId string, rewardCustomerModel model.RewardCustomer) (bool, error) {
	log.Println("##  Calling repository to update rewardCustomer ##")

	if rewardCustomerModel.Id == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and rewardCustomer id can not be empty")
	}

	return x.rewardCustomerRepository.UpdateRewardCustomer(dropshipId, rewardCustomerModel)
}

func (x rewardCustomerServiceImpl) DeleteRewardCustomer(dropshipId string, customerId string) (bool, error) {
	log.Println("##  Calling repository to delete rewardCustomer ##")

	if customerId == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and rewardCustomer id can not be empty")
	}

	return x.rewardCustomerRepository.DeleteRewardCustomer(dropshipId, customerId)
}

func (x rewardCustomerServiceImpl) GetRewardCustomer(dropshipId string, customerId string) (*model.RewardCustomer, error) {
	log.Println("##  Calling repository to fetch rewardCustomer ##")

	if dropshipId == "" || customerId == "" {
		return nil, fmt.Errorf("dropshipId and rewardCustomer id can not be empty")
	}

	return x.rewardCustomerRepository.GetRewardCustomer(dropshipId, customerId)
}

func (x rewardCustomerServiceImpl) GetRewardCustomerList(dropshipId string) (*[]model.RewardCustomer, error) {
	log.Println("##  Calling repository to fetch rewardCustomers ##")

	if dropshipId == "" {
		return nil, fmt.Errorf("dropshipId can not be empty")
	}

	return x.rewardCustomerRepository.GetRewardCustomerList(dropshipId)
}

func (x rewardCustomerServiceImpl) DeleteRewardCustomerContent(dropshipId string, orderId string) (bool, error) {
	log.Println("##  Calling repository to delete reward customer content ##")

	if dropshipId == "" || orderId == "" {
		return false, fmt.Errorf("dropshipId and rewardCustomer id can not be empty")
	}

	return x.rewardCustomerRepository.DeleteRewardCustomerContent(dropshipId, orderId)
}

func (x rewardCustomerServiceImpl) CalculateRewardCustomer(dropshipId string, orderId string) (bool, error) {
	log.Println("##  Calling repository to fetch rewardCustomer ##")

	if dropshipId == "" || orderId == "" {
		return false, fmt.Errorf("dropshipId and rewardCustomer id can not be empty")
	}

	return x.rewardCustomerRepository.CalculateRewardCustomer(dropshipId, orderId)
}