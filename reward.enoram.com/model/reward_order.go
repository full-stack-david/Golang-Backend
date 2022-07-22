package model

import "time"

type RewardOrder struct {
	Id              string    `firestore:"-" json:"id,omitempty"`
	InvoiceNo       string    `firestore:"invoiceNo" json:"invoiceNo,omitempty"`
	InvoiceDate     time.Time `firestore:"invoiceDate" json:"invoiceDate,omitempty"`
	OrderNo         string    `firestore:"orderNo" json:"orderNo,omitempty"`
	OrderDate       time.Time `firestore:"orderDate" json:"orderDate,omitempty"`
	CustomerId      string    `firestore:"customerId" json:"customerId,omitempty"`
	CustomerCode    string    `firestore:"customerCode" json:"customerCode,omitempty"`
	CustomerName    string    `firestore:"customerName" json:"customerName,omitempty"`
	CreatedBy       string    `firestore:"createdBy" json:"createdBy,omitempty"`
	Total           float32   `firestore:"total" json:"total,omitempty"`
	User1           string    `firestore:"user1" json:"user1,omitempty"`
	Level1          float32   `firestore:"level1" json:"level1,omitempty"`
	User1Id         string    `firestore:"user1id" json:"user1id,omitempty"`
	User1Name       string    `firestore:"user1name" json:"user1name,omitempty"`
	User1Email      string    `firestore:"user1email" json:"user1email,omitempty"`
	User2           string    `firestore:"user2" json:"user2,omitempty"`
	Level2          float32   `firestore:"level2" json:"level2,omitempty"`
	User2Id         string    `firestore:"user2id" json:"user2id,omitempty"`
	User2Name       string    `firestore:"user2name" json:"user2name,omitempty"`
	User2Email      string    `firestore:"user2email" json:"user2email,omitempty"`
	User3           string    `firestore:"user3" json:"user3,omitempty"`
	Level3          float32   `firestore:"level3" json:"level3,omitempty"`
	User3Id         string    `firestore:"user3id" json:"user3id,omitempty"`
	User3Name       string    `firestore:"user3name" json:"user3name,omitempty"`
	User3Email      string    `firestore:"user4email" json:"user3email,omitempty"`
	DanaSegar       float32   `firestore:"danaSegar" json:"danaSegar,omitempty"`
	UsedRewardPoint float32   `firestore:"usedRewardPoint" json:"usedRewardPoint,omitempty"`
}
