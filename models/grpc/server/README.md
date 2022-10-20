grpc
1、server目录是rpc的服务端
2、client目录是rpc的客户端
3、cd models/grpc/server; go run main.go, 启动rpc服务端
4、cd models/grpc/client; go run main.go, 执行rpc的客户端验证效果
5、编写employee.proto文件
6、go get google.golang.org/grpc/cmd/protoc-gen-go-grpc, 下载protoc的go rpc插件
7、go install google.golang.org/grpc/cmd/protoc-gen-go-grpc, 编译protoc的go rpc插件, 最终生成的protoc-gen-go-grpc.exe移动到了GOPATH下
8、通过protoc和proto的go插件生成xxx.pb.go, go使用proto. cd /d/program/godemo/conf/proto/; /d/Program\ Files/protoc/bin/protoc.exe --plugin=protoc-gen-go=/c/Users/zhanglei16312/go/bin/protoc-gen-go.exe --go_out=./ employee.proto
9、通过protoc和proto的go插件生成xxx.pb.go, go使用proto. cd /d/program/godemo/conf/proto/; /d/Program\ Files/protoc/bin/protoc.exe --plugin=protoc-gen-go=/c/Users/zhanglei16312/go/bin/protoc-gen-go.exe --go_out=./ search.proto
10、通过protoc和proto的go插件生成grpc的xxx_grpc.pb.go代码. grpc的代码, plugin使用的是go grpc的插件. cd /d/program/godemo/conf/proto/; /d/Program\ Files/protoc/bin/protoc.exe --plugin=protoc-gen-go=/c/Users/zhanglei16312/go/bin/protoc-gen-go-grpc.exe --go-grpc_out=./ grpc.proto
11、mv godemo/models/protobuf_package/employee ../../models/protobuf_package, 将包放到合适的地方去
12、编写go代码. see models/grpc/server/main.go and models/grpc/client/main.go