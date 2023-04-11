package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "grpc-simple/route"

	"google.golang.org/grpc"
)

func readIntFromCommendLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runfunc(client pb.LocalGuideClient) {
	stream, err := client.GetLocation(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// this routine listen to the server stream
	go func() {
		for {
			location, err2 := stream.Recv()
			if err2 == io.EOF {
				break
			}
			if err2 != nil {
				log.Fatalln(err2)
			}
			fmt.Println("二维拼接的存储地址:", location)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	// 实现一个交互，问client要信息，收到信息就给server进行处理
	for {
		request := pb.IniLoc{IniLocation: " "}
		var location string
		fmt.Print("Enter the inital location: ")
		readIntFromCommendLine(reader, &location)

		request.IniLocation = location
		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(10 * time.Second)
	}

}

func Clientup(path string) {
	// 拨向端口，忽略证书验证（服务器没有提供证书），让dial变成blocking的，不要往下走
	conn, err := grpc.Dial(path, grpc.WithInsecure(), grpc.WithBlock())
	fmt.Println("connected!")
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Close()

	// 新建一个client
	client := pb.NewLocalGuideClient(conn)

	runfunc(client)

}

func main() {
	Clientup("localhost:30031")
}
