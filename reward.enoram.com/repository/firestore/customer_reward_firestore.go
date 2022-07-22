package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type customerRewardFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewCustomerRewardFirestoreRepository() (repository.CustomerRewardRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &customerRewardFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a customerRewardFirestoreRepository) GetCustomerRewardList(dropshipId string, customerId string) (*[]model.CustomerReward, error) {
	log.Printf("## Get reward order list of dropshipId = %s ##", dropshipId)

	customerRewardPath := "rewards/" + dropshipId + "/customers/" + customerId + "/rewardPoints"

	var customerRewardArray []model.CustomerReward

	iter := a.client.Collection(customerRewardPath).Documents(a.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var customerReward model.CustomerReward
		doc.DataTo(&customerReward)
		customerReward.Id = doc.Ref.ID
		customerRewardArray = append(customerRewardArray, customerReward)
	}

	return &customerRewardArray, nil
}

func (a customerRewardFirestoreRepository) GetCustomerReward(dropshipId string, customerId string, orderId string) (*model.CustomerReward, error) {
	log.Printf("## Get CustomerReward By dropshipId and orderId = %s and %s ##", dropshipId, orderId)

	customerRewardCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/rewardPoints"
	dsnap, err := a.client.Collection(customerRewardCollection).Doc(orderId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var customerReward model.CustomerReward
	dsnap.DataTo(&customerReward)
	customerReward.Id = dsnap.Ref.ID
	return &customerReward, nil
}

func (a customerRewardFirestoreRepository) CreateCustomerReward(dropshipId string, customerId string, customerRewardModel model.CustomerReward) (bool, error) {
	customerRewardCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/rewardPoints"
	_, err := a.client.Collection(customerRewardCollection).Doc(customerRewardModel.Id).Set(a.ctx, customerRewardModel)

	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("customerReward created with key : %s", customerRewardModel.Id)
	return true, nil
}

func (a customerRewardFirestoreRepository) UpdateCustomerReward(dropshipId string, customerId string, customerReward model.CustomerReward) (bool, error) {
	customerRewardCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/rewardPoints"
	_, err := a.client.Collection(customerRewardCollection).Doc(customerReward.Id).Set(a.ctx, customerReward)
	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerReward updated with key : %s", customerReward.Id)
	return true, nil
}

func (a customerRewardFirestoreRepository) DeleteCustomerReward(dropshipId string, customerId string, orderId string) (bool, error) {
	customerRewardCollection := "rewards/" + dropshipId + "/customers/" + customerId + "/rewardPoints"
	_, err := a.client.Collection(customerRewardCollection).Doc(orderId).Delete(a.ctx)
	if err != nil {
		log.Printf("An error while deleting data occurred: %s", err)
		return false, err
	}
	log.Printf("CustomerReward deleted")
	return true, nil
}
