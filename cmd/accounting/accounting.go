package main

import (
	"company-finance-service/protos"
	"context"
	"fmt"
)

type AccountServer struct{
	protos.UnimplementedAccountingServiceServer
}

func NewAccountServer() *AccountServer{
	return &AccountServer{}
}

func (acc *AccountServer) CreateLedgerEntry(ctx context.Context, cle*protos.CreateLedgerEntryRequest) (*protos.CreateLedgerEntryResponse, error) {
	fmt.Println("[CreateLedgerEntry]", cle)

	listLen := len(listLedger)
	new_id := fmt.Sprintf("%v", listLen + 1)

	newEntry := &protos.Ledger{
		Id:          new_id,
		InvoiceId:   cle.InvoiceId,
		TypeId: 		 cle.TypeId,
		Description: cle.Description,
		Amount:      cle.Amount,
		Date:        cle.Date,
	}
	
	listLedger = append(listLedger, newEntry)

	return &protos.CreateLedgerEntryResponse{Id: new_id}, nil
}

func (acc *AccountServer) GetLedgerEntry(ctx context.Context, gle*protos.GetLedgerEntryRequest) (*protos.GetLedgerEntryResponse, error) {
	fmt.Println("[GetLedgerEntry]", gle)

	newEntry := &protos.Ledger{
		Id:          gle.Id,
		TypeId: 		 "2",
		Description: "Tools",
		Amount:      1032.20,
		Date:        "02022024",
	}

	return &protos.GetLedgerEntryResponse{Ledger: newEntry}, nil
}

func (acc *AccountServer) ListLedgerEntries(ctx context.Context, lle*protos.ListLedgerEntriesRequest) (*protos.ListLedgerEntriesResponse, error) {
	fmt.Println("[GetLedgerList]", "Start Date:", lle.GetStartDate(), "End Date:", lle.GetEndDate())
	return &protos.ListLedgerEntriesResponse{Entries: listLedger}, nil
}

var listLedger = []*protos.Ledger{
	{
		Id:          "1",
		Description: "Salary Payment",
		Amount:      5000.0,
		Date:        "2024-07-07",
	},
	{
		Id:          "2",
		Description: "Office Rent",
		Amount:      1200.0,
		Date:        "2024-07-06",
	},
}