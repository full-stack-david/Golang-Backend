package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"mauappa-go/model"
	"mauappa-go/repository"
)

const storeProfileCollection = "storeProfiles"

type storeProfileFirestoreRepository struct {
	client   *firestore.Client
	ctx      context.Context
}

func NewStoreProfileFirestoreRepository() (repository.StoreProfileRepository, error) {
	client, ctx, err := GetFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &storeProfileFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a storeProfileFirestoreRepository) GetStoreProfile(storeId string) (*model.StoreProfile, error) {
	log.Printf("## Get Store Profile detail of = %s ##", storeId)

	dsnap, err := a.client.Collection(storeProfileCollection).Doc(storeId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var storeProfile model.StoreProfile
	err = dsnap.DataTo(&storeProfile)
	if err != nil {
		return nil, err
	}
	storeProfile.Id = storeId

	return &storeProfile, nil
}
