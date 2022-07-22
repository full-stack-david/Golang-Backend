package service

import (
	"fmt"
	"log"
	"math"
	"mauappa-go/model"
	"mauappa-go/repository"
)

type rewardOrderServiceImpl struct {
	rewardOrderRepository  repository.RewardOrderRepository
	invoiceRepository      repository.InvoiceRepository
	customerRepository     repository.CustomerRepository
	storeProfileRepository repository.StoreProfileRepository
}

func NewRewardOrderService(
	rewardOrderRepo repository.RewardOrderRepository,
	invoiceRepo repository.InvoiceRepository,
	customerRepo repository.CustomerRepository,
	storeProfileRepo repository.StoreProfileRepository,
) RewardOrderService {
	return &rewardOrderServiceImpl{rewardOrderRepo, invoiceRepo, customerRepo, storeProfileRepo}
}

func (x rewardOrderServiceImpl) CreateRewardOrder(dropshipId string, rewardOrderModel model.RewardOrder) (bool, error) {
	log.Println("##  Calling repository to create reward order ##")

	if rewardOrderModel.Id == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and reward order id can not be empty")
	}

	return x.rewardOrderRepository.CreateRewardOrder(dropshipId, rewardOrderModel)
}

func (x rewardOrderServiceImpl) UpdateRewardOrder(dropshipId string, rewardOrderModel model.RewardOrder) (bool, error) {
	log.Println("##  Calling repository to update rewardOrder ##")

	if rewardOrderModel.Id == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and rewardOrder id can not be empty")
	}

	return x.rewardOrderRepository.UpdateRewardOrder(dropshipId, rewardOrderModel)
}

func (x rewardOrderServiceImpl) DeleteRewardOrder(dropshipId string, orderId string) (bool, error) {
	log.Println("##  Calling repository to delete rewardOrder ##")

	if orderId == "" || dropshipId == "" {
		return false, fmt.Errorf("dropshipId and rewardOrder id can not be empty")
	}

	return x.rewardOrderRepository.DeleteRewardOrder(dropshipId, orderId)
}

func (x rewardOrderServiceImpl) GetRewardOrder(dropshipId string, orderId string) (*model.RewardOrder, error) {
	log.Println("##  Calling repository to fetch rewardOrder ##")

	if dropshipId == "" || orderId == "" {
		return nil, fmt.Errorf("dropshipId and rewardOrder id can not be empty")
	}

	return x.rewardOrderRepository.GetRewardOrder(dropshipId, orderId)
}

func (x rewardOrderServiceImpl) GetRewardOrderList(dropshipId string) (*[]model.RewardOrder, error) {
	log.Println("##  Calling repository to fetch rewardOrders ##")

	if dropshipId == "" {
		return nil, fmt.Errorf("dropshipId can not be empty")
	}

	return x.rewardOrderRepository.GetRewardOrderList(dropshipId)
}

func (x rewardOrderServiceImpl) CalculateRewardOrder(storeId string, dropshipId string, orderId string) (bool, error) {
	log.Printf("##  Calling repository to fetch invoices ##")

	storeProfile, err := x.storeProfileRepository.GetStoreProfile(storeId)

	if err != nil {
		return false, err
	}

	log.Print(storeProfile.ShipmentCostGE)

	log.Print(dropshipId)
	invoice, err := x.invoiceRepository.GetInvoice(dropshipId, orderId)

	if err != nil {
		return false, err
	}

	calAmount := invoice.Total
	if storeProfile.ShipmentCostGE != 0 {
		if invoice.Total <= storeProfile.ShipmentCostGE {
			calAmount = invoice.Total - storeProfile.ShipmentCostValue
		}
	}

	level1 := float32(math.Round(float64(calAmount * 1.5 / 100)))
	level2 := float32(math.Round(float64(calAmount * 1 / 100)))
	level3 := float32(math.Round(float64(calAmount * 0.5 / 100)))

	user1, user1Id, user1Name, user1Email := "", "", "", ""
	user2, user2Id, user2Name, user2Email := "", "", "", ""
	user3, user3Id, user3Name, user3Email := "", "", "", ""

	customer, err := x.customerRepository.GetCustomerById(storeId, invoice.CustomerId)

	if err != nil {
		log.Fatalf("failed to get customer by id: %s", err)
	}

	if customer != nil && customer.ReferBy != "" {
		user1 = customer.ReferBy
		level1Customer, err := x.customerRepository.GetCustomerByReferalCode(storeId, customer.ReferBy)
		if err != nil {
			log.Fatalf("failed to get customer by referalCode: %s", err)
		}
		user1Id = level1Customer.Id
		user1Name = level1Customer.Name
		user1Email = level1Customer.Email

		if level1Customer != nil && err == nil && level1Customer.ReferBy != "" {
			user2 = level1Customer.ReferBy
			level2Customer, err := x.customerRepository.GetCustomerByReferalCode(storeId, level1Customer.ReferBy)
			if err != nil {
				log.Fatalf("failed to get customer by referalCode: %s", err)
			}
			user2Id = level2Customer.Id
			user2Name = level2Customer.Name
			user2Email = level2Customer.Email

			if level2Customer != nil && err == nil && level2Customer.ReferBy != "" {
				user3 = level2Customer.ReferBy
				level3Customer, err := x.customerRepository.GetCustomerByReferalCode(storeId, level2Customer.ReferBy)
				if err != nil {
					log.Fatalf("failed to get customer by referalCode: %s", err)
				}
				user3Id = level3Customer.Id
				user3Name = level3Customer.Name
				user3Email = level3Customer.Email
			}
		}
	}

	var rewardOrder model.RewardOrder = model.RewardOrder{
		Id:              orderId,
		InvoiceNo:       invoice.InvoiceNo,
		InvoiceDate:     invoice.InvoiceDate,
		OrderNo:         invoice.OrderNo,
		OrderDate:       invoice.OrderDate,
		CustomerId:      invoice.CustomerId,
		CustomerName:    invoice.CustomerName,
		CustomerCode:    customer.ReferalCode,
		CreatedBy:       invoice.CreatedBy,
		Total:           invoice.Total,
		User1:           user1,
		Level1:          level1,
		User1Id:         user1Id,
		User1Name:       user1Name,
		User1Email:      user1Email,
		User2:           user2,
		Level2:          level2,
		User2Id:         user2Id,
		User2Name:       user2Name,
		User2Email:      user2Email,
		User3:           user3,
		Level3:          level3,
		User3Id:         user3Id,
		User3Name:       user3Name,
		User3Email:      user3Email,
		DanaSegar:       invoice.DanaSegar,
		UsedRewardPoint: invoice.UsedRewardPoint,
	}

	return x.rewardOrderRepository.CreateRewardOrder(dropshipId, rewardOrder)
}
