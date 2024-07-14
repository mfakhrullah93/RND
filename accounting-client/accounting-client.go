package accountingclient

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

func ServiceClient() (protos.AccountingServiceClient, *grpc.ClientConn){
	conn, err := grpc.NewClient("localhost:9901", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
			log.Fatalf("AccountingServiceClient did not connect: %v", err)
	}

	return protos.NewAccountingServiceClient(conn), conn
}

func HTTPMethod(e *echo.Echo, client protos.AccountingServiceClient) {

	// TODO CreateLedgerEntry
	e.POST("/accounting", func(c echo.Context) error {
				
		param := new(protos.CreateLedgerEntryRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.CreateLedgerEntry(context.TODO(), &protos.CreateLedgerEntryRequest{
			InvoiceId:   param.InvoiceId,
			TypeId: 		 param.TypeId,
			Description: param.Description,
			Amount:      param.Amount,
			Date:        param.Date,
		})
		
		if err != nil {
			fmt.Println("[CreateLedgerEntry]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})

	// TODO GetLedgerEntry
	e.GET("/accounting", func(c echo.Context) error {
				
		param := new(protos.GetLedgerEntryRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.GetLedgerEntry(context.TODO(), &protos.GetLedgerEntryRequest{
			Id: param.Id,
		})
		
		if err != nil {
			fmt.Println("[GetLedgerEntry]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})
	
	// TODO ListLedgerEntries
	e.POST("/accounting/list-entries", func(c echo.Context) error {
				
		param := new(protos.ListLedgerEntriesRequest)
		if err := c.Bind(param); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		res, err := client.ListLedgerEntries(context.TODO(), &protos.ListLedgerEntriesRequest{
			StartDate: param.StartDate, EndDate: param.EndDate,
		})
		
		if err != nil {
			fmt.Println("[ListLedgerEntries]", err)
			return c.JSON(http.StatusInternalServerError, err.Error());
		}

		return c.JSON(http.StatusOK, res);
	})
}