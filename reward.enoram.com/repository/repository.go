package repository

import (
	"mauappa-go/model"
)

type CustomerRepository interface {
	GetCustomerByReferalCode(storeId string, referalCode string) (*model.Customer, error)
	GetCustomerById(storeId string, customerId string) (*model.Customer, error)
}

type InvoiceRepository interface {
	GetInvoice(dropshipId string, orderId string) (*model.Invoice, error)
}

type StoreProfileRepository interface {
	GetStoreProfile(storeId string) (*model.StoreProfile, error)
}

type RewardOrderRepository interface {
	GetRewardOrderList(dropshipId string) (*[]model.RewardOrder, error)
	GetRewardOrder(dropshipId string, orderId string) (*model.RewardOrder, error)
	CreateRewardOrder(dropshipId string, RewardOrderModel model.RewardOrder) (bool, error)
	UpdateRewardOrder(dropshipId string, RewardOrderModel model.RewardOrder) (bool, error)
	DeleteRewardOrder(dropshipId string, orderId string) (bool, error)
}

type RewardCustomerRepository interface {
	GetRewardCustomerList(dropshipId string) (*[]model.RewardCustomer, error)
	GetRewardCustomer(dropshipId string, customerId string) (*model.RewardCustomer, error)
	CreateRewardCustomer(dropshipId string, RewardCustomerModel model.RewardCustomer) (bool, error)
	UpdateRewardCustomer(dropshipId string, RewardCustomerModel model.RewardCustomer) (bool, error)
	DeleteRewardCustomer(dropshipId string, customerId string) (bool, error)
	DeleteRewardCustomerContent(dropshipId string, orderId string) (bool, error)
	CalculateRewardCustomer(dropshipId string, orderId string) (bool, error)
}

type CustomerRewardRepository interface {
	GetCustomerRewardList(dropshipId string, customerId string) (*[]model.CustomerReward, error)
	GetCustomerReward(dropshipId string, customerId string, orderId string) (*model.CustomerReward, error)
	CreateCustomerReward(dropshipId string, customerId string, CustomerRewardModel model.CustomerReward) (bool, error)
	UpdateCustomerReward(dropshipId string, customerId string, CustomerRewardModel model.CustomerReward) (bool, error)
	DeleteCustomerReward(dropshipId string, customerId string, orderId string) (bool, error)
}

type CustomerSpentRepository interface {
	GetCustomerSpentList(dropshipId string, customerId string) (*[]model.CustomerSpent, error)
	GetCustomerSpent(dropshipId string, customerId string, orderId string) (*model.CustomerSpent, error)
	CreateCustomerSpent(dropshipId string, customerId string, CustomerSpentModel model.CustomerSpent) (bool, error)
	UpdateCustomerSpent(dropshipId string, customerId string, CustomerSpentModel model.CustomerSpent) (bool, error)
	DeleteCustomerSpent(dropshipId string, customerId string, orderId string) (bool, error)
}

type CustomerDanaSegarRepository interface {
	GetCustomerDanaSegarList(dropshipId string, customerId string) (*[]model.CustomerDanaSegar, error)
	GetCustomerDanaSegar(dropshipId string, customerId string, orderId string) (*model.CustomerDanaSegar, error)
	CreateCustomerDanaSegar(dropshipId string, customerId string, CustomerDanaSegarModel model.CustomerDanaSegar) (bool, error)
	UpdateCustomerDanaSegar(dropshipId string, customerId string, CustomerDanaSegarModel model.CustomerDanaSegar) (bool, error)
	DeleteCustomerDanaSegar(dropshipId string, customerId string, orderId string) (bool, error)
}
