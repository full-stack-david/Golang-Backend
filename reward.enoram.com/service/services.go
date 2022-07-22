package service

import (
	"mauappa-go/model"
)

type RewardOrderService interface {
	GetRewardOrderList(dropshipId string) (*[]model.RewardOrder, error)
	GetRewardOrder(dropshipId string, orderId string) (*model.RewardOrder, error)
	CreateRewardOrder(dropshipId string, RewardOrderModel model.RewardOrder) (bool, error)
	UpdateRewardOrder(dropshipId string, RewardOrderModel model.RewardOrder) (bool, error)
	DeleteRewardOrder(dropshipId string, orderId string) (bool, error)
	CalculateRewardOrder(storeId string, dropshipId string, orderId string) (bool, error)
}

type RewardCustomerService interface {
	GetRewardCustomerList(dropshipId string) (*[]model.RewardCustomer, error)
	GetRewardCustomer(dropshipId string, CustomerId string) (*model.RewardCustomer, error)
	CreateRewardCustomer(dropshipId string, RewardCustomerModel model.RewardCustomer) (bool, error)
	UpdateRewardCustomer(dropshipId string, RewardCustomerModel model.RewardCustomer) (bool, error)
	DeleteRewardCustomer(dropshipId string, CustomerId string) (bool, error)
	DeleteRewardCustomerContent(dropshipId string, storeId string) (bool, error)
	CalculateRewardCustomer(dropshipId string, storeId string) (bool, error)
}

type RewardService interface {
	CalculateReward(storeId string, dropshipId string, orderId string) (bool, error)
}
