package model

type Customer struct {
	Id           string `firestore:"-" json:"id",omitempty"`
	Name         string `firestore:"name" json:"name",omitempty"`
	Email        string `firestore:"email" json:"email",omitempty"`
	PhoneNumber  string `firestore:"phoneNumber" json:"phoneNumber",omitempty"`
	ReferalCode  string `firestore:"referalCode" json:"referalCode",omitempty"`
	ReferBy      string `firestore:"referBy" json:"referBy",omitempty"`
	RewardPoints int    `firestore:"rewardPoints" json:"rewardPoints",omitempty"`
}
