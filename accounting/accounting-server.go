package accounting

import (
	"log"
	"net"

	"company-finance-service/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartAccountingServer() {
	
	lis, err := net.Listen("tcp", ":1993")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	
	s := grpc.NewServer()
	accountServer := NewAccountServer()

	protos.RegisterAccountingServiceServer(s, accountServer)	

	//! Development only
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}