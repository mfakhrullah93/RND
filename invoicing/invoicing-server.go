package invoicing

import (
	"company-finance-service/protos"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartInvoicingServer() {

	lis, err := net.Listen("tcp", ":1993")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	
	s := grpc.NewServer()
	invoiceServer := NewInvoicingServer()

	protos.RegisterInvoicingServiceServer(s, invoiceServer)	

	//! Development only
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
