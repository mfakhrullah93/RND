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

	listLen := len(listInvoice)
	new_id := fmt.Sprintf("%v", listLen + 1)

	newEntry := &protos.Invoice{
		Id: new_id,
		InvoiceNumber: ci.InvoiceNumber,
		Items: ci.Items,
		InvoiceDate: ci.InvoiceDate,
		DueDate: ci.DueDate,
		TotalAmount: ci.TotalAmount,
		TaxAmount: ci.TaxAmount,
		GrandTotal: ci.GrandTotal,
		Status: ci.Status,
	}
	
	listInvoice = append(listInvoice, newEntry)

	return &protos.CreateInvoiceResponse{Id: new_id}, nil
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

var listInvoice = []*protos.Invoice{
	{
		Id: "1",
		InvoiceNumber: "INV0001",
		Items: []string{"item1", "item2"},
		InvoiceDate:"02022024",
		DueDate: "02022024",
		TotalAmount: 88.5,
		TaxAmount: 6.25,
		GrandTotal: 94.75,
		Status: "Paid",
	},
	{
		Id: "2",
		InvoiceNumber: "INV0002",
		Items: []string{"item1", "item2"},
		InvoiceDate:"02022024",
		DueDate: "02022024",
		TotalAmount: 88.5,
		TaxAmount: 6.25,
		GrandTotal: 94.75,
		Status: "Paid",
	},
}