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
	// 模拟二维拼接
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
	fmt.Println("new server2...")
	return &localGuideServer{
		location: []*pb.FinLoc{
			{FinLocation: "the first final location"},
			{FinLocation: "the second final location"},
		},
	}
}

// runNumGoroutineMonitor 协程数量监控
// func runNumGoroutineMonitor() {
// 	log.Printf("协程数量->%d\n", runtime.NumGoroutine())

// 	for {
// 		select {
// 		case <-time.After(time.Second):
// 			log.Printf("协程数量->%d\n", runtime.NumGoroutine())
// 		}
// 	}
// }

func main() {
	//runNumGoroutineMonitor()
	// 新建一个服务的listener
	lis, err := net.Listen("tcp", "localhost:30032")
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
