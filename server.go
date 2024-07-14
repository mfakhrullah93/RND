package main

import (
	accounting "company-finance-service/accounting-client"
	invoicing "company-finance-service/invoicing-client"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//TODO Configure CORS to allow requests from http://wails.localhost:34115
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:4200",
	},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	accServerClient, accConnection := accounting.ServiceClient()	
	defer accConnection.Close()
	invoiceServerClient, accConnection := invoicing.ServiceClient()	
	defer accConnection.Close()
	
	accounting.HTTPMethod(e, accServerClient)
	invoicing.HTTPMethod(e, invoiceServerClient)

	//TODO END HTTP METHOD

	e.Logger.Fatal(e.Start("0.0.0.0:9900"))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
	}
}