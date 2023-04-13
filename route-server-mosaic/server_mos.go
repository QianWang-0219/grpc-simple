package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"grpc-simple/image"
	pb "grpc-simple/route"

	"google.golang.org/grpc"
)

type localGuideServer struct {
	location []*pb.FinLoc
	pb.UnimplementedLocalGuideServer
}

func (s *localGuideServer) finalLocationOnce(request *pb.IniLoc) (*pb.FinLoc, error) {
	// 长图片拼接
	s.location[0].FinLocation = image.Image_mosaic(request.IniLocation)
	fmt.Println("...........")
	return s.location[0], nil
}

func (s *localGuideServer) GetLocation(stream pb.LocalGuide_GetLocationServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		finalLocation, err := s.finalLocationOnce(request)
		if err != nil {
			return err
		}
		err = stream.Send(finalLocation)
		if err != nil {
			return err
		}
	}
}

func newServer() *localGuideServer {
	fmt.Println("new server...")
	return &localGuideServer{
		location: []*pb.FinLoc{
			{FinLocation: "the first final location"},
			{FinLocation: "the second final location"},
		},
	}
}

func main() {
	//runNumGoroutineMonitor()
	// 新建一个服务的listener
	lis, err := net.Listen("tcp", "localhost:30031")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}
	// 新建一个grpc的server
	grpcServer := grpc.NewServer()
	// 注册一些服务
	pb.RegisterLocalGuideServer(grpcServer, newServer())
	// 起server
	log.Fatalln(grpcServer.Serve(lis))

}
