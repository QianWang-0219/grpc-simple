package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	pb "grpc-simple/route"

	encode "grpc-simple/encodeImage"

	"google.golang.org/grpc"
)

const (
	task1_ipPort = "192.168.83.162:30033"
	task2_ipPort = "localhost:30031"
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
	runfunc(client, path)

}

func IsIPv4(ipAddr string) (string, error) {
	ip := net.ParseIP(ipAddr)
	if ip != nil && strings.Contains(ipAddr, ".") {
		return ipAddr, nil
	}
	return ipAddr, errors.New("IP address not valid")
}

func IsPort(input string) (string, error) {
	pattern := "\\d+"
	res, _ := regexp.MatchString(pattern, input)
	if res {
		return input, nil
	}
	return input, errors.New("Port number not valid")
}

func main() {
	var ipPort string
	fmt.Println("请选择您需要的服务:")
	fmt.Println("1. 手写数字识别")
	fmt.Println("2. 图像拼接")
	fmt.Println("3. 其他")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// 对输入进行判断
	if input == "1" {
		ipPort = task1_ipPort
	} else if input == "2" {
		ipPort = task2_ipPort
	} else {
		fmt.Println("您未选择预设任务!")
		fmt.Print("若您选择自定义任务，请输入自定义任务模块的ip地址(eg.127.0.0.1):")
		scanner.Scan()
		input = scanner.Text()
		ip, err := IsIPv4(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("请输入自定义任务模块的端口号:")
		scanner.Scan()
		input = scanner.Text()
		port, err := IsPort(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		ipPort = ip + ":" + port
	}
	//fmt.Println("YOU ENTERED:", ipPort)
	time.Sleep(2 * time.Second)
	Clientup(ipPort)
}
