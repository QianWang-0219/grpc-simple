package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "grpc-simple/route"

	"google.golang.org/grpc"
)

const (
	task_ipPort = "localhost:30030"
	chunk_size  = 1024
)

func readIntFromCommendLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runUpfilefunc(client pb.LocalGuideClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	stream, err := client.UploadFile(ctx)
	if err != nil {
		log.Fatal("cannot upload file: ", err)
	}

	reader := bufio.NewReader(os.Stdin)
	// 实现一个交互，问client要信息，收到信息就给server进行处理

	var path string
	fmt.Print("请输入待拼接图像文件夹地址: ")
	readIntFromCommendLine(reader, &path)
	fmt.Println(path) // path输入的地址string

	fmt.Println(filepath.Ext(path))
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	fmt.Println("dddddd")
	defer file.Close()

	// metadata
	req := &pb.UploadFileRequest{
		Request: &pb.UploadFileRequest_Metadata{
			Metadata: &pb.MetaData{
				Filename:  "aaa",
				Extension: filepath.Ext(path),
			},
		},
	}
	// fmt.Println("eeeeee")
	// if err := stream.Send(req); err != nil {
	// 	log.Fatalln(err)
	// }
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send meta info to server: ", err, stream.RecvMsg(nil))
	}

	// chunk_data
	reader_file := bufio.NewReader(file)
	buffer := make([]byte, chunk_size)
	for {
		n, err := reader_file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadFileRequest{
			Request: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	fmt.Println(res)

	time.Sleep(20 * time.Second)
	//}
}

func Clientup(path string) {
	// 拨向端口，忽略证书验证（服务器没有提供证书），让dial变成blocking的，不要往下走
	conn, err := grpc.Dial(path, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	} else {
		fmt.Println("client dialed grpc server")
	}
	defer conn.Close()

	// 新建一个client
	client := pb.NewLocalGuideClient(conn)
	runUpfilefunc(client)

}

func main() {
	//var ipPort string
	ipPort := task_ipPort
	time.Sleep(2 * time.Second)
	Clientup(ipPort)
}
