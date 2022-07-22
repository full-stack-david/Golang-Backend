package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type customerDanaSegarFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewCustomerDanaSegarFirestoreRepository() (repository.CustomerDanaSegarRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &customerDanaSegarFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a customerDanaSegarFirestoreRepository) GetCustomerDanaSegarList(dropshipId string, customerId string) (*[]model.CustomerDanaSegar, error) {
	log.Printf("## Get DanaSegar order list of dropshipId = %s ##", dropshipId)

	customerDanaSegarPath := "rewards/" + dropshipId + "/customers/" + customerId + "/danaSegars"

	var customerDanaSegarArray []model.CustomerDanaSegar

	iter := a.client.Collection(customerDanaSegarPath).Documents(a.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var customerDanaSegar model.CustomerDanaSegar
		doc.DataTo(&customerDanaSegar)
		customerDanaSegar.Id = doc.Ref.ID
		customerDanaSegarArray = append(customerDanaSegarArray, customerDanaSegar)
	}

	return &customerDanaSegarArray, nil
}

func (a customerDanaSegarFirestoreRepository) GetCustomerDanaSegar(dropshipId string, customerId string, orderId string) (*model.CustomerDanaSegar, error) {
	log.Printf("## Get CustomerDanaSegar By dropshipId and orderId = %s and %s ##", dropshipId, orderId)

	customerDanaSegarCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/danaSegars"
	dsnap, err := a.client.Collection(customerDanaSegarCollection).Doc(orderId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var customerDanaSegar model.CustomerDanaSegar
	dsnap.DataTo(&customerDanaSegar)
	customerDanaSegar.Id = dsnap.Ref.ID
	return &customerDanaSegar, nil
}

func (a customerDanaSegarFirestoreRepository) CreateCustomerDanaSegar(dropshipId string, customerId string, customerDanaSegarModel model.CustomerDanaSegar) (bool, error) {
	customerDanaSegarCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/danaSegars"
	_, err := a.client.Collection(customerDanaSegarCollection).Doc(customerDanaSegarModel.Id).Set(a.ctx, customerDanaSegarModel)

	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("customerDanaSegar created with key : %s", customerDanaSegarModel.Id)
	return true, nil
}

func (a customerDanaSegarFirestoreRepository) UpdateCustomerDanaSegar(dropshipId string, customerId string, customerDanaSegar model.CustomerDanaSegar) (bool, error) {
	customerDanaSegarCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/danaSegars"
	_, err := a.client.Collection(customerDanaSegarCollection).Doc(customerDanaSegar.Id).Set(a.ctx, customerDanaSegar)
	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerDanaSegar updated with key : %s", customerDanaSegar.Id)
	return true, nil
}

func (a customerDanaSegarFirestoreRepository) DeleteCustomerDanaSegar(dropshipId string, customerId string, orderId string) (bool, error) {
	customerDanaSegarCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/danaSegars"
	_, err := a.client.Collection(customerDanaSegarCollection).Doc(orderId).Delete(a.ctx)
	if err != nil {
		log.Printf("An error while deleting data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerDanaSegar deleted")
	return true, nil
}
