package main

import (
	"log"
	"os"

	"github.com/xvbnm48/learn-hexagonal/internal/adapters/app/api"
	"github.com/xvbnm48/learn-hexagonal/internal/adapters/core/arithmetic"
	"github.com/xvbnm48/learn-hexagonal/internal/adapters/framework/left/grpc"
	"github.com/xvbnm48/learn-hexagonal/internal/adapters/framework/right/db"
	"github.com/xvbnm48/learn-hexagonal/internal/ports"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DBPort
	var core ports.ArithmaticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbaseAdapterDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DB_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseAdapterDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate DB adapter: %v", err)
	}
	defer dbaseAdapter.CloseDBConnection()
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	grpcAdapter = grpc.NewAdapter(appAdapter)
	grpcAdapter.Run()
}
