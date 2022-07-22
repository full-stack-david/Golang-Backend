package model

import "time"

type Invoice struct {
	Id              string    `firestore:"-" json:"id,omitempty"`
	InvoiceNo       string    `firestore:"invoiceNo" json:"invoiceNo,omitempty"`
	OrderNo         string    `firestore:"orderNo" json:"orderNo,omitempty"`
	PoNumber        string    `firestore:"poNumber" json:"poNumber,omitempty"`
	CustomerId      string    `firestore:"customerId" json:"customerId,omitempty"`
	CustomerName    string    `firestore:"customerName" json:"customerName,omitempty"`
	CreatedBy       string    `firestore:"createdBy" json:"createdBy,omitempty"`
	Created         time.Time `firestore:"created" json:"created,omitempty"`
	CreatedStr      string    `firestore:"createdStr" json:"createdStr,omitempty"`
	InvoiceBy       string    `firestore:"invoiceBy" json:"invoiceBy,omitempty"`
	OrderDate       time.Time `firestore:"orderDate" json:"orderDate,omitempty"`
	InvoiceDate     time.Time `firestore:"invoiceDate" json:"invoiceDate,omitempty"`
	DeliveryCost    float32   `firestore:"deliveryCost" json:"deliveryCost,omitempty"`
	Total           float32   `firestore:"total" json:"total,omitempty"`
	DanaSegar       float32   `firestore:"danaSegar" json:"danaSegar,omitempty"`
	UsedRewardPoint float32   `firestore:"usedRewardPoint" json:"usedRewardPoint,omitempty"`
}
