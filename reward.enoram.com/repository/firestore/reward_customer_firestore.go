package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type rewardCustomerFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewRewardCustomerFirestoreRepository() (repository.RewardCustomerRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &rewardCustomerFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a rewardCustomerFirestoreRepository) GetRewardCustomerList(dropshipId string) (*[]model.RewardCustomer, error) {
	log.Printf("## Get reward Customer list of dropshipId = %s ##", dropshipId)

	rewardCustomerPath := "rewards/" + dropshipId + "/customers"

	var rewardCustomerArray []model.RewardCustomer

	iter := a.client.Collection(rewardCustomerPath).Documents(a.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var rewardCustomer model.RewardCustomer
		doc.DataTo(&rewardCustomer)
		rewardCustomer.Id = doc.Ref.ID
		rewardCustomerArray = append(rewardCustomerArray, rewardCustomer)
	}

	return &rewardCustomerArray, nil
}

func (a rewardCustomerFirestoreRepository) GetRewardCustomer(dropshipId string, customerId string) (*model.RewardCustomer, error) {
	log.Printf("## Get RewardCustomer By dropshipId and CustomerId = %s and %s ##", dropshipId, customerId)

	rewardCustomerCollection := "rewards/" + dropshipId + "/customers"
	dsnap, err := a.client.Collection(rewardCustomerCollection).Doc(customerId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var rewardCustomer model.RewardCustomer
	dsnap.DataTo(&rewardCustomer)
	rewardCustomer.Id = dsnap.Ref.ID
	return &rewardCustomer, nil
}

func (a rewardCustomerFirestoreRepository) CreateRewardCustomer(dropshipId string, rewardCustomerModel model.RewardCustomer) (bool, error) {
	rewardCustomerCollection := "rewards/" + dropshipId + "/customers"
	_, err := a.client.Collection(rewardCustomerCollection).Doc(rewardCustomerModel.Id).Set(a.ctx, rewardCustomerModel)

	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardCustomer created with key : %s", rewardCustomerModel.Id)
	return true, nil
}

func (a rewardCustomerFirestoreRepository) UpdateRewardCustomer(dropshipId string, rewardCustomer model.RewardCustomer) (bool, error) {
	rewardCustomerCollection := "rewards/" + dropshipId + "/customers"
	_, err := a.client.Collection(rewardCustomerCollection).Doc(rewardCustomer.Id).Set(a.ctx, rewardCustomer)
	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardCustomer updated with key : %s", rewardCustomer.Id)
	return true, nil
}

func (a rewardCustomerFirestoreRepository) DeleteRewardCustomer(dropshipId string, customerId string) (bool, error) {
	rewardCustomerCollection := "rewards/" + dropshipId + "/customers"
	_, err := a.client.Collection(rewardCustomerCollection).Doc(customerId).Delete(a.ctx)
	if err != nil {
		log.Printf("An error while deleting data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardCustomer deleted")
	return true, nil
}

func (a rewardCustomerFirestoreRepository) DeleteRewardCustomerContent(dropshipId string, orderId string) (bool, error) {
	orderRef := a.client.Doc("rewards/" + dropshipId + "/orders/" + orderId)
	err := a.client.RunTransaction(a.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		orderDoc, err := tx.Get(orderRef) // tx.Get, NOT ny.Get!
		if err != nil {
			if orderDoc != nil && !orderDoc.Exists() {
				return nil
			} else {
				return err
			}
		}
		var rewardOrder model.RewardOrder
		orderDoc.DataTo(&rewardOrder)
		rewardOrder.Id = orderDoc.Ref.ID

		updatedCustomerMap := make(map[string]model.RewardCustomer)
		var deleteArray []string

		if rewardOrder.UsedRewardPoint != 0 {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.CustomerId)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			if customerDoc.Exists() {
				var customer model.RewardCustomer
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalSpentPoints -= rewardOrder.UsedRewardPoint
				customer.BalancePoints += rewardOrder.UsedRewardPoint
				updatedCustomerMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId] = customer
				deleteArray = append(deleteArray, "rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId+"/spentPoints/"+rewardOrder.Id)
			}
		}

		if rewardOrder.DanaSegar != 0 {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.CustomerId)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			if customerDoc.Exists() {
				var customer model.RewardCustomer
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalDanaSegars -= rewardOrder.DanaSegar
				customer.BalancePoints -= rewardOrder.DanaSegar
				updatedCustomerMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId] = customer
				deleteArray = append(deleteArray, "rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId+"/danaSegars/"+rewardOrder.Id)
			}
		}

		if rewardOrder.User1 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User1Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			if customerDoc.Exists() {
				var customer model.RewardCustomer
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints -= rewardOrder.Level1
				customer.BalancePoints -= rewardOrder.Level1
				updatedCustomerMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User1Id] = customer
				deleteArray = append(deleteArray, "rewards/"+dropshipId+"/customers/"+rewardOrder.User1Id+"/rewardPoints/"+rewardOrder.Id)
			}
		}

		if rewardOrder.User2 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User2Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			if customerDoc.Exists() {
				var customer model.RewardCustomer
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints -= rewardOrder.Level2
				customer.BalancePoints -= rewardOrder.Level2
				updatedCustomerMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User2Id] = customer
				deleteArray = append(deleteArray, "rewards/"+dropshipId+"/customers/"+rewardOrder.User2Id+"/rewardPoints/"+rewardOrder.Id)
			}
		}

		if rewardOrder.User3 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User3Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			if customerDoc.Exists() {
				var customer model.RewardCustomer
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints -= rewardOrder.Level3
				customer.BalancePoints -= rewardOrder.Level3
				updatedCustomerMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User3Id] = customer
				deleteArray = append(deleteArray, "rewards/"+dropshipId+"/customers/"+rewardOrder.User3Id+"/rewardPoints/"+rewardOrder.Id)
			}
		}

		for key, value := range updatedCustomerMap {
			docRef := a.client.Doc(key)
			err = tx.Set(docRef, value)
			if err != nil {
				return err
			}
		}

		for _, value := range deleteArray {
			docRef := a.client.Doc(value)
			err = tx.Delete(docRef)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a rewardCustomerFirestoreRepository) CalculateRewardCustomer(dropshipId string, orderId string) (bool, error) {
	log.Printf("CalculateRewardCustomer started")
	orderRef := a.client.Doc("rewards/" + dropshipId + "/orders/" + orderId)
	err := a.client.RunTransaction(a.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		orderDoc, err := tx.Get(orderRef) // tx.Get, NOT ny.Get!
		if err != nil {
			if orderDoc != nil && !orderDoc.Exists() {
				return nil
			} else {
				return err
			}
		}
		var rewardOrder model.RewardOrder
		orderDoc.DataTo(&rewardOrder)
		rewardOrder.Id = orderDoc.Ref.ID

		setRewardMap := make(map[string]interface{})

		log.Printf("Before Calculating")

		if rewardOrder.UsedRewardPoint != 0 {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.CustomerId)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			var customer model.RewardCustomer
			if customerDoc.Exists() {
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalSpentPoints += rewardOrder.UsedRewardPoint
				customer.BalancePoints -= rewardOrder.UsedRewardPoint
			} else {
				customer = model.RewardCustomer{
					Id:                rewardOrder.CustomerId,
					Name:              rewardOrder.CustomerName,
					Email:             rewardOrder.CreatedBy,
					ReferalCode:       rewardOrder.CustomerCode,
					TotalSpentPoints:  rewardOrder.UsedRewardPoint,
					BalancePoints:     -rewardOrder.UsedRewardPoint,
				}
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId] = customer
			spent := model.CustomerSpent{
				Id:          rewardOrder.Id,
				InvoiceNo:   rewardOrder.InvoiceNo,
				SpentPoints: rewardOrder.UsedRewardPoint,
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId+"/spentPoints/"+rewardOrder.Id] = spent
		}

		log.Printf("After calculating UsedPoints")

		if rewardOrder.DanaSegar != 0 {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.CustomerId)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			var customer model.RewardCustomer
			if customerDoc.Exists() {
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalDanaSegars += rewardOrder.DanaSegar
				customer.BalancePoints += rewardOrder.DanaSegar
			} else {
				customer = model.RewardCustomer{
					Id:                rewardOrder.CustomerId,
					Name:              rewardOrder.CustomerName,
					Email:             rewardOrder.CreatedBy,
					ReferalCode:       rewardOrder.CustomerCode,
					TotalDanaSegars:   rewardOrder.DanaSegar,
					BalancePoints:     rewardOrder.DanaSegar,
				}
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId] = customer
			danaSegar := model.CustomerDanaSegar{
				Id:          rewardOrder.Id,
				InvoiceNo:   rewardOrder.InvoiceNo,
				DanaSegar: rewardOrder.DanaSegar,
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.CustomerId+"/danaSegars/"+rewardOrder.Id] = danaSegar
		}

		log.Printf("After calculating DanaSegars")

		if rewardOrder.User1 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User1Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			var customer model.RewardCustomer
			if customerDoc.Exists() {
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints += rewardOrder.Level1
				customer.BalancePoints += rewardOrder.Level1
			} else {
				customer = model.RewardCustomer{
					Id:                rewardOrder.User1Id,
					Name:              rewardOrder.User1Name,
					Email:             rewardOrder.User1Email,
					ReferalCode:       rewardOrder.User1,
					TotalRewardPoints: rewardOrder.Level1,
					BalancePoints:     rewardOrder.Level1,
				}
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User1Id] = customer
			reward := model.CustomerReward{
				Id:           rewardOrder.Id,
				InvoiceNo:    rewardOrder.InvoiceNo,
				RewardPoints: rewardOrder.Level1,
				Level:        1,
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User1Id+"/rewardPoints/"+rewardOrder.Id] = reward
		}

		log.Printf("After calculating rewardPoints 1")

		if rewardOrder.User2 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User2Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			var customer model.RewardCustomer
			if customerDoc.Exists() {
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints += rewardOrder.Level2
				customer.BalancePoints += rewardOrder.Level2
			} else {
				customer = model.RewardCustomer{
					Id:                rewardOrder.User2Id,
					Name:              rewardOrder.User2Name,
					Email:             rewardOrder.User2Email,
					ReferalCode:       rewardOrder.User2,
					TotalRewardPoints: rewardOrder.Level2,
					BalancePoints:     rewardOrder.Level2,
				}
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User2Id] = customer
			reward := model.CustomerReward{
				Id:           rewardOrder.Id,
				InvoiceNo:    rewardOrder.InvoiceNo,
				RewardPoints: rewardOrder.Level2,
				Level:        2,
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User2Id+"/rewardPoints/"+rewardOrder.Id] = reward
		}

		log.Printf("After calculating rewardPoints 2")

		if rewardOrder.User3 != "" {
			customerRef := a.client.Doc("rewards/" + dropshipId + "/customers/" + rewardOrder.User3Id)
			customerDoc, err := tx.Get(customerRef)
			if err != nil && orderDoc == nil {
				return err
			}
			var customer model.RewardCustomer
			if customerDoc.Exists() {
				customerDoc.DataTo(&customer)
				customer.Id = customerDoc.Ref.ID
				customer.TotalRewardPoints += rewardOrder.Level3
				customer.BalancePoints += rewardOrder.Level3
			} else {
				customer = model.RewardCustomer{
					Id:                rewardOrder.User3Id,
					Name:              rewardOrder.User3Name,
					Email:             rewardOrder.User3Email,
					ReferalCode:       rewardOrder.User3,
					TotalRewardPoints: rewardOrder.Level3,
					BalancePoints:     rewardOrder.Level3,
				}
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User3Id] = customer
			reward := model.CustomerReward{
				Id:           rewardOrder.Id,
				InvoiceNo:    rewardOrder.InvoiceNo,
				RewardPoints: rewardOrder.Level3,
				Level:        3,
			}
			setRewardMap["rewards/"+dropshipId+"/customers/"+rewardOrder.User3Id+"/rewardPoints/"+rewardOrder.Id] = reward
		}

		log.Printf("After calculating rewardPoints 3")

		for key, value := range setRewardMap {
			docRef := a.client.Doc(key)
			err = tx.Set(docRef, value)
			if err != nil {
				return err
			}
		}

		log.Printf("Before Return")

		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
