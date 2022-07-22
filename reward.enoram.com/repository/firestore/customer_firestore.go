package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
)

const productCollection = "products"

type customerFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewCustomerFirestoreRepository() (repository.CustomerRepository, error) {
	client, ctx, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &customerFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a customerFirestoreRepository) GetCustomerByReferalCode(storeId string, referalCode string) (*model.Customer, error) {
	log.Printf("## Get Customer By Referal code = %s ##", referalCode)

	customerPath := "customers/" + storeId + "/actives"

	iter, err := a.client.Collection(customerPath).Where("referalCode", "==", referalCode).Documents(a.ctx).Next()
	if err != nil {
		return nil, err
	}
	var customer model.Customer
	err = iter.DataTo(&customer)
	if err != nil {
		return nil, err
	}
	customer.Id = iter.Ref.ID

	return &customer, nil
}

func (a customerFirestoreRepository) GetCustomerById(storeId string, customerId string) (*model.Customer, error) {
	log.Printf("## Get Customer by ID = %s ##", customerId)

	customerPath := "customers/" + storeId + "/actives"

	dsnap, err := a.client.Collection(customerPath).Doc(customerId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var customer model.Customer
	err = dsnap.DataTo(&customer)
	if err != nil {
		return nil, err
	}
	customer.Id = dsnap.Ref.ID

	return &customer, nil
}
