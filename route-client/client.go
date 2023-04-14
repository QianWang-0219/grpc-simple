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
	ck "grpc-simple/utils"

	encode "grpc-simple/encodeImage"

	"google.golang.org/grpc"
)

const (
	//HOST         = "localhost"
	task1_ipPort = "30033" //图像识别
	task2_ipPort = "30031" //图像拼接
	task_ipPort  = "30030" //文件传输
	chunk_size   = 1024
)

func readIntFromCommendLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runfunc(client pb.LocalGuideClient, path string) {

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
			fmt.Println("处理结果:", location)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	// 实现一个交互，问client要信息，收到信息就给server进行处理
	for {
		request := pb.IniLoc{IniLocation: " "}
		var location string
		if path == task1_ipPort {
			// 对图像进行编码
			fmt.Print("请输入待识别图像: ")
			readIntFromCommendLine(reader, &location)
			req, _ := encode.EncodeImage(location)
			request.IniLocation = req
		} else {
			// 地址，无需编码
			fmt.Print("请输入待拼接图像文件夹地址: ")
			readIntFromCommendLine(reader, &location)
			request.IniLocation = location
		}
		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(60 * time.Second)
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
	fmt.Print("请输入待上传到服务端的文件地址: ")
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
	if path == task_ipPort {
		runUpfilefunc(client)
	} else {
		runfunc(client, path)
	}
}

func main() {
	var ipPort string
	fmt.Println("请选择您需要的服务:")
	fmt.Println("1. 手写数字识别")
	fmt.Println("2. 长图像拼接")
	fmt.Println("3. 上传文件")
	fmt.Println("4. 其他")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// 对输入进行判断
	if input == "1" {
		//ipPort = task1_ipPort
		fmt.Print("请输入服务器ip地址:")
		scanner.Scan()
		input := scanner.Text()
		HOST, err := ck.IsIPv4(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		ipPort = HOST + ":" + task1_ipPort
	} else if input == "2" {
		//ipPort = task2_ipPort
		fmt.Print("请输入服务器ip地址:")
		scanner.Scan()
		input := scanner.Text()
		HOST, err := ck.IsIPv4(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		ipPort = HOST + ":" + task2_ipPort
	} else if input == "3" {
		//ipPort = task_ipPort
		fmt.Print("请输入服务器ip地址:")
		scanner.Scan()
		input := scanner.Text()
		HOST, err := ck.IsIPv4(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		ipPort = HOST + ":" + task_ipPort
	} else {
		fmt.Println("您未选择预设任务!")
		fmt.Print("若您选择自定义任务，请输入自定义任务模块的ip地址(eg.127.0.0.1):")
		scanner.Scan()
		input = scanner.Text()
		HOST, err := ck.IsIPv4(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("请输入自定义任务模块的端口号:")
		scanner.Scan()
		input = scanner.Text()
		port, err := ck.IsPort(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		ipPort = HOST + ":" + port
	}
	//fmt.Println("YOU ENTERED:", ipPort)
	time.Sleep(2 * time.Second)
	Clientup(ipPort)
}
