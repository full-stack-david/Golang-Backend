package model

type CustomerSpent struct {
	Id          string  `firestore:"-" json:"id,omitempty"`
	InvoiceNo   string  `firestore:"invoiceNo" json:"invoiceNo,omitempty"`
	SpentPoints float32 `firestore:"spentPoints" json:"spentPoints,omitempty"`
}
