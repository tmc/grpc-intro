package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	_ "golang.org/x/net/trace"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tmc/grpc-intro/apidocs"
	"github.com/tmc/grpc-intro/protos/echoservice"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	httpAddr  = flag.String("addr", ":7000", "listen http addr")
	grpcAddr  = flag.String("grpcaddr", ":7001", "listen grpc addr")
	debugAddr = flag.String("debugaddr", ":7002", "listen debug addr")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := listenGRPC(*grpcAddr); err != nil {
		return err
	}

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := echoservice.RegisterEchoHandlerFromEndpoint(ctx, gwMux, *grpcAddr, opts)
	if err != nil {
		return err
	}
	go http.ListenAndServe(*debugAddr, nil)
	fmt.Println("listening")
	mux := http.NewServeMux()
	mux.Handle(apidocs.URLPrefixUI, apidocs.HandlerUI)
	mux.Handle(apidocs.URLPrefix, apidocs.Handler)
	mux.Handle("/", wsproxy.WebsocketProxy(gwMux))
	return http.ListenAndServe(*httpAddr, mux)
}

func listenGRPC(listenAddr string) error {
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	echoservice.RegisterEchoServer(grpcServer, &Server{})
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Println("serveGRPC err:", err)
		}
	}()
	return nil

}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
