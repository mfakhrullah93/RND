package main

import (
	"company-finance-service/accounting"
	"company-finance-service/invoicing"
)

func main() {
	accounting.StartAccountingServer()
	invoicing.StartInvoicingServer()
}