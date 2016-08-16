package main

import (
	"fmt"
	"time"

	"github.com/tmc/grpc-intro/protos/echoservice"
)

type Server struct{}

func (s *Server) Stream(_ *echoservice.Empty, stream echoservice.Echo_StreamServer) error {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		if err := stream.Send(&echoservice.EchoResponse{
			Message: "hello there!" + fmt.Sprint(time.Now().Sub(start)),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) Echo(srv echoservice.Echo_EchoServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			return err
		}
		if err := srv.Send(&echoservice.EchoResponse{
			Message: req.Message + "!",
		}); err != nil {
			return err
		}
	}
}

/*
func (s *Server) Heartbeats(srv echoservice.Echo_HeartbeatsServer) error {
	go func() {
		for {
			_, err := srv.Recv()
			if err != nil {
				log.Println("Recv() err:", err)
				return
			}
			log.Println("got hb from client")
		}
	}()
	t := time.NewTicker(time.Second * 1)
	for {
		log.Println("sending hb")
		hb := &echoservice.Heartbeat{
			Status: echoservice.Heartbeat_OK,
		}
		b := new(bytes.Buffer)
		if err := (&jsonpb.Marshaler{}).Marshal(b, hb); err != nil {
			log.Println("marshal err:", err)
		}
		log.Println(string(b.Bytes()))
		if err := srv.Send(hb); err != nil {
			return err
		}
		<-t.C
	}
	return nil
}
*/
