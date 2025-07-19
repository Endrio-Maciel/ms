package types

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Pickup struct {
	PickupGUID uuid.UUID
	CompanyID  uuid.UUID
	CreatedAt  time.Time
	Status     string
}

type PickupInvoice struct {
	InvoiceID  uuid.UUID
	PickupGUID uuid.UUID
	NFNumber   string
	NFSeries   string
	EmittedAt  time.Time
	Value      float64
}

type PickupOrigin struct {
	OriginID   uuid.UUID
	PickupGUID uuid.UUID
	Address    string
	City       string
	State      string
}

type PickupDestination struct {
	DestinationID uuid.UUID
	PickupGUID    uuid.UUID
	Address       string
	City          string
	State         string
}

type PickupDB struct {
	Pickup             []Pickup
	PickupInvoices     []PickupInvoice
	PickupOrigins      []PickupOrigin
	PickupDestinations []PickupDestination
}

func (p Pickup) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"pickup_guid": p.PickupGUID.String(),
		"company_id":  p.CompanyID.String(),
		"created_at":  p.CreatedAt.Format(time.RFC3339),
		"status":      p.Status,
	}
}

func (pi PickupInvoice) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"invoice_id":  pi.InvoiceID.String(),
		"pickup_guid": pi.PickupGUID.String(),
		"nf_number":   pi.NFNumber,
		"nf_series":   pi.NFSeries,
		"emitted_at":  pi.EmittedAt.Format(time.RFC3339),
		"value":       pi.Value,
	}
}

func (po PickupOrigin) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"origin_id":   po.OriginID.String(),
		"pickup_guid": po.PickupGUID.String(),
		"address":     po.Address,
		"city":        po.City,
		"state":       po.State,
	}
}

func (pd PickupDestination) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"destination_id": pd.DestinationID.String(),
		"pickup_guid":    pd.PickupGUID.String(),
		"address":        pd.Address,
		"city":           pd.City,
		"state":          pd.State,
	}
}

func InsertAll(ctx context.Context, token string, db PickupDB) error {
	for _, p := range db.Pickup {
		if err := executaQuerier(token, ctx, "insertPickupSQL", p.ToMap()); err != nil {
			return err
		}
	}
	for _, i := range db.PickupInvoices {
		if err := executaQuerier(token, ctx, "insertPickupInvoiceSQL", i.ToMap()); err != nil {
			return err
		}
	}
	for _, o := range db.PickupOrigins {
		if err := executaQuerier(token, ctx, "insertPickupOriginSQL", o.ToMap()); err != nil {
			return err
		}
	}
	for _, d := range db.PickupDestinations {
		if err := executaQuerier(token, ctx, "insertPickupDestinationSQL", d.ToMap()); err != nil {
			return err
		}
	}
	return nil
}

func executaQuerier(t string, ctx context.Context, querie string, coisas any) error {

	return nil
}
