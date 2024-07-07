package accounting

import (
	protos "company-finance-service/protos"
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
	newId := fmt.Sprintf("%v", listLen + 1)

	// Append a new element
	newEntry := &protos.GetLedgerEntryResponse{
		Id:          newId,
		Type: 			 cle.Type,
		Description: cle.Description,
		Amount:      cle.Amount,
		Date:        cle.Date,
	}
	
	listLedger = append(listLedger, newEntry)

	return &protos.CreateLedgerEntryResponse{Id: newId}, nil
}

func (acc *AccountServer) GetLedgerEntry(ctx context.Context, gle*protos.GetLedgerEntryRequest) (*protos.GetLedgerEntryResponse, error) {
	fmt.Println("[GetLedgerEntry]", gle)
	return &protos.GetLedgerEntryResponse{Id: gle.Id, Type: "1", Description: "Tools", Amount: 1032.20, Date: "02022024"}, nil
}

func (acc *AccountServer) ListLedgerEntries(ctx context.Context, lle*protos.ListLedgerEntriesRequest) (*protos.ListLedgerEntriesResponse, error) {
	fmt.Println("[GetLedgerList]", "Start Date:", lle.GetStartDate(), "End Date:", lle.GetEndDate())
	return &protos.ListLedgerEntriesResponse{Entries: listLedger}, nil
}

var listLedger = []*protos.GetLedgerEntryResponse{
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