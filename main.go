package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/log"

	"github.com/smallken/wallect-chain-account2/chaindispatcher"
	"github.com/smallken/wallect-chain-account2/config"
	wallet2 "github.com/smallken/wallect-chain-account2/rpc/account"
)

func main() {
	var f = flag.String("c", "config.yml", "config path")
	flag.Parse()
	conf, err := config.New(*f)
	if err != nil {
		panic(err)
	}
	dispatcher, err := chaindispatcher.New(conf)
	log.Info("dispatcher:", dispatcher)
	if err != nil {
		log.Error("Setup dispatcher failed", "err", err)
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(dispatcher.Interceptor))
	defer grpcServer.GracefulStop()

	wallet2.RegisterWalletAccountServiceServer(grpcServer, dispatcher)

	listen, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		log.Error("net listen failed", "err", err)
		panic(err)
	}
	reflection.Register(grpcServer)

	log.Info("dragon fucking greatest wallet rpc services start success", "port", conf.Server.Port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Error("grpc server serve failed", "err", err)
		panic(err)
	}
}
