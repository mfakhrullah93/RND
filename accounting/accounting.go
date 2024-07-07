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
	fmt.Println("CreateLedgerEntry: ", cle)
	return &protos.CreateLedgerEntryResponse{Id: "1"}, nil
}

func (acc *AccountServer) ListLedgerEntries(ctx context.Context, lle*protos.ListLedgerEntriesRequest) (*protos.ListLedgerEntriesResponse, error) {
	fmt.Println("GetLedgerList: ", "Start Date: ", lle.GetStartDate(), "End Date: ", lle.GetEndDate())
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