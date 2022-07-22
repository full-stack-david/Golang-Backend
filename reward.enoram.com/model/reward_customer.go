package model

type RewardCustomer struct {
	Id                string  `firestore:"-" json:"id,omitempty"`
	Name              string  `firestore:"name" json:"name,omitempty"`
	Email             string  `firestore:"email" json:"email,omitempty"`
	ReferalCode       string  `firestore:"referalCode" json:"referalCode,omitempty"`
	TotalRewardPoints float32 `firestore:"totalRewardPoints" json:"totalRewardPoints,omitempty"`
	TotalSpentPoints  float32 `firestore:"totalSpentPoints" json:"totalSpentPoints,omitempty"`
	TotalDanaSegars   float32 `firestore:"totalDanaSegars" json:"totalDanarSegars,omitempty"`
	BalancePoints     float32 `firestore:"balancePoints" json:"balancePoints,omitempty"`
}
