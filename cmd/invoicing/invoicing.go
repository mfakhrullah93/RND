package main

import (
	"company-finance-service/protos"
	"context"
	"fmt"
)

type InvoiceServer struct {
	protos.UnimplementedInvoicingServiceServer
}

func NewInvoicingServer() *InvoiceServer{
	return &InvoiceServer{}
}

func (is *InvoiceServer) CreateInvoice(ctx context.Context, ci*protos.CreateInvoiceRequest) (*protos.CreateInvoiceResponse, error){
	fmt.Println("[CreateInvoice]", ci)
	return &protos.CreateInvoiceResponse{}, nil
}

func (is *InvoiceServer) GetInvoice(ctx context.Context, gi*protos.GetInvoiceRequest) (*protos.GetInvoiceResponse, error){
	fmt.Println("[GetInvoice]", gi)
	invoice := &protos.Invoice{
		Id: gi.Id,
		InvoiceNumber: "INV0001",
		Items: []string{"item1", "item2"},
		InvoiceDate:"02022024",
		DueDate: "02022024",
		TotalAmount: 88.5,
		TaxAmount: 6.25,
		GrandTotal: 94.75,
		Status: "Paid",
	}

	return &protos.GetInvoiceResponse{Invoice: invoice}, nil
}

func (is *InvoiceServer) ReceivePayment(ctx context.Context, rp*protos.ReceivePaymentRequest) (*protos.ReceivePaymentResponse, error){
	fmt.Println("[ReceivePayment]", rp)
	// TODO Handle Create Ledger Entry 
	return &protos.ReceivePaymentResponse{}, nil
}