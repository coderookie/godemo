package main

import (
	"context"
	"encoding/json"
	"fmt"
	employee "godemo/models/protobuf_package/employee"
	searchProto "godemo/models/protobuf_package/search"
	"os"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	godemoGrpc "godemo/models/protobuf_package/grpc"
)

func main() {
	// 启动grpc的客户端
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc client, grpc.Dial error, error[%s]\n", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	// 初始化grpc服务的客户端
	client := godemoGrpc.NewGodemoServiceClient(conn)

	// 初始化context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()

	// 获取职员信息
	var id int64 = 1
	resp, err := client.GetEmployeeInfo(ctx, &employee.EmployeeRequest{
		Id: &id,
	})
	if err != nil {
		fmt.Printf("grpc client, client.GetEmployeeInfo error, error[%s]\n", err.Error())
		os.Exit(2)
	}
	var returnData []byte
	returnData, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		fmt.Printf("grpc client, json.Marshar1 employee response error, error[%s]\n", jsonErr.Error())
		os.Exit(3)
	}
	fmt.Printf("grpc client. GetEmployeeInfo, data[%s]\n", string(returnData))

	// 获取检索结果
	query := "query1"
	// query := ""
	searchResp, searchErr := client.Search(ctx, &searchProto.SearchRequest{
		Query: &query,
	})
	if searchErr != nil {
		fmt.Printf("grpc client. client.Search error, error[%s]\n", searchErr)
		os.Exit(3)
	}
	var searchReturnData []byte
	searchReturnData, searchJsonErr := json.Marshal(searchResp)
	if searchJsonErr != nil {
		fmt.Printf("grpc client, json.Marshal search response error, error[%s]\n", searchJsonErr.Error())
		os.Exit(4)
	}
	fmt.Printf("grpc client. search, data[%s]\n", string(searchReturnData))
}