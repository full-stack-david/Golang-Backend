package model

type CustomerDanaSegar struct {
	Id        string  `firestore:"-" json:"id,omitempty"`
	InvoiceNo string  `firestore:"invoiceNo" json:"invoiceNo,omitempty"`
	DanaSegar float32 `firestore:"danaSegar" json:"danaSegar,omitempty"`
}
