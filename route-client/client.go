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
	ck "grpc-simple/utils"

	"google.golang.org/grpc"
)

const (
	//HOST         = "localhost"
	task_ipPort  = "30030" //文件上传&下载
	task1_ipPort = "30031" //图像识别
	task2_ipPort = "30032" //图像拼接
	chunk_size   = 1024
)

func readIntFromCommendLine(reader *bufio.Reader, target *string) {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runfunc(client pb.LocalGuideClient, input_num string) {
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
		switch input_num {
		case "3":
			// 图像识别任务，对图像进行编码
			fmt.Print("请输入待识别图像: ")
			readIntFromCommendLine(reader, &location)
			req, _ := ck.EncodeImage(location)
			request.IniLocation = req
		default:
			// 图像拼接任务，或自定义任务，传入地址无需编码
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
	var path string
	fmt.Print("请输入待上传到服务端的文件地址: ")
	readIntFromCommendLine(reader, &path)
	//fmt.Println(path) // path输入的地址string

	//fmt.Println(filepath.Ext(path))
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	_, fileType, fileNameOnly := ck.ReadFilename(path)
	// metadata
	req := &pb.UploadFileRequest{
		Request: &pb.UploadFileRequest_Metadata{
			Metadata: &pb.MetaData{
				Filename:  fileNameOnly,
				Extension: fileType,
			},
		},
	}
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
	time.Sleep(10 * time.Second)
}

func runDownfilefunc(client pb.LocalGuideClient) {
	reader := bufio.NewReader(os.Stdin)
	// 实现一个交互，问client要信息，收到信息就给server进行处理
	var path string
	fmt.Print("请输入待从服务器下载的文件地址: ")
	readIntFromCommendLine(reader, &path)
	fileNameWithSuffix, fileType, fileNameOnly := ck.ReadFilename(path)

	serverstream, err := client.DownloadFile(context.Background(), &pb.MetaData{
		Filename:  fileNameOnly,
		Extension: fileType,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//file, err := os.Create("download_file/" + fileNameWithSuffix)
	file, err := os.Create("download_file/" + fileNameWithSuffix)
	if err != nil {
		fmt.Println("文件创建失败 ", err.Error())
		return
	}
	defer file.Close()

	for {
		FileResponse, err := serverstream.Recv()
		if err == io.EOF { // stream关闭的时候
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(FileResponse)
		_, err = file.Write(FileResponse.ChunkData)
		if err != nil {
			fmt.Println("写入失败", err.Error())
			return
		}
	}
	fmt.Println("文件成功下载到", "download_file/", fileNameWithSuffix, "目录下！")
	time.Sleep(10 * time.Second)
}

func Clientup(path string, input_num string) {
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
	switch input_num {
	case "1": // 上传文件
		runUpfilefunc(client)
	case "2": // 下载文件
		runDownfilefunc(client)
	default:
		runfunc(client, input_num)
	}
}

func main() {
	var ipPort string
	var port string
	for {
		fmt.Println("服务列表:")
		fmt.Println("1. 上传文件")
		fmt.Println("2. 下载文件")
		fmt.Println("3. 手写数字识别")
		fmt.Println("4. 长图像拼接")
		fmt.Println("5. 其他")
		fmt.Println("6. 退出服务")

		fmt.Print("请输入你的选择: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_num := scanner.Text()

		// 对输入进行判断
		switch input_num {
		case "1", "2":
			// fmt.Print("请输入服务器ip地址:")
			// scanner.Scan()
			// input := scanner.Text()
			// HOST, err := ck.IsIPv4(input)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			HOST := "localhost"
			port = task_ipPort
			ipPort = HOST + ":" + port
		case "3":
			fmt.Print("请输入服务器ip地址:")
			scanner.Scan()
			input := scanner.Text()
			HOST, err := ck.IsIPv4(input)
			if err != nil {
				fmt.Println("输入ip地址格式错误！")
				return
			}
			port = task1_ipPort
			ipPort = HOST + ":" + port
		case "4":
			fmt.Print("请输入服务器ip地址:")
			scanner.Scan()
			input := scanner.Text()
			HOST, err := ck.IsIPv4(input)
			if err != nil {
				fmt.Println(err)
				return
			}
			port = task2_ipPort
			ipPort = HOST + ":" + port
		case "5":
			fmt.Println("您未选择预设任务!")
			fmt.Print("若您选择自定义任务，请输入自定义任务模块的ip地址(eg.127.0.0.1):")
			scanner.Scan()
			input := scanner.Text()
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
		case "6":
			return
		}
		time.Sleep(2 * time.Second)
		Clientup(ipPort, input_num)
	}
	fmt.Println("exit")
}
