package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type customerSpentFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewCustomerSpentFirestoreRepository() (repository.CustomerSpentRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &customerSpentFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a customerSpentFirestoreRepository) GetCustomerSpentList(dropshipId string, customerId string) (*[]model.CustomerSpent, error) {
	log.Printf("## Get Spent order list of dropshipId = %s ##", dropshipId)

	customerSpentPath := "rewards/" + dropshipId + "/customers/" + customerId + "/spentPoints"

	var customerSpentArray []model.CustomerSpent

	iter := a.client.Collection(customerSpentPath).Documents(a.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var customerSpent model.CustomerSpent
		doc.DataTo(&customerSpent)
		customerSpent.Id = doc.Ref.ID
		customerSpentArray = append(customerSpentArray, customerSpent)
	}

	return &customerSpentArray, nil
}

func (a customerSpentFirestoreRepository) GetCustomerSpent(dropshipId string, customerId string, orderId string) (*model.CustomerSpent, error) {
	log.Printf("## Get CustomerSpent By dropshipId and orderId = %s and %s ##", dropshipId, orderId)

	customerSpentCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/spentPoints"
	dsnap, err := a.client.Collection(customerSpentCollection).Doc(orderId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var customerSpent model.CustomerSpent
	dsnap.DataTo(&customerSpent)
	customerSpent.Id = dsnap.Ref.ID
	return &customerSpent, nil
}

func (a customerSpentFirestoreRepository) CreateCustomerSpent(dropshipId string, customerId string, customerSpentModel model.CustomerSpent) (bool, error) {
	customerSpentCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/spentPoints"
	_, err := a.client.Collection(customerSpentCollection).Doc(customerSpentModel.Id).Set(a.ctx, customerSpentModel)

	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("customerSpent created with key : %s", customerSpentModel.Id)
	return true, nil
}

func (a customerSpentFirestoreRepository) UpdateCustomerSpent(dropshipId string, customerId string, customerSpent model.CustomerSpent) (bool, error) {
	customerSpentCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/spentPoints"
	_, err := a.client.Collection(customerSpentCollection).Doc(customerSpent.Id).Set(a.ctx, customerSpent)
	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerSpent updated with key : %s", customerSpent.Id)
	return true, nil
}

func (a customerSpentFirestoreRepository) DeleteCustomerSpent(dropshipId string, customerId string, orderId string) (bool, error) {
	customerSpentCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/spentPoints"
	_, err := a.client.Collection(customerSpentCollection).Doc(orderId).Delete(a.ctx)
	if err != nil {
		log.Printf("An error while deleting data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerSpent deleted")
	return true, nil
}
