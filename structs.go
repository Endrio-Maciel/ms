package main

// --- structs auxiliares ---
type ResultCrudPickup struct {
	PickupGUID string
	CompanyID  string
}

type ResultCrudPickupInvoice struct {
	NFId             string
	PickupGUID       string
	Valor            float64
	NomeDestinatario string
	NomeRemetente    string
}

type ResultCrudPickupOrigin struct {
	NFId          string
	Cidade        string
	Bairro        string
	CEP           string
	NomeRemetente string
}

type ResultCrudPickupDestination struct {
	NFId             string
	Cidade           string
	Bairro           string
	CEP              string
	NomeDestinatario string
}

// --- modelos de domínio ---
type Pickup struct {
	GUID      string
	CompanyID string
}

type PickupInvoice struct {
	GUID             string
	PickupGUID       string
	NFId             string
	Valor            float64
	NomeDestinatario string
	NomeRemetente    string
}

type PickupOrigin struct {
	GUID          string
	PickupGUID    string
	Cidade        string
	Bairro        string
	CEP           string
	NomeRemetente string
}

type PickupDestination struct {
	GUID             string
	PickupGUID       string
	Cidade           string
	Bairro           string
	CEP              string
	NomeDestinatario string
}

type PickupDB struct {
	Pickups      []Pickup
	Invoices     []PickupInvoice
	Origins      []PickupOrigin
	Destinations []PickupDestination
}
