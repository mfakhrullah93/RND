package invoicingclient

import (
	"company-finance-service/protos"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceClient() (protos.InvoicingServiceClient, *grpc.ClientConn){
	conn, err := grpc.NewClient("localhost:9902", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
			log.Fatalf("AccountingServiceClient did not connect: %v", err)
	}

	return protos.NewInvoicingServiceClient(conn), conn
}

func HTTPMethod(e *echo.Echo, client protos.InvoicingServiceClient) {
	// TODO CreateInvoice
	e.POST("/invoicing", func(c echo.Context) error {
				
		param := new(protos.CreateInvoiceRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.CreateInvoice(context.TODO(), &protos.CreateInvoiceRequest{
			InvoiceNumber: param.InvoiceNumber,
			Items: param.Items,
			InvoiceDate: param.InvoiceDate,
			DueDate: param.DueDate,
			TotalAmount: param.TotalAmount,
			TaxAmount: param.TaxAmount,
			GrandTotal: param.GrandTotal,
		})
		
		if err != nil {
			fmt.Println("[CreateInvoice]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})

	// TODO GetInvoice
	e.GET("/invoicing", func(c echo.Context) error {
				
		param := new(protos.GetInvoiceRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.GetInvoice(context.TODO(), &protos.GetInvoiceRequest{
			Id: param.Id,
		})
		
		if err != nil {
			fmt.Println("[GetInvoice]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})

	// TODO GetInvoice
	e.POST("/invoicing/receive-payment", func(c echo.Context) error {
				
		param := new(protos.ReceivePaymentRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.ReceivePayment(context.TODO(), &protos.ReceivePaymentRequest{
			InvoiceId: param.InvoiceId,
			Amount: param.Amount,
			PaymentDate: param.PaymentDate,
		})
		
		if err != nil {
			fmt.Println("[ReceivePayment]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})
}