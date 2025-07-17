package main

import "github.com/google/uuid"

func GetInfoClientDB(tenant, token string, pickupDB *PickupDB) error {
	pickupsRaw, _ := fetchPickup(tenant, token)
	invoicesRaw, _ := fetchPickupInvoice(tenant, token)
	originsRaw, _ := fetchPickupOrigin(tenant, token)
	destsRaw, _ := fetchPickupDestination(tenant, token)

	originMap := make(map[string]ResultCrudPickupOrigin)
	for _, o := range originsRaw {
		originMap[o.NFId] = o
	}

	destMap := make(map[string]ResultCrudPickupDestination)
	for _, d := range destsRaw {
		destMap[d.NFId] = d
	}

	type groupKey string
	pickupMap := make(map[groupKey]*struct {
		Pickup      Pickup
		Invoices    []PickupInvoice
		Origin      PickupOrigin
		Destination PickupDestination
	})

	for _, inv := range invoicesRaw {
		if inv.NFId == "" {
			continue
		}
		o := originMap[inv.NFId]
		d := destMap[inv.NFId]

		key := groupKey(o.CEP + "|" + d.CEP)

		item, exists := pickupMap[key]
		if !exists {
			guid := uuid.NewString()
			for _, pr := range pickupsRaw {
				if pr.PickupGUID != "" {
					guid = pr.PickupGUID
					break
				}
			}
			item = &struct {
				Pickup      Pickup
				Invoices    []PickupInvoice
				Origin      PickupOrigin
				Destination PickupDestination
			}{
				Pickup: Pickup{
					GUID:      guid,
					CompanyID: pickupsRaw[0].CompanyID,
				},
				Origin: PickupOrigin{
					GUID:          uuid.NewString(),
					PickupGUID:    guid,
					Cidade:        o.Cidade,
					Bairro:        o.Bairro,
					CEP:           o.CEP,
					NomeRemetente: o.NomeRemetente,
				},
				Destination: PickupDestination{
					GUID:             uuid.NewString(),
					PickupGUID:       guid,
					Cidade:           d.Cidade,
					Bairro:           d.Bairro,
					CEP:              d.CEP,
					NomeDestinatario: d.NomeDestinatario,
				},
			}
			pickupMap[key] = item
		}

		item.Invoices = append(item.Invoices, PickupInvoice{
			GUID:             uuid.NewString(),
			PickupGUID:       item.Pickup.GUID,
			NFId:             inv.NFId,
			Valor:            inv.Valor,
			NomeRemetente:    inv.NomeRemetente,
			NomeDestinatario: inv.NomeDestinatario,
		})
	}

	for _, v := range pickupMap {
		pickupDB.Pickups = append(pickupDB.Pickups, v.Pickup)
		pickupDB.Origins = append(pickupDB.Origins, v.Origin)
		pickupDB.Destinations = append(pickupDB.Destinations, v.Destination)
		pickupDB.Invoices = append(pickupDB.Invoices, v.Invoices...)
	}

	return nil
}

func fetchPickup(tenant, token string) ([]ResultCrudPickup, error) {
	return []ResultCrudPickup{
		{PickupGUID: "", CompanyID: "simfrete-log-2025"},
	}, nil
}

func fetchPickupInvoice(tenant, token string) ([]ResultCrudPickupInvoice, error) {
	return []ResultCrudPickupInvoice{
		{NFId: "NF001", Valor: 120.50, NomeDestinatario: "Supermercado ABC", NomeRemetente: "Distribuidora XPTO"},
		{NFId: "NF002", Valor: 89.90, NomeDestinatario: "Supermercado ABC", NomeRemetente: "Distribuidora XPTO"},
		{NFId: "NF003", Valor: 300.00, NomeDestinatario: "Loja Mega", NomeRemetente: "Armazém Central"},
		{NFId: "NF004", Valor: 150.75, NomeDestinatario: "Loja Mega", NomeRemetente: "Armazém Central"},
	}, nil
}

func fetchPickupOrigin(tenant, token string) ([]ResultCrudPickupOrigin, error) {
	return []ResultCrudPickupOrigin{
		{NFId: "NF001", Cidade: "São Paulo", Bairro: "Lapa", CEP: "05075-010", NomeRemetente: "Distribuidora XPTO"},
		{NFId: "NF002", Cidade: "São Paulo", Bairro: "Lapa", CEP: "05075-010", NomeRemetente: "Distribuidora XPTO"},
		{NFId: "NF003", Cidade: "Campinas", Bairro: "Centro", CEP: "13010-000", NomeRemetente: "Armazém Central"},
		{NFId: "NF004", Cidade: "Campinas", Bairro: "Centro", CEP: "13010-000", NomeRemetente: "Armazém Central"},
	}, nil
}

func fetchPickupDestination(tenant, token string) ([]ResultCrudPickupDestination, error) {
	return []ResultCrudPickupDestination{
		{NFId: "NF001", Cidade: "Osasco", Bairro: "Centro", CEP: "06010-060", NomeDestinatario: "Supermercado ABC"},
		{NFId: "NF002", Cidade: "Osasco", Bairro: "Centro", CEP: "06010-060", NomeDestinatario: "Supermercado ABC"},
		{NFId: "NF003", Cidade: "Jundiaí", Bairro: "Anhangabaú", CEP: "13208-000", NomeDestinatario: "Loja Mega"},
		{NFId: "NF004", Cidade: "Jundiaí", Bairro: "Anhangabaú", CEP: "13208-000", NomeDestinatario: "Loja Mega"},
	}, nil
}
