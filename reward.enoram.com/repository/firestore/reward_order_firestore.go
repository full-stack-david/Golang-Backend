package firestore

import (
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type rewardOrderFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewRewardOrderFirestoreRepository() (repository.RewardOrderRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &rewardOrderFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a rewardOrderFirestoreRepository) GetRewardOrderList(dropshipId string) (*[]model.RewardOrder, error) {
	log.Printf("## Get reward order list of dropshipId = %s ##", dropshipId)

	rewardOrderPath := "rewards/" + dropshipId + "/orders"

	var rewardOrderArray []model.RewardOrder

	iter := a.client.Collection(rewardOrderPath).Documents(a.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var rewardOrder model.RewardOrder
		doc.DataTo(&rewardOrder)
		rewardOrder.Id = doc.Ref.ID
		rewardOrderArray = append(rewardOrderArray, rewardOrder)
	}

	return &rewardOrderArray, nil
}

func (a rewardOrderFirestoreRepository) GetRewardOrder(dropshipId string, orderId string) (*model.RewardOrder, error) {
	log.Printf("## Get RewardOrder By dropshipId and orderId = %s and %s ##", dropshipId, orderId)

	rewardOrderCollection := "rewards/" + dropshipId + "/orders"
	dsnap, err := a.client.Collection(rewardOrderCollection).Doc(orderId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var rewardOrder model.RewardOrder
	dsnap.DataTo(&rewardOrder)
	rewardOrder.Id = dsnap.Ref.ID
	return &rewardOrder, nil
}

func (a rewardOrderFirestoreRepository) CreateRewardOrder(dropshipId string, rewardOrderModel model.RewardOrder) (bool, error) {
	rewardOrderCollection := "rewards/" + dropshipId + "/orders"
	_, err := a.client.Collection(rewardOrderCollection).Doc(rewardOrderModel.Id).Set(a.ctx, rewardOrderModel)

	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardOrder created with key : %s", rewardOrderModel.Id)
	return true, nil
}

func (a rewardOrderFirestoreRepository) UpdateRewardOrder(dropshipId string, rewardOrder model.RewardOrder) (bool, error) {
	rewardOrderCollection := "rewards/" + dropshipId + "/orders"
	_, err := a.client.Collection(rewardOrderCollection).Doc(rewardOrder.Id).Set(a.ctx, rewardOrder)
	if err != nil {
		log.Printf("An error while storing data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardOrder updated with key : %s", rewardOrder.Id)
	return true, nil
}

func (a rewardOrderFirestoreRepository) DeleteRewardOrder(dropshipId string, orderId string) (bool, error) {
	rewardOrderCollection := "rewards/" + dropshipId + "/orders"
	_, err := a.client.Collection(rewardOrderCollection).Doc(orderId).Delete(a.ctx)
	if err != nil {
		log.Printf("An error while deleting data occurred: %s", err)
		return false, err
	}
	log.Printf("RewardOrder deleted")
	return true, nil
}
