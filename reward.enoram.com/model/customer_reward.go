package model

type CustomerReward struct {
	Id           string  `firestore:"-" json:"id,omitempty"`
	InvoiceNo    string  `firestore:"invoiceNo" json:"invoiceNo,omitempty"`
	RewardPoints float32 `firestore:"rewardPoints" json:"rewardPoints,omitempty"`
	Level        int     `firestore:"level" json:"level,omitempty"`
}
