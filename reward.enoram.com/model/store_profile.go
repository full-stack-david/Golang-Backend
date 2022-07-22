package model

type StoreProfile struct {
	Id           		string `firestore:"-" json:"id",omitempty"`
	ShipmentCostGE      float32 `firestore:"shipmentCostGE" json:"shipmentCostGE",omitempty"`
	ShipmentCostValue      float32 `firestore:"shipmentCostValue" json:"shipmentCostValue",omitempty"`
}
