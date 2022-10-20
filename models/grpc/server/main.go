package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"google.golang.org/grpc"
	employeeProto "godemo/models/protobuf_package/employee" // 基于employee.proto生成的grpc代码
	searchProto "godemo/models/protobuf_package/search"
	godemoGrpc "godemo/models/protobuf_package/grpc"
	userProto "godemo/models/protobuf_package/user"
)

// employee的grpc服务
type GrpcServer struct {
	godemoGrpc.UnimplementedGodemoServiceServer
}

// 获取员工信息
func (s *GrpcServer) GetEmployeeInfo(ctx context.Context, in *employeeProto.EmployeeRequest) (*employeeProto.EmployeeResponse, error) {
	// employee.proto中getEmployeeInfo方法参数和返回值中仅有一个, 这里通过context测试一下
	if ctx.Value("TestError") == 1 {
		return &employeeProto.EmployeeResponse{}, fmt.Errorf("getEmployeeInfo return error")
	}

	var response = &employeeProto.EmployeeResponse{}
	var id = in.Id
	response.Id = id
	username := "张磊"
	language := "golang"
	companyName := "汽车之家"
	companyGroup := "商业职能部-智能搜索团队-搜索工程组"
	companyPosition := "高级系统开发工程师"
	companyAddressProvince := "北京市"
	companyAddressCity := "北京市"
	companyAddressArea := "海淀区"
	companyAddressDetail := "北京市海淀区中国电子大厦"
	response.Username = &username
	response.Language = &language
	var company userProto.Company
	company.Name = &companyName
	company.Group = &companyGroup
	company.Position = &companyPosition
	company.Address = &userProto.Address{
		Province: &companyAddressProvince,
		City: &companyAddressCity,
		Area: &companyAddressArea,
		Address: &companyAddressDetail,
	}
	response.Company = &company
	return response, nil
}

// 搜索
func (s *GrpcServer) Search(ctx context.Context, request *searchProto.SearchRequest) (*searchProto.SearchResponse, error) {
	returnData := &searchProto.SearchResponse{}
	if len(*request.Query) == 0 {
		var code int32 = 100001
		message := "query为空"
		returnData.Code = &code
		returnData.Message = &message
		returnData.Data = nil
		return returnData, nil
	}


	var code int32 = 0
	message := "success"
	responseData := make([]*searchProto.SearchResponseData, 0)

	var id1 int64 = 1
	title1 := "title1"
	content1 := "content1"
	data1 := &searchProto.SearchResponseData{
		Id: &id1,
		Title: &title1,
		Content: &content1,
	}
	var id2 int64 = 1
	title2 := "title1"
	content2 := "content1"
	data2 := &searchProto.SearchResponseData{
		Id: &id2,
		Title: &title2,
		Content: &content2,
	}
	responseData = append(responseData, data1)
	responseData = append(responseData, data2)

	returnData.Code = &code
	returnData.Message = &message
	returnData.Data = responseData

	return returnData, nil
}

func main() {
	// 监听tcp的8081端口
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("grpc server, net.Listen error, error[%s]\n", err.Error())
		os.Exit(1)
	}

	// 启动一个rpc的server, 但是还未注册任何服务, 没有启动来接受请求
	grpc := grpc.NewServer()

	// 注册rpc的service
	godemoGrpc.RegisterGodemoServiceServer(grpc, &GrpcServer{})

	// 启动rpc的服务, 监听tcp的连接
	if err := grpc.Serve(listener); err != nil {
		fmt.Printf("grpc server, grpc.Serve error, error[%s]\n", err.Error())
		os.Exit(2)
	}
}