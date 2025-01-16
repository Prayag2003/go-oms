package main

import "github.com/Prayag2003/go-oms/services/orders/types"

func main() {
	grpcServer := types.NewGRPCServer(":9000")
	err := grpcServer.Start()
	if err != nil {
		return
	}
}
