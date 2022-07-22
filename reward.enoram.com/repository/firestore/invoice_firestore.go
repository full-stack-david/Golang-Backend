package firestore

import (
	"context"
	"mauappa-go/model"
	"mauappa-go/repository"

	"cloud.google.com/go/firestore"
)

type invoiceFirestoreRepository struct {
	client *firestore.Client
	ctx    context.Context
}

func NewInvoiceFirestoreRepository() (repository.InvoiceRepository, error) {
	client, ctx, err := GetDropshipDllFirestoreClient()
	if err != nil {
		return nil, err
	}
	return &invoiceFirestoreRepository{client: client, ctx: ctx}, nil
}

func (a invoiceFirestoreRepository) GetInvoice(dropshipId string, orderId string) (*model.Invoice, error) {
	invoiceCollection := "invoices/" + dropshipId + "/orders"
	dsnap, err := a.client.Collection(invoiceCollection).Doc(orderId).Get(a.ctx)
	if err != nil {
		return nil, err
	}
	var invoice model.Invoice
	dsnap.DataTo(&invoice)
	invoice.Id = dsnap.Ref.ID
	return &invoice, nil
}
